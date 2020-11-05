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
//BenchmarkFact1_1-4      302576580                3.97 ns/op
//BenchmarkFact1_2-4      93612594                12.4 ns/op
//BenchmarkFact1_3-4      59797405                19.0 ns/op
//BenchmarkFact1_4-4      33646185                36.0 ns/op
//BenchmarkFact1_5-4      18706471                66.9 ns/op
func BenchmarkFact1_1(b *testing.B) {benchmarkFactorial1(1, b)}
func BenchmarkFact1_2(b *testing.B) {benchmarkFactorial1(5, b)}
func BenchmarkFact1_3(b *testing.B) {benchmarkFactorial1(8, b)}
func BenchmarkFact1_4(b *testing.B) {benchmarkFactorial1(10, b)}
func BenchmarkFact1_5(b *testing.B) {benchmarkFactorial1(20, b)}

// 尾递归基准测试
//BenchmarkFact2_1-4      308755232                3.84 ns/op
//BenchmarkFact2_2-4      100000000               11.9 ns/op
//BenchmarkFact2_3-4      66315529                17.4 ns/op
//BenchmarkFact2_4-4      46351069                24.9 ns/op
//BenchmarkFact2_5-4      22874569                52.9 ns/op
func BenchmarkFact2_1(b *testing.B) {benchmarkFactorial2(1, b)}
func BenchmarkFact2_2(b *testing.B) {benchmarkFactorial2(5, b)}
func BenchmarkFact2_3(b *testing.B) {benchmarkFactorial2(8, b)}
func BenchmarkFact2_4(b *testing.B) {benchmarkFactorial2(10, b)}
func BenchmarkFact2_5(b *testing.B) {benchmarkFactorial2(20, b)}

func benchmarkFactorial1(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial1(n)
	}
}

func benchmarkFactorial2(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial2(n, 1)
	}
}