package main

import (
	"fmt"
	"time"
)

func calculate_py(n_term int) float64 {
	numerator := 4.0
	denominator := 1.0
	operation := 1.0
	v_pi := 0.0

	for i := 1; i < n_term; i++ {
		v_pi += operation * (numerator / denominator)
		denominator += 2.0
		operation *= -1.0
	}

	return v_pi
}

func main() {
	start := time.Now()
	fmt.Print("PI                  : ")
	fmt.Println(calculate_py(100000000))
	elapsed := time.Since(start)
	fmt.Printf("PI calculation took : %s\n\n", elapsed)
}
