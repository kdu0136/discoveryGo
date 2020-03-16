package capter3

import (
	"fmt"
	"strings"
)

func ExampleNewIntGenerator() {
	gen := NewIntGenerator()
	fmt.Println(gen())
	fmt.Println(gen())
	// Output:
	// 1
	// 2
}

func ExampleBinOpToBInSub() {
	sub := BinOpToBInSub(func(a, b int) int {
		return a + b
	})
	sub(5, 7)
	sub(5, 7)
	count := sub(5, 7)
	fmt.Println("count:", count)
	// Output:
	// 12
	// 12
	// 12
	// count: 3
}

func Example_f() {
	r := strings.NewReader("hello\nhi")
	m := NewMultiset()
	if err := ReadFrom(r, BindMap(Insert, m)); err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
	// Output:
	// map[hello:1 hi:1]
}