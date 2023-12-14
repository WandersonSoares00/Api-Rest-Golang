package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Entity interface {
	Scan(row *sql.Rows) error
	New() Entity
	SqlCreate() string
	SqlUpdate() string
	SqlDelete() string
	SqlQuery(filter string) string
	SqlQueryJoin(filter string) string
}

func init() {
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func MessageToJson(success bool, msg string) []byte {
	response := Response{
		Success: success,
		Message: msg,
	}

	jsData, _ := json.Marshal(response)

	return jsData
}

func ReturnJsonResponse(w http.ResponseWriter, httpCode int, resMessage []byte) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(resMessage)
}
