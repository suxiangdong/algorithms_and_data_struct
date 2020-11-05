package code

import "testing"

func TestFactorial(t *testing.T) {
	t.Parallel()
	datas := []struct {
		in, out int
	}{
		{in: 5, out: 120},
		{in: 10, out: 3628800},
		{in: 0, out: 1},
	}

	for _, v := range datas {
		if r := Factorial1(v.in); r != v.out {
			t.Fatal("Factorial1 test fail")
		}
	}
}
