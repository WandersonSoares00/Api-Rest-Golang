package schema

import "time"

type Telefone_gravadora struct {
	Cod    int    `json:"cod"`
	Numero string `json:"numero"`
}

type Gravadora struct {
	Cod     int    `json:"cod"`
	CodFone int    `json:"codFone"`
	Nome    string `json:"nome"`
	Ender   string `json:"endereco"`
	HomeP   string `json:"home page"`
}

type Album struct {
	Cod     int64     `json:"cod"`
	CodDown int64     `json:"codDown"`
	Desc    string    `json:"descricao"`
	DtGrav  time.Time `json:"dtGrav"`
	PrComp  float32   `json:"prCompra"`
	PrVen   float32   `json:"prVenda"`
}

type Composicao struct {
	Cod  int    `json:"cod"`
	Desc string `json:"descricao"`
	Tipo string `json:"tipo"`
}

type Faixa struct {
	Cod     int           `json:"cod"`
	CodCd   int64         `json:"codCd"`
	CodVin  int64         `json:"codVinil"`
	CodDown int64         `json:"codDown"`
	CodComp int64         `json:"codComp"`
	Num     int           `json:"numero"`
	Desc    string        `json:"descricao"`
	TExec   time.Duration `json:"tempexec"`
	TpGrav  string        `json:"tipogravacao"`
}

type Compositor struct {
	Cod    int       `json:"cod"`
	CodPm  int64     `json:"codPm"`
	Nome   string    `json:"nome"`
	DtNasc time.Time `json:"dtnascimento"`
	DtMort time.Time `json:"dtmorte,omitempty"`
}

type Playlist struct {
	Cod     int           `json:"cod"`
	Nome    string        `json:"nome"`
	TempTot time.Duration `json:"tempotot"`
	DtCriac time.Duration `json:"dtcriacao"`
}

type Interprete struct {
	Cod  int    `json:"cod"`
	Nome string `json:"nome"`
	Tipo string `json:"tipo"`
}
