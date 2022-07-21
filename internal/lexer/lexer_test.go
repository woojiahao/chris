package lexer

import "testing"

type lexerCase struct {
	expression     Expression
	expectedTokens []Token
}

func nextOnCases(cases []lexerCase, t *testing.T) {
	for _, c := range cases {
		l := New(c.expression)

		for _, expected := range c.expectedTokens {
			result := l.Next()
			if result.Value != expected.Value {
				t.Errorf("Expected %f, got %f instead", result.Value, expected.Value)
			}

			if result.Text != expected.Text {
				t.Errorf("Expected %s, got %s instead", result.Text, expected.Text)
			}
		}
	}
}

func TestLexer_Next_NumbersOnly(t *testing.T) {
	cases := []lexerCase{
		{"24 39", []Token{
			{Number, 24, ""},
			{Number, 39, ""},
		}},
		{"  24   39", []Token{
			{Number, 24, ""},
			{Number, 39, ""},
		}},
		{"  24   39     ", []Token{
			{Number, 24, ""},
			{Number, 39, ""},
		}},
		{"abc", []Token{
			{Variable, 0, "a"},
			{Variable, 0, "b"},
			{Variable, 0, "c"},
		}},
	}

	nextOnCases(cases, t)
}

func TestLexer_Next_Operators(t *testing.T) {
	cases := []lexerCase{
		{"1+2", []Token{
			*NewNumber(1),
			*NewOperator('+'),
			*NewNumber(2),
		}},
		{"  1 +    2  ", []Token{
			*NewNumber(1),
			*NewOperator('+'),
			*NewNumber(2),
		}},
		{"6899*17", []Token{
			*NewNumber(6899),
			*NewOperator('*'),
			*NewNumber(17),
		}},
		{"6899/17", []Token{
			*NewNumber(6899),
			*NewOperator('/'),
			*NewNumber(17),
		}},
		{"2^n", []Token{
			*NewNumber(2),
			*NewOperator('^'),
			*NewVariable("n"),
		}},
		{"15- 10", []Token{
			*NewNumber(15),
			*NewOperator('-'),
			*NewNumber(10),
		}},
		{"a = 10", []Token{
			*NewVariable("a"),
			*NewOperator('='),
			*NewNumber(10),
		}},
	}

	nextOnCases(cases, t)
}
