package capter6

import "fmt"

func F() int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover in a")
		}
	}()
	a()
	return 100
}

func a() {
	b()
}

func b() {
	c()
}

func c() {
	panic("panic in c")
}
