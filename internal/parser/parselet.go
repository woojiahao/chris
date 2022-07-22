package parser

import (
	"woojiahao.com/chris/internal/lexer"
)

// PrefixParselet is anything that has no left Node
type PrefixParselet interface {
	Parse(parser *Parser, token *lexer.Token) Node
}

type VariableParselet struct{}

func (vp VariableParselet) Parse(parser *Parser, token *lexer.Token) Node {
	return VariableNode(token.Text)
}

type KeywordParselet struct{}

func (kp KeywordParselet) Parse(parser *Parser, token *lexer.Token) Node {
	return KeywordNode(token.Text)
}

type NumberParselet struct{}

func (np NumberParselet) Parse(parser *Parser, token *lexer.Token) Node {
	return NumberNode(token.Value)
}

// GroupParselet creates a sub-expression group using the prefix version of (
type GroupParselet struct{}

func (gp GroupParselet) Parse(parser *Parser, token *lexer.Token) Node {
	// As the GroupParselet needs to make the inner sub-expression more important, we start parsing with a precedence
	// of 0
	subExpression := parser.parseExpression(0)
	if !parser.expect(lexer.RightParenthesis) {
		// If the next token in the expression is not ), we can just panic
		panic("Expected ), did not receive it at the end of the sub-expression")
	}
	parser.consume()
	return subExpression
}

type PrefixOperatorParselet struct{}

func (pop PrefixOperatorParselet) Parse(parser *Parser, token *lexer.Token) Node {
	right := parser.parseExpression(token.TokenType.Precedence)
	return PrefixNode{token, right}
}

// InfixParselet is anything that has a left Node (may not have a right Node like the last character but it is still an
// infix
type InfixParselet interface {
	Parse(parser *Parser, left Node, token *lexer.Token) Node
}

type BinaryOperatorParselet struct {
	isRight bool
}

func (bop BinaryOperatorParselet) Parse(parser *Parser, left Node, token *lexer.Token) Node {
	// TODO: Parse according to type of binary operator
	tick := 0
	if bop.isRight {
		tick = 1
	}
	right := parser.parseExpression(token.TokenType.Precedence - tick)

	return OperatorNode{left, right, token.TokenType}
}

// FunctionCallParselet creates a function call that parses the arguments within (). It is the infix ( operator
type FunctionCallParselet struct{}

func (fcp FunctionCallParselet) Parse(parser *Parser, left Node, token *lexer.Token) Node {
	// left is a KeywordNode while token is the ( infix operator
	keywordNode, _ := left.(KeywordNode)

	var args []Node

	// We check if the next token is ) and if it is, we consume and don't enter the loop
	// If it is not ), we don't consume and instead, start parsing the arguments
	if !parser.expectAndConsume(lexer.RightParenthesis) {
		for {
			// Parse the fist argument first before any checks
			args = append(args, parser.parseExpression(0))

			// If the next token is not a comma, we don't parse it
			// Exit from the loop and check for )
			if !parser.expectAndConsume(lexer.Comma) {
				break
			}
		}

		// If it is ), consume, if not, don't consume
		parser.expectAndConsume(lexer.RightParenthesis)
	}

	return FunctionNode{keywordNode, args}
}

type AssignmentParselet struct{}

func (ap AssignmentParselet) Parse(parser *Parser, left Node, token *lexer.Token) Node {
	variableNode, _ := left.(VariableNode)

	right := parser.parseExpression(lexer.Assignment.Precedence - 1)
	return AssignmentNode{variableNode, right}
}

type PostfixOperatorParselet struct{}

func (pop PostfixOperatorParselet) Parse(parser *Parser, left Node, token *lexer.Token) Node {
	return PostfixNode{left, token.TokenType}
}
