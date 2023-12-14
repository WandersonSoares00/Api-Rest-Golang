package handlers

import (
	"database/sql"
	"fmt"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/schema"
)

type Faixa_playlist schema.Faixa_playlist

func (p *Faixa_playlist) Scan(row *sql.Rows) error {
	return row.Scan(&p.Codfaixa, &p.CodAlb, &p.Meio, &p.CodPlay)
}

func (p *Faixa_playlist) New() Entity {
	return &Faixa_playlist{}
}

func (p *Faixa_playlist) SqlCreate() string {
	return fmt.Sprintf(`INSERT INTO faixa_playlist (nro_faixa, cod_alb, meio, cod_play)
		VALUES (%d, %d, '%s', %d) RETURNING cod_play`, p.Codfaixa, p.CodAlb, p.Meio, p.CodPlay)
}

func (p *Faixa_playlist) SqlUpdate() string {
	return fmt.Sprintf(`UPDATE faixa_playlist SET dt_ult_repr='%s', qtd_repr=%d WHERE nro_faixa = %d AND
	 cod_alb = %d AND meio = '%s' AND cod_play = %d`, p.DtUltRep.Time.Format("15:04:05"), p.QtdRep, p.Codfaixa, p.CodAlb, p.Meio, p.CodPlay)
}

func (p *Faixa_playlist) SqlDelete() string {
	return fmt.Sprintf(`DELETE FROM faixa_playlist WHERE nro_faixa = %d AND cod_alb = %d AND meio = '%s' AND cod_play = %d`,
		p.Codfaixa, p.CodAlb, p.Meio, p.CodPlay)
}

func (p *Faixa_playlist) SqlQuery(filter string) string {
	return fmt.Sprintf(`SELECT * FROM faixa_playlist WHERE %s`, filter)
}

func (p *Faixa_playlist) SqlQueryJoin(filter string) string {
	return fmt.Sprintf(`SELECT f.* FROM faixa_playlist
					JOIN faixa f   USING(nro_faixa, cod_alb, meio) WHERE %s`, filter)
}
