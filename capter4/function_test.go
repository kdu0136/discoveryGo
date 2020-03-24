package capter4

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
	r := strings.NewReader("hello\nhi\nhi")
	m := NewMultiset()
	if err := ReadFrom(r, BindMap(Insert, m)); err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
	// Output:
	// {hello hi hi }
}

func ExampleVertexID_print() {
	i := VertexID(100)
	fmt.Println(i)
	// Output:
	// VertexID(100)
}

func ExampleMultiSet() {
	m := NewMultiset()
	m.Insert("a")
	m.Insert("b")
	m.Insert("b")
	m.Insert("c")
	m.Erase("a")
	m.Erase("b")
	fmt.Println(m.Count("a"))
	fmt.Println(m)
	// Output:
	// 0
	// {b c }
}
