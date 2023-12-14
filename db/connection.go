package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetConfs()

	strcon := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		conf.DbHost, conf.DbPort, conf.DbUser, conf.DbPass, conf.Dbname, conf.SSL)

	conn, err := sql.Open("postgres", strcon)
	if err != nil {
		return conn, err
	}

	err = conn.Ping()

	return conn, err
}
