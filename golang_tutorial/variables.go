package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i, j int = 1, 2
	fmt.Println(i, j)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(j))
	fmt.Println()

	c, python, java := true, false, "no!"
	fmt.Println(c, python, java)
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(python))
	fmt.Println(reflect.TypeOf(java))
	fmt.Println()

	// Inference Type
	k := 3
	var zero_value int
	const pi = 3.14
	// pi = 0 // cannot assign to pi (untyped float constant 3.14)
	fmt.Printf("Type: %T Value: %v\n", k, k)
	fmt.Println(reflect.TypeOf(k))
	fmt.Printf("Type: %T Value: %v\n", zero_value, zero_value)
	fmt.Printf("Type: %T Value: %v\n", pi, pi)
	fmt.Println()
}
