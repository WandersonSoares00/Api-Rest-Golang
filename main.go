package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/configs"
	"github.com/WandersonSoares00/Api-Rest-Golang.git/routes"
)

func main() {
	logger := configs.GetLogger("main")

	mux := http.NewServeMux()

	mux.Handle("/api/v1", http.HandlerFunc(routes.Serve))
	mux.Handle("/api/v1/", http.HandlerFunc(routes.Serve))

	p := configs.GetServerPort()

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", p))

	if err != nil {
		logger.Errorf("error starting server: %s\n", err.Error())
		return
	}

	logger.Infof("server listening in port %d", p)

	http.Serve(ln, mux)

}
