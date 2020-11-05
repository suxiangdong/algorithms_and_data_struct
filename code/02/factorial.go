package code

// 阶乘的基本递归实现
func Factorial1(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial1(n-1)
}
