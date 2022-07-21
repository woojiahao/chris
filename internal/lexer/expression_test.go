package lexer

import (
	"testing"
)

type expressionCase struct {
	exp          Expression
	start        int
	expected     float64
	expectedNext int
}

func TestExpression_readNumber(t *testing.T) {
	cases := []expressionCase{
		{"5", 0, 5, 1},
		{"123 ", 0, 123, 3},
		{" 123,", 1, 123, 4},
	}

	for _, c := range cases {
		result, next := c.exp.readNumber(c.start)
		if *result != c.expected {
			t.Errorf("Expected %f, but got %f instead", c.expected, *result)
		}

		if next != c.expectedNext {
			t.Errorf("Expected next %d, but got %d instead", c.expectedNext, next)
		}
	}
}
