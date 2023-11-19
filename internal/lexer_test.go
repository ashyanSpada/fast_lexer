package internal

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
					source: []rune("ha你好ha"),
				},
				&IdentifierToken{
					source: []rune("haha"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := getTestLexer(tt.fields.input)
			var tokens []Token
			for {
				token, err := l.Next()
				if err != nil || token == nil {
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

func TestLexer_Peek(t *testing.T) {
	type fields struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		want   Token
	}{
		{
			name: "test",
			fields: fields{
				input: "true false True False ( )  { } [ ] -124.123 123e-03 \"ha你好ha\" haha",
			},
			want: &BoolToken{
				source: []rune("true"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := getTestLexer(tt.fields.input)
			token, _ := l.Peek()
			if !reflect.DeepEqual(token, tt.want) {
				t.Errorf("Lexer.Peek() got = %v, want %v", token, tt.want)
			}
		})
	}
}

func TestLexer_IsEnd(t *testing.T) {
	type fields struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			fields: fields{
				input: "",
			},
			want: true,
		},
		{
			fields: fields{
				input: "   \t",
			},
			want: true,
		},
		{
			fields: fields{
				input: "   \n",
			},
			want: false,
		},
		{
			fields: fields{
				input: "   \r\n",
			},
			want: false,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			l := getTestLexer(tt.fields.input)
			if got := l.IsEnd(); got != tt.want {
				t.Errorf("Lexer.IsEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getTestLexer(input string) *Lexer {
	l := NewLexer(input)
	l.RegisterToken(
		new(BoolToken),
		new(BracketToken),
		new(NumberToken),
		new(StringToken),
		new(IdentifierToken),
	)
	return l
}
