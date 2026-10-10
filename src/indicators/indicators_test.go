package indicators_test

import (
	"math"
	"testing"

	tu "github.com/Bitvested/ta.go/internal/testutil"
	ind "github.com/Bitvested/ta.go/src/indicators"
	"github.com/stretchr/testify/assert"
)

// --- v2.0.0 NaN edge cases ---------------------------------------------------

func TestRocNaNOnZeroReference(t *testing.T) {
	got := ind.Roc([]float64{0, 0, 0, 0, 0}, 3)
	for i, v := range got {
		if !math.IsNaN(v) {
			t.Errorf("Roc on flat zero data should be NaN at %d, got %v", i, v)
		}
	}
}

func TestPrNaNOnFlatWindow(t *testing.T) {
	got := ind.Pr([]float64{5, 5, 5, 5, 5}, 4)
	for i, v := range got {
		if !math.IsNaN(v) {
			t.Errorf("Pr on flat window should be NaN at %d, got %v", i, v)
		}
	}
}

func TestRsiNaNOnFlatWindow(t *testing.T) {
	got := ind.Rsi([]float64{5, 5, 5, 5, 5, 5, 5, 5}, 6)
	if len(got) == 0 || !math.IsNaN(got[0]) {
		t.Errorf("Rsi on flat window should be NaN, got %v", got)
	}
}

func TestRsi(t *testing.T) {
	// v2.0.0: Wilder smoothing (was Cutler simple-average).
	have := ind.Rsi([]float64{1, 2, 3, 4, 5, 6, 7, 5, 6, 4}, 6)
	want := []float64{100, 71.42857142857143, 75.60975609756098, 55.95667870036102}
	tu.Slice(t, have, want, "Rsi")
}

func TestWrsi(t *testing.T) {
	have := ind.Wrsi([]float64{1, 2, 3, 4, 5, 6, 7, 5, 6}, 6)
	want := []float64{100, 71.42857142857143, 75.60975609756098}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Wrsi Failed!")
	}
}

func TestMacd(t *testing.T) {
	have := ind.Macd([]float64{1, 2, 3, 4, 5, 6, 14}, 3, 6)
	want := []float64{1.5, 3}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Macd Failed!")
	}
}

func TestAroonUp(t *testing.T) {
	have := ind.AroonUp([]float64{5, 4, 5, 2}, 3)
	want := []float64{100, 50}
	if !assert.Equal(t, want, have) {
		t.Fatalf("AroonUp Failed!")
	}
}

func TestAroonDown(t *testing.T) {
	have := ind.AroonDown([]float64{2, 5, 4, 5}, 3)
	want := []float64{0, 50}
	if !assert.Equal(t, want, have) {
		t.Fatalf("AroonDown Failed!")
	}
}

func TestMfi(t *testing.T) {
	// v2.0.0: input is [H, L, C, V] tuples (TA-Lib convention).
	have := ind.Mfi([][]float64{{10, 8, 9, 100}, {12, 10, 11, 150}, {11, 9, 10, 200}, {13, 11, 12, 120}, {12, 10, 11, 80}}, 3)
	want := []float64{60.70726915520628, 33.333333333333336}
	tu.Slice(t, have, want, "Mfi")
}

func TestRoc(t *testing.T) {
	// v2.0.0: returns percent (×100).
	have := ind.Roc([]float64{1, 2, 3, 4}, 3)
	want := []float64{200, 100}
	tu.Slice(t, have, want, "Roc")
}

func TestCop(t *testing.T) {
	// v2.0.0: cascades from Roc percent change (×100).
	have := ind.Cop([]float64{3, 4, 5, 3, 4, 5, 6, 4, 7, 5, 4, 7, 5}, 4, 6, 5)
	want := []float64{37.55555555555556, 23.666666666666664}
	tu.Slice(t, have, want, "Cop")
}

func TestKst(t *testing.T) {
	// v2.0.0: cascades from Roc ×100 convention.
	have := ind.Kst([]float64{8, 6, 7, 6, 8, 9, 7, 5, 6, 7, 6, 8, 6, 7, 6, 8, 9, 9, 8, 6, 4, 6, 5, 6, 7, 8, 9}, 5, 7, 10, 15, 5, 5, 5, 7, 4)
	want := [][]float64{{-68.28231292517006, -51.74886621315193}, {-29.393424036281182, -57.86281179138322}, {35.17800453514739, -35.96882086167801}}
	tu.Slice2D(t, have, want, "Kst")
}

