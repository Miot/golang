package main

import (
	"fmt"
	"errors"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.values) {
		return 0, errors.New(fmt.Sprintf("no element at index %v", index))
	}
	return s.values[index], nil
}
	

func main() {
	stuff := Stuff{values: []int{1, 2, 3}}
	value, err := stuff.Get(6)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value is", value)
	}
}

