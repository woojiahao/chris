package parser

import "woojiahao.com/chris/internal/lexer"

var prefixParselets = map[lexer.TokenType]PrefixParselet{
	lexer.Variable:        VariableParselet{},
	lexer.Number:          NumberParselet{},
	lexer.Minus:           PrefixOperatorParselet{},
	lexer.LeftParenthesis: GroupParselet{},
	lexer.Keyword:         KeywordParselet{},
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

func getPrefixParselet(tokenType lexer.TokenType) PrefixParselet {
	if prefixParselet, ok := prefixParselets[tokenType]; !ok {
		panic("Invalid prefix token")
	} else {
		return prefixParselet
	}
}

func getInfixParselet(tokenType lexer.TokenType) InfixParselet {
	if infixParselet, ok := infixParselets[tokenType]; !ok {
		panic("Invalid prefix token")
	} else {
		return infixParselet
	}
}

type Parser struct {
	lexer *lexer.Lexer
}

func New(expression lexer.Expression) *Parser {
	return &Parser{lexer.New(expression)}
}

func (p *Parser) ParseExpression(precedence int) Node {
	token := p.consume()

	// Begin parsing the body to the right
	left := getPrefixParselet(token.TokenType).Parse(p, token)

	for precedence < p.nextPrecedence() {
		token = p.consume()

		left = getInfixParselet(token.TokenType).Parse(p, left, token)
	}

	return left
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

func (p *Parser) consume() *lexer.Token {
	token := p.lexer.Next()
	if token == nil {
		panic("Unable to parse expression")
	}

	return token
}

// expect checks the next token inline and returns whether it is the same as the target token type
func (p *Parser) expect(target lexer.TokenType) bool {
	nextToken, _ := p.lexer.Peek()
	if nextToken == nil {
		return false
	}

	return target == nextToken.TokenType
}

func (p *Parser) expectAndConsume(target lexer.TokenType) bool {
	if !p.expect(target) {
		return false
	}

	p.consume()
	return true
}
