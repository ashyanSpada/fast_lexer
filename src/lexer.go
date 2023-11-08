package src

import (
	"io"
	"regexp"
	"strings"
)

type Lexer struct {
	cur    int
	reader *strings.Reader
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
		reader: strings.NewReader(input),
	}
}

func (l *Lexer) parseLiteral(token Token) (Token, bool) {
	patterns := token.Config().Literal
	for _, pattern := range patterns {
		patternReader := strings.NewReader(pattern)
		tmpReader := *l.reader
		matched := false
		cnt := 0
		for a, _, err := patternReader.ReadRune(); err == nil; {
			b, _, err1 := tmpReader.ReadRune()
			if err1 != nil || a != b {
				break
			}
			cnt++
		}
		b, _, err := tmpReader.ReadRune()
		if err != nil || isWhiteSpaceChar(b) {
			matched = true
		}
		if matched {
			var ans []rune
			for i := 0; i < cnt; i++ {
				a, _, _ := l.reader.ReadRune()
				ans = append(ans, a)
			}
			result := token.New()
			result.Source(ans)
			return result, true
		}
	}
	return nil, false
}

func (l *Lexer) parseRegexp(token Token) (Token, bool) {
	pattern, err := regexp.Compile(*token.Config().Regexp)
	if err != nil {
		return nil, false
	}

	ans := pattern.FindReaderIndex(l.reader)
	if len(ans) != 2 {
		return nil, false
	}
	a, b := fffff========================================
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
