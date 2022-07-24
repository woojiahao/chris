package compiler

import (
	"github.com/woojiahao/chris/pkg/lexer"
	"github.com/woojiahao/chris/pkg/parser"
	"math"
)

var keywords = []string{"sin", "cos", "tan", "sec", "csc", "cot"}
var constants = []string{"pi"}

var keywordFnMap = map[string]func(float64) float64{
	"sin": math.Sin,
	"cos": math.Cos,
	"tan": math.Tan,
	"sec": math.Acos,
	"csc": math.Asin,
	"cot": math.Atan,
}

var constantValueMap = map[string]float64{
	"pi": math.Pi,
}

var operatorFnMap = map[string]func(float64, float64) float64{
	"+": func(a, b float64) float64 { return a + b },
	"-": func(a, b float64) float64 { return a - b },
	"*": func(a, b float64) float64 { return a * b },
	"/": func(a, b float64) float64 { return a / b },
	"^": func(a, b float64) float64 { return math.Pow(a, b) },
}

type compiler struct {
	l *lexer.Lexer
	p *parser.Parser
}

func New(exp string) *compiler {
	l := lexer.New(exp, keywords, constants)
	p := parser.New(l)
	return &compiler{l, p}
}

// GenerateFunction generates a function that receives the variable and returns the evaluated body of the equation
func (c *compiler) GenerateFunction() func(float64) float64 {
	ast, err := c.p.Parse()
	if err != nil {
		panic(err)
	}
	fn := func(x float64) float64 {
		return recursiveDescent(ast, x)
	}
	return fn
}

func recursiveDescent(node parser.Node, variable float64) float64 {
	switch n := node.(type) {
	case parser.NumberNode:
		return float64(n)
	case parser.VariableNode:
		return variable
	case parser.ConstantNode:
		constant := constantValueMap[string(n)]
		return constant
	case parser.PrefixNode:
		// For now only minus
		right := recursiveDescent(n.Right, variable)
		if n.PrefixToken.Symbol == "-" {
			return -right
		}

	case parser.OperatorNode:
		left := recursiveDescent(n.Left, variable)
		right := recursiveDescent(n.Right, variable)
		return operatorFnMap[n.Operator.Symbol](left, right)
	case parser.FunctionNode:
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
