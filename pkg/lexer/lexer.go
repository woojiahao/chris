package lexer

import (
	"github.com/woojiahao/chris/internal/utils"
	"strings"
)

// TODO Use regex to check if the current character is an operator? Benchmark this
var operators = []rune{'+', '-', '/', '*', '^', '=', '(', ')', ','}

type Lexer struct {
	expression     Expression
	nextExpression Expression
	token          *Token
	keywords       []string
	constants      []string
}

func New(exp string, keywords []string, constants []string) *Lexer {
	expression := Expression(exp)

	// TODO: Ensure that constants and keywords do not mix up

	return &Lexer{
		expression,
		expression,
		nil,
		keywords,
		constants,
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

func (l *Lexer) tokenizeKeyword(keyword string) []*Token {
	cur := ""
	var tokens []*Token

	for i := 0; i < len(keyword); i++ {
		cur += string(keyword[i])

		if utils.In(l.constants, cur) {
			// If the whole cur is a constant, return the constant
			tokens = append(tokens, NewConstant(cur))
			cur = ""
			break
		} else if utils.In(l.keywords, cur) {
			// If the whole cur is a function, return the function
			tokens = append(tokens, NewKeyword(cur))
			cur = ""
			break
		} else {
			// Else, see if it's about to form a constant/function using fuzzy matching from the start
			if utils.Any(l.constants, func(c string) bool { return utils.FuzzyMatch(c, cur) }) {
				continue
			} else if utils.Any(l.keywords, func(c string) bool { return utils.FuzzyMatch(c, cur) }) {
				continue
			} else {
				// If it's really forming nothing, convert everything thus far to a variable
				for _, ch := range cur {
					tokens = append(tokens, NewVariable(string(ch)))
				}
				cur = ""
			}
		}
	}

	return tokens
}
