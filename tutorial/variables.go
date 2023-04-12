package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(j))
	fmt.Println(reflect.TypeOf(k))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(python))
	fmt.Println(reflect.TypeOf(java))
}
