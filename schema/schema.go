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
	Cod    int       `json:"cod"`
	Desc   string    `json:"descricao"`
	DtGrav time.Time `json:"dtgrav"`
	PrComp float32   `json:"prcompra"`
	PrVen  float32   `json:"prvenda"`
}

type Composicao struct {
	Cod  int    `json:"cod"`
	Desc string `json:"descricao"`
	Tipo string `json:"tipo"`
}

type Faixa struct {
	Cod    int           `json:"cod"`
	Num    int           `json:"numero"`
	Desc   string        `json:"descricao"`
	TExec  time.Duration `json:"tempexec"`
	TpGrav string        `json:"tipogravacao"`
}

type Compositor struct {
	Cod    int       `json:"cod"`
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

func GetTable(t string) interface{} {
	switch t {
	case "gravadoras":
		return Gravadora{}
	case "albuns":
		return Album{}
	case "composicoes":
		return Composicao{}
	case "faixa":
		return Faixa{}
	case "compositores":
		return Compositor{}
	case "playlists":
		return Playlist{}
	case "interpretes":
		return Interprete{}
	default:
		return nil
	}
}
