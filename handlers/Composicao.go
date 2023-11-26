package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/db"
	"github.com/WandersonSoares00/Api-Rest-Golang.git/schema"
)

type Composicao schema.Composicao

func (c Composicao) GetAll(w http.ResponseWriter) error {
	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT cod_composicao, descricao, tipo FROM composicao`)

	if err != nil {
		return err
	}

	var composicoes []Composicao

	for rows.Next() {

		err = rows.Scan(&c.Cod, &c.Desc, &c.Tipo)

		if err != nil {
			return err
		}
		composicoes = append(composicoes, c)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	if composicoesJSON, err := json.Marshal(&composicoes); err != nil {
		HandlerMessage := MessageToJson(false, "Error parsing the composicao data")
		ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return err
	} else {
		ReturnJsonResponse(w, http.StatusOK, composicoesJSON)
	}

	return nil
}

func (c Composicao) Get(w http.ResponseWriter, id int64) error {
	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT cod_composicao, descricao, tipo FROM composicao WHERE cod_composicao=$1`, id)

	//var c schema.Composicao
	err = row.Scan(&c.Cod, &c.Cod, &c.Desc, &c.Tipo)

	if err != nil {
		ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, "No data"))
		return nil
	}

	if err = row.Err(); err != nil {
		return err
	}

	if composicaoJSON, err := json.Marshal(&c); err != nil {
		HandlerMessage := MessageToJson(false, "Error parsing the composicao data")
		ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return err
	} else {
		ReturnJsonResponse(w, http.StatusOK, composicaoJSON)
	}

	return nil
}

func (c Composicao) Create(w http.ResponseWriter, r *http.Request) error {
	//var c schema.Composicao

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		ReturnJsonResponse(w, http.StatusBadRequest, MessageToJson(false, "invalid input data"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	sql := `INSERT INTO composicao (cod_composicao, descricao, tipo)
			VALUES ($1, $2, $3, $4, $5) RETURNING cod_composicao`

	ret := conn.QueryRow(sql, c.Cod, c.Desc, c.Tipo)

	if ret.Err() != nil {
		ReturnJsonResponse(w, http.StatusBadRequest, MessageToJson(false, "erro ao tentar inserir dados apresentados"))
		return ret.Err()
	}

	ReturnJsonResponse(w, http.StatusBadRequest, MessageToJson(true, fmt.Sprintf("%d inserido com sucesso!", c.Cod)))

	return nil
}

func (c Composicao) Update(w http.ResponseWriter, r *http.Request, id int64) error {
	//var c schema.Composicao

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "Error decoding json"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	sql := `UPDATE composicao SET descricao=$2, tipo=$3 WHERE cod_composicao=$1`

	ret, err := conn.Exec(sql, id, c.Desc, c.Tipo)

	if err != nil {
		return err
	}

	var qtd int64

	if qtd, err = ret.RowsAffected(); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao atualizar dados apresentados."))
		return err
	} else if qtd > 1 {
		err = fmt.Errorf("unexpected number of rows affected (%d) during insert operation for ID %d in the composicao table", qtd, id)
	}

	ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, fmt.Sprintf("%d atualizado com sucesso!", id)))

	return err
}

func (c Composicao) Delete(w http.ResponseWriter, id int64) error {

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	ret, err := conn.Exec(`DELETE FROM composicao WHERE cod_composicao=$1`, id)

	if err != nil {
		return err
	}

	var qtd int64

	if qtd, err = ret.RowsAffected(); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao atualizar dados apresentados."))
		return err
	} else if qtd > 1 {
		err = fmt.Errorf("unexpected number of rows affected (%d) during delete operation for ID %d in the composicao table", qtd, id)
	}

	ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, fmt.Sprintf("%d removido com sucesso!", id)))

	return err
}
