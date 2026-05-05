package movingaverages_test

import (
	"math/rand"
	"testing"

	ma "github.com/Bitvested/ta.go/src/moving_averages"
)

func benchData(n int) []float64 {
	rng := rand.New(rand.NewSource(1))
	out := make([]float64, n)
	for i := range out {
		out[i] = 100 + rng.Float64()*10
	}
	return out
}

func BenchmarkSma1k(b *testing.B) {
	d := benchData(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ma.Sma(d, 14)
	}
}

func BenchmarkSma10k(b *testing.B) {
	d := benchData(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ma.Sma(d, 50)
	}
}

func BenchmarkSma100k(b *testing.B) {
	d := benchData(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ma.Sma(d, 200)
	}
}

func BenchmarkEma1k(b *testing.B) {
	d := benchData(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ma.Ema(d, 14)
	}
}

func BenchmarkEma10k(b *testing.B) {
	d := benchData(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ma.Ema(d, 50)
	}
}

func BenchmarkEma100k(b *testing.B) {
	d := benchData(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ma.Ema(d, 200)
	}
}

func BenchmarkWma10k(b *testing.B) {
	d := benchData(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ma.Wma(d, 50)
	}
}
