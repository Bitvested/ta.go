package statistics_test

import (
	"math"
	"testing"

	tu "github.com/Bitvested/ta.go/internal/testutil"
	stats "github.com/Bitvested/ta.go/src/statistics"
	"github.com/stretchr/testify/assert"
)

// --- v2.0.0 NaN edge cases ---------------------------------------------------

func TestDifNaNOnZeroDenominator(t *testing.T) {
	got := stats.Dif(1, 0)
	if !math.IsNaN(got) {
		t.Fatalf("Dif(_, 0) should return NaN, got %v", got)
	}
}

// --- v2.0.0 mutation-free guarantees -----------------------------------------

func TestStdDoesNotMutateInput(t *testing.T) {
	in := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	snap := append([]float64(nil), in...)
	_ = stats.Std(in, 5)
	if !assert.Equal(t, snap, in) {
		t.Fatalf("Std mutated its input")
	}
}

func TestPercentileDoesNotMutateInput(t *testing.T) {
	in := [][]float64{{6, 4, 7}, {5, 3, 6}, {7, 5, 8}}
	snap := make([][]float64, len(in))
	for i, r := range in {
		snap[i] = append([]float64(nil), r...)
	}
	_ = stats.Percentile(in, 0.5)
	for i, r := range in {
		if !assert.Equal(t, snap[i], r) {
			t.Fatalf("Percentile mutated row %d", i)
		}
	}
}

func TestMedian(t *testing.T) {
	// v2.0.0: even-window length returns avg of two middle sorted values.
	have := stats.Median([]float64{4, 6, 3, 1, 2, 5}, 4)
	want := []float64{3.5, 2.5, 2.5}
	tu.Slice(t, have, want, "Median")
}

func TestNormalize(t *testing.T) {
	have := stats.Normalize([]float64{5, 4, 9, 4}, 0.1)
	want := []float64{0.2222222222222222, 0.06349206349206349, 0.8571428571428571, 0.06349206349206349}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Normalize Failed!")
	}
}

func TestDenormalize(t *testing.T) {
	have := stats.Denormalize([]float64{0.2222222222222222, 0.06349206349206349, 0.8571428571428571, 0.06349206349206349, 0.4444444444444444}, []float64{5, 4, 9, 4}, 0.1)
	want := []float64{5, 4, 9, 4, 6.4}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Denormalize Failed!")
	}
}

func TestStandardize(t *testing.T) {
	have := stats.Standardize([]float64{6, 4, 6, 8, 6})
	want := []float64{0, -1.5811388300841895, 0, 1.5811388300841895, 0}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Standardize Failed!")
	}
}

func TestMad(t *testing.T) {
	// v2.0.0: cascades from Median's avg-of-middles convention.
	have := stats.Mad([]float64{3, 7, 5, 4, 3, 8, 9}, 6)
	want := []float64{1.5, 2}
	tu.Slice(t, have, want, "Mad")
}

func TestAad(t *testing.T) {
	have := stats.Aad([]float64{4, 6, 8, 6, 8, 9, 10, 11}, 7)
	want := []float64{1.6734693877551021, 1.469387755102041}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Aad Failed!")
	}
}

func TestSsd(t *testing.T) {
	have := stats.Ssd([]float64{7, 6, 5, 7, 9, 8, 3, 5, 4}, 7)
	want := []float64{4.869731585445518, 4.9856938190329, 5.3718844791323335}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Ssd Failed!")
	}
}

func TestStd(t *testing.T) {
	have := stats.Std([]float64{1, 2, 3}, 3)
	want := 0.816496580927726
	if !assert.Equal(t, want, have) {
		t.Fatalf("Std Failed!")
	}
}

func TestNormsinv(t *testing.T) {
	have := stats.Normsinv(0.4732)
	want := -0.06722824471054376
	if !assert.Equal(t, want, have) {
		t.Fatalf("Normsinv Failed!")
	}
}

func TestCor(t *testing.T) {
	have := stats.Cor([]float64{1, 2, 3, 4, 5, 2}, []float64{1, 3, 2, 4, 6, 3})
	want := 0.8808929232684737
	if !assert.Equal(t, want, have) {
		t.Fatalf("Cor Failed!")
	}
}

func TestDif(t *testing.T) {
	have := stats.Dif(0.75, 0.5)
	want := 0.5
	if !assert.Equal(t, want, have) {
		t.Fatalf("Dif Failed!")
	}
}

