package experimental_test

import (
	"testing"

	exp "github.com/Bitvested/ta.go/src/experimental"
	stats "github.com/Bitvested/ta.go/src/statistics"
	"github.com/stretchr/testify/assert"
)

func TestSupport(t *testing.T) {
	hl := stats.RecentLow([]float64{4, 3, 2, 5, 7, 6, 5, 4, 7, 8, 5, 4, 6, 7, 5}, 20)
	have := exp.Support([]float64{4, 3, 2, 5, 7, 6, 5, 4, 7, 8, 5, 4, 6, 7, 5}, hl)
	want := exp.Line{Slope: 0.2222222222222222, Value: 2, Index: 2}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Support Failed!")
	}
}

func TestResistance(t *testing.T) {
	hl := stats.RecentHigh([]float64{5, 7, 5, 5, 4, 6, 5, 4, 6, 5, 4, 3, 2, 4, 3, 2, 1}, 20)
	have := exp.Resistance([]float64{5, 7, 5, 5, 4, 6, 5, 4, 6, 5, 4, 3, 2, 4, 3, 2, 1}, hl)
	want := exp.Line{Slope: -0.14285714285714285, Value: 7, Index: 1}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Resistance Failed!")
	}
	have2 := exp.LineCalc(4, have)
	want2 := 6.428571428571429
	if !assert.Equal(t, want2, have2) {
		t.Fatalf("LineCalc Failed!")
	}
}

func TestDivergenceState(t *testing.T) {
	have := exp.DivergenceState(
		[]float64{48, 34, 43, 54, 56, 64, 43, 51, 52, 53, 55, 51, 48, 45, 40, 42, 44, 45},
		[]float64{76, 74, 43, 55, 34, 32, 45, 47, 48, 53, 54, 54, 50, 52, 49, 47, 48, 46},
		12, 3, 1, 0.03, 0.01,
	)
	want := [][]string{
		{"convergence"},
		{"divergence"},
		{"convergence"},
		{"divergence"},
		{"convergence"},
		{"exaggerated_bearish"},
	}
	if !assert.Equal(t, want, have) {
		t.Fatalf("DivergenceState Failed!")
	}
}
