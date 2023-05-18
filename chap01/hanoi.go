package main

import "fmt"

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

func hanoi(begin *Stack, end *Stack, temp *Stack, n int) {
	if n == 1 {
		push(end, pop(begin))
	} else {
		hanoi(begin, temp, end, n-1)
		hanoi(begin, end, temp, 1)
		hanoi(temp, end, begin, n-1)
	}
}

func main() {
	fmt.Println("Hanoi Tower")
	num_discs := 3
	var tower_a Stack
	var tower_b Stack
	var tower_c Stack

	for i := 1; i <= num_discs; i++ {
		push(&tower_a, Data(i))
	}

	hanoi(&tower_a, &tower_c, &tower_b, num_discs)

	fmt.Print("Tower A: ")
	for is_empty(&tower_a) != true {
		fmt.Printf("%d ", pop(&tower_a))
	}
	fmt.Println()

	fmt.Print("Tower B: ")
	for is_empty(&tower_b) != true {
		fmt.Printf("%d ", pop(&tower_b))
	}
	fmt.Println()

	fmt.Print("Tower C: ")
	for is_empty(&tower_c) != true {
		fmt.Printf("%d ", pop(&tower_c))
	}
	fmt.Println()
}
