package capter4

import "fmt"

func ExampleNewEvaluator() {
	eval := NewEvaluator()

	fmt.Println(eval("5"))
	fmt.Println(eval("1 + 2"))
	fmt.Println(eval("1 - 2 - 4"))
	// Output:
	// 5 <nil>
	// 3 <nil>
	// -5 <nil>
}
