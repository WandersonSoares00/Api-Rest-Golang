package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/db"
	"github.com/WandersonSoares00/Api-Rest-Golang.git/schema"
)

type Faixa schema.Faixa

func (f Faixa) GetAll(w http.ResponseWriter) error {
	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	sql := `SELECT cod_faixa, cod_cd, cod_vinil, cod_down, cod_composicao, numero, decricao, tempo_exec, tipo_grav FROM faixa`
	rows, err := conn.Query(sql)

	if err != nil {
		return err
	}

	var faixas []Faixa

	for rows.Next() {

		err = rows.Scan(&f.Cod, &f.CodCd, &f.CodVin, &f.CodDown, &f.CodComp, &f.Num, &f.Num, &f.Desc, &f.TExec, &f.TpGrav)

		if err != nil {
			return err
		}
		faixas = append(faixas, f)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	if faixasJSON, err := json.Marshal(&faixas); err != nil {
		HandlerMessage := MessageToJson(false, "Error parsing the faixa data")
		ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return err
	} else {
		ReturnJsonResponse(w, http.StatusOK, faixasJSON)
	}

	return nil
}

func (f Faixa) Get(w http.ResponseWriter, id int64) error {
	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT cod_faixa, cod_cd, cod_vinil, cod_down, cod_composicao, numero, decricao, tempo_exec, tipo_grav FROM faixa WHERE cod_faixa=$1`, id)

	//var f schema.Faixa
	err = row.Scan(&f.Cod, &f.CodCd, &f.CodVin, &f.CodDown, &f.CodComp, &f.Num, &f.Desc, &f.TExec, &f.TpGrav)

	if err != nil {
		ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, "No data"))
		return nil
	}

	if err = row.Err(); err != nil {
		return err
	}

	if faixaJSON, err := json.Marshal(&f); err != nil {
		HandlerMessage := MessageToJson(false, "Error parsing the faixa data")
		ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return err
	} else {
		ReturnJsonResponse(w, http.StatusOK, faixaJSON)
	}

	return nil
}

func (f Faixa) Create(w http.ResponseWriter, r *http.Request) error {
	//var f schema.Faixa

	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		ReturnJsonResponse(w, http.StatusBadRequest, MessageToJson(false, "invalid input data"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	sql := `INSERT INTO faixa (cod_faixa, cod_cd, cod_vinil, cod_down, cod_composicao, numero, decricao, tempo_exec, tipo_grav)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING cod_faixa`

	ret := conn.QueryRow(sql, f.Cod, f.CodCd, f.CodVin, f.CodDown, f.CodComp, f.Num, f.Desc, f.TExec, f.TpGrav)

	if ret.Err() != nil {
		ReturnJsonResponse(w, http.StatusBadRequest, MessageToJson(false, "erro ao tentar inserir dados apresentados"))
		return ret.Err()
	}

	ReturnJsonResponse(w, http.StatusBadRequest, MessageToJson(true, fmt.Sprintf("%d inserido com sucesso!", f.Cod)))

	return nil
}

func (f Faixa) Update(w http.ResponseWriter, r *http.Request, id int64) error {
	//var f schema.Faixa

	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "Error decoding json"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	sql := `UPDATE faixa SET numero=$2, descricao=$3, tempo_exec=$4, tipo_grav=$5 WHERE cod_faixa=$1`

	ret, err := conn.Exec(sql, id, f.Desc, f.TExec, f.TpGrav)

	if err != nil {
		return err
	}

	var qtd int64

	if qtd, err = ret.RowsAffected(); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao atualizar dados apresentados."))
		return err
	} else if qtd > 1 {
		err = fmt.Errorf("unexpected number of rows affected (%d) during insert operation for ID %d in the faixa table", qtd, id)
	}

	ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, fmt.Sprintf("%d atualizado com sucesso!", id)))

	return err
}

func (f Faixa) Delete(w http.ResponseWriter, id int64) error {

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	ret, err := conn.Exec(`DELETE FROM faixa WHERE cod_faixa=$1`, id)

	if err != nil {
		return err
	}

	var qtd int64

	if qtd, err = ret.RowsAffected(); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao atualizar dados apresentados."))
		return err
	} else if qtd > 1 {
		err = fmt.Errorf("unexpected number of rows affected (%d) during delete operation for ID %d in the faixa table", qtd, id)
	}

	ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, fmt.Sprintf("%d removido com sucesso!", id)))

	return err
}
