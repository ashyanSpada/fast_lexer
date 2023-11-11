package src

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	input := "true false True False ( ) { } [ ] -124.123 123e-03 \"ha你好ha\""
	lexer := NewLexer(input)

	lexer.RegisterToken(new(BoolToken), new(BracketToken), new(NumberToken), new(StringToken))

	for i := 0; i < 100; i++ {
		token, ok := lexer.Next()
		fmt.Println(token, "parseOk", ok)
		if !ok {
			break
		}
	}

}
