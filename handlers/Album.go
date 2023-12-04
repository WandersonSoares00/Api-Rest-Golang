package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/db"
	"github.com/WandersonSoares00/Api-Rest-Golang.git/schema"
)

type Album schema.Album

func (alb Album) Get(w http.ResponseWriter, filter ...string) error {

	len := len(filter)
	if len == 0 {
		return alb.GetAll(w, `SELECT cod_alb, cod_meio, cod_grav, nome, descricao, data_grav, pr_compra, pr_venda, meio FROM album`)
	}
	if len == 2 {
		return alb.GetAll(w, fmt.Sprintf(`SELECT cod_alb, cod_meio, cod_grav, nome, descricao, data_grav, pr_compra, pr_venda, meio FROM album WHERE %s = %s`, filter[0], filter[1]))
	}

	sql := fmt.Sprintf(`SELECT cod_alb, cod_meio, cod_grav, nome, descricao, data_grav, pr_compra, pr_venda, meio FROM album WHERE cod_alb = %s`, filter[0])

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	row := conn.QueryRow(sql)

	//var alb schema.Album
	err = row.Scan(&alb.Cod, &alb.CodMeio, &alb.CodGrav, &alb.Nome, &alb.Desc, &alb.DtGrav, &alb.PrComp, &alb.PrVen, &alb.Meio)

	if err != nil {
		ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, "No data"))
		return nil
	}

	if err = row.Err(); err != nil {
		return err
	}

	if albunsJSON, err := json.Marshal(&alb); err != nil {
		HandlerMessage := MessageToJson(false, "Error parsing the album data")
		ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return err
	} else {
		ReturnJsonResponse(w, http.StatusOK, albunsJSON)
	}

	return nil
}

func (alb Album) GetAll(w http.ResponseWriter, sql string) error {
	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	rows, err := conn.Query(sql)

	if err != nil {
		return err
	}

	var albuns []Album

	for rows.Next() {

		err = rows.Scan(&alb.Cod, &alb.CodMeio, &alb.CodGrav, &alb.Nome, &alb.Desc, &alb.DtGrav, &alb.PrComp, &alb.PrVen, &alb.Meio)

		if err != nil {
			return err
		}
		albuns = append(albuns, alb)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	if albunsJSON, err := json.Marshal(&albuns); err != nil {
		HandlerMessage := MessageToJson(false, "Error parsing the album data")
		ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return err
	} else {
		ReturnJsonResponse(w, http.StatusOK, albunsJSON)
	}

	return nil
}

func (alb Album) Create(w http.ResponseWriter, r *http.Request) error {
	//var alb schema.Album

	if err := json.NewDecoder(r.Body).Decode(&alb); err != nil {
		ReturnJsonResponse(w, http.StatusBadRequest, MessageToJson(false, "invalid input data"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	sql := `INSERT INTO album (cod_meio, cod_grav, nome, descricao, data_grav, pr_compra, pr_venda, meio)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING cod_alb`

	ret := conn.QueryRow(sql, &alb.CodMeio, &alb.CodGrav, &alb.Nome, &alb.Desc, &alb.DtGrav, &alb.PrComp, &alb.PrVen, &alb.Meio)

	if ret.Err() != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao tentar inserir dados apresentados"))
		return ret.Err()
	}

	ReturnJsonResponse(w, http.StatusCreated, MessageToJson(true, fmt.Sprintf("%d inserido com sucesso!", alb.Cod)))

	return nil
}

func (alb Album) Update(w http.ResponseWriter, r *http.Request, id int) error {
	//var alb schema.Album

	if err := json.NewDecoder(r.Body).Decode(&alb); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "Error decoding json"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	sql := `UPDATE album SET nome= $2 descricao=$3, data_grav=$4, pr_compra=$5 pr_venda=$6 WHERE cod_alb=$1`

	ret, err := conn.Exec(sql, id, alb.Nome, alb.Desc, alb.DtGrav, alb.PrComp, alb.PrVen)

	if err != nil {
		return err
	}

	var qtd int64

	if qtd, err = ret.RowsAffected(); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao atualizar dados apresentados."))
		return err
	} else if qtd > 1 {
		err = fmt.Errorf("unexpected number of rows affected (%d) during insert operation for ID %d in the album table", qtd, id)
	}

	ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, fmt.Sprintf("%d atualizado com sucesso!", id)))

	return err
}

func (alb Album) Delete(w http.ResponseWriter, r *http.Request) error {

	if err := json.NewDecoder(r.Body).Decode(&alb); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "Error decoding json"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	ret, err := conn.Exec(`DELETE FROM album WHERE cod_alb=$1 AND cod_grav = $2 AND cod_meio = $3`, alb.Cod, alb.CodGrav, alb.CodMeio)

	if err != nil {
		return err
	}

	var qtd int64

	if qtd, err = ret.RowsAffected(); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao atualizar dados apresentados."))
		return err
	} else if qtd > 1 {
		err = fmt.Errorf("unexpected number of rows affected (%d) during delete operation for ID %d in the album table", qtd, alb.Cod)
	}

	ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, fmt.Sprintf("%d removido com sucesso!", alb.Cod)))

	return err
}
