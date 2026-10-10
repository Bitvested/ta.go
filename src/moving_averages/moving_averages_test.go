package movingaverages_test

import (
	"testing"

	tu "github.com/Bitvested/ta.go/internal/testutil"
	ma "github.com/Bitvested/ta.go/src/moving_averages"
	"github.com/stretchr/testify/assert"
)

func TestSma(t *testing.T) {
	have := ma.Sma([]float64{1, 2, 3, 4, 5, 6, 10}, 6)
	want := []float64{3.5, 5}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Sma Failed!")
	}
}

func TestEma(t *testing.T) {
	have := ma.Ema([]float64{1, 2, 3, 4, 5, 6, 10}, 6)
	want := []float64{3.5, 5.357142857142857}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Ema Failed!")
	}
}

func TestSmma(t *testing.T) {
	// v2.0.0: Wilder recurrence; output length = N - L + 1.
	have := ma.Smma([]float64{1, 2, 3, 4, 5, 6, 10}, 5)
	want := []float64{3, 3.6, 4.88}
	tu.Slice(t, have, want, "Smma")
}

func TestWma(t *testing.T) {
	have := ma.Wma([]float64{69, 68, 66, 70, 68}, 4)
	want := []float64{68.3, 68.2}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Wma Failed!")
	}
}

func TestPwma(t *testing.T) {
	have := ma.Pwma([]float64{17, 26, 23, 29, 20}, 4)
	want := []float64{24.090909090909093, 25.18181818181818}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Pwma Failed!")
	}
}

func TestHwma(t *testing.T) {
	have := ma.Hwma([]float64{54, 51, 86, 42, 47}, 4)
	want := []float64{56.2, 55}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Hwma Failed!")
	}
}

func TestVwma(t *testing.T) {
	have := ma.Vwma([][]float64{{1, 59}, {1.1, 82}, {1.21, 27}, {1.42, 73}, {1.32, 42}}, 4)
	want := []float64{1.184771784232365, 1.258794642857143}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Vwma Failed!")
	}
}

func TestVwwma(t *testing.T) {
	have := ma.Vwwma([][]float64{{1, 59}, {1.1, 82}, {1.21, 27}, {1.42, 73}, {1.32, 42}}, 4)
	want := []float64{1.2618288590604028, 1.3160229445506693}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Vwwma Failed!")
	}
}

func TestWsma(t *testing.T) {
	have := ma.Wsma([]float64{1, 2, 3, 4, 5, 6, 10}, 6)
	want := []float64{3.5, 4.583333333333333}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Wsma Failed!")
	}
}

func TestLsma(t *testing.T) {
	have := ma.Lsma([]float64{5, 6, 6, 3, 4, 6, 7}, 6)
	want := []float64{4.714285714285714, 5.761904761904762}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Lsma Failed!")
	}
}

func TestHull(t *testing.T) {
	have := ma.Hull([]float64{6, 7, 5, 6, 7, 4, 5, 7}, 6)
	want := []float64{4.761904761904762, 5.476190476190476}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Hull Failed!")
	}
}

func TestKama(t *testing.T) {
	have := ma.Kama([]float64{8, 7, 8, 9, 7, 9}, 2, 4, 8)
	want := []float64{7, 7.0493827160493829, 7.3614814814814817, 7.3328406442107408, 7.415169501286754}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Kama Failed!")
	}
}

func TestCwma(t *testing.T) {
	have := ma.Cwma([]float64{69, 68, 66, 70, 68, 69}, []float64{1, 2, 3, 5, 8})
	want := []float64{68.26315789473684, 68.52631578947368}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Cwma Failed!")
	}
}

func TestDema(t *testing.T) {
	have := ma.Dema([]float64{1, 2, 3, 4, 5, 6, 10}, 3)
	want := []float64{5, 6, 9.25}
	tu.Slice(t, have, want, "Dema")
}

func TestTema(t *testing.T) {
	have := ma.Tema([]float64{1, 2, 3, 4, 5, 6, 10, 8, 11}, 3)
	want := []float64{9.5, 8.5, 10.71875}
	tu.Slice(t, have, want, "Tema")
}

func TestT3(t *testing.T) {
	have := ma.T3([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 2, 0.7)
	want := []float64{6.549999999999997, 7.550000000000001, 8.550000000000004, 9.549999999999997, 10.550000000000004, 11.550000000000004}
	tu.Slice(t, have, want, "T3")
}

func TestTrima(t *testing.T) {
	have := ma.Trima([]float64{1, 2, 3, 4, 5, 6, 7}, 5)
	want := []float64{3, 4, 5}
	tu.Slice(t, have, want, "Trima")
}

func TestVidya(t *testing.T) {
	have := ma.Vidya([]float64{2, 4, 3, 6, 5, 7, 8, 6, 9, 10}, 2, 5, 0.2)
	want := []float64{6, 5.929289321881345, 6.080710678118655, 6.192267103986121, 6.154560422638152, 6.758169308820896, 6.987401355340037}
	tu.Slice(t, have, want, "Vidya")
}

func TestZlema(t *testing.T) {
	have := ma.Zlema([]float64{10, 11, 9, 12, 8, 13, 7}, 4)
	want := []float64{10, 10.8, 9.280000000000001, 11.568000000000001, 8.5408, 12.324480000000001, 7.794688000000001}
	tu.Slice(t, have, want, "Zlema")
}
