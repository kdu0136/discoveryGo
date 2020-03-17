package capter3

import (
	"fmt"
	"math"
)

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

// Func y = f(x). 실수값을 하나 받아서 실수값 하나를 돌려줌
type Func func(float64) float64
// Transform Func를 받아서 Func를 돌려주는 함수. 함수 하나를 다른 함수로 변환하는 함수의 형태
type Transform func(Func) Func

const (
	tolerance = 0.00001
	dx        = 0.0001
)

// Square 제곱을 구하는 함수
func Square(x float64) float64 {
	fmt.Println("Square")
	return x * x
}

// FixedPoint 어떤 함수 f를 계속해서 적용했을 때, 어떤 값으로 수렴하는 경우 그 수렴 값을 찾는 함수
func FixedPoint(f Func, firstGuess float64) float64 {
	fmt.Println("FixedPoint")
	// 두 수가 서로 tolerance 이하로 가까워졌을 경우 true
	closeEnough := func(v1, v2 float64) bool {
		fmt.Println("closeEnough:", v1, v2)
		return math.Abs(v1-v2) < tolerance
	}
	// guess에 반복적으로 함수 f를 적용시키다가 그 변화가 충분히 작으면 수렴된 것으로 보고 그 값을 반환하고 종료
	var try Func
	try = func(guess float64) float64 {
		fmt.Println("guess:", guess)
		next := f(guess)
		if closeEnough(guess, next) {
			return next
		} else {
			return try(next)
		}
	}
	return try(firstGuess)
}

// FixedPointOfTransform FiexedPoint의 아이디어를 함수로 변환하여 적용시켜 수렴 값을 찾는 함수.
// transform에 어떤 함수를 수렴하는 성질을 갖는 함수로 변환시키는 함수를 넘겨주면 동작
func FixedPointOfTransform(g Func, transform Transform, guess float64) float64 {
	fmt.Println("FixedPointOfTransform")
	return FixedPoint(transform(g), guess)
}

// Deriv 어떤 함수를 받아서 다른 함수로 돌려주는 함수
// dx 는 매우 작은 숫자인데, 각 점에서 기울기를 구하고 있으므로, g 함수를 받아서 도함수(미분된 함수)로 돌려주는 함수
func Deriv(g Func) Func {
	return func(x float64) float64 {
		fmt.Println("Deriv")
		return (g(x+dx) - g(x)) / dx
	}
}

// NewtonTranform f(x)=0이 되는 x를 찾는 방법 중 하나.
// 함수 g를 받아서 반복적으로 적용하면 g(x)=0에 점점 다가가는 x 값을 반환하는 함수를 돌려주는 함수
func NewtonTransform(g Func) Func {
	return func(x float64) float64 {
		fmt.Println("NewtonTransform")
		return x - (g(x) / Deriv(g)(x))
	}
}

// Sqrt x의 제곱근을 구하는 함수
func Sqrt(x float64) float64 {
	fmt.Println("Sqrt")
	return FixedPointOfTransform(func(y float64) float64 {
		return Square(y) - x
	}, NewtonTransform, 1.0)
}