func TestDrawdown(t *testing.T) {
	have := stats.Drawdown([]float64{1, 2, 3, 4, 2, 3})
	want := -0.5
	if !assert.Equal(t, want, have) {
		t.Fatalf("Drawdown Failed!")
	}
}

func TestVariance(t *testing.T) {
	have := stats.Variance([]float64{6, 7, 2, 3, 5, 8, 6, 2}, 7)
	want := []float64{3.918367346938776, 5.061224489795919}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Variance Failed!")
	}
}

func TestPercentile(t *testing.T) {
	have := stats.Percentile([][]float64{{6, 4, 7}, {5, 3, 6}, {7, 5, 8}}, 0.5)
	want := []float64{6, 4, 7}
	tu.Slice(t, have, want, "Percentile")
}

func TestRecentHigh(t *testing.T) {
	have := stats.RecentHigh([]float64{4, 5, 6, 7, 8, 9, 8, 7, 8, 9, 10, 3, 2, 1}, 3)
	want := stats.RecentHighLow{Index: 10, Value: 10}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Recent High Failed!")
	}
}

func TestRecentLow(t *testing.T) {
	have := stats.RecentLow([]float64{1, 4, 5, 6, 4, 3, 2, 3, 4, 3, 5, 7, 8, 8, 5}, 4)
	want := stats.RecentHighLow{Index: 6, Value: 2}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Recent Low Failed!")
	}
}

func TestEr(t *testing.T) {
	have := stats.Er([]float64{0.02, -0.01, 0.03, 0.05, -0.03})
	want := 0.011934565489708282
	if !assert.Equal(t, want, have) {
		t.Fatalf("Er Failed!")
	}
}

func TestAr(t *testing.T) {
	have := stats.Ar([]float64{0.02, -0.01, 0.03, 0.05, -0.03}, 3)
	want := []float64{0.03667479679633267, -0.053301281310417566}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Ar Failed!")
	}
}

func TestAvgwin(t *testing.T) {
	have := stats.Avgwin([]float64{0.01, 0.02, -0.01, -0.03, -0.015, 0.005})
	want := 0.011666666666666665
	if !assert.Equal(t, want, have) {
		t.Fatalf("Avgwin Failed!")
	}
}

func TestAvgloss(t *testing.T) {
	have := stats.Avgloss([]float64{0.01, 0.02, -0.01, -0.03, -0.015, 0.005})
	want := -0.018333333333333333
	if !assert.Equal(t, want, have) {
		t.Fatalf("Avgloss Failed!")
	}
}

func TestSe(t *testing.T) {
	have := stats.Se([]float64{34, 54, 45, 43, 57, 38, 49}, 10)
	want := 2.4243661069253055
	if !assert.Equal(t, want, have) {
		t.Fatalf("Se Failed!")
	}
}

func TestKelly(t *testing.T) {
	have := stats.Kelly([]float64{0.01, 0.02, -0.01, -0.03, -0.015, 0.045, 0.005})
	want := 0.14434748152632182
	if !assert.Equal(t, want, have) {
		t.Fatalf("Kelly Failed!")
	}
}

func TestZscore(t *testing.T) {
	have := stats.Zscore([]float64{34, 54, 45, 43, 57, 38, 49}, 5)
	want := []float64{1.2664106627730554, -1.3314928442246727, 0.4078462733398033}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Zscore Failed!")
	}
}

func TestNormalizePair(t *testing.T) {
	have := stats.NormalizePair([]float64{10, 12, 11, 13}, []float64{100, 130, 100, 140})
	want := [][]float64{{55, 55}, {66, 71.5}, {60.5, 54.99999999999999}, {71.5, 76.99999999999999}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("NormalizePair Failed!")
	}
}

func TestNormalizeFrom(t *testing.T) {
	have := stats.NormalizeFrom([]float64{8, 12, 10, 11}, 100)
	want := []float64{100, 150, 125, 137.5}
	if !assert.Equal(t, want, have) {
		t.Fatalf("NormalizeFrom Failed!")
	}
}

func TestLog(t *testing.T) {
	have := stats.Log([]float64{5, 14, 18, 28, 68, 103})
	want := []float64{1.6094379124341003, 2.639057329615259, 2.8903717578961645, 3.332204510175204, 4.219507705176107, 4.634728988229636}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Log Failed!")
	}
}

func TestExp(t *testing.T) {
	// Tolerant comparator handles cross-runtime last-digit drift in math.Exp.
	have := stats.Exp([]float64{1.6, 2.63, 2.89, 3.33, 4.22, 4.63})
	want := []float64{4.953032424395115, 13.873769902129904, 17.993309601550315, 27.938341703236507, 68.03348428941966, 102.51406411049346}
	tu.Slice(t, have, want, "Exp")
}

