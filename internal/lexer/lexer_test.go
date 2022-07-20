package lexer

import "testing"

func TestLexer_Next(t *testing.T) {
	l := New("24 39")
	token := l.Next()
	if token.Value != 24 {
		t.Errorf("token.Value is %f, not 24", token.Value)
	}

	token = l.Next()
	if token.Value != 39 {
		t.Errorf("token.Value is %f, not 39", token.Value)
	}
}
