package main

import "fmt"

type Student struct {
	name string
}

func newStudent() *Student {
	return &Student{
		name: "test",
	}
}

// go build -gcflags='-m' main.co
//
// # command-line-arguments
//./main.go:9:6: can inline newStudent
//./main.go:16:23: inlining call to newStudent
//./main.go:17:13: inlining call to fmt.Println
//./main.go:10:9: &Student literal escapes to heap
//./main.go:16:23: &Student literal escapes to heap
//./main.go:17:13: []interface {} literal does not escape
//<autogenerated>:1: .this does not escape
func main() {
	student := newStudent()
	fmt.Println(student)
}