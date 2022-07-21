package lexer

type TokenType int

const (
	Assignment TokenType = iota
	Number
	Variable
	Function
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

func NewFunction(function string) *Token {
	return &Token{Variable, 0, function}
}
