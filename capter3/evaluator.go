package capter3

import (
	"fmt"
	"strconv"
	"strings"
)

// String set type
type StrSet map[string]struct{}
// Return a new StrSet
func NewStrset(strs ...string) StrSet {
	m := StrSet{}
	for _, str := range strs {
		m[str] = struct{}{}
	}
	return m
}
// Map keyed by operator to set of higher precedence operators
type PrecMap map[string]StrSet

func NewEvaluator(opMap map[string]BinOp, prec PrecMap) func(expr string) int {
	return func(expr string) int {
		return Eval(opMap, prec, expr)
	}
}

// Calculator returns the evaluation result of the given expr.
// The expression can hav +,-,*,/,(,) operators and decimal integers.
// Operators and operands should be space delimited.
func Eval(opMap map[string]BinOp, prec PrecMap, expr string) int {
	var ops []string
	var nums []int
	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}
	reduce := func(nextOp string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]
			if _, higher := prec[nextOp][op]; nextOp != ")" && !higher {
				// 더 낮은 순위 연산자이므로 여기서 계산 종료
				return
			}
			ops = ops[:len(ops)-1]
			if op == "(" {
				// 괄호를 제거하였으므로 종료
				return
			}
			b, a := pop(), pop()
			if f := opMap[op]; f != nil {
				nums = append(nums, f(a,b))
			}
		}
	}

	for _, token := range strings.Split(expr, " ") {
		fmt.Println("token:", token)
		if token == "(" {
			ops = append(ops, token)
		} else if _, ok := prec[token]; ok {
			reduce(token)
			ops = append(ops, token)
		} else if token == ")" {
			// 닫는 괄호는 여는 괄호까지 계산하고 제거
			reduce(token)
		} else {
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}
	reduce(")") // 초기의 여는 괄호까지 모두 계산
	return nums[0]
}
