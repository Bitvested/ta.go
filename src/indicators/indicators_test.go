package indicators_test

import (
	"testing"

	tu "github.com/Bitvested/ta.go/internal/testutil"
	ind "github.com/Bitvested/ta.go/src/indicators"
	"github.com/stretchr/testify/assert"
)

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
	// v2.0.0: cascades from Smma's Wilder recurrence + length-N+1 output.
	have := ind.Alligator([]float64{8, 7, 8, 9, 7, 8, 9, 6, 7, 8, 6, 8, 10, 8, 7, 9, 8, 7, 9, 6, 7, 9}, 13, 8, 5, 8, 5, 3)
	want := [][]float64{{7.726445152480656, 7.813665029359981, 7.6378735707815935}, {7.78698224852071, 7.929902890697122, 8.047341963476992}, {7.769230769230769, 7.919889017939568, 7.809177454346241}}
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
	if !assert.Equal(t, want, have) {
		t.Fatalf("Divergence Failed!")
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
