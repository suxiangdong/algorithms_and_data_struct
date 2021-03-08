package dqueue

import "testing"

func TestDQ(t *testing.T) {
	dq := New()

	if _, err := dq.FirstPop(); err == nil {
		t.Fatal("DQueue empty firstPop test failed")
	}
	if _, err := dq.LastPop(); err == nil {
		t.Fatal("DQueue empty lastPop test failed")
	}

	dq.FirstPush(1)
	if v, _ := dq.FirstPop(); v != 1 {
		t.Fatal("DQueue firstPop test failed")
	}

	dq.FirstPush(2)
	dq.FirstPush(1)

	if v, _ := dq.LastPop(); v != 2 {
		t.Fatal("DQueue lastPop test failed")
	}

	if dq.Len() != 1 {
		t.Fatal("DQueue len test failed")
	}

	dq.LastPush(3)
	if v, _ := dq.LastPop(); v != 3 {
		t.Fatal("DQueue lastPop test failed")
	}
}