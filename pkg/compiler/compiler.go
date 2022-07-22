package compiler

import (
	"fmt"
	"math"
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

func (c *Compiler) GenerateSyntax() func(float64) float64 {
	ast := c.parser.Parse()
	fmt.Println(ast.Print())
	fn := generate(ast)
	return fn
}

func generate(root parser.Node) func(float64) float64 {
	fn := func(x float64) float64 {
		var exp float64 = 0
		exp += recursiveDescent(root, x)
		return exp
	}

	return fn
}

var validKeywords = []string{"sin", "cos", "tan", "sec", "csc", "cot"}
var keywordFnMap = map[string]func(float64) float64{
	"sin": math.Sin,
	"cos": math.Cos,
	"tan": math.Tan,
	"sec": math.Acos,
	"csc": math.Asin,
	"cot": math.Atan,
}

func recursiveDescent(node parser.Node, variable float64) float64 {
	switch n := node.(type) {
	case parser.NumberNode:
		return float64(n)
	case parser.VariableNode:
		return variable
	case parser.PrefixNode:
		// For now only minus
		right := recursiveDescent(n.Right, variable)
		if n.PrefixToken.TokenType.Symbol == "-" {
			return -right
		}
	case parser.KeywordNode:
		// For now only pi
		if n == "pi" {
			return math.Pi
		}

	case parser.OperatorNode:
		left := recursiveDescent(n.Left, variable)
		right := recursiveDescent(n.Right, variable)
		op := dispatchOperator(n.Operator.Symbol)
		return op(left, right)
	case parser.FunctionNode:
		if !isValidKeyword(string(n.Keyword)) {
			panic("Invalid keyword")
		}

		if len(n.Arguments) != 1 {
			panic("Supported keywords only require 1 argument")
		}

		arg := recursiveDescent(n.Arguments[0], variable)

		fn := keywordFnMap[string(n.Keyword)]
		return fn(arg)

	default:
		panic("Invalid node to parse")
	}

	return -1
}

func isValidKeyword(keyword string) bool {
	for _, k := range validKeywords {
		if k == keyword {
			return true
		}
	}

	return false
}

func dispatchOperator(symbol string) func(float64, float64) float64 {
	var op func(float64, float64) float64
	switch symbol {
	case "+":
		op = add
	case "-":
		op = minus
	case "/":
		op = divide
	case "*":
		op = multiply
	case "^":
		op = exponent
	}

	return op
}

func add(a, b float64) float64 {
	return a + b
}

func minus(a, b float64) float64 {
	return a - b
}

func divide(a, b float64) float64 {
	return a / b
}

func multiply(a, b float64) float64 {
	return a * b
}

func exponent(a, b float64) float64 {
	return math.Pow(a, b)
}
