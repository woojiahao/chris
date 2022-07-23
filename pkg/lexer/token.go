package lexer

import "fmt"

type TokenType struct {
	Name       string
	Precedence int
	Symbol     string
}

var (
	Number           = TokenType{"NUMBER", -1, ""}
	Variable         = TokenType{"VARIABLE", -1, ""}
	Keyword          = TokenType{"KEYWORD", -1, ""}
	Constant         = TokenType{"CONSTANT", -1, ""}
	RightParenthesis = TokenType{"RIGHT PARENTHESIS", -1, ")"}
	Comma            = TokenType{"COMMA", -1, ","}
	EndOfExpression  = TokenType{"END OF EXPRESSION", -9999, ""}
	Assignment       = TokenType{"ASSIGNMENT", 1, "="}
	Minus            = TokenType{"MINUS", 2, "-"}
	Add              = TokenType{"ADD", 2, "+"}
	Divide           = TokenType{"DIVIDE", 3, "/"}
	Multiply         = TokenType{"MULTIPLY", 3, "*"}
	Exponent         = TokenType{"EXPONENT", 4, "^"}
	LeftParenthesis  = TokenType{"LEFT PARENTHESIS", 5, "("}
)

type Token struct {
	TokenType TokenType
	Value     float64
	Text      string
}

func NewNumber(value float64) *Token {
	return &Token{Number, value, fmt.Sprintf("%f", value)}
}

func NewVariable(variable string) *Token {
	return &Token{Variable, 0, variable}
}

func NewKeyword(function string) *Token {
	return &Token{Keyword, 0, function}
}

func NewConstant(constant string) *Token {
	return &Token{Constant, 0, constant}
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
