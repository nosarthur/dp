package main

import (
	"fmt"
	"testing"
)

func BenchmarkFib(b *testing.B) {
	fibFuncs := []struct {
		name string
		f    func(int) int
	}{
		{"recursive", fibRecu},
		{"iterative", fibIter},
	}
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
