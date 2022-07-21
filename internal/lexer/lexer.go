package lexer

import (
	"unicode"
)

type Lexer struct {
	expression Expression
	index      int
	token      *Token
}

func New(exp Expression) *Lexer {
	return &Lexer{exp, 0, nil}
}

// Peek reads the next non-whitespace token and returns the token and the index after reading the token
func (l *Lexer) Peek() (*Token, int) {
	i := l.index
	for i < l.expression.Len() {
		cur := l.expression.Get(i)

		if cur == nil {
			// Unable to retrieve character at given index
			return nil, -1
		}

		if unicode.IsSpace(*cur) {
			// Move pointer forward if cur is whitespace since we aren't reading whitespaces
			i++
			continue
		}

		if unicode.IsDigit(*cur) {
			// Read the next few characters for the number
			num, next := l.expression.readNumber(i)
			l.token = &Token{Number, *num, *cur}
			return l.token, next
		}

		l.token = &Token{Variable, 0, *cur}
		return l.token, i + 1
	}

	return nil, -1
}

// Next returns token at the current position of index and moves the head to the next character for further calls
func (l *Lexer) Next() *Token {
	token, next := l.Peek()
	l.index = next
	return token
}
