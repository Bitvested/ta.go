package oscillators_test

import (
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
	have := osc.ChaikinOsc([][]float64{{2, 3, 4, 6}, {5, 5, 5, 4}, {5, 4, 3, 7}, {4, 3, 3, 4}, {6, 5, 4, 6}, {7, 4, 3, 6}}, 2, 4)
	want := []float64{-1.6666666666666665, -0.28888888888888886, -0.7362962962962962}
	if !assert.Equal(t, want, have) {
		t.Fatalf("ChaikinOsc Failed!")
	}
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
	// v2.0.0: cascades from Smma's Wilder fix.
	have := osc.Gator([]float64{8, 7, 8, 9, 7, 8, 9, 6, 7, 8, 6, 8, 10, 8, 7, 9, 8, 7, 9, 6, 7, 9}, 13, 8, 5, 8, 5, 3)
	want := [][]float64{{-0.08721987687932575, -0.17579145857838796}, {-0.14292064217641176, -0.1174390727798702}, {-0.1506582487087984, -0.1107115635933269}}
	tu.Slice2D(t, have, want, "Gator")
}
