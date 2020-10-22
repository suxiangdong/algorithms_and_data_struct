package code

import (
	"errors"
)

func Max(s ...int) (int, error) {
	if s == nil {
		return 0, errors.New("length must greater than 0")
	}

	max := s[0]

	for _, v := range s {
		if v > max {
			max = v
		}
	}

	return max, nil
}