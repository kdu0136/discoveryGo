package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(src)
	src = append(src[:4], src[3:]...)
	src[4] = -1
	fmt.Println(src)
}
