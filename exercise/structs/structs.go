//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type rectangle struct {
	length int
	width int
}

func area(rectangle rectangle) int {
	return rectangle.length * rectangle.width
}

func perimeter(rectangle rectangle) int {
	return 2 * (rectangle.length + rectangle.width)
}

func main() {
	rect := rectangle{}
	rect.length = 10
	rect.width = 5
	fmt.Println("Area:", area(rect))
	fmt.Println("Perimeter:", perimeter(rect))
	rect.length *= 2
	rect.width *= 2
	fmt.Println("Area:", area(rect))
	fmt.Println("Perimeter:", perimeter(rect))
}

