package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/db"
	"github.com/WandersonSoares00/Api-Rest-Golang.git/schema"
)

type Interprete schema.Interprete

func (i Interprete) Get(w http.ResponseWriter, filter ...string) error {

	len := len(filter)

	if len == 0 {
		return i.GetAll(w, `SELECT cod_inter, nome, tipo FROM interprete`)
	}
	if len == 2 {
		return i.GetAll(w, fmt.Sprintf(`SELECT cod_inter, nome, tipo FROM interprete WHERE %s = %s`, filter[0], filter[1]))
	}

	sql := fmt.Sprintf(`SELECT cod_inter, nome, tipo FROM interprete WHERE cod_inter = %s`, filter[0])

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	row := conn.QueryRow(sql)

	//var i schema.Interprete
	err = row.Scan(&i.Cod, &i.Nome, &i.Tipo)

	if err != nil {
		ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, "No data"))
		return nil
	}

	if err = row.Err(); err != nil {
		return err
	}

	if interpreteJSON, err := json.Marshal(&i); err != nil {
		HandlerMessage := MessageToJson(false, "Error parsing the interprete data")
		ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return err
	} else {
		ReturnJsonResponse(w, http.StatusOK, interpreteJSON)
	}

	return nil
}

func (i Interprete) GetAll(w http.ResponseWriter, sql string) error {
	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT cod_inter, nome, tipo FROM interprete`)

	if err != nil {
		return err
	}

	var composicoes []Interprete

	for rows.Next() {

		err = rows.Scan(&i.Cod, &i.Nome, &i.Tipo)

		if err != nil {
			return err
		}
		composicoes = append(composicoes, i)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	if interpretesJSON, err := json.Marshal(&composicoes); err != nil {
		HandlerMessage := MessageToJson(false, "Error parsing the interprete data")
		ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return err
	} else {
		ReturnJsonResponse(w, http.StatusOK, interpretesJSON)
	}

	return nil
}

func (i Interprete) Create(w http.ResponseWriter, r *http.Request) error {
	//var i schema.Interprete

	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		ReturnJsonResponse(w, http.StatusBadRequest, MessageToJson(false, "invalid input data"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	sql := `INSERT INTO interprete (nome, tipo)
			VALUES ($1, $2) RETURNING cod_inter`

	ret := conn.QueryRow(sql, i.Nome, i.Tipo)

	if ret.Err() != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao tentar inserir dados apresentados"))
		return ret.Err()
	}

	ReturnJsonResponse(w, http.StatusCreated, MessageToJson(true, fmt.Sprintf("%d inserido com sucesso!", i.Cod)))

	return nil
}

func (i Interprete) Update(w http.ResponseWriter, r *http.Request, id int) error {
	//var i schema.Interprete

	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "Error decoding json"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	sql := `UPDATE interprete SET nome=$2, tipo=$3 WHERE cod_inter=$1`

	ret, err := conn.Exec(sql, id, i.Nome, i.Tipo)

	if err != nil {
		return err
	}

	var qtd int64

	if qtd, err = ret.RowsAffected(); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao atualizar dados apresentados."))
		return err
	} else if qtd > 1 {
		err = fmt.Errorf("unexpected number of rows affected (%d) during insert operation for ID %d in the interprete table", qtd, id)
	}

	ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, fmt.Sprintf("%d atualizado com sucesso!", id)))

	return err
}

func (i Interprete) Delete(w http.ResponseWriter, r *http.Request) error {

	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		ReturnJsonResponse(w, http.StatusBadRequest, MessageToJson(false, "invalid input data"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	ret, err := conn.Exec(`DELETE FROM interprete WHERE cod_inter=$1`, i.Cod)

	if err != nil {
		return err
	}

	var qtd int64

	if qtd, err = ret.RowsAffected(); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao atualizar dados apresentados."))
		return err
	} else if qtd > 1 {
		err = fmt.Errorf("unexpected number of rows affected (%d) during delete operation for ID %d in the interprete table", qtd, i.Cod)
	}

	ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, fmt.Sprintf("%d removido com sucesso!", i.Cod)))

	return err
}
