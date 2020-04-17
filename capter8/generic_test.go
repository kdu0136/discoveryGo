package capter8

import "fmt"

func ExampleAssertEqual() {
	var a float64 = 3
	var b float64 = 3
	fmt.Println(AssertEqual(a, b))
	// Output:
	// true
}
