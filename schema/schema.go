package schema

type Gravadora struct {
}

type Album struct {
}

type Composicao struct {
}

type Faixa struct {
}

type Compositor struct {
}

type Playlist struct {
}

type Interprete struct {
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
