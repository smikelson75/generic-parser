package tokens

type TokenType int

const (
	IDENTIFIER TokenType = iota
	COMPLEXIDENTIFIER
	KEYWORD
	OPERATOR
	SYMBOL
	NUMBER
	STRING
	COMMENT
	NEWLINE
	WHITESPACE
)

func (t TokenType) String() string {
	switch t {
	case IDENTIFIER:
		return "IDENTIFIER"
	case COMPLEXIDENTIFIER:
		return "COMPLEXIDENTIFIER"
	case KEYWORD:
		return "KEYWORD"
	case OPERATOR:
		return "OPERATOR"
	case SYMBOL:
		return "SYMBOL"
	case NUMBER:
		return "NUMBER"
	case STRING:
		return "STRING"
	case COMMENT:
		return "COMMENT"
	case NEWLINE:
		return "NEWLINE"
	case WHITESPACE:
		return "WHITESPACE"
	default:
		return "UNKNOWN"
	}
}