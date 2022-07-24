package main

import (
	"fmt"
	"github.com/woojiahao/chris/example/compiler"
)

func main() {
	c := compiler.New("x(1 + sin(pi/4))")      // Generates an equation x * (1 + sin((pi / 4))) where x is my variable
	fn := c.GenerateFunction()                 // fn is fn x := x * (1 + sin((pi / 4)))
	fmt.Printf("fn on x = 2 is %f\n", fn(2))   // Should return 3.41421
	fmt.Printf("fn on x = 16 is %f\n", fn(16)) // Should return 27.3137
}
