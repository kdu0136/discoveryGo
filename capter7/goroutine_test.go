package capter7

import "fmt"

func ExampleMin() {
	fmt.Println(Min([]int {
		15,68,32,11,77,
		23,55,33,25,75,
		24,37,36,
	}))
	// Output:
	// 11
}

func ExampleParallelMin() {
	fmt.Println(ParallelMin([]int {
		15,68,32,11,77,
		23,55,33,25,75,
		24,37,36,
	}, 4))
	// Output:
	// 11
}