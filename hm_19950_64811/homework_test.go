package homework

import "testing"

func BenchmarkFibSlow(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FibSlow(50)
	}
}

func BenchmarkFibFast(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FibFast(50)
	}
}

func BenchmarkSumOfSquaresSlow(b *testing.B) {
	const ln = 100000
	sl := make([]int, ln)
	for i := 0; i < ln; i++ {
		sl[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumOfSquaresSlow(sl)
	}
}

func BenchmarkSumOfSquaresFast(b *testing.B) {
	const ln = 100000
	sl := make([]int, ln)
	for i := 0; i < ln; i++ {
		sl[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumOfSquaresFast(sl)
	}
}
