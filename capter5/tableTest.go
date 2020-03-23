package capter5

import "testing"

func TestFib(t *testing.T) {
	cases := []struct {
		in, want int
	} {
		{0, 0},
		{5, 5},
		{6, 8},
	}
	for _, c := range cases {
		got := DPFunc(c.in, Fibonacci)
		if got != c.want {
			t.Errorf("Fib(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}

var DP []int
// Fibonacci function
func DPFunc(n int, f func(n int) int) int {
	DP = make([]int, n+1)
	f(n)
	return DP[n]
}
// Fibonacci function
func Fibonacci(n int) int {
	if n == 0 {
		return n
	}

	if DP[n] == 0 {
		if n == 1 {
			DP[n] = n
		} else {
			DP[n] = Fibonacci(n-1) + Fibonacci(n-2)
		}
	}

	return DP[n]
}
