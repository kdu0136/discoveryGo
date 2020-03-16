package capter3

import "fmt"

func NewIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

type BinOp func(int, int) int
type BinSub func(int, int) int

func BinOpToBInSub(f BinOp) BinSub {
	var count int
	return func(a, b int) int {
		fmt.Println(f(a, b))
		count++
		return count
	}
}

type MultiSet map[string]int
type SetOp func(m MultiSet, val string)

func NewMultiset() MultiSet {
	return make(MultiSet)
}
func Insert(m MultiSet, val string) {
	m[val]++
}
func BindMap(f SetOp, m MultiSet) func(val string) {
	return func(val string) {
		f(m, val)
	}
}