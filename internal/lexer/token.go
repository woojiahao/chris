package lexer

import "fmt"

type TokenType struct {
	Precedence int
}

var (
	Assignment       = TokenType{1}
	Number           = TokenType{6}
	Variable         = TokenType{6}
	Keyword          = TokenType{8}
	Minus            = TokenType{3}
	Add              = TokenType{3}
	Divide           = TokenType{4}
	Multiply         = TokenType{4}
	Exponent         = TokenType{6}
	LeftParenthesis  = TokenType{0}
	RightParenthesis = TokenType{0}
	Comma            = TokenType{-1}
)

type Token struct {
	TokenType TokenType
	Value     float64
	Text      string
}

func NewNumber(value float64) *Token {
	return &Token{Number, value, fmt.Sprintf("%f", value)}
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
	case ',':
		tokenType = Comma
	}

	return &Token{tokenType, 0, string(operator)}
}
