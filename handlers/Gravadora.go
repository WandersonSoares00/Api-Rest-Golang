package handlers

import (
	"database/sql"
	"fmt"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/schema"
)

type Gravadora schema.Gravadora

func (g *Gravadora) Scan(row *sql.Rows) error {
	return row.Scan(&g.Cod, &g.Nome, &g.Cidade, &g.Pais, &g.HomeP)
}

func (g *Gravadora) New() Entity {
	return &Gravadora{}
}

func (g *Gravadora) SqlCreate() string {
	return fmt.Sprintf(`INSERT INTO gravadora (cod_grav, nome, cidade, pais, end_homep)
			VALUES (%d, '%s', '%s', '%s', '%s') RETURNING cod_grav`, g.Cod, g.Nome, g.Cidade, g.Pais, g.HomeP)
}

func (g Gravadora) SqlUpdate() string {
	return fmt.Sprintf(`UPDATE gravadora SET nome='%s', cidade='%s', pais='%s', end_homep='%s' WHERE cod_grav=%d`,
		g.Nome, g.Cidade, g.Pais, g.HomeP, g.Cod)
}

func (g *Gravadora) SqlDelete() string {
	return fmt.Sprintf(`DELETE FROM gravadora WHERE cod_grav = %d`, g.Cod)
}

func (g *Gravadora) SqlQuery(filter string) string {
	return fmt.Sprintf(`SELECT * FROM gravadora WHERE %s`, filter)
}

func (g *Gravadora) SqlQueryJoin(filter string) string {
	return fmt.Sprintf(`SELECT a.* FROM gravadora g JOIN album a USING(cod_grav) WHERE %s`, filter)
}
