package parser

import (
	"github.com/woojiahao/chris/pkg/lexer"
)

// PrefixParselet is anything that has no left Node
type PrefixParselet interface {
	Parse(parser *Parser, token *lexer.Token) (Node, error)
}

type VariableParselet struct{}

func (vp VariableParselet) Parse(parser *Parser, token *lexer.Token) (Node, error) {
	return VariableNode(token.Text), nil
}

type KeywordParselet struct{}

func (kp KeywordParselet) Parse(parser *Parser, token *lexer.Token) (Node, error) {
	return KeywordNode(token.Text), nil
}

type NumberParselet struct{}

func (np NumberParselet) Parse(parser *Parser, token *lexer.Token) (Node, error) {
	return NumberNode(token.Value), nil
}

// GroupParselet creates a sub-expression group using the prefix version of (
type GroupParselet struct{}

func (gp GroupParselet) Parse(parser *Parser, token *lexer.Token) (Node, error) {
	// As the GroupParselet needs to make the inner sub-expression more important, we start parsing with a precedence
	// of 0
	subExpression, err := parser.parseExpression(0)
	if err != nil {
		return nil, err
	}

	if !parser.expect(lexer.RightParenthesis) {
		// If the next token in the expression is not ), we can just panic
		return nil, &ParseError{token.TokenType, invalidEndOfGroup}
	}

	_, err = parser.consume()
	if err != nil {
		return nil, err
	}

	return subExpression, nil
}

type PrefixOperatorParselet struct{}

func (pop PrefixOperatorParselet) Parse(parser *Parser, token *lexer.Token) (Node, error) {
	right, err := parser.parseExpression(token.TokenType.Precedence)
	if err != nil {
		return nil, err
	}
	return PrefixNode{token.TokenType, right}, nil
}

// InfixParselet is anything that has a left Node (may not have a right Node like the last character but it is still an
// infix
type InfixParselet interface {
	Parse(parser *Parser, left Node, token *lexer.Token) (Node, error)
}

type BinaryOperatorParselet struct {
	isRight bool
}

func (bop BinaryOperatorParselet) Parse(parser *Parser, left Node, token *lexer.Token) (Node, error) {
	// TODO: Parse according to type of binary operator
	tick := 0
	if bop.isRight {
		tick = 1
	}
	right, err := parser.parseExpression(token.TokenType.Precedence - tick)
	if err != nil {
		return nil, err
	}

	return OperatorNode{left, right, token.TokenType}, nil
}

// FunctionCallParselet creates a function call that parses the arguments within (). It is the infix ( operator
type FunctionCallParselet struct{}

func (fcp FunctionCallParselet) Parse(parser *Parser, left Node, token *lexer.Token) (Node, error) {
	// left is a KeywordNode while token is the ( infix operator
	keywordNode, ok := left.(KeywordNode)
	if !ok {
		return nil, &ParseError{token.TokenType, invalidKeywordInFunctionCall}
	}

	var args []Node

	// We check if the next token is ) and if it is, we consume and don't enter the loop
	// If it is not ), we don't consume and instead, start parsing the arguments
	// TODO: Enforce checks strictly
	if !parser.expectAndConsume(lexer.RightParenthesis) {
		for {
			// Parse the fist argument first before any checks
			arg, err := parser.parseExpression(0)
			if err != nil {
				return nil, err
			}
			args = append(args, arg)

			// If the next token is not a comma, we don't parse it
			// Exit from the loop and check for )
			if !parser.expectAndConsume(lexer.Comma) {
				break
			}
		}

		// If it's not a ), return an error
		if !parser.expect(lexer.RightParenthesis) {
			return nil, &ParseError{token.TokenType, invalidEndOfFunction}
		}
		// Guaranteed to consume a ) token
		parser.consume()
	}

	return FunctionNode{keywordNode, args}, nil
}

type AssignmentParselet struct{}

func (ap AssignmentParselet) Parse(parser *Parser, left Node, token *lexer.Token) (Node, error) {
	variableNode, ok := left.(VariableNode)
	if !ok {
		return nil, &ParseError{token.TokenType, invalidVariableInAssignment}
	}

	right, err := parser.parseExpression(lexer.Assignment.Precedence - 1)
	if err != nil {
		return nil, err
	}
	return AssignmentNode{variableNode, right}, nil
}

type PostfixOperatorParselet struct{}

func (pop PostfixOperatorParselet) Parse(parser *Parser, left Node, token *lexer.Token) Node {
	return PostfixNode{left, token.TokenType}
}
