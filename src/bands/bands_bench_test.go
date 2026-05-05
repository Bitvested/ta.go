package bands_test

import (
	"math/rand"
	"testing"

	bnd "github.com/Bitvested/ta.go/src/bands"
)

func benchData(n int) []float64 {
	rng := rand.New(rand.NewSource(1))
	out := make([]float64, n)
	for i := range out {
		out[i] = 100 + rng.Float64()*10
	}
	return out
}

func benchHL(n int) [][]float64 {
	rng := rand.New(rand.NewSource(1))
	out := make([][]float64, n)
	for i := range out {
		c := 100 + rng.Float64()*10
		out[i] = []float64{c + 0.5, c - 0.5}
	}
	return out
}

func BenchmarkBands1k(b *testing.B) {
	d := benchData(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bnd.Bands(d, 20, 2)
	}
}

func BenchmarkBands10k(b *testing.B) {
	d := benchData(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bnd.Bands(d, 20, 2)
	}
}

func BenchmarkBands100k(b *testing.B) {
	d := benchData(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bnd.Bands(d, 20, 2)
	}
}

func BenchmarkDon10k(b *testing.B) {
	d := benchHL(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bnd.Don(d, 20)
	}
}
