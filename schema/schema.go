package schema

import (
	"database/sql"
	"time"
)

type Gravadora struct {
	Cod    int    `json:"id"`
	Nome   string `json:"nome"`
	Cidade string `json:"cidade"`
	Pais   string `json:"pais"`
	HomeP  string `json:"home page"`
}

type Album struct {
	Cod     int             `json:"id"`
	CodMeio int             `json:"codMeio"`
	CodGrav int             `json:"codGrav"`
	Nome    string          `json:"nome"`
	Desc    string          `json:"descricao"`
	DtGrav  time.Time       `json:"dtGrav"`
	PrComp  float32         `json:"prCompra"`
	PrVen   sql.NullFloat64 `json:"prVenda"`
	Meio    string          `json:"meio"`
}

type Interprete struct {
	Cod  int    `json:"id"`
	Nome string `json:"nome"`
	Tipo string `json:"tipo"`
}

type Compositor struct {
	Cod    int       `json:"id"`
	CodPm  int       `json:"codPeriodoMusical"`
	Nome   string    `json:"nome"`
	DtNasc time.Time `json:"dtnascimento"`
	DtMort time.Time `json:"dtmorte,omitempty"`
}

type Playlist struct {
	Cod     int       `json:"id"`
	Nome    string    `json:"nome"`
	TempTot time.Time `json:"tempoTot"`
	DtCriac time.Time `json:"dtCriacao"`
}

type Faixa struct {
	Cod     int       `json:"id"`
	CodAlb  int       `json:"codAlbum"`
	CodMeio int       `json:"codMeio"`
	CodComp int       `json:"codComposicao"`
	Num     int       `json:"numero"`
	Desc    string    `json:"descricao"`
	TExec   time.Time `json:"tempExec"`
	TpGrav  string    `json:"tipogravacao"`
}

type Faixa_playlist struct {
	Codfaixa int          `json:"codFaixa"`
	CodAlb   int          `json:"codAlbum"`
	CodMeio  int          `json:"codMeio"`
	CodPlay  int          `json:"codPlay"`
	UltRep   sql.NullTime `json:"UtimaRep"`
	QtdRep   int          `json:"qtdRep"`
}
