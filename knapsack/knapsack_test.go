package main

import (
	"fmt"
	"testing"
)

var dpFuncs = []struct {
	name string
	f    func(int, int) int
}{
	{"bottom up", dpBottomUp},
	{"top down", dpTopDown},
}

func TestKnapsack(t *testing.T) {
	vs = []int{12, 10, 20, 15}
	ws = []int{2, 1, 3, 2}
	cap := 5
	for _, dp := range dpFuncs {
		t.Run(fmt.Sprintf("%s", dp.name), func(t *testing.T) {
			v := dp.f(len(vs), cap)
			if v != 37 {
				t.Errorf("Expected value 37, got %d", v)
			}
		})
	}
}

func BenchmarkKnapsack(b *testing.B) {
	for _, dp := range dpFuncs {
		for k := 10; k < 1001; k *= 10 {
			b.Run(fmt.Sprintf("%s size %d", dp.name, k), func(b *testing.B) {
				b.StopTimer()
				capacity := NewRandomCase(k)
				for n := 0; n < b.N; n++ {
					// reset memo
					memoTopDown = nil
					memoTopDown = make(map[pair]int)
					b.StartTimer()
					dp.f(k, capacity)
				}
			})
		}
	}
}
