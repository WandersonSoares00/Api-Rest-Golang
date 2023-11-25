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
	//var h http.Handler

	//var s string
	var id int64
	var err error = nil
	p := r.URL.Path

	//[a-z]+\b

	start := time.Now()

	switch {
	case match(p, "^gravadoras/?$") && r.Method == "GET":
		err = handlers.GetAllGravadora(w)
	case match(p, "^gravadoras/\\d+$", &id) && r.Method == "GET":
		err = handlers.GetGravadora(w, id)
	case match(p, "^gravadoras/?$") && r.Method == http.MethodPost:
		err = handlers.CreateGravadora(w, r)
	case match(p, "^gravadoras/\\d+$", &id) && r.Method == http.MethodDelete:
		err = handlers.DeleteGravadora(w, id)
	case match(p, "^gravadoras/\\d+$", &id) && r.Method == http.MethodPut:
		err = handlers.UpdateGravadora(w, r, id)

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

	if len(vars) >= 1 {
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
	}

	return true
}
