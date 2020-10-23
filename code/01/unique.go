package code

import "sort"

// O(n^2)
func Unique1(S []int) bool {
	n := len(S)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if S[i] == S[j] {
				return false
			}
		}
	}
	return true
}

// O(nlog(n))
func Unique2(S []int) bool {
	sort.Ints(S) // 排序，保证最差时间复杂度为O(n(log(n))
	n := len(S)
	for i := 0; i < n - 1; i++ {
		if S[i] == S[i+1] {
			return false
		}
	}
	return true
}
