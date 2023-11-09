package src

type TokenConfig struct {
	Regexp  *string
	Literal []string
}

type Token interface {
	Config() TokenConfig
	WithConfig(config TokenConfig) Token
	New([]rune) Token
}

type CommonToken struct {
	TokenConfig
	source []rune
}

func (c *CommonToken) Config() TokenConfig {
	return c.TokenConfig
}

func (c *CommonToken) New(input []rune) CommonToken {
	return CommonToken{
		TokenConfig: c.TokenConfig,
		source:      input,
	}
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

func (b *BoolToken) New(input []rune) Token {
	return &BoolToken{
		CommonToken: b.CommonToken.New(input),
	}
}

func (b *BoolToken) WithConfig(config TokenConfig) Token {
	b.TokenConfig = config
	return b
}

type BracketToken struct {
	CommonToken
}

func (b *BracketToken) New(input []rune) Token {
	return &BracketToken{
		CommonToken: b.CommonToken.New(input),
	}
}

func (b *BracketToken) WithConfig(config TokenConfig) Token {
	b.TokenConfig = config
	return b
}

func (b *BracketToken) IsLeftBracket() bool {
	return string(b.source) == "["
}

func (b *BracketToken) IsRightBracket() bool {
	return string(b.source) == "]"
}

func (b *BracketToken) IsLeftCurly() bool {
	return string(b.source) == "{"
}

func (b *BracketToken) IsRightCurly() bool {
	return string(b.source) == "}"
}

func (b *BracketToken) IsLeftBrace() bool {
	return string(b.source) == "("
}

func (b *BracketToken) IsRightBrace() bool {
	return string(b.source) == ")"
}
