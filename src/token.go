package src

import "fmt"

type TokenConfig struct {
	Regexp  *string
	Literal []string
}

type Token interface {
	Config() TokenConfig
	New([]rune) Token
}

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
		source: input,
	}
}

func (s *StringToken) Config() TokenConfig {
	// reg := "(^\"[.|\\\"]*\"$)|(^'[.|\\']*'$)"
	reg := "^\"[.|(\\\")]*\"$"
	return TokenConfig{
		Regexp: &reg,
	}
}

func (s *StringToken) String() string {
	return fmt.Sprintf("StringToken: %s", string(s.source))
}
