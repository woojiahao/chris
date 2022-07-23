package lexer

import (
	"github.com/woojiahao/chris/internal/utils"
)

// TODO Use regex to check if the current character is an operator? Benchmark this
var operators = []rune{'+', '-', '/', '*', '^', '=', '(', ')', ','}

// TODO: Accept functions and constants and differentiate between the two
type Lexer struct {
	expression     Expression
	nextExpression Expression
	token          *Token
	keywords       []string
}

func New(exp string, keywords []string) *Lexer {
	expression := Expression(exp)
	return &Lexer{
		expression,
		expression,
		nil,
		keywords,
	}
}

// Peek reads the next non-whitespace token and returns the token and the index after reading the token
func (l *Lexer) Peek() (*Token, string) {
	// Get the very first token in the nextExpression
	token, j := l.nextExpression.lookAhead(0)
	// Get the very next token in the nextExpression
	next, _ := l.nextExpression.lookAhead(j)

	nextExpression := l.nextExpression.substring(j, l.nextExpression.Len()-1)

	acceptedPairs := map[TokenType][]TokenType{
		//Keyword: LeftParenthesis, // pi( (not supported yet until word parsing works properly)
		Number:   {Variable, Keyword, LeftParenthesis}, // 3x, 3sin, 3(
		Variable: {Keyword, LeftParenthesis},           // xsin, x(
		Keyword:  {Variable},                           // pi4
	}
	if suffix, ok := acceptedPairs[token.TokenType]; ok && utils.In(suffix, next.TokenType) {
		nextExpression = "*" + nextExpression
	}

	return token, nextExpression
}

// Next returns token at the current position of index and moves the head to the next character for further calls
func (l *Lexer) Next() *Token {
	token, next := l.Peek()
	l.token = token
	l.nextExpression = Expression(next)
	return token
}
