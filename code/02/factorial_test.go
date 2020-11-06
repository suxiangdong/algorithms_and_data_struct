package code

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	t.Parallel()
	datas := []struct {
		in, out int
	}{
		{in: 5, out: 120},
		{in: 10, out: 3628800},
		{in: 0, out: 1},
	}

	for _, v := range datas {
		if r := Factorial1(v.in); r != v.out {
			t.Fatal("Factorial1 test fail")
		}
	}
	for _, v := range datas {
		if r := Factorial2(v.in, 1); r != v.out {
			t.Fatal("Factorial2 test fail")
		}
	}
}

// 普通递归基准测试
//
//BenchmarkFact1_1-4      295683223                3.99 ns/op            0 B/op          0 allocs/op
//BenchmarkFact1_2-4      96102536                12.6 ns/op             0 B/op          0 allocs/op
//BenchmarkFact1_3-4      61382451                19.0 ns/op             0 B/op          0 allocs/op
//BenchmarkFact1_4-4      33726433                35.5 ns/op             0 B/op          0 allocs/op
//BenchmarkFact1_5-4      18646712                64.9 ns/op             0 B/op          0 allocs/op
func BenchmarkFact1_1(b *testing.B) {benchmarkFactorial1(1, b)}
func BenchmarkFact1_2(b *testing.B) {benchmarkFactorial1(5, b)}
func BenchmarkFact1_3(b *testing.B) {benchmarkFactorial1(8, b)}
func BenchmarkFact1_4(b *testing.B) {benchmarkFactorial1(10, b)}
func BenchmarkFact1_5(b *testing.B) {benchmarkFactorial1(20, b)}

// 尾递归基准测试
//BenchmarkFact2_1-4      307067715                3.93 ns/op            0 B/op          0 allocs/op
//BenchmarkFact2_2-4      99202470                11.9 ns/op             0 B/op          0 allocs/op
//BenchmarkFact2_3-4      68360604                17.7 ns/op             0 B/op          0 allocs/op
//BenchmarkFact2_4-4      48932998                24.7 ns/op             0 B/op          0 allocs/op
//BenchmarkFact2_5-4      22987033                52.8 ns/op             0 B/op          0 allocs/op
func BenchmarkFact2_1(b *testing.B) {benchmarkFactorial2(1, b)}
func BenchmarkFact2_2(b *testing.B) {benchmarkFactorial2(5, b)}
func BenchmarkFact2_3(b *testing.B) {benchmarkFactorial2(8, b)}
func BenchmarkFact2_4(b *testing.B) {benchmarkFactorial2(10, b)}
func BenchmarkFact2_5(b *testing.B) {benchmarkFactorial2(20, b)}

func benchmarkFactorial1(n int, b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Factorial1(n)
	}
}

func benchmarkFactorial2(n int, b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Factorial2(n, 1)
	}
}