//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

type LineCallback func(line string)

func lineIterator(lines []string, callback LineCallback) {
	for i := 0; i < len(lines); i++ {
		callback(lines[i])
	}
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	letters := 0
	numbers := 0
	punctuations := 0
	spaces := 0

	lineIterator(lines, func(line string) {
		for _, r := range line {
			if unicode.IsLetter(r) {
				letters++
			} else if unicode.IsNumber(r) {
				numbers++
			} else if unicode.IsSpace(r) {
				spaces++
			} else if unicode.IsPunct(r) {
				punctuations++
			}
		}
	})
	fmt.Println("Letters:", letters)
	fmt.Println("Numbers:", numbers)
	fmt.Println("Punctuations:", punctuations)
	fmt.Println("Spaces:", spaces)
}
