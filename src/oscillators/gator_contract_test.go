package oscillators_test

import (
	"math"
	"testing"

	tu "github.com/Bitvested/ta.go/internal/testutil"
	ind "github.com/Bitvested/ta.go/src/indicators"
	osc "github.com/Bitvested/ta.go/src/oscillators"
)

// Contract tests for Gator. It is derived from Alligator's displaced lines, so it
// inherits the bar alignment, and its upper histogram is an absolute difference:
// jaw - teeth comes out negative, which is the sign bug this pins.

var gatorData = []float64{8, 7, 8, 9, 7, 8, 9, 6, 7, 8, 6, 8, 10, 8, 7, 9, 8, 7, 9, 6, 7, 9}

func TestGatorDerivedFromDisplacedAlligatorLines(t *testing.T) {
	for _, shifts := range [][3]int{{8, 5, 3}, {0, 0, 0}, {3, 2, 1}} {
		lines := ind.Alligator(gatorData, 13, 8, 5, shifts[0], shifts[1], shifts[2])
		hist := osc.Gator(gatorData, 13, 8, 5, shifts[0], shifts[1], shifts[2])
		if len(lines) != len(hist) {
			t.Fatalf("shifts %v: %d alligator rows vs %d gator rows", shifts, len(lines), len(hist))
		}
		for i := range lines {
			jaw, teeth, lips := lines[i][0], lines[i][1], lines[i][2]
			upper, lower := hist[i][0], hist[i][1]

			if got, want := math.IsNaN(upper), math.IsNaN(jaw) || math.IsNaN(teeth); got != want {
				t.Errorf("shifts %v row %d: upper NaN=%v, want %v", shifts, i, got, want)
			}
			if !math.IsNaN(upper) {
				tu.Scalar(t, upper, math.Abs(jaw-teeth), "upper histogram")
			}

			if got, want := math.IsNaN(lower), math.IsNaN(teeth) || math.IsNaN(lips); got != want {
				t.Errorf("shifts %v row %d: lower NaN=%v, want %v", shifts, i, got, want)
			}
			if !math.IsNaN(lower) {
				tu.Scalar(t, lower, -math.Abs(teeth-lips), "lower histogram")
			}
		}
	}
}

func TestGatorUpperHistogramIsNeverNegative(t *testing.T) {
	for i, row := range osc.Gator(gatorData, 13, 8, 5, 8, 5, 3) {
		if !math.IsNaN(row[0]) && row[0] < 0 {
			t.Errorf("upper[%d] = %v, want >= 0", i, row[0])
		}
	}
}

func TestGatorDataShorterThanJawLengthIsAllNaN(t *testing.T) {
	rows := osc.Gator([]float64{1, 1, 1, 1, 1}, 13, 8, 5, 8, 5, 3)
	if len(rows) != 5 {
		t.Fatalf("got %d rows, want 5", len(rows))
	}
	for i, row := range rows {
		for col, v := range row {
			if !math.IsNaN(v) {
				t.Errorf("row %d column %d: got %v, want NaN", i, col, v)
			}
		}
	}
}