func TestObv(t *testing.T) {
	have := ind.Obv([][]float64{{25200, 10}, {30000, 10.15}, {25600, 10.17}, {32000, 10.13}})
	want := []float64{0, 30000, 55600, 23600}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Obv Failed!")
	}
}

func TestVwap(t *testing.T) {
	have := ind.Vwap([][]float64{{127.21, 89329}, {127.17, 16137}, {127.16, 23945}}, 2)
	want := []float64{127.20387973375304, 127.16402599670675}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Vwap Failed!")
	}
}

func TestMom(t *testing.T) {
	have := ind.Mom([]float64{1, 1.1, 1.2, 1.24, 1.34}, 4, false)
	want := []float64{0.24, 0.24}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Mom Failed!")
	}
}

func TestBop(t *testing.T) {
	have := ind.Bop([][]float64{{4, 5, 4, 5}, {5, 6, 5, 6}, {6, 8, 5, 6}}, 2)
	want := []float64{1, 0.5}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Bop Failed!")
	}
}

func TestFi(t *testing.T) {
	have := ind.Fi([][]float64{{1.4, 200}, {1.5, 240}, {1.1, 300}, {1.2, 240}, {1.5, 400}}, 4)
	want := []float64{12.00000000000001}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Fi Failed!")
	}
}

func TestAsi(t *testing.T) {
	have := ind.Asi([][]float64{{7, 6, 4}, {9, 7, 5}, {9, 8, 6}})
	want := []float64{0, -12.5}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Asi Failed!")
	}
}

func TestPr(t *testing.T) {
	have := ind.Pr([]float64{2, 1, 3, 1, 2}, 4)
	want := []float64{-100, -50}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Pr Failed!")
	}
}

func TestIchimoku(t *testing.T) {
	have := ind.Ichimoku([][]float64{{6, 3, 2}, {5, 4, 2}, {5, 4, 3}, {6, 4, 3}, {7, 6, 4}, {6, 5, 3}, {7, 6, 5}, {7, 5, 3}, {8, 6, 5}, {9, 7, 6}, {8, 7, 6}, {7, 5, 5}, {6, 5, 4}, {6, 5, 3}, {6, 3, 2}, {5, 4, 2}}, 2, 4, 6, 4)
	want := [][]float64{{7, 6, 10.5, 6, 5}, {7.5, 6, 7.5, 5.5, 6}, {6.5, 7, 8, 5, 5}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Ichimoku Failed!")
	}
}

func TestStoch(t *testing.T) {
	// v2.0.0: monotonic-deque tracker emits one extra leading value.
	have := ind.Stoch([][]float64{{3, 2, 1}, {2, 2, 1}, {4, 3, 1}, {2, 2, 1}}, 2, 1, 1)
	want := [][]float64{{50, 50}, {66.66666666666667, 66.66666666666667}, {33.333333333333336, 33.333333333333336}}
	tu.Slice2D(t, have, want, "Stoch")
}

func TestAtr(t *testing.T) {
	// v2.0.0: TR uses |H-Cprev| and |L-Cprev|; output length = N - L + 1.
	have := ind.Atr([][]float64{{3, 2, 1}, {2, 2, 1}, {4, 3, 1}, {2, 2, 1}}, 3)
	want := []float64{2.111111111111111, 2.074074074074074}
	tu.Slice(t, have, want, "Atr")
}

func TestAlligator(t *testing.T) {
	// Rows align to the input bars: one row per bar, each line displaced right by
	// its own shift and NaN until its first bar (jaw 12+8=20, teeth 7+5=12,
	// lips 4+3=7). Spot-checked by hand: jaw[20] = 101/13 = 7.76923 is the raw
	// SMMA's seed value landing 8 bars later, teeth[12] = 62/8 = 7.75 likewise.
	have := ind.Alligator([]float64{8, 7, 8, 9, 7, 8, 9, 6, 7, 8, 6, 8, 10, 8, 7, 9, 8, 7, 9, 6, 7, 9}, 13, 8, 5, 8, 5, 3)
	nan := math.NaN()
	want := [][]float64{
		{nan, nan, nan}, {nan, nan, nan}, {nan, nan, nan}, {nan, nan, nan},
		{nan, nan, nan}, {nan, nan, nan}, {nan, nan, nan},
		{nan, nan, 7.8}, {nan, nan, 7.84}, {nan, nan, 8.072},
		{nan, nan, 7.6575999999999995}, {nan, nan, 7.526079999999999},
		{nan, 7.75, 7.620863999999999}, {nan, 7.65625, 7.2966912},
		{nan, 7.69921875, 7.43735296}, {nan, 7.48681640625, 7.949882368},
		{nan, 7.55096435546875, 7.9599058944}, {nan, 7.857093811035156, 7.76792471552},
		{nan, 7.874957084655762, 8.014339772416001}, {nan, 7.7655874490737915, 8.0114718179328},
		{7.769230769230769, 7.919889017939568, 7.809177454346241},
		{7.78698224852071, 7.929902890697122, 8.047341963476992},
	}
	tu.Slice2D(t, have, want, "Alligator")
}

