package indicators_test

import (
	"fmt"
	"math"
	"testing"

	tu "github.com/Bitvested/ta.go/internal/testutil"
	ind "github.com/Bitvested/ta.go/src/indicators"
)

// Contract tests for HalfTrend, encoding everget's Pine source
// (tradingview.com/script/U1SJ8ubc-HalfTrend/), which the port was transliterated
// from but did not reproduce.
//
// Three properties of that Pine are what a values fixture alone cannot pin, and
// what the v2.0.0 port got wrong: one row per input bar, with the ATR channel NaN
// until the ATR window is full while the trend line is defined from bar 0; price
// extremes taken over `amplitude` bars rather than the ATR window; and a trend
// that actually flips. The old code read the SMA at the start of its window
// instead of at the current bar, so both flip conditions compared values around
// `atrlen` bars old and never fired — on 4901 bars of 5-minute data it produced
// zero trend changes.
//
// The series is generated, not pasted, so the JS and Python ports pin the
// identical numbers.

// assertHalfTrendRows compares 4-tuple rows cell by cell: floats with the shared
// tolerance (NaN treated as a value, since the channel is undefined during the
// ATR warmup), direction strings exactly.
func assertHalfTrendRows(t *testing.T, got, want [][]interface{}, name string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("%s: got %d rows, want %d", name, len(got), len(want))
	}
	for i := range got {
		if len(got[i]) != 4 {
			t.Fatalf("%s: row %d has %d cells, want 4", name, i, len(got[i]))
		}
		for c := 0; c < 4; c++ {
			switch w := want[i][c].(type) {
			case float64:
				g, ok := got[i][c].(float64)
				if !ok {
					t.Errorf("%s[%d][%d]: got %T, want float64", name, i, c, got[i][c])
					continue
				}
				if math.IsNaN(w) {
					if !math.IsNaN(g) {
						t.Errorf("%s[%d][%d]: got %v, want NaN", name, i, c, g)
					}
					continue
				}
				tu.Scalar(t, g, w, fmt.Sprintf("%s[%d][%d]", name, i, c))
			case string:
				if got[i][c] != w {
					t.Errorf("%s[%d][%d]: got %v, want %q", name, i, c, got[i][c], w)
				}
			default:
				t.Fatalf("%s: unexpected fixture cell type %T", name, w)
			}
		}
	}
}

func halfTrendSawtooth() [][]float64 {
	bars := [][]float64{}
	price := 100.0
	for _, leg := range []float64{3, -3, 4, -4, 5, -5, 6, -6, 4, -4, 2, -2} {
		for i := 0; i < 4; i++ {
			price += leg
			bars = append(bars, []float64{price + 1, price, price - 1})
		}
	}
	return bars
}

func TestHalfTrendOneRowPerInputBar(t *testing.T) {
	all := halfTrendSawtooth()
	for _, n := range []int{6, 12, 48} {
		if got := len(ind.HalfTrend(all[:n], 6, 3, 2)); got != n {
			t.Errorf("n=%d: got %d rows, want %d", n, got, n)
		}
	}
}

func TestHalfTrendChannelIsNaNUntilTheATRWindowIsFull(t *testing.T) {
	out := ind.HalfTrend(halfTrendSawtooth(), 6, 3, 2)
	for i, row := range out {
		for _, c := range []int{0, 2} {
			v, ok := row[c].(float64)
			if !ok {
				t.Fatalf("row %d col %d: got %T, want float64", i, c, row[c])
			}
			if gotNaN := math.IsNaN(v); gotNaN != (i < 5) {
				t.Errorf("atrHigh/atrLow at row %d: NaN=%v, want %v", i, gotNaN, i < 5)
			}
		}
		if v, _ := row[1].(float64); math.IsNaN(v) {
			t.Errorf("trend line undefined at row %d", i)
		}
	}
}

func TestHalfTrendFlipsAtThePineBars(t *testing.T) {
	out := ind.HalfTrend(halfTrendSawtooth(), 6, 3, 2)
	flips := []int{}
	for i := 1; i < len(out); i++ {
		if out[i][3] != out[i-1][3] {
			flips = append(flips, i)
		}
	}
	want := []int{6, 10, 14, 18, 22, 26, 30, 35, 38}
	if fmt.Sprint(flips) != fmt.Sprint(want) {
		t.Errorf("flip bars: got %v, want %v", flips, want)
	}

	line := []float64{}
	for _, i := range []int{5, 6, 20, 24, 47} {
		v, _ := out[i][1].(float64)
		line = append(line, v)
	}
	tu.Slice(t, line, []float64{108, 108, 114, 107, 105}, "trend line")

	atrHigh := []float64{}
	for _, i := range []int{5, 6, 24} {
		v, _ := out[i][0].(float64)
		atrHigh = append(atrHigh, v)
	}
	tu.Slice(t, atrHigh, []float64{111.66666666666667, 111.72222222222223, 112.91735310176668}, "atrHigh")
}
