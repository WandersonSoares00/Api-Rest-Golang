package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/db"
	"github.com/WandersonSoares00/Api-Rest-Golang.git/schema"
)

type Faixa_playlist schema.Faixa_playlist

func (f Faixa_playlist) Get(w http.ResponseWriter, filter ...string) error {

	len := len(filter)

	if len == 0 {
		return f.GetAll(w, `SELECT cod_faixa, cod_alb, cod_meio, cod_play, dt_ult_repr, qtd_repr FROM faixa_playlist`)
	}
	if len == 2 {
		return f.GetAll(w, fmt.Sprintf(`SELECT cod_faixa, cod_alb, cod_meio, cod_play, dt_ult_repr, qtd_repr FROM faixa_playlist WHERE %s = %s`, filter[0], filter[1]))
	}

	sql := fmt.Sprintf(`SELECT cod_faixa, cod_alb, cod_meio, cod_play, dt_ult_repr, qtd_repr FROM faixa_playlist WHERE cod_faixa = %s`, filter[0])

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	row := conn.QueryRow(sql)

	err = row.Scan(&f.Codfaixa, &f.CodAlb, &f.CodMeio, &f.CodPlay, &f.UltRep, &f.QtdRep)

	if err != nil {
		ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, "No data"))
		return nil
	}

	if err = row.Err(); err != nil {
		return err
	}

	if faixaJSON, err := json.Marshal(&f); err != nil {
		HandlerMessage := MessageToJson(false, "Error parsing the faixa_playlist data")
		ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return err
	} else {
		ReturnJsonResponse(w, http.StatusOK, faixaJSON)
	}

	return nil
}

func (f Faixa_playlist) GetAll(w http.ResponseWriter, sql string) error {
	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	rows, err := conn.Query(sql)

	if err != nil {
		return err
	}

	var faixas_play []Faixa_playlist

	for rows.Next() {

		err = rows.Scan(&f.Codfaixa, &f.CodAlb, &f.CodMeio, &f.CodPlay, &f.UltRep, &f.QtdRep)

		if err != nil {
			return err
		}
		faixas_play = append(faixas_play, f)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	if faixasJSON, err := json.Marshal(&faixas_play); err != nil {
		HandlerMessage := MessageToJson(false, "Error parsing the faixa_playlist data")
		ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return err
	} else {
		ReturnJsonResponse(w, http.StatusOK, faixasJSON)
	}

	return nil
}

func (f Faixa_playlist) Create(w http.ResponseWriter, r *http.Request) error {

	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		ReturnJsonResponse(w, http.StatusBadRequest, MessageToJson(false, "invalid input data"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	sql := `INSERT INTO faixa_playlist (cod_faixa, cod_alb, cod_meio, cod_play, dt_ult_repr, qtd_repr)
			VALUES ($1, $2, $3, $4, $5, $6) RETURNING cod_faixa`

	ret := conn.QueryRow(sql, f.Codfaixa, f.CodAlb, f.CodMeio, f.CodPlay, f.UltRep, f.QtdRep)

	if ret.Err() != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao tentar inserir dados apresentados"))
		return ret.Err()
	}

	ReturnJsonResponse(w, http.StatusCreated, MessageToJson(true, fmt.Sprintf("%d inserido com sucesso!", f.CodPlay)))

	return nil
}

func (f Faixa_playlist) Update(w http.ResponseWriter, r *http.Request, id int) error {

	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "Error decoding json"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	sql := `UPDATE faixa_playlist SET dt_ult-repr=$2, qtd_repr=$3 WHERE cod_faixa=$1`

	ret, err := conn.Exec(sql, id, f.UltRep, f.QtdRep)

	if err != nil {
		return err
	}

	var qtd int64

	if qtd, err = ret.RowsAffected(); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao atualizar dados apresentados."))
		return err
	} else if qtd > 1 {
		err = fmt.Errorf("unexpected number of rows affected (%d) during insert operation for ID %d in the faixa_playlist table", qtd, id)
	}

	ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, fmt.Sprintf("%d atualizado com sucesso!", id)))

	return err
}

func (f Faixa_playlist) Delete(w http.ResponseWriter, r *http.Request) error {

	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		ReturnJsonResponse(w, http.StatusBadRequest, MessageToJson(false, "invalid input data"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	ret, err := conn.Exec(`DELETE FROM faixa_playlist WHERE cod_faixa=$1 AND cod_alb=$2 AND cod_meio=$3 AND cod_play=$4`, f.Codfaixa, f.CodAlb, f.CodMeio, f.CodPlay)

	if err != nil {
		return err
	}

	var qtd int64

	if qtd, err = ret.RowsAffected(); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao atualizar dados apresentados."))
		return err
	} else if qtd > 1 {
		err = fmt.Errorf("unexpected number of rows affected (%d) during delete operation in the faixa_playlist table", qtd)
	}

	ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, "removido com sucesso!"))

	return err
}