func TestFractals(t *testing.T) {
	have := ind.Fractals([][]float64{{7, 6}, {8, 6}, {9, 6}, {8, 5}, {7, 4}, {6, 3}, {7, 4}, {8, 5}})
	want := [][]bool{{false, false}, {false, false}, {true, false}, {false, false}, {false, false}, {false, true}, {false, false}, {false, false}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Fractals Failed!")
	}
}

func TestFractalsPrice(t *testing.T) {
	// v2.0.0 new variant: emits highs at up-fractals, lows at down-fractals, -1 elsewhere.
	have := ind.FractalsPrice([][]float64{{7, 6}, {8, 6}, {9, 6}, {8, 5}, {7, 4}, {6, 3}, {7, 4}, {8, 5}})
	want := [][]float64{{-1, -1}, {-1, -1}, {9, -1}, {-1, -1}, {-1, -1}, {-1, 3}, {-1, -1}, {-1, -1}}
	tu.Slice2D(t, have, want, "FractalsPrice")
}

func TestFib(t *testing.T) {
	have := ind.Fib(1, 2)
	want := []float64{1, 1.236, 1.3820000000000001, 1.5, 1.6179999999999999, 1.786, 2, 2.6180000000000003, 3.618, 4.618, 5.236}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Fib Failed!")
	}
}

func TestTsi(t *testing.T) {
	have := ind.Tsi([]float64{1.32, 1.27, 1.42, 1.47, 1.42, 1.45, 1.59}, 3, 2, 2)
	want := [][]float64{{0.3268608414239478, 0.32038834951456274}, {0.5795418491021003, 0.7058823529411765}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Tsi Failed!")
	}
}

func TestZigZag(t *testing.T) {
	have := ind.ZigZag([][]float64{{10, 9}, {12, 10}, {14, 12}, {15, 13}, {16, 15}, {11, 10}, {18, 15}}, 0.25)
	want := []float64{9, 10.75, 12.5, 14.25, 16, 10, 18}
	if !assert.Equal(t, want, have) {
		t.Fatalf("ZigZag Failed!")
	}
}

func TestPsar(t *testing.T) {
	// v2.0.0: downtrend extreme tracks the LOW (data[i][1]), not high.
	// Fixture extends past the original last bug-quiet bar so the corrected branch is exercised.
	have := ind.Psar([][]float64{{82.15, 81.29}, {81.89, 80.64}, {83.03, 81.31}, {83.30, 82.65}, {83.85, 83.07}, {83.90, 83.11}, {83.33, 82.49}, {84.30, 82.3}, {84.84, 84.15}, {85, 84.11}, {75.9, 74.03}, {76.58, 75.39}, {76.98, 75.76}, {78, 77.17}, {70.87, 70.01}, {70.50, 69.50}, {69.00, 68.00}, {68.50, 67.50}}, 0.02, 0.2)
	want := []float64{81.29, 82.15, 80.64, 80.64, 80.7464, 80.932616, 81.17000672, 81.3884061824, 81.67956556416, 82.0588176964608, 85, 85, 84.7806, 84.565588, 84.35487624000001, 83.78108119040002, 82.92421631897602, 81.73027901345795}
	tu.Slice(t, have, want, "Psar")
}

func TestSupertrend(t *testing.T) {
	// v2.0.0: cascades from Atr fix (TR uses absolute values).
	have := ind.Supertrend([][]float64{{3, 2, 1}, {2, 2, 1}, {4, 3, 1}, {2, 2, 1}}, 3, 0.5)
	want := [][]float64{{3.5555555555555554, 1.4444444444444444}, {2.537037037037037, 0.462962962962963}}
	tu.Slice2D(t, have, want, "Supertrend")
}

func TestElderray(t *testing.T) {
	have := ind.Elderray([]float64{6, 5, 4, 7, 8, 9, 6, 8}, 7)
	want := [][]float64{{2.571428571428571, -2.428571428571429}, {2.2857142857142856, -2.7142857142857144}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Elderray Failed!")
	}
}

