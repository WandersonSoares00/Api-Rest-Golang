package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/db"
)

func Get(e *Entity, w http.ResponseWriter, sql string) error {
	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	rows, err := conn.Query(sql)

	if err != nil {
		return err
	}

	var entities []Entity
	var countRows int = 0

	for rows.Next() {
		entity := (*e).New()

		err = entity.Scan(rows)

		if err != nil {
			return err
		}

		entities = append(entities, entity)
		countRows++
	}

	if err = rows.Err(); err != nil {
		return err
	}

	if countRows == 0 {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(true, "No data"))
	} else if entitiesJSON, err := json.Marshal(&entities); err != nil {
		HandlerMessage := MessageToJson(false, "Error parsing data")
		ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return err
	} else {
		ReturnJsonResponse(w, http.StatusOK, entitiesJSON)
	}

	return nil
}

func Create(e *Entity, w http.ResponseWriter, r *http.Request) error {

	if err := json.NewDecoder(r.Body).Decode(e); err != nil {
		ReturnJsonResponse(w, http.StatusBadRequest, MessageToJson(false, "invalid input data"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao tentar inserir dados apresentados"))
		return err
	}

	defer conn.Close()

	ret := conn.QueryRow((*e).SqlCreate())

	if ret.Err() != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao tentar inserir dados apresentados"))
		return ret.Err()
	}

	ReturnJsonResponse(w, http.StatusCreated, MessageToJson(true, "Dados inseridos com sucesso!"))

	return nil
}

func Update(e *Entity, w http.ResponseWriter, r *http.Request) error {

	if err := json.NewDecoder(r.Body).Decode(e); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "Error decoding json"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	ret, err := conn.Exec((*e).SqlUpdate())

	if err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao atualizar dados apresentados."))
		return err
	}

	if qtd, er := ret.RowsAffected(); er != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao atualizar dados apresentados."))
		return err
	} else if qtd > 1 {
		err = fmt.Errorf("unexpected number of rows affected (%d) during update operation", qtd)
	}

	ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, "Atualização realizada com sucesso!"))

	return err
}

func Delete(e *Entity, w http.ResponseWriter, r *http.Request) error {

	if err := json.NewDecoder(r.Body).Decode(e); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "Error decoding json"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	ret, err := conn.Exec((*e).SqlDelete())

	if err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao remover os dados apresentados."))
		return err
	}

	if qtd, er := ret.RowsAffected(); er != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao remover os dados apresentados."))
		return err
	} else if qtd > 1 {
		err = fmt.Errorf("unexpected number of rows affected (%d) during delete operation", qtd)
	}

	ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, "Remoção realizada com sucesso!"))

	return err
}
