// Package testutil provides float-tolerant assertion helpers used by every
// subpackage's _test.go files. Mirrors the close() comparator at
// ta.js-1/test/helpers/close.js so the Go tests reuse the same v2.0.0
// fixtures with the same numerical tolerance.
package testutil

import (
	"math"
	"testing"
)

// AbsTol and RelTol match the JS reference defaults.
const (
	AbsTol = 1e-12
	RelTol = 1e-12
)

// Near reports whether a and b are equal within the configured tolerances.
// NaN==NaN counts as equal so test fixtures can encode NaN outputs directly.
func Near(a, b float64) bool {
	if a == b {
		return true
	}
	if math.IsNaN(a) && math.IsNaN(b) {
		return true
	}
	if math.IsInf(a, 0) || math.IsInf(b, 0) {
		return a == b
	}
	diff := math.Abs(a - b)
	if diff <= AbsTol {
		return true
	}
	maxAbs := math.Max(math.Abs(a), math.Abs(b))
	return diff <= RelTol*maxAbs
}

// Scalar fails the test if got and want are not Near.
func Scalar(t *testing.T, got, want float64, msg string) {
	t.Helper()
	if !Near(got, want) {
		t.Errorf("%s: got %.17g, want %.17g (diff %.3g)", msg, got, want, math.Abs(got-want))
	}
}

// Slice fails the test if any element of got differs from want by more than
// the configured tolerance, or the lengths differ.
func Slice(t *testing.T, got, want []float64, msg string) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("%s: length mismatch — got %d, want %d\n  got:  %v\n  want: %v", msg, len(got), len(want), got, want)
		return
	}
	for i := range got {
		if !Near(got[i], want[i]) {
			t.Errorf("%s: element %d differs — got %.17g, want %.17g (diff %.3g)", msg, i, got[i], want[i], math.Abs(got[i]-want[i]))
		}
	}
}

// Slice2D applies Slice element-wise to nested slices.
func Slice2D(t *testing.T, got, want [][]float64, msg string) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("%s: outer length mismatch — got %d, want %d", msg, len(got), len(want))
		return
	}
	for i := range got {
		Slice(t, got[i], want[i], msg+" row "+itoa(i))
	}
}

func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	neg := i < 0
	if neg {
		i = -i
	}
	var buf [20]byte
	pos := len(buf)
	for i > 0 {
		pos--
		buf[pos] = byte('0' + i%10)
		i /= 10
	}
	if neg {
		pos--
		buf[pos] = '-'
	}
	return string(buf[pos:])
}
