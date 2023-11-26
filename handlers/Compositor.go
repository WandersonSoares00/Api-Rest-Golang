package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/db"
	"github.com/WandersonSoares00/Api-Rest-Golang.git/schema"
)

type Compositor schema.Compositor

func (c Compositor) GetAll(w http.ResponseWriter) error {
	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT cod_compositor, cod_pm, nome, dt_nasc, dt_morte FROM compositor`)

	if err != nil {
		return err
	}

	var compositors []Compositor

	for rows.Next() {

		err = rows.Scan(&c.Cod, &c.CodPm, &c.Nome, &c.DtNasc, &c.DtMort)

		if err != nil {
			return err
		}
		compositors = append(compositors, c)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	if compositorsJSON, err := json.Marshal(&compositors); err != nil {
		HandlerMessage := MessageToJson(false, "Error parsing the compositor data")
		ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return err
	} else {
		ReturnJsonResponse(w, http.StatusOK, compositorsJSON)
	}

	return nil
}

func (c Compositor) Get(w http.ResponseWriter, id int64) error {
	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT cod_compositor, cod_pm, nome, dt_nasc, dt_morte FROM compositor WHERE cod_compositor=$1`, id)

	//var c schema.Compositor
	err = row.Scan(&c.Cod, &c.CodPm, &c.Nome, &c.DtNasc, &c.DtMort)

	if err != nil {
		ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, "No data"))
		return nil
	}

	if err = row.Err(); err != nil {
		return err
	}

	if compositorJSON, err := json.Marshal(&c); err != nil {
		HandlerMessage := MessageToJson(false, "Error parsing the compositor data")
		ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return err
	} else {
		ReturnJsonResponse(w, http.StatusOK, compositorJSON)
	}

	return nil
}

func (c Compositor) Create(w http.ResponseWriter, r *http.Request) error {
	//var c schema.Compositor

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		ReturnJsonResponse(w, http.StatusBadRequest, MessageToJson(false, "invalid input data"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	sql := `INSERT INTO compositor (cod_compositor, cod_pm, nome, dt_nasc, dt_morte)
			VALUES ($1, $2, $3, $4, $5) RETURNING cod_compositor`

	ret := conn.QueryRow(sql, c.Cod, c.CodPm, c.Nome, c.DtNasc, c.DtMort)

	if ret.Err() != nil {
		ReturnJsonResponse(w, http.StatusBadRequest, MessageToJson(false, "erro ao tentar inserir dados apresentados"))
		return ret.Err()
	}

	ReturnJsonResponse(w, http.StatusBadRequest, MessageToJson(true, fmt.Sprintf("%d inserido com sucesso!", c.Cod)))

	return nil
}

func (c Compositor) Update(w http.ResponseWriter, r *http.Request, id int64) error {
	//var c schema.Compositor

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "Error decoding json"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	sql := `UPDATE compositor SET nome=$2, dt_nasc=$3, dt_morte=$4 WHERE cod_compositor=$1`

	ret, err := conn.Exec(sql, id, c.Nome, c.DtNasc, c.DtMort)

	if err != nil {
		return err
	}

	var qtd int64

	if qtd, err = ret.RowsAffected(); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao atualizar dados apresentados."))
		return err
	} else if qtd > 1 {
		err = fmt.Errorf("unexpected number of rows affected (%d) during insert operation for ID %d in the compositor table", qtd, id)
	}

	ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, fmt.Sprintf("%d atualizado com sucesso!", id)))

	return err
}

func (c Compositor) Delete(w http.ResponseWriter, id int64) error {

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	ret, err := conn.Exec(`DELETE FROM compositor WHERE cod_compositor=$1`, id)

	if err != nil {
		return err
	}

	var qtd int64

	if qtd, err = ret.RowsAffected(); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao atualizar dados apresentados."))
		return err
	} else if qtd > 1 {
		err = fmt.Errorf("unexpected number of rows affected (%d) during delete operation for ID %d in the compositor table", qtd, id)
	}

	ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, fmt.Sprintf("%d removido com sucesso!", id)))

	return err
}
