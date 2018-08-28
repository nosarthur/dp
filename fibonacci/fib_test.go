package main

import (
	"fmt"
	"testing"
)

var fibFuncs = []struct {
	name string
	f    func(int) int
}{
	{"recursive", fibRecu},
	{"iterative", fibIter},
}

func TestFib(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}
	for _, tc := range testCases {
		for _, fibFunc := range fibFuncs {
			t.Run(fmt.Sprintf("%d for %s", tc.n, fibFunc.name), func(t *testing.T) {
				got := fibFunc.f(tc.n)
				if got != tc.expected {
					t.Errorf("%s with input %d: expect %d, got %d",
						fibFunc.name, tc.n, tc.expected, got)
				}
			})
		}
	}
}

func BenchmarkFib(b *testing.B) {
	for _, fibFunc := range fibFuncs {
		// calculate k'th Fibonacci number
		for k := 10; k < 1001; k *= 10 {
			b.Run(fmt.Sprintf("%s Fib %v", fibFunc.name, k), func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					//					b.StopTimer()
					// reset the memo
					memo = map[int]int{0: 0, 1: 1, 2: 1}
					//					b.StartTimer()
					fibFunc.f(k)
				}
			})
		}
	}
}
