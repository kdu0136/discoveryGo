package capter3

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	zeroUnicode = 48
	nineUnicode = 57
)

// Calculator returns the evaluation result of the given expr.
// The expression can hav +,-,*,/,(,) operators and decimal integers.
// Operators and operands should be space delimited.
func Calculator(expr string) int {
	expr = strings.Replace(strings.TrimSpace(expr), " ", "", -1)
	var ops []string
	var nums []int
	fmt.Println(expr)
	fmt.Println(operators)

	for i, r := range expr {
		if r >= zeroUnicode && r <= nineUnicode { // number
			n, _ := strconv.Atoi(string(r))
			nums = append(nums, n)
		} else { // others
			if vo := isValidOperator(r); !vo {
				panic("include invalid operator in expression")
			} else {
				ops = append(ops, string(r))
				if i == 0 {
					nums = append(nums, 0)
				}
			}
		}
	}
	fmt.Println(ops)
	fmt.Println(nums)
	//// nums stack 에 있는 마지막 number 가지고 오고 pop
	//pop := func() (last int) {
	//	last = nums[len(nums)-1]
	//	nums = nums[:len(nums)-1]
	//	return
	//}
	//reduce := func(higher string) {
	//	for len(ops) > 0 {
	//		op := ops[len(ops)-1]
	//		if strings.Index(higher, op) < 0 {
	//			return
	//		}
	//	}
	//}
	return 0
}

var operators = "+-*/()"
func isValidOperator(o int32) bool {
	for _, r := range operators {
		if o == r {
			return true
		}
	}
	return false
}
