package handlers

import (
	"database/sql"
	"fmt"

	"github.com/WandersonSoares00/Api-Rest-Golang.git/schema"
)

type Playlist schema.Playlist

func (p *Playlist) Scan(row *sql.Rows) error {
	return row.Scan(&p.Cod, &p.Nome, &p.TempTot, &p.DtCriac)
}

func (p *Playlist) New() Entity {
	return &Playlist{}
}

func (p *Playlist) SqlCreate() string {
	return fmt.Sprintf(`INSERT INTO playlist (cod_play, nome) VALUES (%d, '%s') RETURNING cod_play`, p.Cod, p.Nome)
}

func (p *Playlist) SqlUpdate() string {
	fmt.Println(p.TempTot.Date())
	return fmt.Sprintf(`UPDATE playlist SET nome='%s', tempo_tot='%s' WHERE cod_play=%d`, p.Nome, p.TempTot.Format("15:04:05"), p.Cod)
}

func (p *Playlist) SqlDelete() string {
	return fmt.Sprintf(`DELETE FROM playlist WHERE cod_play = %d`, p.Cod)
}

func (p *Playlist) SqlQuery(filter string) string {
	return fmt.Sprintf(`SELECT * FROM playlist WHERE %s`, filter)
}

func (p *Playlist) SqlQueryJoin(filter string) string {
	return fmt.Sprintf(`SELECT f.* FROM playlist
					JOIN faixa_playlist USING(cod_play)
					JOIN faixa f		USING(nro_faixa, cod_alb, meio) WHERE %s`, filter)
}
