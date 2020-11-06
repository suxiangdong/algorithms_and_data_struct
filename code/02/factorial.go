package code

import (
	stack2 "github.com/golang-collections/collections/stack"
)

// 阶乘的基本递归实现
func Factorial1(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial1(n-1)
}

// 尾递归实现
func Factorial2(n, r int) int {
	if n == 0 {
		return r
	}
	return Factorial2(n - 1, r * n)
}

// 手工压栈、非递归实现
func Factorial3(n int) int {
	res := 1
	stack := stack2.New()
	for n > 0 {
		stack.Push(n)
		n--
	}
	for stack.Len() > 0 {
		v := stack.Pop()
		// 确定v一定是int，否则可用switch v.(type)处理
		res *= v.(int)
	}
	return res
}

// 循环实现
func Factorial4(n int) int {
	res := 1
	for n > 0 {
		res *= n
		n--
	}
	return res
}