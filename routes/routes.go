package routes

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/handlers"
)

func Serve(w http.ResponseWriter, r *http.Request) {
	//var h http.Handler

	var table string

	p := r.URL.Path
	//[a-z]+\b
	switch {
	//case match(p, "/api/v1"):
	//h = showschema()
	case match(p, "/api/v1/([a-z]+)", &table):
		handlers.Get(table)
		/*
			case match(p, "/api/v1/gravadoras\\/?"):
				//h = get(gravadora)
			case match(p, "/api/vi/Albuns\\/?"):

			case match(p, "/api/vi/Composicoes\\/?"):

			case match(p, "/api/vi/Faixas\\/?"):

			case match(p, "/api/vi/Compositores\\/?"):

			case match(p, "/api/vi/Playlists\\/?"):

			case match(p, "/api/vi/Interpretes\\/?"):
		*/
	default:
		fmt.Println("não matchiou")
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

	for i, match := range matches[0:] {
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
