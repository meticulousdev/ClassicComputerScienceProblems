package main

import "fmt"

func add01(x int, y int) int {
	return x + y
}

func add02(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

// naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	x := 3
	y := 4
	fmt.Println(add01(x, y))
	fmt.Println(add02(x, y))
	fmt.Println(swap(x, y))
	fmt.Println(split(21))
}
