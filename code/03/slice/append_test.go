package slice

import "testing"

func TestAppend(t *testing.T) {
	if [2]int{4, 6} != append1() {
		t.Fatal("append1 test failed")
	}

	if [2]int{7, 8} != append2() {
		t.Fatal("append2 test failed")
	}
}