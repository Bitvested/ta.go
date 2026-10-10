package indicators_test

import (
	"math"
	"testing"

	ind "github.com/Bitvested/ta.go/src/indicators"
	ma "github.com/Bitvested/ta.go/src/moving_averages"
)

// Contract tests for Occ, encoding the Pine it was ported from:
//
//	closeSeries = variant(basisType, close[delayOffset], basisLen, ...)
//	openSeries  = variant(basisType, open[delayOffset],  basisLen, ...)
//	xlong  = ta.crossover(closeSeries, openSeries)
//	xshort = ta.crossunder(closeSeries, openSeries)
//
// Most of these pin alignment rather than values, because a wrong warm-up assumption
// yields plausible numbers at the wrong bars.

var occBasisTypes = []string{"SMA", "EMA", "DEMA", "TEMA", "WMA", "SMMA", "HullMA", "TMA"}

func occData(n int, step float64) [][]float64 {
	out := make([][]float64, n)
	open := 100.0
	for i := range out {
		out[i] = []float64{open, open + step}
		open += step
	}
	return out
}

func occCell(row []interface{}, col int) float64 {
	if v, ok := row[col].(float64); ok {
		return v
	}
	return math.NaN()
}

func occDir(row []interface{}) string {
	if s, ok := row[2].(string); ok {
		return s
	}
	return ""
}

func TestOccRowsAlignToInputBars(t *testing.T) {
	for _, n := range []int{6, 60, 200} {
		have := ind.Occ(occData(n, 1), 3, "SMA", 0)
		if len(have) != n {
			t.Errorf("n=%d: got %d rows, want %d", n, len(have), n)
		}
		for i, row := range have {
			if len(row) != 3 {
				t.Errorf("n=%d row %d: got %d cells, want 3", n, i, len(row))
			}
		}
	}
}

func TestOccSMABasisIsTheRollingMeanByHand(t *testing.T) {
	// arithmetic done here, not by the library
	data := [][]float64{{10, 12}, {11, 13}, {12, 14}, {13, 15}, {14, 16}, {15, 17}}
	have := ind.Occ(data, 3, "SMA", 0)
	for i, row := range have {
		if i < 2 {
			if !math.IsNaN(occCell(row, 0)) || !math.IsNaN(occCell(row, 1)) {
				t.Errorf("row %d should be NaN, got %v", i, row)
			}
			continue
		}
		var sumClose, sumOpen float64
		for k := i - 2; k <= i; k++ {
			sumOpen += data[k][0]
			sumClose += data[k][1]
		}
		if got, want := occCell(row, 0), sumClose/3; math.Abs(got-want) > 1e-12 {
			t.Errorf("row %d close basis: got %v, want %v", i, got, want)
		}
		if got, want := occCell(row, 1), sumOpen/3; math.Abs(got-want) > 1e-12 {
			t.Errorf("row %d open basis: got %v, want %v", i, got, want)
		}
	}
}

func TestOccDirectionIsWhicheverBasisIsOnTop(t *testing.T) {
	for _, basis := range occBasisTypes {
		for i, row := range ind.Occ(occData(120, 0.7), 8, basis, 0) {
			closeBasis, openBasis := occCell(row, 0), occCell(row, 1)
			if math.IsNaN(closeBasis) {
				if !math.IsNaN(openBasis) {
					t.Errorf("%s row %d: close NaN but open is not", basis, i)
				}
				continue
			}
			want := "short"
			if closeBasis > openBasis {
				want = "long"
			}
			if got := occDir(row); got != want {
				t.Errorf("%s row %d: direction %q, want %q", basis, i, got, want)
			}
		}
	}
}

func TestOccBothBasesAreDefinedOnTheSameBars(t *testing.T) {
	// Warm-up differs per type (DEMA needs 2*len-1 bars, TEMA 3*len-2, Hull its own),
	// and a row is only usable where both lines exist.
	for _, basis := range occBasisTypes {
		for _, basisLen := range []int{3, 8, 13} {
			have := ind.Occ(occData(200, 0.5), basisLen, basis, 0)
			first := -1
			for i, row := range have {
				if math.IsNaN(occCell(row, 0)) {
					if first >= 0 {
						t.Errorf("%s len=%d: gap at row %d", basis, basisLen, i)
					}
					continue
				}
				if first < 0 {
					first = i
				}
				if math.IsNaN(occCell(row, 1)) {
					t.Errorf("%s len=%d row %d: close defined but open is not", basis, basisLen, i)
				}
			}
			if first < 0 {
				t.Errorf("%s len=%d: nothing defined", basis, basisLen)
			}
		}
	}
}

func TestOccBasisMatchesTheUnderlyingPrimitive(t *testing.T) {
	data := occData(200, 0.5)
	have := ind.Occ(data, 8, "SMA", 0)
	closes := make([]float64, len(data))
	for i, row := range data {
		closes[i] = row[1]
	}
	want := ma.Sma(closes, 8)
	first := len(data) - len(want)
	for i := first; i < len(data); i++ {
		if got := occCell(have[i], 0); math.Abs(got-want[i-first]) > 1e-12 {
			t.Errorf("row %d: got %v, want %v", i, got, want[i-first])
		}
	}
}

func TestOccDelayShiftsBothBasesRightByExactlyDelay(t *testing.T) {
	data := occData(120, 0.7)
	for _, basis := range occBasisTypes {
		base := ind.Occ(data, 8, basis, 0)
		delayed := ind.Occ(data, 8, basis, 3)
		for i, row := range delayed {
			if math.IsNaN(occCell(row, 0)) {
				continue
			}
			for col := 0; col < 2; col++ {
				got, want := occCell(row, col), occCell(base[i-3], col)
				if math.Abs(got-want) > 1e-12 {
					t.Errorf("%s row %d col %d: got %v, want %v", basis, i, col, got, want)
				}
			}
		}
		for i := 0; i < 3; i++ {
			if !math.IsNaN(occCell(delayed[i], 0)) {
				t.Errorf("%s row %d: expected NaN with delay 3", basis, i)
			}
		}
	}
}

func TestOccRejectsInputItCannotUse(t *testing.T) {
	data := occData(40, 1)
	for _, basis := range []string{"ALMA", "LSMA", "VWMA", "SSMA", "nonsense", ""} {
		if got := ind.Occ(data, 8, basis, 0); got != nil {
			t.Errorf("basisType %q: got %d rows, want nil", basis, len(got))
		}
	}
	if got := ind.Occ(data, 0, "SMA", 0); got != nil {
		t.Errorf("basisLen 0: got %d rows, want nil", len(got))
	}
	if got := ind.Occ(data, 8, "SMA", -1); got != nil {
		t.Errorf("delay -1: got %d rows, want nil", len(got))
	}
}

func TestOccDataShorterThanTheBasisIsAllNaN(t *testing.T) {
	have := ind.Occ(occData(4, 1), 13, "SMA", 0)
	if len(have) != 4 {
		t.Fatalf("got %d rows, want 4", len(have))
	}
	for i, row := range have {
		for col := 0; col < 2; col++ {
			if !math.IsNaN(occCell(row, col)) {
				t.Errorf("row %d col %d: got %v, want NaN", i, col, row[col])
			}
		}
	}
}
