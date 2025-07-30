package main

import "fmt"

const (
	Low    = iota
	Medium
	High
)

type PrriorityQueue[P comparable, V any] struct {
	items map[P][]V
	priorities []P
}

func NewPriorityQueue[P comparable, V any](priorities []P) PrriorityQueue[P, V] {
	return PrriorityQueue[P, V]{
		items: make(map[P][]V),
		priorities: priorities,
	}
}

func (pq *PrriorityQueue[P, V]) Add(priority P, value V) {
	pq.items[priority] = append(pq.items[priority], value)
}

func (pq *PrriorityQueue[P, V]) Next() (V, bool) {
	for _, priority := range pq.priorities {
		if len(pq.items[priority]) > 0 {
			next := pq.items[priority][0]
			pq.items[priority] = pq.items[priority][1:]
			return next, true
		}
	}
	var empty V
	return empty, false
}

func main() {
	pq := NewPriorityQueue[int, string]([]int{High, Medium, Low})
	pq.Add(Low, "Low 1")
	pq.Add(High, "High 1")
	fmt.Println(pq.Next())

	pq.Add(Medium, "Medium 1")
	pq.Add(High, "High 2")
	pq.Add(High, "High 3")
	fmt.Println(pq.Next())
	fmt.Println(pq.Next())
	fmt.Println(pq.Next())
	fmt.Println(pq.Next())
	fmt.Println(pq.Next())
}

