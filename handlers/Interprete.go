package handlers

import (
	"database/sql"
	"fmt"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/schema"
)

type Interprete schema.Interprete

func (i *Interprete) Scan(row *sql.Rows) error {
	return row.Scan(&i.Cod, &i.Nome, &i.Tipo)
}

func (i *Interprete) New() Entity {
	return &Interprete{}
}

func (i *Interprete) SqlCreate() string {
	return fmt.Sprintf(`INSERT INTO interprete (cod_inter, nome, tipo)
			VALUES (%d, '%s', '%s') RETURNING cod_inter`, i.Cod, i.Nome, i.Tipo)
}

func (i *Interprete) SqlUpdate() string {
	return fmt.Sprintf(`UPDATE interprete SET nome='%s', tipo='%s' WHERE cod_inter=%d`, i.Nome, i.Tipo, i.Cod)
}

func (i *Interprete) SqlDelete() string {
	return fmt.Sprintf(`DELETE FROM interprete WHERE cod_inter = %d`, i.Cod)
}

func (i *Interprete) SqlQuery(filter string) string {
	return fmt.Sprintf(`SELECT * FROM interprete WHERE %s`, filter)
}

func (i *Interprete) SqlQueryJoin(filter string) string {
	return fmt.Sprintf(`SELECT i.* FROM interprete i
				JOIN faixa_interprete USING(cod_inter)
				JOIN faixa 			  USING(nro_faixa, cod_alb, meio) WHERE %s`, filter)
}
