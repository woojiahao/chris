package parser

import (
	"fmt"
	"strings"
	"woojiahao.com/chris/internal/lexer"
)

type Node interface {
	Print() string
}

type NumberNode struct {
	Value float64
}

func (nn NumberNode) Print() string {
	return fmt.Sprintf("%f", nn.Value)
}

type AssignmentNode struct {
	Variable VariableNode
	Right    Node
}

func (an AssignmentNode) Print() string {
	return fmt.Sprintf("%s = %s", an.Variable.Print(), an.Right.Print())
}

type VariableNode struct {
	Variable string
}

func (vn VariableNode) Print() string {
	return fmt.Sprintf("%s", vn.Variable)
}

type OperatorNode struct {
	Left     Node
	Right    Node
	Operator lexer.TokenType
}

func (on OperatorNode) Print() string {
	return fmt.Sprintf("(%s %s %s)", on.Left.Print(), on.Operator.Symbol, on.Right.Print())
}

type PostfixNode struct {
	Left     Node
	Operator lexer.TokenType
}

func (pn PostfixNode) Print() string {
	return fmt.Sprintf("(%s%s)", pn.Left.Print(), pn.Operator.Symbol)
}

type FunctionNode struct {
	Keyword   KeywordNode
	Arguments []Node
}

func (fn FunctionNode) Print() string {
	var arguments []string
	for _, arg := range fn.Arguments {
		arguments = append(arguments, arg.Print())
	}
	return fmt.Sprintf("%s(%s)", fn.Keyword.Print(), strings.Join(arguments, ", "))
}

type KeywordNode struct {
	Keyword string
}

func (kn KeywordNode) Print() string {
	return fmt.Sprintf("%s", kn.Keyword)
}

type PrefixNode struct {
	PrefixToken *lexer.Token
	Right       Node
}

func (pn PrefixNode) Print() string {
	return fmt.Sprintf("%s%s", pn.PrefixToken.Text, pn.Right.Print())
}