func TestHv(t *testing.T) {
	have := ind.Hv([]float64{7, 6, 5, 7, 8, 9, 7, 6, 5}, 8)
	want := []float64{0.6420403501307599, 0.6823595190987881}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Hv Failed!")
	}
}

func TestRvi(t *testing.T) {
	have := ind.Rvi([][]float64{{4, 6, 3, 3}, {3, 5, 2, 2}, {2, 5, 2, 4}, {4, 6, 4, 5}, {5, 7, 4, 4}, {4, 6, 3, 4}, {4, 7, 3, 5}, {5, 7, 5, 6}, {6, 8, 6, 6}, {6, 9, 5, 6}, {6, 8, 6, 7}, {7, 9, 5, 6}, {6, 7, 4, 5}, {5, 6, 5, 6}, {6, 8, 5, 5}, {5, 7, 2, 6}}, 8)
	want := []float64{0.29878048780487804, 0.21951219512195122, 0.1589403973509934, 0.16083916083916086, 0.09859154929577463, 0.05109489051094891}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Rvi Failed!")
	}
}

func TestRviSignal(t *testing.T) {
	have := ind.Rvi_signal([]float64{0.29878048780487804, 0.21951219512195122, 0.1589403973509934, 0.16083916083916086, 0.09859154929577463, 0.05109489051094891})
	want := []float64{0.2027541389316547, 0.1596104767996724, 0.12148278468863555}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Rvi Signal Failed!")
	}
}

func TestRsiDivergence(t *testing.T) {
	have := ind.Rsi_divergence([]float64{74, 83, 66, 78, 69, 70, 84, 73, 74, 73, 83}, 5, ind.Wrsi)
	want := []float64{0, 0, 1, 0, 1, 0}
	if !assert.Equal(t, want, have) {
		t.Fatalf("RSI Divergence Failed!")
	}
}

func TestDivergence(t *testing.T) {
	have := ind.Divergence([]float64{48, 34, 43, 54, 56, 64, 43}, []float64{76, 74, 43, 55, 34, 32, 45, 47})
	want := []float64{0, 0, 1, 1, 0, 1}
	tu.Slice(t, have, want, "Divergence")
}

func TestDivergenceDoesNotMutateInputs(t *testing.T) {
	d1 := []float64{48, 34, 43, 54, 56, 64, 43}
	d2 := []float64{76, 74, 43, 55, 34, 32, 45, 47}
	s1 := append([]float64(nil), d1...)
	s2 := append([]float64(nil), d2...)
	_ = ind.Divergence(d1, d2)
	if !assert.Equal(t, s1, d1) {
		t.Fatalf("Divergence mutated data1")
	}
	if !assert.Equal(t, s2, d2) {
		t.Fatalf("Divergence mutated data2")
	}
}

func TestHalfTrend(t *testing.T) {
	// v2.0.0: cascades from Atr fix; loop now `i < len(data)` (6 entries instead of 7).
	have := ind.HalfTrend([][]float64{{100, 97, 90}, {101, 98, 94}, {103, 96, 92}, {106, 100, 95}, {110, 101, 100}, {112, 110, 105}, {110, 100, 90}, {103, 100, 97}, {95, 90, 85}, {94, 80, 80}, {90, 82, 81}, {85, 80, 70}}, 6, 3, 2)
	want := [][]interface{}{{110.13773148148148, 100.0, 89.86226851851852, "long"}, {115.77674897119341, 105.0, 94.22325102880659, "long"}, {111.32021604938272, 100.0, 88.67978395061728, "long"}, {116.9335133744856, 105.0, 93.0664866255144, "long"}, {111.94302983539094, 100.0, 88.05697016460906, "long"}, {115.346579218107, 105.0, 94.653420781893, "long"}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("HalfTrend Failed!")
	}
}

// --- Phase 3b: ports from ta.js v2.0.0 --------------------------------------

func TestCross(t *testing.T) {
	have := ind.Cross([]float64{3, 4, 5, 4, 3}, []float64{4, 3, 2, 3, 4})
	want := []ind.CrossSignal{{Index: 1, Cross: true}, {Index: 4, Cross: false}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Cross Failed!")
	}
}

func TestMacdSignal(t *testing.T) {
	have := ind.MacdSignal([]float64{1, 2, 3, 4, 5, 6, 14, 8, 10, 11}, 3, 6, 3)
	want := []float64{2.107142857142857, 1.8380102040816328, 1.6534256559766765}
	tu.Slice(t, have, want, "MacdSignal")
}

