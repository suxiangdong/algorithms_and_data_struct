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
	for _, v := range datas {
		if r := Factorial3(v.in); r != v.out {
			t.Fatal("Factorial3 test fail")
		}
	}
	for _, v := range datas {
		if r := Factorial4(v.in); r != v.out {
			t.Fatal("Factorial4 test fail")
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

// 手动压栈
//BenchmarkFact3_1-4      22467722                47.4 ns/op            40 B/op          2 allocs/op
//BenchmarkFact3_2-4       5100942               236 ns/op             200 B/op         10 allocs/op
//BenchmarkFact3_3-4       3214581               375 ns/op             320 B/op         16 allocs/op
//BenchmarkFact3_4-4       2504330               475 ns/op             400 B/op         20 allocs/op
//BenchmarkFact3_5-4       1285077               940 ns/op             800 B/op         40 allocs/op
func BenchmarkFact3_1(b *testing.B) {benchmarkFactorial3(1, b)}
func BenchmarkFact3_2(b *testing.B) {benchmarkFactorial3(5, b)}
func BenchmarkFact3_3(b *testing.B) {benchmarkFactorial3(8, b)}
func BenchmarkFact3_4(b *testing.B) {benchmarkFactorial3(10, b)}
func BenchmarkFact3_5(b *testing.B) {benchmarkFactorial3(20, b)}

// 循环实现
//BenchmarkFact4_1-4      483686618                2.29 ns/op            0 B/op          0 allocs/op
//BenchmarkFact4_2-4      349090507                3.51 ns/op            0 B/op          0 allocs/op
//BenchmarkFact4_3-4      272061386                4.51 ns/op            0 B/op          0 allocs/op
//BenchmarkFact4_4-4      244875709                4.86 ns/op            0 B/op          0 allocs/op
//BenchmarkFact4_5-4      143601668                8.20 ns/op            0 B/op          0 allocs/op
func BenchmarkFact4_1(b *testing.B) {benchmarkFactorial4(1, b)}
func BenchmarkFact4_2(b *testing.B) {benchmarkFactorial4(5, b)}
func BenchmarkFact4_3(b *testing.B) {benchmarkFactorial4(8, b)}
func BenchmarkFact4_4(b *testing.B) {benchmarkFactorial4(10, b)}
func BenchmarkFact4_5(b *testing.B) {benchmarkFactorial4(20, b)}

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

func benchmarkFactorial3(n int, b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Factorial3(n)
	}
}

func benchmarkFactorial4(n int, b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Factorial4(n)
	}
}