package src

import (
	"reflect"
	"testing"
)

func TestLexer_Next(t *testing.T) {
	type fields struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		want   []Token
	}{
		{
			name: "test",
			fields: fields{
				input: "true false True False ( )  { } [ ] -124.123 123e-03 \"ha你好ha\" haha",
			},
			want: []Token{
				&BoolToken{
					source: []rune("true"),
				},
				&BoolToken{
					source: []rune("false"),
				},
				&BoolToken{
					source: []rune("True"),
				},
				&BoolToken{
					source: []rune("False"),
				},
				&BracketToken{
					source: []rune("("),
				},
				&BracketToken{
					source: []rune(")"),
				},
				&BracketToken{
					source: []rune("{"),
				},
				&BracketToken{
					source: []rune("}"),
				},
				&BracketToken{
					source: []rune("["),
				},
				&BracketToken{
					source: []rune("]"),
				},
				&NumberToken{
					source: []rune("-124.123"),
				},
				&NumberToken{
					source: []rune("123e-03"),
				},
				&StringToken{
					source: []rune("\"ha你好ha\""),
				},
				&LiteralToken{
					source: []rune("haha"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.fields.input)
			l.RegisterToken(
				new(BoolToken),
				new(BracketToken),
				new(NumberToken),
				new(StringToken),
				new(LiteralToken),
			)
			var tokens []Token
			for {
				token, ok := l.Next()
				if !ok {
					break
				}
				tokens = append(tokens, token)
			}

			if !reflect.DeepEqual(tokens, tt.want) {
				t.Errorf("Lexer.NextAll() got = %v, want %v", tokens, tt.want)
			}
		})
	}
}
