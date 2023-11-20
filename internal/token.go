package internal

type TokenConfig struct {
	Regexp  *string
	Literal []string
}

type Token interface {
	Config() TokenConfig
	New([]rune) Token
}

type TokenCounter struct {
	Cnt   int
	Token Token
}

type SingletonTokenWrapper struct {
	Token
	m map[string]*TokenCounter
}

func (s *SingletonTokenWrapper) New(input []rune, singleton bool) Token {
	if singleton {
		key := string(input)
		config, ok := s.m[key]
		if ok {
			config.Cnt++
			return config.Token
		}
		token := s.Token.New(input)
		s.m[key] = &TokenCounter{
			Cnt:   1,
			Token: token,
		}
		return token
	}
	return s.Token.New(input)
}

func (s *SingletonTokenWrapper) Config() TokenConfig {
	return s.Token.Config()
}

func NewSingletonWrapper(token Token) *SingletonTokenWrapper {
	return &SingletonTokenWrapper{
		Token: token,
		m:     make(map[string]*TokenCounter),
	}
}
