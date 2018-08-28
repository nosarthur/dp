package main

// n is the number of items, cap is the maximum capacity
type pair struct {
	n, cap int
}

var memoTopDown = make(map[pair]int)

// Return the item indices and total value using dynamic programming
func dpTopDown(n int, cap int) int {
	if cap == 0 || n == 0 {
		return 0
	}
	p := pair{n, cap}
	if v, ok := memoTopDown[p]; ok {
		return v
	}
	n--              // make it 0-based index
	if ws[n] > cap { // The n'th item cannot fit in
		return dpTopDown(n, cap)
	}
	toInclude := vs[n] + dpTopDown(n, cap-ws[n])
	notInclude := dpTopDown(n, cap)
	v := max(toInclude, notInclude)
	memoTopDown[p] = v
	return v
}

// n is the number of items, cap is the maximum capacity
func dpBottomUp(n int, cap int) int {
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, cap+1)
	}
	// first row of memo: one item problem
	for j := 1; j <= cap; j++ {
		if ws[0] > j { // first item is heavier than j
			memo[0][j] = 0
		} else {
			memo[0][j] = vs[0]
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= cap; j++ {
			if ws[i] > j { // not enough capacity for the last item
				memo[i][j] = memo[i-1][j]
			} else {
				toInclude := vs[i] + memo[i-1][j-ws[i]]
				notInclude := memo[i-1][j]
				memo[i][j] = max(toInclude, notInclude)
			}
		}
	}
	return memo[n-1][cap]
}

func dpBranchAndBound(n int, cap int) int {
	return 0
}
