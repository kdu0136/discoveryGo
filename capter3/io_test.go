package capter3

import (
	"fmt"
	"os"
	"strings"
)

func ExampleWriteTo() {
	lines := []string {
		"hello",
		"hi",
	}
	if err := WriteTo(os.Stdout, lines); err != nil {
		fmt.Println(err)
	}
	// Output:
	// hello
	// hi
}

func ExampleReadFrom() {
	r := strings.NewReader("hello\nhi")
	var lines []string
	if err := ReadFrom(r, func(line string){
		lines = append(lines, line)
	}); err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
	// Output:
	// [hello hi]
}
