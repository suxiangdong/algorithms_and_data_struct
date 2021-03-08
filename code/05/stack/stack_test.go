package stack

import (
	"testing"
)

func TestSliceStack(t *testing.T) {
	s := NewSliceStack()
	s1 := NewLinkedListStack()

	in := []interface{}{1, "a", `b`, [2]int{1, 2}}
	for _, i := range in {
		s.Push(i)
		s1.Push(i)
	}

	if s.Len() != len(in) || s1.Len() != len(in) {
		t.Fatal("length test failed")
	}

	length := len(in)
	for i := 0; i < length; i++ {
		if v, _ := s.Pop(); v != in[s.Len()] {
			t.Fatal("sliceStack pop test failed")
		}

		if v, _ := s1.Pop(); v != in[s1.Len()] {
			t.Fatal("linkedListStack pop test failed")
		}
	}

	if _, err := s.Pop(); err == nil {
		t.Fatal("sliceStack empty pop failed")
	}

	if _, err := s1.Pop(); err == nil {
		t.Fatal("linkedListStack empty pop failed")
	}
}