package queue

import "testing"

func TestQueue(t *testing.T) {
	in := []elem{1, 2, 3, 4, 5, "a"}
	queue := New()
	for _, i := range in {
		queue.Push(i)
	}

	if len(in) != queue.Len() {
		t.Fatal("queue len test failed")
	}

	for i := 0; i < len(in); i++ {
		if v, _ := queue.Pop(); v != in[i] {
			t.Fatal("queue pop test failed")
		}
	}

	if _, err := queue.Pop(); err == nil {
		t.Fatal("empty queue test failed")
	}
}