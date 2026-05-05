package statistics_test

import (
	"math/rand"
	"testing"

	stats "github.com/Bitvested/ta.go/src/statistics"
)

func benchData(n int) []float64 {
	rng := rand.New(rand.NewSource(1))
	out := make([]float64, n)
	for i := range out {
		out[i] = 100 + rng.Float64()*10
	}
	return out
}

func BenchmarkStd(b *testing.B) {
	d := benchData(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = stats.Std(d, 50)
	}
}

func BenchmarkVariance10k(b *testing.B) {
	d := benchData(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = stats.Variance(d, 50)
	}
}

func BenchmarkVariance100k(b *testing.B) {
	d := benchData(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = stats.Variance(d, 50)
	}
}

func BenchmarkSum100k(b *testing.B) {
	d := benchData(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = stats.Sum(d)
	}
}
