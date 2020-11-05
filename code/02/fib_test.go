package code

import (
	"testing"
)

func TestFib(t *testing.T) {
	t.Parallel()
	datas := []struct {
		in, out int
	}{
		{in: 0, out: 0},
		{in: 1, out: 1},
		{in: 10, out: 55},
		{in: 15, out: 610},
	}

	for _, v := range datas {
		if r := Fib1(v.in); r != v.out {
			t.Fatal("Fib1 test fail")
		}
	}
}