package internal

type TokenConfig struct {
	Regexp  *string
	Literal []string
}

type Token interface {
	Config() TokenConfig
	New([]rune) Token
}
