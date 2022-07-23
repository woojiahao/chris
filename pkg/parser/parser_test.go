package parser

import (
	"fmt"
	"github.com/woojiahao/chris/pkg/lexer"
	"reflect"
	"testing"
)

func TestTerminalNodes(t *testing.T) {
	cases := []parserCase{
		assertParserCase("3", NumberNode(3)),
		assertParserCase("17.6", NumberNode(17.6)),
		assertParserCase("a", VariableNode("a")),
		assertParserCase("sin", KeywordNode("sin")),
	}
	assertCases(t, cases)
}

func TestPrefixNode_Assert(t *testing.T) {
	cases := []parserCase{
		assertParserCase("-3", PrefixNode{lexer.Minus, NumberNode(3)}),
		assertParserCase("-a", PrefixNode{lexer.Minus, VariableNode("a")}),
		assertParserCase("-sin", PrefixNode{lexer.Minus, KeywordNode("sin")}),
	}
	assertCases(t, cases)
}

func TestPrefixNode_Expect(t *testing.T) {
	cases := []parserCase{
		expectParserCase("^3", invalidPrefixToken),
	}
	expectCases(t, cases)
}

func TestOperatorNode_Assert(t *testing.T) {
	cases := []parserCase{
		assertParserCase("3 + 3", OperatorNode{NumberNode(3), NumberNode(3), lexer.Add}),
		assertParserCase("3 - 3", OperatorNode{NumberNode(3), NumberNode(3), lexer.Minus}),
		assertParserCase("3 * 3", OperatorNode{NumberNode(3), NumberNode(3), lexer.Multiply}),
		assertParserCase("3 / 3", OperatorNode{NumberNode(3), NumberNode(3), lexer.Divide}),
		assertParserCase("3 ^ 3", OperatorNode{NumberNode(3), NumberNode(3), lexer.Exponent}),
	}
	assertCases(t, cases)
}

func TestFunctionNode_Assert(t *testing.T) {
	cases := []parserCase{
		assertParserCase("sin()", FunctionNode{"sin", []Node{}}),
		assertParserCase("cos(x)", FunctionNode{"cos", []Node{VariableNode("x")}}),
		assertParserCase("tan(x, y, z)", FunctionNode{"tan", []Node{VariableNode("x"), VariableNode("y"), VariableNode("z")}}),
		assertParserCase("csc(1)", FunctionNode{"csc", []Node{NumberNode(1)}}),
		assertParserCase("sec(1, 2, 17)", FunctionNode{"sec", []Node{NumberNode(1), NumberNode(2), NumberNode(17)}}),
		assertParserCase("cot(pi)", FunctionNode{"cot", []Node{KeywordNode("pi")}}),
	}
	assertCases(t, cases)
}

func TestFunctionNode_Expect(t *testing.T) {
	cases := []parserCase{
		expectParserCase("sin(", invalidEndOfFunction),
		expectParserCase("a(", invalidKeywordInFunctionCall),
	}
	expectCases(t, cases)
}

func TestAssignmentNode_Assert(t *testing.T) {
	cases := []parserCase{
		assertParserCase("a = 1", AssignmentNode{"a", NumberNode(1)}),
		assertParserCase("a = b", AssignmentNode{"a", VariableNode("b")}),
		assertParserCase("a = sin", AssignmentNode{"a", KeywordNode("sin")}),
	}
	assertCases(t, cases)
}

func TestAssignmentNode_Expect(t *testing.T) {
	cases := []parserCase{
		expectParserCase("sin = 2", invalidVariableInAssignment),
		expectParserCase("3 = 2", invalidVariableInAssignment),
		expectParserCase("a =", invalidEndOfAssignment),
	}
	expectCases(t, cases)
}

