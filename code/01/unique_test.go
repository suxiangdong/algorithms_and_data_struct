package code

import "testing"

func TestUnique(t *testing.T) {
	datas := []struct{
		in []int
		out bool
	}{
		{in: nil, out: true},
		{in: []int{2, 5, 9, 10, 7, 4, 14, 52, 5}, out: false},
		{in: []int{32, 15, 9, 10, 7, 4, 14, 52, 5}, out: true},
		{in: []int{52, 5, 9, 130, 7, 4, 14, 52, 5}, out: false},
		{in: []int{42, 85, 9, 10, 7, 4, 14, 52, 5}, out: true},
	}

	for _, v := range datas {
		if ok := Unique1(v.in); ok != v.out {
			t.Errorf("unique1 test fail")
		}
	}

	for _, v := range datas {
		if ok := Unique2(v.in); ok != v.out {
			t.Errorf("unique2 test fail")
		}
	}
}
