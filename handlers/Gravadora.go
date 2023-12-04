package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/db"
	"github.com/WandersonSoares00/Api-Rest-Golang.git/schema"
)

type Gravadora schema.Gravadora

func (g Gravadora) Get(w http.ResponseWriter, filter ...string) error {

	len := len(filter)
	if len == 0 {
		return g.GetAll(w, `SELECT cod_grav, nome, cidade, pais, end_homep FROM gravadora`)
	}
	if len == 2 {
		return g.GetAll(w, fmt.Sprintf(`SELECT cod_grav, nome, cidade, pais, end_homep FROM gravadora WHERE %s = %s`, filter[0], filter[1]))
	}

	sql := fmt.Sprintf(`SELECT cod_grav, nome, cidade, pais, end_homep FROM gravadora WHERE cod_grav = %s`, filter[0])

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	row := conn.QueryRow(sql)

	//var g schema.Gravadora
	err = row.Scan(&g.Cod, &g.Nome, &g.Cidade, &g.Pais, &g.HomeP)

	if err != nil {
		ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, "No data"))
		return nil
	}

	if err = row.Err(); err != nil {
		return err
	}

	if gravadoraJSON, err := json.Marshal(&g); err != nil {
		HandlerMessage := MessageToJson(false, "Error parsing the gravadora data")
		ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return err
	} else {
		ReturnJsonResponse(w, http.StatusOK, gravadoraJSON)
	}

	return nil
}

func (g Gravadora) GetAll(w http.ResponseWriter, sql string) error {
	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	rows, err := conn.Query(sql)

	if err != nil {
		return err
	}

	var gravadoras []Gravadora

	for rows.Next() {

		err = rows.Scan(&g.Cod, &g.Nome, &g.Cidade, &g.Pais, &g.HomeP)

		if err != nil {
			return err
		}
		gravadoras = append(gravadoras, g)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	if gravadorasJSON, err := json.Marshal(&gravadoras); err != nil {
		HandlerMessage := MessageToJson(false, "Error parsing the gravadora data")
		ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return err
	} else {
		ReturnJsonResponse(w, http.StatusOK, gravadorasJSON)
	}

	return nil
}

func (g Gravadora) Create(w http.ResponseWriter, r *http.Request) error {
	//var g schema.Gravadora

	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		ReturnJsonResponse(w, http.StatusBadRequest, MessageToJson(false, "invalid input data"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	sql := `INSERT INTO gravadora (nome, cidade, pais, end_homep)
			VALUES ($1, $2, $3, $4) RETURNING cod_grav`

	ret := conn.QueryRow(sql, g.Nome, g.Cidade, g.Pais, g.HomeP)

	if ret.Err() != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao tentar inserir dados apresentados"))
		return ret.Err()
	}

	ReturnJsonResponse(w, http.StatusCreated, MessageToJson(true, fmt.Sprintf("%d inserido com sucesso!", g.Cod)))

	return nil
}

func (g Gravadora) Update(w http.ResponseWriter, r *http.Request, id int) error {
	//var g schema.Gravadora

	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "Error decoding json"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	sql := `UPDATE gravadora SET nome=$2, cidade=$3, pais=$4 end_homep=$5 WHERE cod_grav=$1`

	ret, err := conn.Exec(sql, id, g.Nome, g.Cidade, g.Pais, g.HomeP)

	if err != nil {
		return err
	}

	var qtd int64

	if qtd, err = ret.RowsAffected(); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao atualizar dados apresentados."))
		return err
	} else if qtd > 1 {
		err = fmt.Errorf("unexpected number of rows affected (%d) during insert operation for ID %d in the gravadora table", qtd, id)
	}

	ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, fmt.Sprintf("%d atualizado com sucesso!", id)))

	return err
}

func (g Gravadora) Delete(w http.ResponseWriter, r *http.Request) error {

	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "Error decoding json"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	ret, err := conn.Exec(`DELETE FROM gravadora WHERE cod_grav=$1`, g.Cod)

	if err != nil {
		return err
	}

	var qtd int64

	if qtd, err = ret.RowsAffected(); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao atualizar dados apresentados."))
		return err
	} else if qtd > 1 {
		err = fmt.Errorf("unexpected number of rows affected (%d) during delete operation for ID %d in the gravadora table", qtd, g.Cod)
	}

	ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, fmt.Sprintf("%d removido com sucesso!", g.Cod)))

	return err
}
