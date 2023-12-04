package handlers

import (
	"encoding/json"
	"net/http"
)

type Entity interface {
	Get(w http.ResponseWriter, filter ...string) error
	Create(w http.ResponseWriter, r *http.Request) error
	Update(w http.ResponseWriter, r *http.Request, id int) error
	Delete(w http.ResponseWriter, r *http.Request) error
}

func GetEntity(str string) Entity {
	switch str {
	case "gravadoras":
		return Gravadora{}
	case "albuns":
		return Album{}
	case "compositores":
		return Compositor{}
	case "faixas":
		return Faixa{}
	case "interpretes":
		return Interprete{}
	case "playlists":
		return Playlist{}
	case "faixasplaylists":
		return Faixa_playlist{}
	default:
		return nil
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
