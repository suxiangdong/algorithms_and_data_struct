package sliceVsArray

import "testing"

// BenchmarkNewArray-4       480286              2465 ns/op               0 B/op          0 allocs/op
func BenchmarkNewArray(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		newArray()
	}
}

// BenchmarkNewSlice-4       355299              3472 ns/op            8192 B/op          1 allocs/op
func BenchmarkNewSlice(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		newSlice()
	}
}