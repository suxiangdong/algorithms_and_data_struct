package code

import (
	"testing"
)

func TestDisjoint(t *testing.T) {
	datas := []struct{
		in [3][]int
		out bool
	}{
		{
			in: [3][]int{
				{}, {}, {},
			},
			out: true,
		},
		{
			in: [3][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			out: true,
		},
		{
			in: [3][]int{
				{15, 22, 30},
				{34, 65, 46},
				{17, 48, 99},
			},
			out: true,
		},
		{
			in: [3][]int{
				{125, 98, 340, 34},
				{34, 645, 446},
				{17, 248, 99, 34, 189},
			},
			out: false,
		},
	}

	for _, v := range datas {
		if ok := Disjoint1(v.in[0], v.in[1], v.in[2]); ok != v.out {
			t.Error("disjoint1 test fail")
		}
	}

	for _, v := range datas {
		if ok := Disjoint2(v.in[0], v.in[1], v.in[2]); ok != v.out {
			t.Error("disjoint2 test fail")
		}
	}
}
