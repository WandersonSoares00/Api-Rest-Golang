package routes

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/configs"
	"github.com/WandersonSoares00/Api-Rest-Golang.git/handlers"
)

var logger *configs.Logger

func init() {
	logger = configs.NewLogger("Routes")
}

func Serve(w http.ResponseWriter, r *http.Request) {

	var entity string
	var id int64
	var err error = nil
	p := r.URL.Path

	start := time.Now()

	switch {
	case match(p, "^([a-z]+)/?$", &entity):
		t := handlers.GetEntity(entity)

		if t == nil {
			msg := fmt.Sprintf("Entidade %s requisitada não definida", entity)
			handlers.ReturnJsonResponse(w, http.StatusBadRequest, handlers.MessageToJson(false, msg))
		} else if r.Method == http.MethodGet {
			err = t.GetAll(w)
		} else if r.Method == http.MethodPost {
			err = t.Create(w, r)
		} else {
			msg := "Método e requisição solicitados são incompatíveis."
			handlers.ReturnJsonResponse(w, http.StatusMethodNotAllowed, handlers.MessageToJson(false, msg))
		}
	case match(p, "^([a-z]+)/\\d+$", &entity, &id):
		t := handlers.GetEntity(entity)

		if t == nil {
			msg := fmt.Sprintf("Entidade %s requisitada não definida", entity)
			handlers.ReturnJsonResponse(w, http.StatusBadRequest, handlers.MessageToJson(false, msg))
		} else if r.Method == http.MethodGet {
			err = t.Get(w, id)
		} else if r.Method == http.MethodDelete {
			err = t.Delete(w, id)
		} else if r.Method == http.MethodPut {
			err = t.Update(w, r, id)
		} else {
			msg := "Método e requisição solicitados são incompatíveis."
			handlers.ReturnJsonResponse(w, http.StatusMethodNotAllowed, handlers.MessageToJson(false, msg))
		}

	default:
		fmt.Println("não matchiou")
	}

	time := time.Since(start).Microseconds()
	requestStr := fmt.Sprintf("Method: %s Url: %s", r.Method, r.URL)

	if err != nil {
		logger.Errorf("%s failed with %dms and %s", requestStr, time, err.Error())
	} else {
		logger.Infof("%s resolved with %dms\n", requestStr, time)
	}

}

func match(path, pattern string, vars ...interface{}) bool {
	regex, err := regexp.Compile(pattern)

	if err != nil {
		println("nil Compile")
		return false
	}

	matches := regex.FindStringSubmatch(path)

	if len(matches) <= 0 {
		return false
	}

	for i, match := range matches[1:] {
		switch p := vars[i].(type) {
		case *string:
			*p = match
		case *int:
			n, err := strconv.Atoi(match)
			if err != nil {
				fmt.Println("type inv")
				return false
			}
			*p = n
		default:
			println("type ummatch")
			return false
		}
	}

	return true
}
