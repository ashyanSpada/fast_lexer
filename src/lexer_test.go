package src

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLexer(t *testing.T) {
	input := "true false True False"
	lexer := NewLexer(input)
	lexer.RegisterToken(new(BoolToken).WithConfig(TokenConfig{
		Literal: []string{"true", "false", "True", "False"},
	}), new(BracketToken).WithConfig(TokenConfig{
		Literal: []string{"(", ")", "{", "}", "[", "]"},
	}))
	for {
		token, ok := lexer.Next()
		fmt.Println(token, ok)
		if !ok {
			break
		}
		fmt.Println(reflect.TypeOf(token))
	}
}
