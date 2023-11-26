package handlers

import (
	"encoding/json"
	"net/http"
)

type Entity interface {
	GetAll(w http.ResponseWriter) error
	Get(w http.ResponseWriter, id int64) error
	Create(w http.ResponseWriter, r *http.Request) error
	Update(w http.ResponseWriter, r *http.Request, id int64) error
	Delete(w http.ResponseWriter, id int64) error
}

func GetEntity(str string) Entity {
	switch str {
	case "gravadoras":	return Gravadora{}
	default:			return nil
	}
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
