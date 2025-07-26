package main

import "fmt"

func main() {
	var myName = "Mio"
	fmt.Println("my name is",myName)

	var name string = "name2"
	fmt.Println("name =", name)

	userName := "admin"
	fmt.Println("username =", userName)

	var sum int
	fmt.Println("sum =", sum)

	part1, part := 1, 5
	fmt.Println("part1 =", part1, "part =", part)

	part2, part := 2, 10
	fmt.Println("part2 =", part2, "part =", part)
	
	sum = part1 + part2
	fmt.Println("sum =", sum)

	var (
		lessonName = "Variables"
		lessontype = "Demo"
	)
	fmt.Println("lessonName =", lessonName)
	fmt.Println("lessontype =", lessontype)

	word1, word2, _ := "a","b","c"
	fmt.Println(word1, word2)
}

