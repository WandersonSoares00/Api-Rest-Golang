package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/db"
	"github.com/WandersonSoares00/Api-Rest-Golang.git/schema"
)

type Playlist schema.Playlist

func (p Playlist) Get(w http.ResponseWriter, filter ...string) error {

	len := len(filter)

	if len == 0 {
		return p.GetAll(w, "SELECT cod_play, nome, tempo_tot, data_criac FROM playlist")
	}
	if len == 2 {
		return p.GetAll(w, fmt.Sprintf("SELECT cod_play, nome, tempo_tot, data_criac FROM playlist WHERE %s = %s", filter[0], filter[1]))
	}

	sql := fmt.Sprintf(`SELECT cod_play, nome, tempo_tot, data_criac FROM playlist WHERE cod_play = %s`, filter[0])

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	row := conn.QueryRow(sql)

	//var p schema.Playlist

	err = row.Scan(&p.Cod, &p.Nome, &p.TempTot, &p.DtCriac)

	if err != nil {
		ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, "No data"))
		return nil
	}

	if err = row.Err(); err != nil {
		return err
	}

	if playlistJSON, err := json.Marshal(&p); err != nil {
		HandlerMessage := MessageToJson(false, "Error parsing the playlist data")
		ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return err
	} else {
		ReturnJsonResponse(w, http.StatusOK, playlistJSON)
	}

	return nil
}

func (p Playlist) GetAll(w http.ResponseWriter, sql string) error {
	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	rows, err := conn.Query(sql)

	if err != nil {
		return err
	}

	var playlists []Playlist

	for rows.Next() {

		err = rows.Scan(&p.Cod, &p.Nome, &p.TempTot, &p.DtCriac)

		if err != nil {
			return err
		}
		playlists = append(playlists, p)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	if playlistsJSON, err := json.Marshal(&playlists); err != nil {
		HandlerMessage := MessageToJson(false, "Error parsing the playlist data")
		ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return err
	} else {
		ReturnJsonResponse(w, http.StatusOK, playlistsJSON)
	}

	return nil
}

func (p Playlist) Create(w http.ResponseWriter, r *http.Request) error {
	//var p schema.Playlist

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		ReturnJsonResponse(w, http.StatusBadRequest, MessageToJson(false, "invalid input data"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	sql := `INSERT INTO playlist (cod_play, nome) VALUES ($1, $2) RETURNING cod_play`

	ret := conn.QueryRow(sql, p.Cod, p.Nome)

	if ret.Err() != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao tentar inserir dados apresentados"))
		return ret.Err()
	}

	ReturnJsonResponse(w, http.StatusCreated, MessageToJson(true, fmt.Sprintf("%d inserido com sucesso!", p.Cod)))

	return nil
}

func (p Playlist) Update(w http.ResponseWriter, r *http.Request, id int) error {
	
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "Error decoding json"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	sql := `UPDATE playlist SET nome=$2, tempo_tot=$3, data_criac=$4 WHERE cod_play=$1`

	ret, err := conn.Exec(sql, id, p.Nome, p.TempTot, p.DtCriac)

	if err != nil {
		return err
	}

	var qtd int64

	if qtd, err = ret.RowsAffected(); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao atualizar dados apresentados."))
		return err
	} else if qtd > 1 {
		err = fmt.Errorf("unexpected number of rows affected (%d) during insert operation for ID %d in the playlist table", qtd, id)
	}

	ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, fmt.Sprintf("%d atualizado com sucesso!", id)))

	return err
}

func (p Playlist) Delete(w http.ResponseWriter, r *http.Request) error {

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "Error decoding json"))
		return err
	}

	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	defer conn.Close()

	ret, err := conn.Exec(`DELETE FROM playlist WHERE cod_play=$1`, p.Cod)

	if err != nil {
		return err
	}

	var qtd int64

	if qtd, err = ret.RowsAffected(); err != nil {
		ReturnJsonResponse(w, http.StatusInternalServerError, MessageToJson(false, "erro ao atualizar dados apresentados."))
		return err
	} else if qtd > 1 {
		err = fmt.Errorf("unexpected number of rows affected (%d) during delete operation for ID %d in the playlist table", qtd, p.Cod)
	}

	ReturnJsonResponse(w, http.StatusOK, MessageToJson(true, fmt.Sprintf("%d removido com sucesso!", p.Cod)))

	return err
}
