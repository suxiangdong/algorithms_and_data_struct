package code

import (
	"testing"
)

func TestMaxWithNil(t *testing.T) {
	t.Parallel()
	_, err := Max()
	if err == nil {
		t.Error("input nil test fail")
	}
}

func TestMax(t *testing.T) {
	t.Parallel()
	for _, data := range []struct {
		in []int
		expected int
	}{
		{
			[]int{3, 4, 9, 20, 17, 29},
			29,
		},
		{
			[]int{1, 4, 87, 12, 90, 34},
			90,
		},
		{
			[]int{82, 24, 9, 20, 7, 43},
			82,
		},
		{
			[]int{53, 41, 91, 28, 14, 23},
			91,
		},
	} {
		max, err := Max(data.in...)
		if  err != nil {
			t.Errorf("input slice test fail")
		}

		if max != data.expected {
			t.Errorf("max and expected not equal")
		}
	}
}
