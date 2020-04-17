package capter8

import "fmt"

func ExampleNewMap() {
	m := NewMap("", "").(map[string]string)
	m["key"] = "value"
	fmt.Println(m)

	m2 := NewMap(0, "").(map[int]string)
	m2[1] = "value"
	fmt.Println(m2)
	// Output:
	// map[key:value]
	// map[1:value]
}
