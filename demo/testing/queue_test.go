package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("Incorrect queue element count: %v, want %v", len.(q.items), i)
		}
	}
}
