package oscillators_test

import (
	"testing"

	osc "github.com/Bitvested/ta.go/src/oscillators"
	"github.com/stretchr/testify/assert"
)

func TestMomOsc(t *testing.T) {
	have := osc.MomOsc([]float64{1, 1.2, 1.3, 1.3, 1.2, 1.4}, 4)
	want := []float64{31.57894736842105, -31.57894736842105, -28.20512820512821}
	if !assert.Equal(t, want, have) {
		t.Fatalf("MomOsc Failed!")
	}
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
	have := osc.Fisher([]float64{8, 6, 8, 9, 7, 8, 9, 8, 7, 8, 6, 7}, 9)
	want := [][]float64{{-0.3181527308248248, -0.1104469157900972}, {-0.4491605923903109, -0.3181527308248248}, {-0.6158050492824234, -0.4491605923903109}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Fisher Failed!")
	}
}

func TestGator(t *testing.T) {
	have := osc.Gator([]float64{8, 7, 8, 9, 7, 8, 9, 6, 7, 8, 6, 8, 10, 8, 7, 9, 8, 7, 9, 6, 7, 9}, 13, 8, 5, 8, 5, 3)
	want := [][]float64{{0.23249042726568714, -0.5289079390282767}, {0.05222951769608297, -0.4002233481486188}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Gator Failed!")
	}
}
