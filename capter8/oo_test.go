package capter8

import "fmt"

func ExampleTotalArea() {
	fmt.Println(TotalArea([]Shape{
		Square{3},
		Rectangle{4, 5},
		Triangle{6, 7},
	}))
	// Output:
	// 50
}
