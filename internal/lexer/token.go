package lexer

type TokenType int

const (
	Assignment TokenType = iota
	Number
	Variable
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
	Symbol    rune
}
