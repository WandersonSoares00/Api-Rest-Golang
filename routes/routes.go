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
	var id int
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
			err = t.Get(w)
		} else if r.Method == http.MethodDelete {
			err = t.Delete(w, r)
		} else if r.Method == http.MethodPost {
			err = t.Create(w, r)
		} else {
			msg := "Método e requisição solicitados são incompatíveis."
			handlers.ReturnJsonResponse(w, http.StatusMethodNotAllowed, handlers.MessageToJson(false, msg))
		}
	case match(p, "^([a-z]+)/([0-9]+)+$", &entity, &id):
		t := handlers.GetEntity(entity)

		if t == nil {
			msg := fmt.Sprintf("Entidade %s requisitada não definida", entity)
			handlers.ReturnJsonResponse(w, http.StatusBadRequest, handlers.MessageToJson(false, msg))
		} else if r.Method == http.MethodGet {
			err = t.Get(w, strconv.Itoa(id))
		} else if r.Method == http.MethodPut {
			err = t.Update(w, r, id)
		} else {
			msg := "Método e requisição solicitados são incompatíveis."
			handlers.ReturnJsonResponse(w, http.StatusMethodNotAllowed, handlers.MessageToJson(false, msg))
		}
	case match(p, "^faixas/albuns/([0-9]+)+$", &id) && r.Method == http.MethodGet:
		t := handlers.GetEntity("faixas")
		err = t.Get(w, "cod_alb", strconv.Itoa(id))
	default:
		handlers.ReturnJsonResponse(w, http.StatusBadRequest, handlers.MessageToJson(false, "url inválida."))
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
				return false
			}
			*p = n
		default:
			return false
		}
	}

	return true
}
