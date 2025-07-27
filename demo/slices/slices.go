package main

import "fmt"

func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := 0; i < len(slice); i++ {
		ele := slice[i]
		fmt.Println(ele)
	}
	
}

func main() {
	route := []string{"Grocery", "Department Store", "Salon"}
	printSlice("Route 1", route)
	route = append(route, "Home")
	printSlice("Route 2", route)

	fmt.Println()
	fmt.Println(route[0], "visted")
	fmt.Println(route[1], "visted")

	route = route[2:]
	printSlice("Remaining locations", route)
}
