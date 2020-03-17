package capter3

import "fmt"

func ExampleNewEvaluator() {
	eval := NewEvaluator(map[string]BinOp{
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
	}, PrecMap{
		"**":  NewStrset(),
		"*":   NewStrset("**", "*", "/", "mod"),
		"/":   NewStrset("**", "*", "/", "mod"),
		"mod": NewStrset("**", "*", "/", "mod"),
		"+":   NewStrset("**", "*", "/", "mod", "+", "-"),
		"-":   NewStrset("**", "*", "/", "mod", "+", "-"),
	})

	fmt.Println(eval("5"))
	fmt.Println(eval("1 + 2"))
	fmt.Println(eval("1 - 2 - 4"))
	// Output:
	// 5
	// 3
	// -5
}
