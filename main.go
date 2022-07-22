package main

import (
	"fmt"
	"math"
	"reflect"
	"woojiahao.com/chris/pkg/compiler"
)

func test(exp string, x float64) {
	fmt.Printf("Compiling %s into a Go function\n", exp)
	c := compiler.New(exp)
	f := c.GenerateSyntax()
	fmt.Printf("Compiled Go function is %v\n", reflect.TypeOf(f))
	fmt.Printf("Evaluating compiled function with x = %f\n", x)
	v := f(x)
	fmt.Printf("f(%f) = %f", x, v)
}

func main() {
	//p := parser.New("a = (1 + 2) * (4 - 3) / 7")
	//fmt.Println(p.parseExpression(0).Print())

	//expression := "y= sin(x,  1 + 2 * 3 ^ (8+ 9))^2 + pi"
	//l := lexer.New(expression)
	//p := parser.New(l)
	//fmt.Printf("Parsing expression: %s\n", expression)
	//fmt.Println(p.Parse().Print())

	//test("-(x + 1) * -2 ^ 7", 2)
	test("x * (3/4) + sin(x) ^ 2 - 0.5", math.Pi/3)
}
