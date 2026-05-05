package multi_test

import (
	"math/rand"
	"testing"

	"github.com/Bitvested/ta.go/multi"
	"github.com/stretchr/testify/assert"
)

func gbmInput(n int, seed int64) []float64 {
	r := rand.New(rand.NewSource(seed))
	out := make([]float64, n)
	out[0] = 100
	for i := 1; i < n; i++ {
		out[i] = out[i-1] * (1 + (r.Float64()-0.5)*0.02)
	}
	return out
}

func TestSimSeededDeterministic(t *testing.T) {
	in := gbmInput(60, 42)
	a := multi.SimSeeded(in, 30, 200, 0.5, 7)
	b := multi.SimSeeded(in, 30, 200, 0.5, 7)
	if !assert.Equal(t, a, b) {
		t.Fatal("SimSeeded must be deterministic for fixed seed")
	}
}

func TestSimSeededShape(t *testing.T) {
	in := gbmInput(40, 11)
	out := multi.SimSeeded(in, 20, 100, 0.5, 99)
	if len(out) != len(in)+20 {
		t.Fatalf("output length %d, want %d", len(out), len(in)+20)
	}
	// The original prefix must be preserved exactly.
	for i, v := range in {
		if out[i] != v {
			t.Fatalf("prefix diverges at %d: got %v want %v", i, out[i], v)
		}
	}
}

func TestSimSeededDifferentSeedsDiffer(t *testing.T) {
	in := gbmInput(40, 1)
	a := multi.SimSeeded(in, 20, 100, 0.5, 1)
	b := multi.SimSeeded(in, 20, 100, 0.5, 2)
	// With 100 sims and 20 future bars, at least one tail value should differ.
	differs := false
	for i := len(in); i < len(a); i++ {
		if a[i] != b[i] {
			differs = true
			break
		}
	}
	if !differs {
		t.Fatal("different seeds produced identical projections")
	}
}

func TestSimSeededPercOrdering(t *testing.T) {
	// p10 ≤ p50 ≤ p90 at every projected bar — ensemble quantiles are
	// monotonic in perc.
	in := gbmInput(50, 3)
	p10 := multi.SimSeeded(in, 20, 500, 0.1, 5)
	p50 := multi.SimSeeded(in, 20, 500, 0.5, 5)
	p90 := multi.SimSeeded(in, 20, 500, 0.9, 5)
	for i := len(in); i < len(p10); i++ {
		if !(p10[i] <= p50[i] && p50[i] <= p90[i]) {
			t.Fatalf("non-monotonic at bar %d: p10=%v p50=%v p90=%v", i, p10[i], p50[i], p90[i])
		}
	}
}

func TestSimSeededInvalidArgs(t *testing.T) {
	in := gbmInput(20, 1)
	if got := multi.SimSeeded(in, 0, 10, 0.5, 1); got != nil {
		t.Fatalf("length=0 should return nil, got %v", got)
	}
	if got := multi.SimSeeded(in, 5, 0, 0.5, 1); got != nil {
		t.Fatalf("sims=0 should return nil, got %v", got)
	}
	if got := multi.SimSeeded([]float64{}, 5, 10, 0.5, 1); got != nil {
		t.Fatalf("empty data should return nil, got %v", got)
	}
}
