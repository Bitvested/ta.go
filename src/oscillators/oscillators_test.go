package oscillators_test

import (
	"math"
	"testing"

	tu "github.com/Bitvested/ta.go/internal/testutil"
	osc "github.com/Bitvested/ta.go/src/oscillators"
	"github.com/stretchr/testify/assert"
)

func TestMomOsc(t *testing.T) {
	// v2.0.0: Chande CMO formula (sum of price changes, not prices).
	have := osc.MomOsc([]float64{1, 1.2, 1.3, 1.3, 1.2, 1.4}, 4)
	want := []float64{50, 50}
	tu.Slice(t, have, want, "MomOsc")
}

func TestChaikinOsc(t *testing.T) {
	// v2.0.0: NaN→0 coercion removed; fixture uses valid H>L bars.
	have := osc.ChaikinOsc([][]float64{{6, 3, 2, 100}, {7, 5, 4, 200}, {8, 6, 5, 150}, {7, 5, 4, 180}, {9, 8, 7, 120}, {8, 6, 5, 90}}, 2, 4)
	want := []float64{-0.9259259259259309, 14.802469135802468, 6.000823045267488}
	tu.Slice(t, have, want, "ChaikinOsc")
}

func TestAroonOsc(t *testing.T) {
	have := osc.AroonOsc([]float64{2, 5, 4, 5}, 3)
	want := []float64{50, 50}
	if !assert.Equal(t, want, have) {
		t.Fatalf("AroonOsc Failed!")
	}
}

func TestAo(t *testing.T) {
	have := osc.Ao([][]float64{{6, 5}, {8, 6}, {7, 4}, {6, 5}, {7, 6}, {9, 8}}, 2, 5)
	want := []float64{0.0, 0.9000000000000004}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Ao Failed!")
	}
}

func TestAc(t *testing.T) {
	have := osc.Ac([][]float64{{6, 5}, {8, 6}, {7, 4}, {6, 5}, {7, 6}, {9, 8}}, 2, 4)
	want := []float64{0.125, 0.5625}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Ac Failed!")
	}
}

func TestFisher(t *testing.T) {
	// v2.0.0: norm = (price - min) / (max - min), output length = N - L + 1.
	have := osc.Fisher([]float64{8, 6, 8, 9, 7, 8, 9, 8, 7, 8, 6, 7}, 9)
	want := [][]float64{{-0.1104469157900972, 0}, {-0.018907501228583555, -0.1104469157900972}, {-0.32522580129662043, -0.018907501228583555}, {-0.4884829106274281, -0.32522580129662043}}
	tu.Slice2D(t, have, want, "Fisher")
}

func TestGator(t *testing.T) {
	// Derived from Alligator's displaced lines, so the rows carry Alligator's bar
	// alignment. The upper histogram is |jaw - teeth|, the lower -|teeth - lips|.
	have := osc.Gator([]float64{8, 7, 8, 9, 7, 8, 9, 6, 7, 8, 6, 8, 10, 8, 7, 9, 8, 7, 9, 6, 7, 9}, 13, 8, 5, 8, 5, 3)
	nan := math.NaN()
	want := [][]float64{
		{nan, nan}, {nan, nan}, {nan, nan}, {nan, nan}, {nan, nan}, {nan, nan},
		{nan, nan}, {nan, nan}, {nan, nan}, {nan, nan}, {nan, nan}, {nan, nan},
		{nan, -0.1291360000000008}, {nan, -0.3595588000000003},
		{nan, -0.2618657899999999}, {nan, -0.4630659617499999},
		{nan, -0.4089415389312503}, {nan, -0.08916909551515584},
		{nan, -0.13938268776023932}, {nan, -0.24588436885900933},
		{0.1506582487087984, -0.1107115635933269}, {0.14292064217641176, -0.1174390727798702},
	}
	tu.Slice2D(t, have, want, "Gator")
}

func TestKvo(t *testing.T) {
	have := osc.Kvo([][]float64{{10, 9, 8, 100}, {11, 10, 9, 150}, {12, 11, 10, 200}, {13, 12, 11, 180}, {12, 11, 10, 120}, {14, 13, 12, 90}}, 3, 5)
	want := []float64{0, 1111.1111111111113, 1685.1851851851852, 95.67901234567853, -450.1028806584368}
	tu.Slice(t, have, want, "Kvo")
}

func TestUlt(t *testing.T) {
	have := osc.Ult([][]float64{{10, 9, 8}, {12, 11, 10}, {11, 10, 9}, {13, 12, 11}, {12, 11, 10}, {14, 13, 12}}, 2, 3, 4)
	want := []float64{59.18367346938775, 60.714285714285715}
	tu.Slice(t, have, want, "Ult")
}