func TestPrecedence(t *testing.T) {
	// Each test case builds on the previous test case's determination of precedence
	// I.e. we start by establishing precedence of + and - are equal, then * and / are equal, then * and / are greater
	// than + and -, and finally ^ is greater than all of them
	cases := []parserCase{
		// + == -
		assertParserCase("1 + 2 + 3", OperatorNode{OperatorNode{NumberNode(1), NumberNode(2), lexer.Add}, NumberNode(3), lexer.Add}),
		assertParserCase("1 - 2 + 3", OperatorNode{OperatorNode{NumberNode(1), NumberNode(2), lexer.Minus}, NumberNode(3), lexer.Add}),
		assertParserCase("1 - 2 - 3", OperatorNode{OperatorNode{NumberNode(1), NumberNode(2), lexer.Minus}, NumberNode(3), lexer.Minus}),

		// * == /
		assertParserCase("1 * 2 * 3", OperatorNode{OperatorNode{NumberNode(1), NumberNode(2), lexer.Multiply}, NumberNode(3), lexer.Multiply}),
		assertParserCase("1 / 2 * 3", OperatorNode{OperatorNode{NumberNode(1), NumberNode(2), lexer.Divide}, NumberNode(3), lexer.Multiply}),
		assertParserCase("1 / 2 / 3", OperatorNode{OperatorNode{NumberNode(1), NumberNode(2), lexer.Divide}, NumberNode(3), lexer.Divide}),

		// */ > +-
		assertParserCase("1 + 2 * 3", OperatorNode{NumberNode(1), OperatorNode{NumberNode(2), NumberNode(3), lexer.Multiply}, lexer.Add}),
		assertParserCase("1 + 2 / 3", OperatorNode{NumberNode(1), OperatorNode{NumberNode(2), NumberNode(3), lexer.Divide}, lexer.Add}),

		// ^ > */ > +-
		assertParserCase("1 * 2 ^ 3", OperatorNode{NumberNode(1), OperatorNode{NumberNode(2), NumberNode(3), lexer.Exponent}, lexer.Multiply}),

		// <variable> == <number>
		assertParserCase("2 + a + b", OperatorNode{OperatorNode{NumberNode(2), VariableNode("a"), lexer.Add}, VariableNode("b"), lexer.Add}),
		assertParserCase("a + a + b", OperatorNode{OperatorNode{VariableNode("a"), VariableNode("a"), lexer.Add}, VariableNode("b"), lexer.Add}),

		// <keyword> == <variable>
		assertParserCase("2 + pi + pi", OperatorNode{OperatorNode{NumberNode(2), KeywordNode("pi"), lexer.Add}, VariableNode("pi"), lexer.Add}),
		assertParserCase("pi + pi + pi", OperatorNode{OperatorNode{KeywordNode("pi"), KeywordNode("pi"), lexer.Add}, VariableNode("pi"), lexer.Add}),
	}
	assertCases(t, cases)
}

type parserCase struct {
	expression     string
	assert         Node
	expectedReason errorReason
}

func assertCases(t *testing.T, cases []parserCase) {
	testParser(t, cases, true)
}

func expectCases(t *testing.T, cases []parserCase) {
	testParser(t, cases, false)
}

func assertParserCase(expression string, assert Node) parserCase {
	return parserCase{expression, assert, ""}
}

func expectParserCase(expression string, expectedReason errorReason) parserCase {
	return parserCase{expression, nil, expectedReason}
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
			assert(t, p, c.assert)
		} else {
			expect(t, p, c.expectedReason)
		}
	}
}

func assert(t *testing.T, p *Parser, expected Node) {
	if result, err := p.Parse(); err != nil {
		t.Errorf("Unexpected error encountered %v", err)
	} else if !equals(result, expected) {
		t.Errorf("Expected %v (%t), got %v (%t) instead", result, result, expected, expected)
	}
}

func expect(t *testing.T, p *Parser, expectedReason errorReason) {
	if result, err := p.Parse(); result != nil || err == nil {
		t.Errorf("Expression should have produced a ParseError")
	} else if parseError, ok := err.(*ParseError); !ok {
		t.Errorf("Expected ParseError, got %t instead", err)
	} else if parseError.reason != expectedReason {
		t.Errorf("Expected error '%s', got '%s' instead", expectedReason, parseError.reason)
	}
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
