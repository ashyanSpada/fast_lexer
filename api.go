package fast_lexer

import "github.com/ashyanSpada/fast_lexer/src"

type Lexer = src.Lexer

type TokenConfig = src.TokenConfig

type (
	BoolToken       = src.BoolToken
	BracketToken    = src.BracketToken
	StringToken     = src.StringToken
	NumberToken     = src.NumberToken
	IdentifierToken = src.IdentifierToken
)
