package tokens

type Token struct {
	token     any
	tokenType TokenType
}

func NewToken(kind TokenType, value any) *Token {
	return &Token{
		token:     value,
		tokenType: kind,
	}
}

func (s Token) Value() any {
	return s.token
}

func (s Token) Type() TokenType {
	return s.tokenType
}