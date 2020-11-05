package code

// 普通递归算法
func Fib1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return Fib1(n - 1) + Fib1(n - 2)
}