package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

type byValue struct {
	ws []int
	vs []int
}

func (b byValue) Len() int {
	return len(b.vs)
}

func (b byValue) Less(i, j int) bool {
	return b.vs[i] < b.vs[j]
}

func (b byValue) Swap(i, j int) {
	b.vs[i], b.vs[j] = b.vs[j], b.vs[i]
	b.ws[i], b.ws[j] = b.ws[j], b.ws[i]
}

// Return the item indices and total value using greedy method
// Put the most valuable item in the knapsack if possible.
func greedy(ws []int, vs []int, c int) (items []int, value int) {
	sort.Sort(sort.Reverse(byValue{ws: ws, vs: vs}))
	fmt.Println("ws", ws, "vs", vs)
	value = 0
	for i, v := range vs {
		if ws[i] <= c {
			c -= ws[i]
			items = append(items, ws[i])
			value += v
		}
	}
	return items, value

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// NewRandomCase generates a random test case
// n is the number of items, returns the knapsack capacity
func NewRandomCase(n int) int {
	// clear the old slices
	ws = nil
	vs = nil
	for i := 0; i < n; i++ {
		ws = append(ws, rand.Intn(100))
		vs = append(vs, rand.Intn(100))
	}
	capacity := rand.Intn(100)
	return capacity
}

// NewUserCase generates a test case from user input
func NewUserCase(scanner *bufio.Scanner) int {
	ws = nil
	vs = nil

	fmt.Println("Weights of the items (integers):")
	scanner.Scan()
	for _, s := range strings.Fields(scanner.Text()) {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("%s is not an integer number", s)
		}
		if n < 1 {
			log.Fatalf("weight cannot be negative: %s", s)
		}
		ws = append(ws, n)
	}

	fmt.Println("Values of the items (integers):")
	scanner.Scan()
	for _, s := range strings.Fields(scanner.Text()) {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("%s is not an integer number", s)
		}
		vs = append(vs, n)
	}
	if len(ws) != len(vs) {
		log.Fatal("Lengths of weights and values do not match.")
	}

	fmt.Println("Weight capacity of the knapsack (integer):")
	scanner.Scan()
	s := scanner.Text()
	capacity, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Input %s is not an integer number.", s)
	}
	if capacity < 1 {
		log.Fatalf("weight cannot be negative: %s", s)
	}
	return capacity
}
