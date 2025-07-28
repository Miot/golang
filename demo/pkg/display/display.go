package display

import "fmt"

func Display(msg string) {
	fmt.Println(msg)
}

// private
func hello(msg string) {
	fmt.Println("hello", msg)
}
