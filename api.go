package fast_lexer

import "github.com/ashyanSpada/fast_lexer/internal"

type Lexer = internal.Lexer

type TokenConfig = internal.TokenConfig

type (
	Token           = internal.Token
	BoolToken       = internal.BoolToken
	BracketToken    = internal.BracketToken
	StringToken     = internal.StringToken
	NumberToken     = internal.NumberToken
	IdentifierToken = internal.IdentifierToken
)
