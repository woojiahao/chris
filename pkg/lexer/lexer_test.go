package lexer

import "testing"

type lexerCase struct {
	expression     string
	expectedTokens []Token
}

func nextOnCases(t *testing.T, cases []lexerCase) {
	for _, c := range cases {
		l := New(c.expression)

		for _, expected := range c.expectedTokens {
			result := l.Next()
			if result.Value != expected.Value {
				t.Errorf("Expected %f, got %f instead", result.Value, expected.Value)
			}

			if expected.Text != "" {
				if result.Text != expected.Text {
					t.Errorf("Expected %s, got %s instead", expected.Text, result.Text)
				}
			}
		}
	}
}

func TestLexer_Next_OperatorlessExpressions(t *testing.T) {
	cases := []lexerCase{
		{"3x", []Token{
			*NewNumber(3),
			*NewOperator('*'),
			*NewVariable("x"),
		}},
		{"3sin(3x)", []Token{
			*NewNumber(3),
			*NewOperator('*'),
			*NewKeyword("sin"),
			*NewOperator('('),
			*NewNumber(3),
			*NewOperator('*'),
			*NewVariable("x"),
			*NewOperator(')'),
		}},
		{"x(1 + 2)", []Token{
			*NewVariable("x"),
			*NewOperator('*'),
			*NewKeyword("sin"),
			*NewOperator('('),
			*NewNumber(1),
			*NewOperator('+'),
			*NewNumber(2),
			*NewOperator(')'),
		}},
		{"xsin(1 + 2)", []Token{
			*NewVariable("x"),
			*NewOperator('*'),
			*NewKeyword("sin"),
			*NewOperator('('),
			*NewNumber(1),
			*NewOperator('+'),
			*NewNumber(2),
			*NewOperator(')'),
		}},
	}

	nextOnCases(t, cases)
}

func TestLexer_Next_WithVariablesAndKeywords(t *testing.T) {
	cases := []lexerCase{
		{"abc", []Token{
			*NewKeyword("abc"),
		}},
		{"a = sin(15)", []Token{
			*NewVariable("a"),
			*NewOperator('='),
			*NewKeyword("sin"),
			*NewOperator('('),
			*NewNumber(15),
			*NewOperator(')'),
		}},
	}

	nextOnCases(t, cases)
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
	}

	nextOnCases(t, cases)
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
		{"(1 + 2) * 3", []Token{
			*NewOperator('('),
			*NewNumber(1),
			*NewOperator('+'),
			*NewNumber(2),
			*NewOperator(')'),
			*NewOperator('*'),
			*NewNumber(3),
		}},
	}

	nextOnCases(t, cases)
}
