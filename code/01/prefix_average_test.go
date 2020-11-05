package code

import (
	"testing"
)

func TestPrefixAverage1(t *testing.T) {
	t.Parallel()
	datas := []struct{
		in []int
		out []float32
	}{
		{
			in: []int{},
			out: []float32{},
		},
		{
			in: []int{1, 2, 3, 4},
			out: []float32{1, 1.5, 2, 2.5},
		},
		{
			in: []int{4, 8, 3, 5},
			out: []float32{4, 6, 5, 5},
		},
	}

	for _, v := range datas {
		A := PrefixAverage1(v.in)
		for i, d := range v.out {
			if d != A[i] {
				t.Errorf("prefix_average1 test fail")
			}
		}
	}

	for _, v := range datas {
		A := PrefixAverage2(v.in)
		for i, d := range v.out {
			if d != A[i] {
				t.Errorf("prefix_average2 test fail")
			}
		}
	}
}
