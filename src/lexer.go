package src

import (
	"regexp"
	"strings"
)

type Lexer struct {
	reader *strings.Reader
	tokens []Token
}

func (l *Lexer) RegisterToken(token ...Token) {
	l.tokens = append(l.tokens, token...)
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		reader: strings.NewReader(input),
	}
}

func (l *Lexer) Next() (Token, bool) {
	l.eatWhitespace()
	for _, tokenConfig := range l.tokens {
		token, ok := l.parseLiteral(tokenConfig)
		if ok {
			return token, true
		}
		token, ok = l.parseRegexp(tokenConfig)
		if ok {
			return token, true
		}
	}
	return nil, false
}

func (l *Lexer) eatWhitespace() {
	for {
		peekRune, err := l.peekRune()
		if err != nil {
			return
		}
		if !isWhiteSpaceChar(peekRune) {
			break
		}
		l.nextRune()
	}

}

func (l *Lexer) parseLiteral(token Token) (Token, bool) {
	if token.Config().Literal == nil {
		return nil, false
	}
	patterns := token.Config().Literal
	for _, pattern := range patterns {
		patternReader := strings.NewReader(pattern)
		tmpReader := *l.reader
		matched := true
		cnt := 0
		for a, _, err := patternReader.ReadRune(); err == nil; a, _, err = patternReader.ReadRune() {
			b, _, err1 := tmpReader.ReadRune()
			if err1 != nil || a != b {
				matched = false
				break
			}
			cnt++
		}
		b, _, err := tmpReader.ReadRune()
		if err == nil && !isWhiteSpaceChar(b) {
			matched = false
		}
		if matched {
			l.reader = &tmpReader
			return token.New([]rune(pattern)), true
		}
	}
	return nil, false
}

func (l *Lexer) parseRegexp(token Token) (Token, bool) {
	if token.Config().Regexp == nil {
		return nil, false
	}
	pattern, err := regexp.Compile(*token.Config().Regexp)
	if err != nil {
		return nil, false
	}
	tmpReader := *l.reader
	loc := pattern.FindReaderIndex(&tmpReader)
	if len(loc) != 2 {
		return nil, false
	}
	a, b := loc[0], loc[1]
	ans, err := eatNByte(l.reader, b-a)
	if err != nil {
		return nil, false
	}
	return token.New(ans), true
}

func (l *Lexer) nextRune() (rune, error) {
	ans, err := l.nextNRune(1)
	if err != nil {
		return 0, err
	}
	return ans[0], nil
}

func (l *Lexer) nextNRune(n int) ([]rune, error) {
	var ans []rune
	for i := 0; i < n; i++ {
		tmp, _, err := l.reader.ReadRune()
		if err != nil {
			return nil, err
		}
		ans = append(ans, tmp)
	}
	return ans, nil
}

func (l *Lexer) peekRune() (rune, error) {
	ans, err := l.peekNRune(1)
	if err != nil {
		return 0, err
	}
	return ans[0], nil
}

func (l *Lexer) peekNRune(n int) ([]rune, error) {
	reader := *l.reader
	var ans []rune
	for i := 0; i < n; i++ {
		tmp, _, err := reader.ReadRune()
		if err != nil {
			return nil, err
		}
		ans = append(ans, tmp)
	}
	return ans, nil
}

func eatNByte(reader *strings.Reader, n int) ([]rune, error) {
	var ans []rune
	for i := 0; i < n; {
		tmp, size, err := reader.ReadRune()
		i += size
		if err != nil {
			return nil, err
		}
		ans = append(ans, tmp)
	}
	return ans, nil
}

func isWhiteSpaceChar(input rune) bool {
	return input == rune('\t') || input == rune(' ') || input == rune('\n')
}
