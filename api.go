package fast_lexer

import "github.com/ashyanSpada/fast_lexer/internal"

type Lexer = internal.Lexer

type TokenConfig = internal.TokenConfig

type Opt = internal.Opt

func EnableSingletonOpt(lexer *Lexer) {
	internal.EnableSingletonOpt(lexer)
}

type (
	Token           = internal.Token
	BoolToken       = internal.BoolToken
	BracketToken    = internal.BracketToken
	StringToken     = internal.StringToken
	NumberToken     = internal.NumberToken
	IdentifierToken = internal.IdentifierToken
)

func NewLexer(input string, opts ...Opt) *Lexer {
	return internal.NewLexer(input, opts...)
}