func TestMacdBars(t *testing.T) {
	have := ind.MacdBars([]float64{1, 2, 3, 4, 5, 6, 14, 8, 10, 11}, 3, 6, 3)
	want := []float64{-0.28571428571428603, -0.26913265306122436, -0.18458454810495661}
	tu.Slice(t, have, want, "MacdBars")
}

func TestAdl(t *testing.T) {
	have := ind.Adl([][]float64{{6, 3, 2, 100}, {7, 5, 4, 200}, {8, 6, 5, 150}, {7, 5, 4, 180}})
	want := []float64{-50, -116.66666666666666, -166.66666666666666, -226.66666666666666}
	tu.Slice(t, have, want, "Adl")
}

func TestApo(t *testing.T) {
	have := ind.Apo([]float64{1, 2, 3, 4, 5, 6, 14}, 3, 6)
	want := []float64{1.5, 3}
	tu.Slice(t, have, want, "Apo")
}

func TestCci(t *testing.T) {
	have := ind.Cci([][]float64{{10, 9, 8}, {12, 11, 10}, {11, 10, 9}, {13, 12, 11}, {12, 11, 10}}, 3)
	want := []float64{0, 100.00000000000001, 0}
	tu.Slice(t, have, want, "Cci")
}

func TestCmf(t *testing.T) {
	have := ind.Cmf([][]float64{{6, 3, 2, 100}, {7, 5, 4, 200}, {8, 6, 5, 150}, {7, 5, 4, 180}, {9, 8, 7, 120}, {8, 6, 5, 90}}, 3)
	want := []float64{-0.37037037037037035, -0.3333333333333333, -0.24444444444444444, -0.23076923076923078}
	tu.Slice(t, have, want, "Cmf")
}

func TestDpo(t *testing.T) {
	have := ind.Dpo([]float64{9, 10, 11, 12, 11, 13, 12, 14, 13, 15}, 4)
	want := []float64{-1.5, -1, -0.75, 0, -1.5, 0, -1.5}
	tu.Slice(t, have, want, "Dpo")
}

func TestPdm(t *testing.T) {
	have := ind.Pdm([][]float64{{10, 9, 8}, {12, 11, 10}, {11, 10, 9}, {13, 12, 11}, {12, 11, 10}, {14, 13, 12}}, 2)
	want := []float64{1, 1.5, 0.75, 1.375}
	tu.Slice(t, have, want, "Pdm")
}

func TestMdm(t *testing.T) {
	have := ind.Mdm([][]float64{{10, 9, 8}, {12, 11, 10}, {11, 10, 9}, {13, 12, 11}, {12, 11, 10}, {14, 13, 12}}, 2)
	want := []float64{0.5, 0.25, 0.625, 0.3125}
	tu.Slice(t, have, want, "Mdm")
}

func TestPdi(t *testing.T) {
	have := ind.Pdi([][]float64{{10, 9, 8}, {12, 11, 10}, {11, 10, 9}, {13, 12, 11}, {12, 11, 10}, {14, 13, 12}}, 2)
	want := []float64{40, 54.54545454545455, 31.57894736842105, 51.16279069767442}
	tu.Slice(t, have, want, "Pdi")
}

func TestMdi(t *testing.T) {
	have := ind.Mdi([][]float64{{10, 9, 8}, {12, 11, 10}, {11, 10, 9}, {13, 12, 11}, {12, 11, 10}, {14, 13, 12}}, 2)
	want := []float64{20, 9.090909090909092, 26.31578947368421, 11.627906976744185}
	tu.Slice(t, have, want, "Mdi")
}

func TestDx(t *testing.T) {
	have := ind.Dx([][]float64{{10, 9, 8}, {12, 11, 10}, {11, 10, 9}, {13, 12, 11}, {12, 11, 10}, {14, 13, 12}}, 2)
	want := []float64{33.333333333333336, 71.42857142857142, 9.090909090909092, 62.96296296296297}
	tu.Slice(t, have, want, "Dx")
}

func TestAdx(t *testing.T) {
	have := ind.Adx([][]float64{{10, 9, 8}, {12, 11, 10}, {11, 10, 9}, {13, 12, 11}, {12, 11, 10}, {14, 13, 12}}, 2)
	want := []float64{52.38095238095238, 30.735930735930737, 46.84944684944685}
	tu.Slice(t, have, want, "Adx")
}

