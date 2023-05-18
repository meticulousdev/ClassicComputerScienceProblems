package main

import (
	"fmt"
)

const STACK_LEN int = 100

type Data int

type ArrayStack struct {
	stackArr [STACK_LEN]Data
	topIndex int
}

type Stack ArrayStack

func stack_init(stack_data Stack) {
	stack_data.topIndex = -1
}

func main() {
	var type_data Data = 1
	fmt.Printf("type_data: %T\n", type_data)
}
