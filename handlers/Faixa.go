package handlers

import (
	"database/sql"
	"fmt"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/schema"
)

type Faixa schema.Faixa

func (f *Faixa) Scan(row *sql.Rows) error {
	return row.Scan(&f.Cod, &f.CodAlb, &f.Meio, &f.CodComp, &f.Desc, &f.TExec, &f.TpGrav)
}

func (f *Faixa) New() Entity {
	return &Faixa{}
}

func (f *Faixa) SqlCreate() string {
	return fmt.Sprintf(`INSERT INTO faixa (cod_alb, meio, nro_faixa, cod_composicao, descricao, tempo_exec)
		VALUES (%d, '%s', %d, %d, '%s', '%s') RETURNING cod_faixa`, f.CodAlb, f.Meio, f.Cod, f.CodComp, f.Desc, f.TExec.Format("15:04:05"))
}

func (f *Faixa) SqlUpdate() string {
	return fmt.Sprintf(`UPDATE faixa SET descricao='%s', tempo_exec='%s', tipo_grav='%s'
				WHERE cod_alb=%d AND nro_faixa=%d AND meio='%s'`, f.Desc, f.TExec.Format("15:04:05"), f.TpGrav.String, f.CodAlb, f.Cod, f.Meio)
}

func (f *Faixa) SqlDelete() string {
	return fmt.Sprintf(`DELETE FROM faixa WHERE cod_alb=%d AND nro_faixa = %d AND meio = '%s'`, f.CodAlb, f.Cod, f.Meio)
}

func (f *Faixa) SqlQuery(filter string) string {
	return fmt.Sprintf(`SELECT * FROM faixa WHERE %s`, filter)
}

func (f *Faixa) SqlQueryJoin(filter string) string {
	return fmt.Sprintf(`SELECT f.* FROM album JOIN faixa f USING(cod_alb, meio) WHERE %s`, filter)
}
