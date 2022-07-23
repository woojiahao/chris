package parser

import (
	"fmt"
	"github.com/woojiahao/chris/pkg/lexer"
	"strings"
)

type Node interface {
	// Print returns the textual representation of the Node
	Print() string
}

type NumberNode float64

func (nn NumberNode) Print() string {
	return fmt.Sprintf("%f", nn)
}

type VariableNode string

func (vn VariableNode) Print() string {
	return fmt.Sprintf("%s", vn)
}

type KeywordNode string

func (kn KeywordNode) Print() string {
	return fmt.Sprintf("%s", kn)
}

// TODO: Load the alternate value?
type ConstantNode string

func (cn ConstantNode) Print() string {
	return fmt.Sprintf("%s", cn)
}

type AssignmentNode struct {
	Variable VariableNode
	Right    Node
}

func (an AssignmentNode) Print() string {
	return fmt.Sprintf("%s = %s", an.Variable.Print(), an.Right.Print())
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

type PrefixNode struct {
	PrefixToken lexer.TokenType
	Right       Node
}

func (pn PrefixNode) Print() string {
	return fmt.Sprintf("%s%s", pn.PrefixToken.Symbol, pn.Right.Print())
}
