package capter6

import "fmt"

func F() int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover in f")
		}
	}()
	var f = a() + 10
	fmt.Println("f",f)
	return f
}

func a() int{
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover in a")
		}
	}()
	var a = b() + 10
	fmt.Println("a :",a)
	return a
}

func b() int{
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("recover in b")
	//	}
	//}()
	var b = c() + 10
	fmt.Println("b :",b)
	return b
}

func c() int{
	panic("panic in c")
	return 10
}