func TestAdxr(t *testing.T) {
	have := ind.Adxr([][]float64{{10, 9, 8}, {12, 11, 10}, {11, 10, 9}, {13, 12, 11}, {12, 11, 10}, {14, 13, 12}}, 2)
	want := []float64{41.55844155844156, 38.79268879268879}
	tu.Slice(t, have, want, "Adxr")
}

func TestEmv(t *testing.T) {
	have := ind.Emv([][]float64{{10, 8, 100}, {11, 9, 150}, {12, 10, 200}, {13, 11, 180}, {12, 10, 120}, {14, 12, 90}})
	want := []float64{133.33333333333334, 100, 111.11111111111111, -166.66666666666666, 444.44444444444446}
	tu.Slice(t, have, want, "Emv")
}

func TestKdj(t *testing.T) {
	have := ind.Kdj([][]float64{{3, 2, 1}, {2, 2, 1}, {4, 3, 1}, {2, 2, 1}}, 2, 1, 1)
	want := [][]float64{{50, 50, 50}, {66.66666666666667, 66.66666666666667, 66.66666666666666}, {33.333333333333336, 33.333333333333336, 33.33333333333333}}
	tu.Slice2D(t, have, want, "Kdj")
}

func TestMass(t *testing.T) {
	data := [][]float64{{10, 8}, {12, 10}, {14, 12}, {13, 11}, {15, 13}, {17, 15}, {16, 14}, {18, 16}, {20, 18}, {19, 17}, {21, 19}, {23, 21}, {22, 20}, {24, 22}, {26, 24}, {25, 23}, {27, 25}, {29, 27}, {28, 26}, {30, 28}, {32, 30}}
	have := ind.Mass(data, 4)
	want := []float64{4, 4}
	tu.Slice(t, have, want, "Mass")
}

func TestNatr(t *testing.T) {
	have := ind.Natr([][]float64{{10, 9, 8}, {12, 11, 10}, {11, 10, 9}, {13, 12, 11}, {12, 11, 10}, {14, 13, 12}}, 2)
	want := []float64{22.727272727272727, 22.5, 21.875, 21.022727272727273, 20.432692307692307}
	tu.Slice(t, have, want, "Natr")
}

func TestNvi(t *testing.T) {
	have := ind.Nvi([][]float64{{9, 100}, {10, 150}, {11, 200}, {12, 180}, {11, 120}, {13, 90}})
	want := []float64{1000, 1000, 1000, 1090.909090909091, 1000, 1181.8181818181818}
	tu.Slice(t, have, want, "Nvi")
}

func TestPvi(t *testing.T) {
	have := ind.Pvi([][]float64{{9, 100}, {10, 150}, {11, 200}, {12, 180}, {11, 120}, {13, 90}})
	want := []float64{1000, 1111.111111111111, 1222.2222222222222, 1222.2222222222222, 1222.2222222222222, 1222.2222222222222}
	tu.Slice(t, have, want, "Pvi")
}

func TestPpo(t *testing.T) {
	have := ind.Ppo([]float64{1, 2, 3, 4, 5, 6, 14}, 3, 6)
	want := []float64{42.857142857142854, 46.15384615384615}
	tu.Slice(t, have, want, "Ppo")
}

func TestStochRsi(t *testing.T) {
	have := ind.StochRsi([]float64{10, 11, 12, 11, 10, 11, 12}, 2, 2, 2, 2)
	want := [][]float64{{50, 25}, {100, 75}}
	tu.Slice2D(t, have, want, "StochRsi")
}

func TestTrix(t *testing.T) {
	have := ind.Trix([]float64{1, 2, 3, 4, 5, 6, 10, 8, 11}, 3)
	want := []float64{29.41176470588235, 23.295454545454547}
	tu.Slice(t, have, want, "Trix")
}

func TestUlcer(t *testing.T) {
	have := ind.Ulcer([]float64{10, 12, 11, 8, 9, 7, 11, 13}, 3)
	want := []float64{4.811252243246882, 19.837301190396804, 22.443514786772667, 25.40024136369333, 16.577149663422453, 12.830005981991683}
	tu.Slice(t, have, want, "Ulcer")
}

func TestVortex(t *testing.T) {
	have := ind.Vortex([][]float64{{10, 9, 8}, {12, 11, 10}, {11, 10, 9}, {13, 12, 11}, {15, 14, 13}}, 2)
	want := [][]float64{{1, 0.6}, {1, 0.6}, {1.3333333333333333, 0}}
	tu.Slice2D(t, have, want, "Vortex")
}
