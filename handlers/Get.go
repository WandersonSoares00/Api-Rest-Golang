package handlers

import (
	"fmt"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/db"
	"github.com/WandersonSoares00/Api-Rest-Golang.git/schema"
)

func Get(table string) interface{} {
	conn, err := db.OpenConnection()

	if err != nil {
		fmt.Println(table)
		return nil
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM $1`, table)

	t := schema.GetTable(table)

	err = row.Scan()

	if err != nil {
		return nil
	}

	return t
}
