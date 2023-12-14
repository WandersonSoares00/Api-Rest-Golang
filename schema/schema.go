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
	Cod     int            `json:"id"`
	CodAlb  int            `json:"codAlbum"`
	Meio    string         `json:"Meio"`
	CodComp int            `json:"codComposicao"`
	Desc    string         `json:"descricao"`
	TExec   time.Time      `json:"tempExec"`
	TpGrav  sql.NullString `json:"tipogravacao"`
}

type Faixa_playlist struct {
	Codfaixa int          `json:"nroFaixa"`
	CodAlb   int          `json:"codAlbum"`
	Meio     string       `json:"meio"`
	CodPlay  int          `json:"codPlay"`
	QtdRep   int          `json:"qtdRep"`
	DtUltRep sql.NullTime `json:"ultimaRepr"`
}
