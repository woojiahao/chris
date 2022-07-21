package lexer

import "testing"

type lexerCase struct {
	expression     Expression
	expectedTokens []Token
}

func TestLexer_Next(t *testing.T) {
	cases := []lexerCase{
		{"24 39", []Token{
			{Number, 24, '2'},
			{Number, 39, '3'},
		}},
		{"  24   39", []Token{
			{Number, 24, '2'},
			{Number, 39, '3'},
		}},
		{"  24   39     ", []Token{
			{Number, 24, '2'},
			{Number, 39, '3'},
		}},
		{"abc", []Token{
			{Variable, 0, 'a'},
			{Variable, 0, 'b'},
			{Variable, 0, 'c'},
		}},
	}

	for _, c := range cases {
		l := New(c.expression)

		for _, expected := range c.expectedTokens {
			result := l.Next()
			if result.Value != expected.Value {
				t.Errorf("Expected %f, got %f instead", expected.Value, result.Value)
			}
		}
	}
}
