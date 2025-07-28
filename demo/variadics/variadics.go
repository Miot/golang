package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	all := append(a, b...)
	fmt.Println(sum(all...))
}