func TestNcdf(t *testing.T) {
	have := stats.Ncdf(13, 10, 2)
	want := 0.9331737996110652
	if !assert.Equal(t, want, have) {
		t.Fatalf("Ncdf Failed!")
	}
}

func TestMartingale(t *testing.T) {
	have := stats.Martingale([]float64{-1, -1, 1, 1, -1, -1}, 5, 200, 2)
	var want float64 = 20
	if !assert.Equal(t, want, have) {
		t.Fatalf("Martingale Failed!")
	}
}

func TestAntimartingale(t *testing.T) {
	have := stats.Antimartingale([]float64{1, 1, -1, -1, 1, 1}, 5, 200, 2)
	var want float64 = 20
	if !assert.Equal(t, want, have) {
		t.Fatalf("Antimartingale Failed!")
	}
}

func TestPermutations(t *testing.T) {
	have := stats.Permutations([]float64{10, 10, 10})
	var want float64 = 1000
	if !assert.Equal(t, want, have) {
		t.Fatalf("Permutations Failed!")
	}
}

func TestMse(t *testing.T) {
	have := stats.Mse([]float64{7, 8, 7, 8, 6, 9}, []float64{6, 8, 8, 9, 6, 8})
	want := 0.6666666666666666
	if !assert.Equal(t, want, have) {
		t.Fatalf("MSE Failed!")
	}
}

func TestSum(t *testing.T) {
	have := stats.Sum([]float64{2, 2, 4, 2})
	want := float64(10)
	if !assert.Equal(t, want, have) {
		t.Fatalf("Sum Failed!")
	}
}

func TestCum(t *testing.T) {
	have := stats.Cum([]float64{3, 5, 7, 5, 10}, 4)
	want := []float64{20, 27}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Cum Failed!")
	}
}

func TestPvalue(t *testing.T) {
	have := stats.Pvalue(2.0, 4)
	want := 0.061018363939899845
	if !assert.Equal(t, want, have) {
		t.Fatalf("Pvalue Failed!")
	}
}

// --- Phase 3d: ports from ta.js v2.0.0 --------------------------------------

func TestCovariance(t *testing.T) {
	have := stats.Covariance([]float64{12, 13, 25, 39}, []float64{67, 45, 32, 21}, 4)
	want := []float64{-165.8125}
	tu.Slice(t, have, want, "Covariance")
}

func TestExpectedTrails(t *testing.T) {
	have := stats.ExpectedTrails(100)
	if have != 519 {
		t.Fatalf("ExpectedTrails got %d want 519", have)
	}
}

func TestLrSlope(t *testing.T) {
	have := stats.LrSlope([]float64{5, 6, 6, 3, 4, 6, 7}, 6)
	want := []float64{-0.11428571428571428, 0.17142857142857143}
	tu.Slice(t, have, want, "LrSlope")
}

func TestLrIntercept(t *testing.T) {
	have := stats.LrIntercept([]float64{5, 6, 6, 3, 4, 6, 7}, 6)
	want := []float64{5.285714285714286, 4.904761904761904}
	tu.Slice(t, have, want, "LrIntercept")
}

func TestLrAngle(t *testing.T) {
	have := stats.LrAngle([]float64{5, 6, 6, 3, 4, 6, 7}, 6)
	want := []float64{-6.519801751656986, 9.727578551401603}
	tu.Slice(t, have, want, "LrAngle")
}

func TestTsf(t *testing.T) {
	have := stats.Tsf([]float64{5, 6, 6, 3, 4, 6, 7}, 6)
	want := []float64{4.6, 5.933333333333333}
	tu.Slice(t, have, want, "Tsf")
}

func TestStdSeries(t *testing.T) {
	have := stats.StdSeries([]float64{1, 2, 3, 4, 5}, 3)
	want := []float64{0.816496580927726, 0.816496580927726, 0.816496580927726}
	tu.Slice(t, have, want, "StdSeries")
}

func TestReturnPositive(t *testing.T) {
	have := stats.ReturnPositive([]float64{0.02, 0.01, -0.03, -0.01, 0.005})
	want := []float64{0.02, 0.01, 0.005}
	tu.Slice(t, have, want, "ReturnPositive")
}

func TestReturnNegative(t *testing.T) {
	have := stats.ReturnNegative([]float64{0.02, 0.01, -0.03, -0.01, 0.005})
	want := []float64{-0.03, -0.01}
	tu.Slice(t, have, want, "ReturnNegative")
}
