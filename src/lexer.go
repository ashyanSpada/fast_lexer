package src

import (
	"regexp"
	"strings"

	"golang.org/x/text/runes"
)

type Lexer struct {
	cur    int
	source []rune
	tokens []Token
}

func (l *Lexer) RegisterToken(token Token) {
	l.tokens = append(l.tokens, token)
}

func (l *Lexer) RegisterTokens(tokens []Token) {
	l.tokens = append(l.tokens, tokens...)
}

func (l *Lexer) Next() Token {

}

func NewLexer(input string) *Lexer {
	return &Lexer{
		source: []rune(input),
	}
}

func (l *Lexer) parseLiteral(token Token) (Token, bool) {
	patterns := token.Config().Literal
	for _, pattern := range patterns {
		m := len(l.source[l.cur:])
		n := len(pattern)
		if m < n {
			continue
		}
		if runeListEqual(l.source[l.cur:l.cur+n], []rune(pattern)) && (m == n || isWhiteSpaceChar(l.source[l.cur+n])) {
			output := token.New()
			output.Source(l.source[l.cur : l.cur+n])
			l.cur += n
			return output, true
		}
	}
	return nil, false
}

func (l *Lexer) parseRegexp(token Token) (Token, bool) {
	pattern, err := regexp.Compile(*token.Config().Regexp)
	if err != nil {
		return nil, false
	}

	pattern.FindReaderIndex()
}

func runeListEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isWhiteSpaceChar(input rune) bool {
	return input == rune('\t') || input == rune(' ') || input == rune('\n')
}
