package main

import (
	"fmt"
	"woojiahao.com/chris/internal/lexer"
)

func main() {
	l := lexer.New("12 34")
	fmt.Printf("lexer.Next step 1 returns %v", l.Next())
	fmt.Printf("lexer.Next step 2 returns %v", l.Next())
}
