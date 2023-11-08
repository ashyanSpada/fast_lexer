package src

type TokenConfig struct {
	Regexp  *string
	Literal []string
}

type Token interface {
	Config() TokenConfig
	Source([]rune)
	New() Token
}

type CommonToken struct {
	TokenConfig
	source []rune
}

func (c *CommonToken) Config() TokenConfig {
	return c.TokenConfig
}

func (c *CommonToken) Source(input []rune) {
	c.source = input
}

type BoolToken struct {
	CommonToken
}

func (b *BoolToken) IsTrue() bool {
	return string(b.source) == "true" || string(b.source) == "True"
}

func (b *BoolToken) IsFalse() bool {
	return string(b.source) == "false" || string(b.source) == "False"
}
