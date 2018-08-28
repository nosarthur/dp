package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// The first element is a place holder. [nil, 1, 1, 2, ...]
var memo = map[int]int{0: 0, 1: 1, 2: 1}

func fibIter(n int) int {
	a, b := 1, 1
	for ; n > 2; n-- {
		a, b = b, a+b
	}
	return b
}

func fibRecu(n int) int {
	if fn, ok := memo[n]; ok {
		return fn
	}
	memo[n] = fibRecu(n-1) + fibRecu(n-2)
	return memo[n]
}

func main() {
	fmt.Println("Get the n'th Fibonacci number. Input n:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Input %s is not a number.", s)
	}
	fmt.Printf("The %s'th Fibonacci number is\n", s)
	fmt.Printf("iterative version: %v\n", fibIter(n))
	fmt.Printf("recursive version: %v\n", fibRecu(n))
}
