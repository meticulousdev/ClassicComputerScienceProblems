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

func stack_init(pstack *Stack) {
	pstack.topIndex = -1
}

func stack_is_empty(pstack *Stack) bool {
	if pstack.topIndex == -1 {
		return true
	} else {
		return false
	}
}

func stack_push(pstack *Stack, data Data) {
	pstack.topIndex += 1
	pstack.stackArr[pstack.topIndex] = data
}

func stack_pop(pstack *Stack) Data {
	if stack_is_empty(pstack) {
		fmt.Println("Stack Memory Error!")
		return -1
	}

	rIdx := pstack.topIndex
	pstack.topIndex -= 1

	return pstack.stackArr[rIdx]
}

func stack_peek(pstack *Stack) Data {
	if stack_is_empty(pstack) {
		fmt.Println("Stack Memory Error!")
		return -1
	}

	return pstack.stackArr[pstack.topIndex]
}

func main() {
	var data Data = 1
	fmt.Printf("data type      : %T\n\n", data)

	var stack Stack
	stack_init(&stack)
	fmt.Printf("stack_init     : %d\n", stack.topIndex)
	fmt.Printf("stack_is_empty : %t\n\n", stack_is_empty(&stack))

	stack_push(&stack, 1)
	stack_push(&stack, 2)
	stack_push(&stack, 3)
	stack_push(&stack, 4)
	stack_push(&stack, 5)
	fmt.Printf("stack_is_empty : %t\n\n", stack_is_empty(&stack))

	fmt.Printf("stack_peek     : %d\n\n", stack_peek(&stack))

	fmt.Print("stack_pop      : ")
	for stack_is_empty(&stack) != true {
		fmt.Printf("%d ", stack_pop(&stack))
	}
	fmt.Println()

	fmt.Printf("stack_is_empty : %t\n\n", stack_is_empty(&stack))
}

// data type      : main.Data

// stack_init     : -1
// stack_is_empty : true

// stack_is_empty : false

// stack_peek     : 5

// stack_pop      : 5 4 3 2 1
// stack_is_empty : true
