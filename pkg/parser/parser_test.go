package parser

import (
	"fmt"
	"github.com/woojiahao/chris/pkg/lexer"
	"reflect"
	"testing"
)

func TestTerminalNodes(t *testing.T) {
	cases := []parserCase{
		{"3", NumberNode(3)},
		{"a", VariableNode("a")},
		{"sin", KeywordNode("sin")},
	}
	testParser(t, cases, true)
}

func TestPrefixNode_Assert(t *testing.T) {
	cases := []parserCase{
		{"-3", PrefixNode{lexer.Minus, NumberNode(3)}},
		{"-a", PrefixNode{lexer.Minus, VariableNode("a")}},
		{"-sin", PrefixNode{lexer.Minus, KeywordNode("sin")}},
	}
	testParser(t, cases, true)
}

func TestOperatorNode_Assert(t *testing.T) {
	cases := []parserCase{
		{"3 + 3", OperatorNode{NumberNode(3), NumberNode(3), lexer.Add}},
		{"3 - 3", OperatorNode{NumberNode(3), NumberNode(3), lexer.Minus}},
		{"3 * 3", OperatorNode{NumberNode(3), NumberNode(3), lexer.Multiply}},
		{"3 / 3", OperatorNode{NumberNode(3), NumberNode(3), lexer.Divide}},
		{"3 ^ 3", OperatorNode{NumberNode(3), NumberNode(3), lexer.Exponent}},
	}
	testParser(t, cases, true)
}

type parserCase struct {
	expression string
	expected   Node
}

func setupParser(exp string) *Parser {
	l := lexer.New(exp)
	p := New(l)
	return p
}

func testParser(t *testing.T, cases []parserCase, isAssert bool) {
	for _, c := range cases {
		p := setupParser(c.expression)
		if isAssert {
			assert(t, p, c.expected)
		}
	}
}

func assert(t *testing.T, p *Parser, expected Node) {
	if result := p.Parse(); !equals(result, expected) {
		t.Errorf("Expected %v (%t), got %v (%t) instead", result, result, expected, expected)
	}
}

// TODO: Add errors to main code to test for error strings
func expect(t *testing.T, p *Parser, error string) {

}

func equals(n1, n2 Node) bool {
	// If n1 != n2 where one is nil, return false, otherwise if both are nil, return true
	if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) {
		return false
	}
	if n1 == nil && n2 == nil {
		return true
	}

	// Ensure that the types are the same to traverse in same order
	if reflect.TypeOf(n1) != reflect.TypeOf(n2) {
		return false
	}

	// Guaranteed to be the same type, so we can traverse in the same order
	switch v1 := n1.(type) {
	// Terminal nodes
	case NumberNode:
		v2 := n2.(NumberNode)
		return v1 == v2
	case VariableNode:
		v2 := n2.(VariableNode)
		return v1 == v2
	case KeywordNode:
		v2 := n2.(KeywordNode)
		return v1 == v2

	// Non-terminal nodes
	case PrefixNode:
		v2 := n2.(PrefixNode)
		return v1.PrefixToken == v2.PrefixToken && equals(v1.Right, v2.Right)
	case PostfixNode:
		v2 := n2.(PostfixNode)
		return v1.Operator == v2.Operator && equals(v1.Left, v2.Left)
	case AssignmentNode:
		v2 := n2.(AssignmentNode)
		return equals(v1.Variable, v2.Variable) && equals(v1.Right, v2.Right)
	case OperatorNode:
		v2 := n2.(OperatorNode)
		return v1.Operator == v2.Operator && equals(v1.Left, v2.Left) && equals(v1.Left, v2.Left)
	case FunctionNode:
		v2 := n2.(FunctionNode)
		if len(v1.Arguments) != len(v2.Arguments) {
			return false
		}

		for i, arg1 := range v1.Arguments {
			arg2 := v2.Arguments[i]
			if !equals(arg1, arg2) {
				return false
			}
		}

		return equals(v1.Keyword, v2.Keyword)
	default:
		panic(fmt.Sprintf("Invalid node: %v", n1))
	}
}
