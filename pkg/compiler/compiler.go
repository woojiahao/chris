package compiler

import (
	"woojiahao.com/chris/internal/lexer"
	"woojiahao.com/chris/internal/parser"
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
