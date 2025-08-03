package main

import "fmt"

func add(lhs, rhs int) int {
	return lhs + rhs
}

func compute(lhs, rhs int, op func(lhs, rhs int) int) int {
	fmt.Printf("Computing %d & %d\n", lhs, rhs)
	return op(lhs, rhs)
}

func main() {
	fmt.Println("2 + 2 =", compute(2, 2, add))
	fmt.Println("2 - 2 =", compute(2, 2, func(lhs, rhs int) int {
		return lhs - rhs
	}))

	mul := func(lhs, rhs int) int {
		fmt.Printf("multiplying %d & %d\n", lhs, rhs)
		return lhs * rhs
	}
	fmt.Println("2 * 2 =", compute(2, 2, mul))
}
