package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// use global variables so we don't need to pass them around
var ws = []int{}
var vs = []int{}

func main() {
	fmt.Println("Generate test case automatically? (Y/n)")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var capacity int
	if scanner.Text() != "n" {
		rand.Seed(time.Now().UnixNano())
		capacity = NewRandomCase(8)
	} else {
		capacity = NewUserCase(scanner)
	}
	fmt.Printf("weights: %v\n", ws)
	fmt.Printf("values: %v\n", vs)
	fmt.Printf("capacity: %v\n", capacity)
	bu := dpBottomUp(len(vs), capacity)
	td := dpTopDown(len(vs), capacity)
	fmt.Println("bottom up: ", bu)
	fmt.Println("top down: ", td)
}
