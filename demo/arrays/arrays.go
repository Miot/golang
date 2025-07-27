package main

import "fmt"

type Room struct {
	name string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "has been cleaned")
		} else {
			fmt.Println(room.name, "has not been cleaned")
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Reception"},
		{name: "Ops"},
	}

	checkCleanliness(rooms)
	rooms[2].cleaned = true
	rooms[3].cleaned = true
	checkCleanliness(rooms)
}

