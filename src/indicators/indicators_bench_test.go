package indicators_test

import (
	"math/rand"
	"testing"

	ind "github.com/Bitvested/ta.go/src/indicators"
)

func benchData(n int) []float64 {
	rng := rand.New(rand.NewSource(1))
	out := make([]float64, n)
	for i := range out {
		out[i] = 100 + rng.Float64()*10
	}
	return out
}

func benchHCL(n int) [][]float64 {
	rng := rand.New(rand.NewSource(1))
	out := make([][]float64, n)
	for i := range out {
		c := 100 + rng.Float64()*10
		out[i] = []float64{c + 0.5, c, c - 0.5}
	}
	return out
}

func BenchmarkRsi1k(b *testing.B) {
	d := benchData(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ind.Rsi(d, 14)
	}
}

func BenchmarkRsi10k(b *testing.B) {
	d := benchData(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ind.Rsi(d, 14)
	}
}

func BenchmarkRsi100k(b *testing.B) {
	d := benchData(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ind.Rsi(d, 14)
	}
}

func BenchmarkAroonUp10k(b *testing.B) {
	d := benchData(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ind.AroonUp(d, 25)
	}
}

func BenchmarkAroonDown10k(b *testing.B) {
	d := benchData(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ind.AroonDown(d, 25)
	}
}

func BenchmarkPr10k(b *testing.B) {
	d := benchData(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ind.Pr(d, 14)
	}
}

func BenchmarkStoch10k(b *testing.B) {
	d := benchHCL(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ind.Stoch(d, 14, 3, 3)
	}
}

func BenchmarkIchimoku5k(b *testing.B) {
	d := benchHCL(5000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ind.Ichimoku(d, 9, 26, 52, 26)
	}
}

func BenchmarkElderray10k(b *testing.B) {
	d := benchData(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ind.Elderray(d, 14)
	}
}
