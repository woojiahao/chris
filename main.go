package main

import (
	"fmt"
	"woojiahao.com/chris/internal/parser"
)

func main() {
	//p := parser.New("a = (1 + 2) * (4 - 3) / 7")
	//fmt.Println(p.ParseExpression(0).Print())

	p := parser.New("y = sin(x,  1 + 2 * 3 ^ (8+ 9))^2 + pi")
	fmt.Println(p.ParseExpression(0).Print())
}
