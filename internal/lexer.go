package internal

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type Opt = func(*Lexer)

type Lexer struct {
	singleton bool
	reader    *strings.Reader
	tokens    []*SingletonTokenWrapper
}

func (l *Lexer) RegisterToken(tokens ...Token) {
	wrappers := make([]*SingletonTokenWrapper, 0, len(tokens))
	for _, token := range tokens {
		wrappers = append(wrappers, NewSingletonWrapper(token))
	}
	l.tokens = append(l.tokens, wrappers...)
}

func EnableSingletonOpt(l *Lexer) {
	l.singleton = true
}

func NewLexer(input string, opts ...Opt) *Lexer {
	l := &Lexer{
		reader: strings.NewReader(input),
	}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

func (l *Lexer) Next() (Token, error) {
	l.eatWhitespace()
	if l.IsEnd() {
		return nil, nil
	}
	for _, tokenWrapper := range l.tokens {
		token, ok := l.parseLiteral(tokenWrapper, l.singleton)
		if ok {
			return token, nil
		}
		token, ok = l.parseRegexp(tokenWrapper, l.singleton)
		if ok {
			return token, nil
		}
	}
	return nil, errors.New("no matched pattern")
}

func (l *Lexer) Peek() (Token, error) {
	tmpReader := *l.reader
	tmpLexer := &Lexer{
		reader: &tmpReader,
		tokens: l.tokens,
	}
	return tmpLexer.Next()
}

func (l *Lexer) IsEnd() bool {
	l.eatWhitespace()
	tmpReader := *l.reader
	if _, err := tmpReader.ReadByte(); err == io.EOF {
		return true
	}
	return false
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

func (l *Lexer) parseLiteral(tokenWrapper *SingletonTokenWrapper, singleton bool) (Token, bool) {
	if tokenWrapper.Config().Literal == nil {
		return nil, false
	}
	patterns := tokenWrapper.Config().Literal
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
			return tokenWrapper.New([]rune(pattern), singleton), true
		}
	}
	return nil, false
}

func (l *Lexer) parseRegexp(tokenWrapper *SingletonTokenWrapper, singleton bool) (Token, bool) {
	if tokenWrapper.Config().Regexp == nil {
		return nil, false
	}
	pattern, err := regexp.Compile(*tokenWrapper.Config().Regexp)
	if err != nil {
		fmt.Println("compile err:", *tokenWrapper.Config().Regexp)
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
	return tokenWrapper.New(ans, singleton), true
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
	return input == rune('\t') || input == rune(' ')
}
