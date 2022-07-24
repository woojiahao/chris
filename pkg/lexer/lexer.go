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

	nextExpression := l.nextExpression.substring(j, l.nextExpression.Len()-1)

	// Parse the current token if it's a keyword since it can tokenizeKeyword into variable, constant, or function (keyword) still
	// TODO: Explore splitting in the actual look ahead
	if token.TokenType == Keyword {
		splitTokens := l.tokenizeKeyword(token.Text)
		if len(splitTokens) > 1 {
			splitTokensText := utils.Map(splitTokens, func(t *Token) string {
				return t.Text
			})
			nextExpression = "*" + strings.Join(splitTokensText[1:], "*") + nextExpression
			return splitTokens[0], nextExpression
		} else {
			// Use the appropriately cast version of token, not the Keyword
			token = splitTokens[0]
		}
	}

	// Get the very next token in the nextExpression
	next, _ := l.nextExpression.lookAhead(j)

	acceptedPairs := map[TokenType][]TokenType{
		Number:   {Variable, Keyword, LeftParenthesis}, // 3x, 3sin, 3(
		Variable: {Keyword, LeftParenthesis},           // xsin, x(
		Constant: {Variable, LeftParenthesis, Keyword}, // pi4, pi(, pisin
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

	for i, ch := range keyword {
		cur += string(ch)

		if utils.In(l.constants, cur) {
			// If the whole cur is a constant, return the constant
			tokens = append(tokens, NewConstant(cur))
			cur = ""
			continue
		} else if utils.In(l.keywords, cur) {
			// If the whole cur is a function, return the function
			tokens = append(tokens, NewKeyword(cur))
			cur = ""
			continue
		} else {
			// Else, see if it's about to form a constant/function using fuzzy matching from the start
			mightBe := func(lst []string) bool {
				anyFits := utils.Any(lst, func(c string) bool {
					// Check if constants have length <= remaining length available
					return len(c)-len(cur) <= len(keyword)-i-1
				})
				anyMatch := utils.Any(lst, func(c string) bool {
					return utils.FuzzyMatch(c, cur)
				})

				return anyFits && anyMatch
			}
			if mightBe(l.constants) || mightBe(l.keywords) {
				continue
			} else {
				// If it's really forming nothing, convert everything thus far to a variable
				for _, c := range cur {
					tokens = append(tokens, NewVariable(string(c)))
				}
				cur = ""
			}
		}
	}

	return tokens
}
