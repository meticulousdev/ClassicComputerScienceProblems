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

func init_stack(pstack *Stack) {
	pstack.topIndex = -1
}

func is_empty(pstack *Stack) bool {
	if pstack.topIndex == -1 {
		return true
	} else {
		return false
	}
}

func push(pstack *Stack, data Data) {
	pstack.topIndex += 1
	pstack.stackArr[pstack.topIndex] = data
}

func pop(pstack *Stack) Data {
	if is_empty(pstack) {
		fmt.Println("Stack Memory Error!")
		return -1
	}

	rIdx := pstack.topIndex
	pstack.topIndex -= 1

	return pstack.stackArr[rIdx]
}

func peek(pstack *Stack) Data {
	if is_empty(pstack) {
		fmt.Println("Stack Memory Error!")
		return -1
	}

	return pstack.stackArr[pstack.topIndex]
}

func main() {
	var data Data = 1
	fmt.Printf("data type      : %T\n\n", data)

	var stack Stack
	init_stack(&stack)
	fmt.Printf("init     : %d\n", stack.topIndex)
	fmt.Printf("is_empty : %t\n\n", is_empty(&stack))

	push(&stack, 1)
	push(&stack, 2)
	push(&stack, 3)
	push(&stack, 4)
	push(&stack, 5)
	fmt.Printf("is_empty : %t\n\n", is_empty(&stack))

	fmt.Printf("peek     : %d\n\n", peek(&stack))

	fmt.Print("pop      : ")
	for is_empty(&stack) != true {
		fmt.Printf("%d ", pop(&stack))
	}
	fmt.Println()

	fmt.Printf("is_empty : %t\n\n", is_empty(&stack))
}

// data type      : main.Data

// stack_init     : -1
// stack_is_empty : true

// stack_is_empty : false

// stack_peek     : 5

// stack_pop      : 5 4 3 2 1
// stack_is_empty : true
