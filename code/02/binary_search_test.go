package code

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 10, 19, 20, 88, 98}
	for _, v := range []struct{
		in int
		out bool
	}{
		{in: 4, out: true},
		{in: 50, out: false},
		{in: 41, out: false},
		{in: 88, out: true},
	} {
		if r := BinarySearch1(s, v.in, 0, len(s)-1); r != v.out {
			fmt.Println(r, v.out, v.in)
			t.Fatal("BinarySearch1 test fail")
		}
	}
}