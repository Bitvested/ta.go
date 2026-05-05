package random_test

import (
	"sort"
	"testing"

	"github.com/Bitvested/ta.go/src/random"
	"github.com/stretchr/testify/assert"
)

func TestPrngReturnsFunction(t *testing.T) {
	r := random.Prng("seed")
	if r == nil {
		t.Fatal("Prng returned nil")
	}
}

func TestPrngDeterministic(t *testing.T) {
	a := random.Prng("abc")
	b := random.Prng("abc")
	for i := 0; i < 10; i++ {
		if a() != b() {
			t.Fatalf("Prng not deterministic at draw %d", i)
		}
	}
}

func TestPrngRange(t *testing.T) {
	r := random.Prng("abc")
	for i := 0; i < 1000; i++ {
		v := r()
		if v < 0 || v >= 1 {
			t.Fatalf("Prng value %v out of [0,1)", v)
		}
	}
}

func TestPrngDifferentSeeds(t *testing.T) {
	a := random.Prng("abc")()
	b := random.Prng("xyz")()
	assert.NotEqual(t, a, b)
}

func TestPick(t *testing.T) {
	r := random.Prng("pick")
	arr := []float64{1, 2, 3, 4, 5}
	for i := 0; i < 50; i++ {
		v := random.Pick(arr, r)
		found := false
		for _, x := range arr {
			if x == v {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("Pick returned %v, not in input", v)
		}
	}
}

func TestRange(t *testing.T) {
	r := random.Prng("range")
	for i := 0; i < 200; i++ {
		v := random.Range(3, 7, r)
		if v < 3 || v >= 7 {
			t.Fatalf("Range %d out of [3,7)", v)
		}
	}
}

func TestFloat(t *testing.T) {
	r := random.Prng("float")
	for i := 0; i < 200; i++ {
		v := random.Float(0, 1, r)
		if v < 0 || v >= 1 {
			t.Fatalf("Float %v out of [0,1)", v)
		}
	}
}

func TestOrderIsPermutationAndDoesNotMutate(t *testing.T) {
	r := random.Prng("order")
	in := []float64{1, 2, 3, 4, 5}
	snapshot := append([]float64(nil), in...)
	out := random.Order(in, r)
	if !assert.Equal(t, snapshot, in) {
		t.Fatal("Order mutated input")
	}
	if len(out) != len(in) {
		t.Fatalf("Order length mismatch: got %d want %d", len(out), len(in))
	}
	sortedOut := append([]float64(nil), out...)
	sort.Float64s(sortedOut)
	if !assert.Equal(t, snapshot, sortedOut) {
		t.Fatal("Order is not a permutation of input")
	}
}
