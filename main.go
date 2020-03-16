package main

import (
	"fmt"
	"github.com/discoveryGo/capter3"
	"strings"
)

func main() {
	r := strings.NewReader("hello\nhi")
	m := NewMultiset()
	if err := capter3.ReadFrom(r, BindMap(Insert, m)); err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}