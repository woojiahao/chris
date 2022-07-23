package parser

import (
	"github.com/woojiahao/chris/pkg/lexer"
)

var prefixParselets = map[lexer.TokenType]PrefixParselet{
	lexer.Variable:        VariableParselet{},
	lexer.Number:          NumberParselet{},
	lexer.Keyword:         KeywordParselet{},
	lexer.Constant:        ConstantParselet{},
	lexer.Minus:           PrefixOperatorParselet{},
	lexer.LeftParenthesis: GroupParselet{},
}

var infixParselets = map[lexer.TokenType]InfixParselet{
	lexer.Add:             BinaryOperatorParselet{false},
	lexer.Minus:           BinaryOperatorParselet{false},
	lexer.Divide:          BinaryOperatorParselet{false},
	lexer.Multiply:        BinaryOperatorParselet{false},
	lexer.Exponent:        BinaryOperatorParselet{true},
	lexer.LeftParenthesis: FunctionCallParselet{},
	lexer.Assignment:      AssignmentParselet{},
}

func getPrefixParselet(tokenType lexer.TokenType) (PrefixParselet, error) {
	if prefixParselet, ok := prefixParselets[tokenType]; !ok {
		return nil, &ParseError{tokenType, invalidPrefixToken}
	} else {
		return prefixParselet, nil
	}
}

func getInfixParselet(tokenType lexer.TokenType) (InfixParselet, error) {
	if infixParselet, ok := infixParselets[tokenType]; !ok {
		return nil, &ParseError{tokenType, invalidInfixToken}
	} else {
		return infixParselet, nil
	}
}

type Parser struct {
	lexer *lexer.Lexer
}

func New(lexer *lexer.Lexer) *Parser {
	return &Parser{lexer}
}

func (p *Parser) Parse() (Node, error) {
	return p.parseExpression(0)
}

func (p *Parser) parseExpression(precedence int) (Node, error) {
	token, err := p.consume()
	if err != nil {
		return nil, err
	}

	// Begin parsing the body to the right
	prefixParselet, err := getPrefixParselet(token.TokenType)
	if err != nil {
		return nil, err
	}

	left, err := prefixParselet.Parse(p, token)
	if err != nil {
		return nil, err
	}

	for precedence < p.nextPrecedence() {
		token, err = p.consume()
		if err != nil {
			return nil, err
		}

		infixParselet, err := getInfixParselet(token.TokenType)
		if err != nil {
			return nil, err
		}

		left, err = infixParselet.Parse(p, left, token)
		if err != nil {
			return nil, err
		}
	}

	return left, nil
}

func (p *Parser) nextPrecedence() int {
	nextToken, _ := p.lexer.Peek()
	if nextToken == nil {
		// If nextToken is nil, it should have reached the end of the expression, which we regard as 0
		// TODO: Create fine tune control over the type of errors with peeking
		return 0
	}

	return nextToken.TokenType.Precedence
}

func (p *Parser) consume() (*lexer.Token, error) {
	token := p.lexer.Next()
	if token == nil {
		return nil, &ParseError{lexer.EndOfExpression, endOfExpression}
	}

	return token, nil
}

// expect checks the next token inline and returns whether it is the same as the target token type
func (p *Parser) expect(target lexer.TokenType) bool {
	nextToken, _ := p.lexer.Peek()
	if nextToken == nil {
		return false
	}

	return target == nextToken.TokenType
}

// TODO: Consider handling errors?
func (p *Parser) expectAndConsume(target lexer.TokenType) bool {
	if !p.expect(target) {
		return false
	}

	_, err := p.consume()
	if err != nil {
		return false
	}
	return true
}
