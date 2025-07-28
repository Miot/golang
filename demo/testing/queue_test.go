package queue

import "testing"

func TestQueue(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("Incorrect queue size")
		}
		if !q.Append(i) {
			t.Errorf("Failed to append item")
		}
	}
	if q.Append(4) {
		t.Errorf("Should not append item")
	}
}

func TestNext(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		q.Append(i)
	}
	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("Should have next item")
		}
		if item != i {
			t.Errorf("Incorrect next item")
		}
	}
	_, ok := q.Next()
	if ok {
		t.Errorf("Should not have next item")
	}
}
