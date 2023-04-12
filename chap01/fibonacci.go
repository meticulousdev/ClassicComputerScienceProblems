package main

import (
	"fmt"
	"time"
)

// (2) Simple Fibonacci
func fib2(n int) int {
	if n < 2 {
		return n
	}

	return fib2(n-2) + fib2(n-1)
}

// (3) Memoization
var memo = make(map[int]int)

func init_fib3() {
	memo[0] = 0
	memo[1] = 1
}

func fib3(n int) int {
	if len(memo) <= n {
		memo[n] = fib3(n-1) + fib3(n-2)
	}
	return memo[n]
}

func main() {
	n := 40

	fmt.Printf("***** Fibonacci : %d *****\n", n)

	fmt.Print("Simple Fibonacci : ")
	start := time.Now()
	fmt.Println(fib2(n))
	elapsed := time.Since(start)
	fmt.Printf("Simple Fibonacci took: %s\n", elapsed)

	fmt.Print("Memoization      : ")
	start = time.Now()
	init_fib3()
	fmt.Println(fib3(n))
	elapsed = time.Since(start)
	fmt.Printf("Simple Fibonacci took: %s\n", elapsed)
}
