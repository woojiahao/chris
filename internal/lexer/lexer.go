package lexer

import (
	"unicode"
)

type Lexer struct {
	expression *Expression
	index      int
	token      *Token
}

func New(exp Expression) *Lexer {
	return &Lexer{&exp, 0, nil}
}

// Next returns token at the current position of index and moves the head to the next character for further calls
func (l *Lexer) Next() *Token {
	if l.index >= l.expression.Len() {
		// If the index to search is at the last character or even further, we just want to return a nil token
		return nil
	}

	for {
		// Create an infinite loop to account for any long characters like digits or sub-expressions like
		// trigonometric expressions
		cur := l.expression.Get(l.index)

		if cur == nil {
			// If the character at the cur position is not a valid, return nil token
			return nil
		}

		if unicode.IsSpace(*cur) {
			// If cur is a space, we will want to keep iterating till we reach a non-space character
			l.index++
			continue
		}

		if unicode.IsDigit(*cur) {
			// If the cur character is a number, we want to create a token of type Number and continue reading the
			// input until we no longer encounter a digit
			num, next := l.expression.readNumber(l.index)
			l.index = next
			l.token = &Token{Number, *num, *cur}
			break
		}

		// For current purposes, if current is neither a whitespace nor number, we just set the symbol to a variable and
		// return the appropriate token
		l.token = &Token{Variable, 0, '0'}
		l.index++
		break
	}

	return l.token
}
