package indicators_test

import (
	"fmt"
	"math"
	"testing"

	tu "github.com/Bitvested/ta.go/internal/testutil"
	ind "github.com/Bitvested/ta.go/src/indicators"
	ma "github.com/Bitvested/ta.go/src/moving_averages"
)

// Contract tests for Alligator, encoding the Pine source from the ta.py issue
// "Williams Alligator not matching the one in trading view" rather than the
// implementation. Two things come off that Pine: smma is seeded with
// ta.sma over the first `length` values, and plot(..., offset=shift) shifts the
// plot right, so bar i carries the value computed at bar i - shift.

var alligatorData = []float64{8, 7, 8, 9, 7, 8, 9, 6, 7, 8, 6, 8, 10, 8, 7, 9, 8, 7, 9, 6, 7, 9}

func repeatAlligatorData(n int) []float64 {
	out := make([]float64, n)
	for i := range out {
		out[i] = alligatorData[i%len(alligatorData)]
	}
	return out
}

func firstDefinedAlligatorBar(rows [][]float64, col int) int {
	for i, row := range rows {
		if !math.IsNaN(row[col]) {
			return i
		}
	}
	return -1
}

func TestAlligatorRowsAlignToInputBars(t *testing.T) {
	for _, n := range []int{22, 60, 200} {
		if got := len(ind.Alligator(repeatAlligatorData(n), 13, 8, 5, 8, 5, 3)); got != n {
			t.Errorf("n=%d: got %d rows, want %d", n, got, n)
		}
	}
}

func TestAlligatorFirstDefinedBar(t *testing.T) {
	have := ind.Alligator(alligatorData, 13, 8, 5, 8, 5, 3)
	for col, want := range []int{13 - 1 + 8, 8 - 1 + 5, 5 - 1 + 3} {
		if got := firstDefinedAlligatorBar(have, col); got != want {
			t.Errorf("line %d first defined at bar %d, want %d", col, got, want)
		}
	}
	for col, v := range have[len(have)-1] {
		if math.IsNaN(v) {
			t.Errorf("line %d is undefined at the last bar", col)
		}
	}
}

func TestAlligatorZeroShiftsDropTheDisplacement(t *testing.T) {
	have := ind.Alligator(alligatorData, 13, 8, 5, 0, 0, 0)
	for col, want := range []int{13 - 1, 8 - 1, 5 - 1} {
		if got := firstDefinedAlligatorBar(have, col); got != want {
			t.Errorf("line %d first defined at bar %d, want %d", col, got, want)
		}
	}
}

func TestAlligatorSmmaSeedLandsShiftBarsLater(t *testing.T) {
	// Hand arithmetic: the fixture's first 13 values sum to 101, its first 8 to
	// 62 and its first 5 to 39, so the seeds are 101/13, 62/8 and 39/5 at bars 12,
	// 7 and 4 -- each drawn `shift` bars later.
	have := ind.Alligator(alligatorData, 13, 8, 5, 8, 5, 3)
	tu.Scalar(t, have[20][0], 101.0/13.0, "jaw seed at bar 20")
	tu.Scalar(t, have[12][1], 62.0/8.0, "teeth seed at bar 12")
	tu.Scalar(t, have[7][2], 39.0/5.0, "lips seed at bar 7")
}

func TestAlligatorJawIsRawSmmaDisplacedByShift(t *testing.T) {
	data := repeatAlligatorData(120)
	for _, c := range []struct{ jl, js int }{{13, 8}, {13, 0}, {5, 3}, {21, 1}} {
		have := ind.Alligator(data, c.jl, 8, 5, c.js, 5, 3)
		raw := ma.Smma(data, c.jl)
		first := c.jl - 1 + c.js
		for i := 0; i < first; i++ {
			if !math.IsNaN(have[i][0]) {
				t.Errorf("jl=%d js=%d: jaw[%d] should be NaN, got %v", c.jl, c.js, i, have[i][0])
			}
		}
		for i := first; i < len(data); i++ {
			tu.Scalar(t, have[i][0], raw[i-first], fmt.Sprintf("jl=%d js=%d bar %d", c.jl, c.js, i))
		}
	}
}

func TestAlligatorDataShorterThanJawLengthIsAllNaN(t *testing.T) {
	have := ind.Alligator([]float64{1, 1, 1, 1, 1}, 13, 8, 5, 8, 5, 3)
	if len(have) != 5 {
		t.Fatalf("got %d rows, want 5", len(have))
	}
	for i, row := range have {
		for col, v := range row {
			if !math.IsNaN(v) {
				t.Errorf("row %d column %d: got %v, want NaN", i, col, v)
			}
		}
	}
}
