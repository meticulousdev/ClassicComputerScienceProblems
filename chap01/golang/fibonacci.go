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

func init_memo() {
	memo[0] = 0
	memo[1] = 1
}

func fib3(n int) int {
	if len(memo) <= n {
		memo[n] = fib3(n-1) + fib3(n-2)
	}
	return memo[n]
}

// (4) Memoization decorator
// None

// (5) Simple for loop
func fib5(n int) int {
	if n == 0 {
		return n
	}

	v_last := 0
	v_next := 1

	for i := 1; i < n; i++ {
		v_last, v_next = v_next, v_last+v_next
	}

	return v_next
}

// (6) Generator and Fibonacci
// fib6
func Fibonacci_Generator() func() int {
	v_last := 0
	v_next := 1
	cnt := 0

	return func() int {
		if cnt == 0 {
			cnt++
			return 0
		} else if cnt == 1 {
			cnt++
			return 1
		} else {
			v_last, v_next = v_next, v_last+v_next
		}
		return v_next
	}
}

func main() {
	n := 40

	fmt.Printf("***** Fibonacci            : %d *****\n\n", n)

	fmt.Print("Simple Fibonacci           : ")
	start := time.Now()
	fmt.Println(fib2(n))
	elapsed := time.Since(start)
	fmt.Printf("Simple Fibonacci took      : %s\n\n", elapsed)

	fmt.Print("Memoization                : ")
	start = time.Now()
	init_memo()
	fmt.Println(fib3(n))
	elapsed = time.Since(start)
	fmt.Printf("Memoization took           : %s\n\n", elapsed)

	fmt.Print("Memoization decorator      : ")
	fmt.Println("None")
	fmt.Printf("Memoization decorator took : %s\n\n", "None")

	fmt.Print("Simple for loop            : ")
	start = time.Now()
	fmt.Println(fib5(n))
	elapsed = time.Since(start)
	fmt.Printf("Simple for loop took       : %s\n\n", elapsed)

	fmt.Print("Generator and Fibonacci     : ")
	start = time.Now()
	fib6 := Fibonacci_Generator()
	for i := 0; i <= n; i++ {
		fmt.Println(fib6())
	}
	elapsed = time.Since(start)
	fmt.Printf("Generator and Fibonacci    : %s\n\n", elapsed)

	// go fmt .\fibonacci.go
	// go build .\fibonacci.go
	// .\fibonacci.exe
	//
	// ***** Fibonacci            : 40 *****

	// Simple Fibonacci           : 102334155
	// Simple Fibonacci took      : 422.915709ms

	// Memoization                : 102334155
	// Memoization took           : 98.625µs

	// Memoization decorator      : None
	// Memoization decorator took : None

	// Simple for loop            : 102334155
	// Simple for loop took       : 917ns

	// Generator and Fibonacci    : 0 ... 102334155
	// Generator and Fibonacci    : 36.125µs
}
