//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type product struct {
	price int
	name string
}

func printStats(list [4]product) {
	var cost, totalItems int

	for i := 0; i < len(list); i++ {
		item := list[i]
		if item.name != "" {
			cost += item.price
			totalItems++
		}
	}

	fmt.Println("Last item:", list[totalItems-1])
	fmt.Println("Total items:", totalItems)
	fmt.Println("Total cost:", cost)
}

func main() {
	var shoppingList = [4]product{
		{price: 1, name: "Milk"},
		{price: 2, name: "Eggs"},
		{price: 3, name: "Cheese"},
	}

	printStats(shoppingList)
	shoppingList[3] = product{price: 4, name: "Bread"}
	printStats(shoppingList)
}

