package main

import (
	"fmt"
	"github.com/discoveryGo/capter3"
)

func main() {
	eval := capter3.NewEvaluator(map[string]capter3.BinOp{
		"**": func(a, b int) int {
			fmt.Println("**")
			if a == 1 {
				return 1
			}
			if b < 0 {
				return 0
			}
			r := 1
			for i := 0; i < b; i++ {
				r *= a
			}
			return r
		},
		"*": func(a, b int) int {
			fmt.Println("*")
			return a * b
		},
		"/": func(a, b int) int {
			fmt.Println("/")
			return a / b
		},
		"mod": func(a, b int) int {
			fmt.Println("%")
			return a % b
		},
		"+": func(a, b int) int {
			fmt.Println("+")
			return a + b
		},
		"-": func(a, b int) int {
			fmt.Println("-")
			return a - b
		},
	}, capter3.PrecMap{
		"**":  capter3.NewStrset(),
		"*":   capter3.NewStrset("**", "*", "/", "mod"),
		"/":   capter3.NewStrset("**", "*", "/", "mod"),
		"mod": capter3.NewStrset("**", "*", "/", "mod"),
		"+":   capter3.NewStrset("**", "*", "/", "mod", "+", "-"),
		"-":   capter3.NewStrset("**", "*", "/", "mod", "+", "-"),
	})

	expr := "3 * ( ( 3 + 1 ) * 3 ) / ( -2 ) ** 3"
	fmt.Println(expr)
	result := eval(expr)
	fmt.Println(result)
}