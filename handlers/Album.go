package handlers

import (
	"database/sql"
	"fmt"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/schema"
)

type Album schema.Album

func (a *Album) Scan(row *sql.Rows) error {
	return row.Scan(&a.Cod, &a.Meio, &a.CodGrav, &a.Nome, &a.Desc, &a.DtGrav, &a.PrComp, &a.PrVen)
}

func (a *Album) New() Entity {
	return &Album{}
}

func (a *Album) SqlCreate() string {
	return fmt.Sprintf(`INSERT INTO album (cod_alb, meio, cod_grav, nome, descricao, data_grav)
			VALUES (%d, '%s', %d, '%s', '%s', '%s') RETURNING cod_grav`, a.Cod, a.Meio, a.CodGrav, a.Nome, a.Desc, a.DtGrav.Format("15:04:05"))
}

func (a *Album) SqlUpdate() string {
	return fmt.Sprintf(`UPDATE album SET nome='%s', descricao='%s', data_grav='%s'
			WHERE cod_alb=%d AND meio='%s' AND cod_grav=%d`, a.Nome, a.Desc, a.DtGrav.Format("15:04:05"), a.Cod, a.Meio, a.CodGrav)
}

func (a *Album) SqlDelete() string {
	return fmt.Sprintf(`DELETE FROM album WHERE cod_alb = %d AND meio = '%s'`, a.Cod, a.Meio)
}

func (a *Album) SqlQuery(filter string) string {
	return fmt.Sprintf(`SELECT * FROM album WHERE %s`, filter)
}

func (a *Album) SqlQueryJoin(filter string) string {
	return fmt.Sprintf(`SELECT a.* FROM album a JOIN gravadora USING(cod_grav) WHERE %s`, filter)
}
