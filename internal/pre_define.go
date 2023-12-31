package internal

import "fmt"

type BoolToken struct {
	source []rune
}

func (b *BoolToken) IsTrue() bool {
	return string(b.source) == "true" || string(b.source) == "True"
}

func (b *BoolToken) IsFalse() bool {
	return string(b.source) == "false" || string(b.source) == "False"
}

func (b *BoolToken) New(input []rune) Token {
	return &BoolToken{
		source: input,
	}
}

func (b *BoolToken) Config() TokenConfig {
	return TokenConfig{
		Literal: []string{"true", "false", "True", "False"},
	}
}

func (b *BoolToken) String() string {
	return fmt.Sprintf("BoolToken: %s", string(b.source))
}

type BracketToken struct {
	source []rune
}

func (b *BracketToken) New(input []rune) Token {
	return &BracketToken{
		source: input,
	}
}

func (b *BracketToken) Config() TokenConfig {
	return TokenConfig{
		Literal: []string{"(", ")", "[", "]", "{", "}"},
	}
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

func (b *BracketToken) String() string {
	return fmt.Sprintf("BracketToken: %s", string(b.source))
}

type NumberToken struct {
	source []rune
}

func (n *NumberToken) New(input []rune) Token {
	return &NumberToken{
		source: input,
	}
}

func (n *NumberToken) Config() TokenConfig {
	reg := "-?(?:0|[1-9]\\d*)(?:\\.\\d+)?(?:[eE][+-]?\\d+)?"
	return TokenConfig{
		Regexp: &reg,
	}
}

func (n *NumberToken) String() string {
	return fmt.Sprintf("NumberToken: %s", string(n.source))
}

type StringToken struct {
	source []rune
}

func (s *StringToken) New(input []rune) Token {
	return &StringToken{
		source: input[1 : len(input)-1],
	}
}

func (s *StringToken) Config() TokenConfig {
	// reg := "(^\"[.|\\\"]*\"$)|(^'[.|\\']*'$)"
	reg := `^"([^"\\\\]*|\\\\["\\\\bfnrt\/]|\\\\u[0-9a-f]{4})*"`
	return TokenConfig{
		Regexp: &reg,
	}
}

func (s *StringToken) String() string {
	return fmt.Sprintf("StringToken: %s", string(s.source))
}

func (s *StringToken) Value() []rune {
	return s.source
}

type IdentifierToken struct {
	source []rune
}

func (i *IdentifierToken) New(input []rune) Token {
	return &IdentifierToken{
		source: input,
	}
}

func (i *IdentifierToken) Config() TokenConfig {
	reg := `^[_a-zA-Z][_a-zA-Z0-9]{0,30}`
	return TokenConfig{
		Regexp: &reg,
	}
}

func (i *IdentifierToken) String() string {
	return fmt.Sprintf("IdentifierToken: %s", string(i.source))
}

func (i *IdentifierToken) Value() []rune {
	return i.source
}
