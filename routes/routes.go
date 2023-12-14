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

	var err error = nil

	start := time.Now()

	entity, sql := parseUrl(r)

	if entity == nil {
		msg := "Entidade requisitada não definida"
		handlers.ReturnJsonResponse(w, http.StatusBadRequest, handlers.MessageToJson(false, msg))
	} else if r.Method == http.MethodGet {
		err = handlers.Get(&entity, w, sql)
	} else if r.Method == http.MethodPost {
		err = handlers.Create(&entity, w, r)
	} else if r.Method == http.MethodPut {
		err = handlers.Update(&entity, w, r)
	} else if r.Method == http.MethodDelete {
		err = handlers.Delete(&entity, w, r)
	} else {
		msg := "Método e requisição solicitados são incompatíveis."
		handlers.ReturnJsonResponse(w, http.StatusMethodNotAllowed, handlers.MessageToJson(false, msg))
	}

	time := time.Since(start).Microseconds()
	requestStr := fmt.Sprintf("Method: %s Url: %s", r.Method, r.URL)

	if err != nil {
		logger.Errorf("%s failed with %dms and %s", requestStr, time, err.Error())
	} else {
		logger.Infof("%s resolved with %dms\n", requestStr, time)
	}

}

func parseUrl(r *http.Request) (entity handlers.Entity, sql string) {
	url := r.URL.Path
	entity = nil
	var id string

	//presume-se, nesse ponto, que toda request(exceto "^playlists/([0-9]+)+/faixas/?$") é Get...
	switch {
	case match(url, "^gravadoras/?$"):
		entity = &handlers.Gravadora{}
		sql = entity.SqlQuery("true")
	case match(url, "^gravadoras/([0-9]+)+/?$", &id):
		entity = &handlers.Gravadora{}
		sql = entity.SqlQuery(fmt.Sprintf("cod_grav = %s", id))
	case match(url, "^albuns/?$"):
		entity = &handlers.Album{}
		sql = entity.SqlQuery("true")
	case match(url, "^albuns/([0-9]+)+/?$", &id):
		entity = &handlers.Album{}
		sql = entity.SqlQuery(fmt.Sprintf("cod_alb = %s", id))
	case match(url, "^albuns/([0-9]+)+/faixas/?$", &id):
		entity = &handlers.Faixa{}
		sql = entity.SqlQueryJoin(fmt.Sprintf("cod_alb = %s", id))
	case match(url, "^faixas/?$"):
		entity = &handlers.Faixa{}
		sql = entity.SqlQuery("true")
	case match(url, "^faixas/([0-9]+)+/?$", &id):
		entity = &handlers.Faixa{}
		sql = entity.SqlQuery(fmt.Sprintf("nro_faixa = %s", id))
	case match(url, "^interpretes/?$"):
		entity = &handlers.Interprete{}
		sql = entity.SqlQuery("true")
	case match(url, "^interpretes/([0-9]+)+/?$", &id):
		entity = &handlers.Interprete{}
		sql = entity.SqlQuery(fmt.Sprintf("cod_inter = %s", id))
	case match(url, "^faixas/([0-9]+)+/interpretes/?$", &id):
		entity = &handlers.Interprete{}
		sql = entity.SqlQueryJoin(fmt.Sprintf("nro_faixa = %s", id))
	case match(url, "^compositores/?$"):
		entity = &handlers.Compositor{}
		sql = entity.SqlQuery("true")
	case match(url, "^compositores/([0-9]+)+/?$", &id):
		entity = &handlers.Compositor{}
		sql = entity.SqlQuery(fmt.Sprintf("cod_compositor = %s", id))
	case match(url, "^faixas/([0-9]+)+/compositores/?$", &id):
		entity = &handlers.Compositor{}
		sql = entity.SqlQueryJoin(fmt.Sprintf("nro_faixa = %s", id))
	case match(url, "^playlists/?$"):
		entity = &handlers.Playlist{}
		sql = entity.SqlQuery("true")
	case match(url, "^playlists/([0-9]+)+/?$", &id):
		entity = &handlers.Playlist{}
		sql = entity.SqlQuery(fmt.Sprintf("cod_play = %s", id))
	case r.Method == "GET" && match(url, "^playlists/([0-9]+)+/faixas/?$", &id):
		entity = &handlers.Playlist{}
		sql = entity.SqlQueryJoin(fmt.Sprintf("cod_play = %s", id))
		entity = &handlers.Faixa{}
	case match(url, "^playlists/([0-9]+)+/faixas/?$", &id):
		entity = &handlers.Faixa_playlist{}
	}

	return
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
