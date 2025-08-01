package main

import (
	"fmt";
	"bufio";
	"strings";
	"strconv";
	"io";
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	sum := 0
	for {
		input, inputErr := r.ReadString(' ')
		n := strings.TrimSpace(input)
		if n == "" {
			continue
		}
		num, convErr := strconv.Atoi(n)
		if convErr != nil {
			fmt.Println("Conversion error:", convErr)
		} else {
			sum += num
		}

		if inputErr == io.EOF {
			break
		}
		if inputErr != nil {
			fmt.Println("Error reading Stdin:", inputErr)
		}
	}

	fmt.Printf("Sum: %v\n", sum)
}

