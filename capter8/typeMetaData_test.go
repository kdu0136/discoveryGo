package capter8

import (
	"errors"
	"fmt"
)

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

func ExampleFieldNames() {
	s := struct {
		id int
		Name string
		Age int
	}{}
	fmt.Println(FieldNames(s))
	// Output:
	// [id Name Age] <nil>
}

func ExampleAppendNilError() {
	f := func() {
		fmt.Println("called")
	}
	f2, err := AppendNilError(f, errors.New("test error"))
	fmt.Println("AppendNilError.err:", err)
	fmt.Println(f2.(func() error)())
	// Output:
	// AppendNilError.err: <nil>
	// called
	// test error
}
