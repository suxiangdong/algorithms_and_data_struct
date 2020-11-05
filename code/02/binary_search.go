package code

// 二分查找递归实现
func BinarySearch1(s []int, target, low, high int) bool {
	if low > high {
		return false
	}

	mid := (low + high) / 2

	if s[mid] == target {
		return true
	} else if s[mid] > target {
		return BinarySearch1(s, target, low, mid - 1)
	} else {
		return BinarySearch1(s, target, mid + 1, high)
	}
}