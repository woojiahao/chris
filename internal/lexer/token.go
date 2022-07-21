package lexer

type TokenType int

const (
	Assignment TokenType = iota
	Number
	Variable
	Keyword
	Add
	Minus
	Divide
	Multiply
	Exponent
	LeftParenthesis
	RightParenthesis
)

type Token struct {
	TokenType TokenType
	Value     float64
	Text      string
}

func NewNumber(value float64) *Token {
	return &Token{Number, value, ""}
}

// TODO: Allow variables and functions to exist in the same name pool instead of having variables just be a single char
func NewVariable(variable string) *Token {
	return &Token{Variable, 0, variable}
}

func NewKeyword(function string) *Token {
	return &Token{Keyword, 0, function}
}

func NewOperator(operator rune) *Token {
	tokenType := Add
	switch operator {
	case '+':
		tokenType = Add
	case '-':
		tokenType = Minus
	case '/':
		tokenType = Divide
	case '*':
		tokenType = Multiply
	case '^':
		tokenType = Exponent
	case '=':
		tokenType = Assignment
	case '(':
		tokenType = LeftParenthesis
	case ')':
		tokenType = RightParenthesis
	}

	return &Token{tokenType, 0, string(operator)}
}
