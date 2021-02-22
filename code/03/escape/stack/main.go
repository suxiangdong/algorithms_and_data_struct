package main
import "math/rand"

// 前言：
// 当分配的内存超出go语言（1.4.2）的限制时，将逃逸到堆上分配。
// 限制大小：64k，源文件地址： /usr/local/Cellar/go/1.14.2_1/libexec/src/cmd/compile/internal/gc/go.go
//  	 //	maximum size of implicit variables that we will allocate on the stack.
//	   	 //	p := new(T)          allocating T on the stack
//	   	 //	p := &T{}            allocating T on the stack
//	   	 //	s := make([]T, n)    allocating [n]T on the stack
//	   	 // s := []byte("...")   allocating [n]byte on the stack
//	   	 // Note: the flag smallframes can update this value.
//	 		maxImplicitStackVarSize = int64(64 * 1024)
//
// 注意：在go其他版本 maxImplicitStackVarSize 变量定义的位置和值可能不同。
//
// 本人机器 64bit 系统，int占8个字节，因此用 64 * 1024 / 8 = 8192 来做分隔测试。
// 如果用byte[]做测试的话，则是 64 * 1024 / 1 = 65535

func generateLt8192() {
	nums := make([]int, 8191) // < 64KB
	for i := 0; i < 8191; i++ {
		nums[i] = rand.Int()
	}
}

func generateEq8192() {
	nums := make([]int, 8192) // = 64KB
	for i := 0; i < 8192; i++ {
		nums[i] = rand.Int()
	}
}

func generateN(n int) {
	nums := make([]int, n) // < 64KB
	for i := 0; i < n; i++ {
		nums[i] = rand.Int()
	}
}

// # command-line-arguments
// ./main.go:21:14: make([]int, 8191) does not escape
// ./main.go:28:14: make([]int, 8192) escapes to heap
// ./main.go:35:14: make([]int, n) escapes to heap
func main() {
	generateLt8192()
	generateEq8192()
	generateN(1)
}
