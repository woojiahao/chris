package compiler

import (
	"github.com/woojiahao/chris/internal/lexer"
	"github.com/woojiahao/chris/internal/parser"
)

type Compiler struct {
	lexer  *lexer.Lexer
	parser *parser.Parser
}

func New(exp string) *Compiler {
	l := lexer.New(exp)
	p := parser.New(l)
	return &Compiler{l, p}
}
