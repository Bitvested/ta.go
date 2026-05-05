// Package random hosts a deterministic seeded PRNG and a small set of
// sampling helpers (Prng, Pick, Range, Float, Order) that mirror the
// random/ subpackage of ta.js v2.0.0.
//
// All sampling helpers take an optional rng function — pass nil to
// fall back to math/rand's package-level source.
package random

import (
	"math"
	"math/rand"
)

// Prng returns a deterministic, seeded PRNG drawn from the same
// 32-bit "mulberry32" construction used by ta.js' random/prng.js.
// The same seed string always yields the same value sequence; values
// are uniformly distributed on [0, 1).
//
// The seed is mixed byte-by-byte. ASCII seeds (which is what ta.js'
// test suite uses) behave identically across the two libraries.
func Prng(seed string) func() float64 {
	var h uint32 = 1779033703 ^ uint32(len(seed))
	for i := 0; i < len(seed); i++ {
		h = (h ^ uint32(seed[i])) * 3432918353
		h = (h << 13) | (h >> 19)
	}
	h = (h ^ (h >> 16)) * 2246822507
	h = (h ^ (h >> 13)) * 3266489909
	h ^= h >> 16
	state := h
	return func() float64 {
		state += 0x6D2B79F5
		t := state
		t = (t ^ (t >> 15)) * (t | 1)
		t ^= t + ((t ^ (t >> 7)) * (t | 61))
		return float64(t^(t>>14)) / 4294967296.0
	}
}

// Pick returns a uniformly-chosen element of `data`. Panics on empty input.
func Pick(data []float64, rng func() float64) float64 {
	r := draw(rng)
	idx := int(math.Floor(r * float64(len(data))))
	return data[idx]
}

// Range returns an integer uniformly drawn from [min, max). Mirrors
// ta.js' floor-rounded range semantics.
func Range(min, max int, rng func() float64) int {
	r := draw(rng)
	return int(math.Floor(r*(float64(max)-float64(min)) + float64(min)))
}

// Float returns a float uniformly drawn from [min, max).
func Float(min, max float64, rng func() float64) float64 {
	return draw(rng)*(max-min) + min
}

// Order returns a uniformly-random permutation of `data`. The input
// slice is not mutated.
func Order(data []float64, rng func() float64) []float64 {
	pool := append([]float64(nil), data...)
	out := make([]float64, 0, len(pool))
	for len(pool) > 0 {
		r := draw(rng)
		idx := int(math.Floor(r * float64(len(pool))))
		out = append(out, pool[idx])
		pool = append(pool[:idx], pool[idx+1:]...)
	}
	return out
}

// draw returns one sample from `rng`, falling back to math/rand's
// package-level source when rng is nil.
func draw(rng func() float64) float64 {
	if rng != nil {
		return rng()
	}
	return rand.Float64()
}
