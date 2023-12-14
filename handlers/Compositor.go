package handlers

import (
	"database/sql"
	"fmt"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/schema"
)

type Compositor schema.Compositor

func (c *Compositor) Scan(row *sql.Rows) error {
	return row.Scan(&c.Cod, &c.CodPm, &c.Nome, &c.DtNasc, &c.DtMort)
}

func (c *Compositor) New() Entity {
	return &Compositor{}
}

func (c *Compositor) SqlCreate() string {
	return fmt.Sprintf(`INSERT INTO compositor (cod_compositor, cod_pm, nome, dt_nasc)
			VALUES (%d, %d, '%s', '%s') RETURNING cod_compositor`, c.Cod, c.CodPm, c.Nome, c.DtNasc.Format("15:04:05"))
}

func (c *Compositor) SqlUpdate() string {
	return fmt.Sprintf(`UPDATE compositor SET nome='%s', dt_nasc='%s', dt_morte='%s'
			WHERE cod_compositor=%d AND cod_pm=%d`, c.Nome, c.DtNasc.Format("15:04:05"), c.DtMort.Format("15:04:05"), c.Cod, c.CodPm)
}

func (c *Compositor) SqlDelete() string {
	return fmt.Sprintf(`DELETE FROM compositor WHERE cod_compositor = %d`, c.Cod)
}

func (c *Compositor) SqlQuery(filter string) string {
	return fmt.Sprintf(`SELECT * FROM compositor WHERE %s`, filter)
}

func (c *Compositor) SqlQueryJoin(filter string) string {
	return fmt.Sprintf(`SELECT c.* FROM compositor c
					JOIN faixa_compositor USING(cod_compositor)
					JOIN faixa 			  USING(nro_faixa, cod_alb, meio) WHERE %s`, filter)
}
