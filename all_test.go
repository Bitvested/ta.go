package ta_test

import (
	ta "github.com/Bitvested/ta.go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMedian(t *testing.T) {
	have := ta.Median([]float64{4, 6, 3, 1, 2, 5}, 4)
	want := []float64{3, 2, 2}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Median Failed!")
	}
}
func TestNormalize(t *testing.T) {
	have := ta.Normalize([]float64{5, 4, 9, 4}, 0.1)
	want := []float64{0.2222222222222222, 0.06349206349206349, 0.8571428571428571, 0.06349206349206349}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Normalize Failed!")
	}
}
func TestDenormalize(t *testing.T) {
	have := ta.Denormalize([]float64{0.2222222222222222, 0.06349206349206349, 0.8571428571428571, 0.06349206349206349, 0.4444444444444444}, []float64{5, 4, 9, 4}, 0.1)
	want := []float64{5, 4, 9, 4, 6.4}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Denormalize Failed!")
	}
}
func TestStandardize(t *testing.T) {
	have := ta.Standardize([]float64{6, 4, 6, 8, 6})
	want := []float64{0, -1.5811388300841895, 0, 1.5811388300841895, 0}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Standardize Failed!")
	}
}
func TestMad(t *testing.T) {
	have := ta.Mad([]float64{3, 7, 5, 4, 3, 8, 9}, 6)
	want := []float64{1, 2}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Mad Failed!")
	}
}
func TestAad(t *testing.T) {
	have := ta.Aad([]float64{4, 6, 8, 6, 8, 9, 10, 11}, 7)
	want := []float64{1.6734693877551021, 1.469387755102041}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Aad Failed!")
	}
}
func TestSma(t *testing.T) {
	have := ta.Sma([]float64{1, 2, 3, 4, 5, 6, 10}, 6)
	want := []float64{3.5, 5}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Sma Failed!")
	}
}
func TestSsd(t *testing.T) {
	have := ta.Ssd([]float64{7, 6, 5, 7, 9, 8, 3, 5, 4}, 7)
	want := []float64{4.869731585445518, 4.9856938190329, 5.3718844791323335}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Ssd Failed!")
	}
}
func TestRsi(t *testing.T) {
	have := ta.Rsi([]float64{1, 2, 3, 4, 5, 6, 7, 5}, 6)
	want := []float64{100, 100, 66.66666666666666}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Rsi Failed!")
	}
}
func TestWsma(t *testing.T) {
	have := ta.Wsma([]float64{1, 2, 3, 4, 5, 6, 10}, 6)
	want := []float64{3.5, 4.583333333333333}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Wsma Failed!")
	}
}
func TestWrsi(t *testing.T) {
	have := ta.Wrsi([]float64{1, 2, 3, 4, 5, 6, 7, 5, 6}, 6)
	want := []float64{100, 71.42857142857143, 75.60975609756098}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Wrsi Failed!")
	}
}
func TestEma(t *testing.T) {
	have := ta.Ema([]float64{1, 2, 3, 4, 5, 6, 10}, 6)
	want := []float64{3.5, 5.357142857142857}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Ema Failed!")
	}
}
func TestSmma(t *testing.T) {
	have := ta.Smma([]float64{1, 2, 3, 4, 5, 6, 10}, 5)
	want := []float64{3.4, 4.92}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Smma Failed!")
	}
}
func TestWma(t *testing.T) {
	have := ta.Wma([]float64{69, 68, 66, 70, 68}, 4)
	want := []float64{68.3, 68.2}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Wma Failed!")
	}
}
func TestPwma(t *testing.T) {
	have := ta.Pwma([]float64{17, 26, 23, 29, 20}, 4)
	want := []float64{24.090909090909093, 25.18181818181818}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Pwma Failed!")
	}
}
func TestHwma(t *testing.T) {
	have := ta.Hwma([]float64{54, 51, 86, 42, 47}, 4)
	want := []float64{56.2, 55}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Hwma Failed!")
	}
}
func TestVwma(t *testing.T) {
	have := ta.Vwma([][]float64{{1, 59}, {1.1, 82}, {1.21, 27}, {1.42, 73}, {1.32, 42}}, 4)
	want := []float64{1.184771784232365, 1.258794642857143}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Vwma Failed!")
	}
}
func TestLsma(t *testing.T) {
	have := ta.Lsma([]float64{5, 6, 6, 3, 4, 6, 7}, 6)
	want := []float64{4.714285714285714, 5.761904761904762}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Lsma Failed!")
	}
}
func TestHull(t *testing.T) {
	have := ta.Hull([]float64{6, 7, 5, 6, 7, 4, 5, 7}, 6)
	want := []float64{4.761904761904762, 5.476190476190476}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Hull Failed!")
	}
}
func TestMacd(t *testing.T) {
	have := ta.Macd([]float64{1, 2, 3, 4, 5, 6, 14}, 3, 6)
	want := []float64{1.5, 3}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Macd Failed!")
	}
}
func TestStd(t *testing.T) {
	have := ta.Std([]float64{1, 2, 3}, 3)
	want := 0.816496580927726
	if !assert.Equal(t, want, have) {
		t.Fatalf("Std Failed!")
	}
}
func TestNormsinv(t *testing.T) {
	have := ta.Normsinv(0.4732)
	want := -0.06722824471054376
	if !assert.Equal(t, want, have) {
		t.Fatalf("Normsinv Failed!")
	}
}
func TestCor(t *testing.T) {
	have := ta.Cor([]float64{1, 2, 3, 4, 5, 2}, []float64{1, 3, 2, 4, 6, 3})
	want := 0.8808929232684737
	if !assert.Equal(t, want, have) {
		t.Fatalf("Cor Failed!")
	}
}
func TestDif(t *testing.T) {
	have := ta.Dif(0.75, 0.5)
	want := 0.5
	if !assert.Equal(t, want, have) {
		t.Fatalf("Dif Failed!")
	}
}
func TestDrawdown(t *testing.T) {
	have := ta.Drawdown([]float64{1, 2, 3, 4, 2, 3})
	want := -0.5
	if !assert.Equal(t, want, have) {
		t.Fatalf("Drawdown Failed!")
	}
}
func TestAroonUp(t *testing.T) {
	have := ta.AroonUp([]float64{5, 4, 5, 2}, 3)
	want := []float64{100, 50}
	if !assert.Equal(t, want, have) {
		t.Fatalf("AroonUp Failed!")
	}
}
func TestAroonDown(t *testing.T) {
	have := ta.AroonDown([]float64{2, 5, 4, 5}, 3)
	want := []float64{0, 50}
	if !assert.Equal(t, want, have) {
		t.Fatalf("AroonDown Failed!")
	}
}
func TestAroonOsc(t *testing.T) {
	have := ta.AroonOsc([]float64{2, 5, 4, 5}, 3)
	want := []float64{50, 50}
	if !assert.Equal(t, want, have) {
		t.Fatalf("AroonOsc Failed!")
	}
}
func TestMfi(t *testing.T) {
	have := ta.Mfi([][]float64{{19, 13}, {14, 38}, {21, 25}, {32, 17}}, 3)
	want := []float64{41.53846153846154, 45.578231292517}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Mfi Failed!")
	}
}
func TestRoc(t *testing.T) {
	have := ta.Roc([]float64{1, 2, 3, 4}, 3)
	want := []float64{2, 1}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Roc Failed!")
	}
}
func TestCop(t *testing.T) {
	have := ta.Cop([]float64{3, 4, 5, 3, 4, 5, 6, 4, 7, 5, 4, 7, 5}, 4, 6, 5)
	want := []float64{0.3755555555555556, 0.23666666666666666}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Cop Failed!")
	}
}
func TestKst(t *testing.T) {
	have := ta.Kst([]float64{8, 6, 7, 6, 8, 9, 7, 5, 6, 7, 6, 8, 6, 7, 6, 8, 9, 9, 8, 6, 4, 6, 5, 6, 7, 8, 9}, 5, 7, 10, 15, 5, 5, 5, 7, 4)
	want := [][]float64{{-0.6828231292517006, -0.5174886621315192}, {-0.2939342403628118, -0.5786281179138322}, {0.3517800453514739, -0.35968820861678}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Kst Failed!")
	}
}
func TestObv(t *testing.T) {
	have := ta.Obv([][]float64{{25200, 10}, {30000, 10.15}, {25600, 10.17}, {32000, 10.13}})
	want := []float64{0, 30000, 55600, 23600}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Obv Failed!")
	}
}
func TestVwap(t *testing.T) {
	have := ta.Vwap([][]float64{{127.21, 89329}, {127.17, 16137}, {127.16, 23945}}, 2)
	want := []float64{127.20387973375304, 127.16402599670675}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Vwap Failed!")
	}
}
func TestMom(t *testing.T) {
	have := ta.Mom([]float64{1, 1.1, 1.2, 1.24, 1.34}, 4, false)
	want := []float64{0.24, 0.24}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Mom Failed!")
	}
}
func TestMomOsc(t *testing.T) {
	have := ta.MomOsc([]float64{1, 1.2, 1.3, 1.3, 1.2, 1.4}, 4)
	want := []float64{31.57894736842105, -31.57894736842105, -28.20512820512821}
	if !assert.Equal(t, want, have) {
		t.Fatalf("MomOsc Failed!")
	}
}
func TestHa(t *testing.T) {
	have := ta.Ha([][]float64{{3, 4, 2, 3}, {3, 6, 3, 5}, {5, 5, 2, 3}})
	want := [][]float64{{3, 4, 2, 3}, {3, 6, 3, 4.25}, {3.625, 5, 2, 3.75}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Ha Failed!")
	}
}
func TestRen(t *testing.T) {
	have := ta.Ren([][]float64{{8, 6}, {9, 7}, {9, 8}, {13, 10}}, 2)
	want := [][]float64{{8, 10, 8, 10}, {10, 12, 10, 12}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Ren Failed!")
	}
}
func TestTsi(t *testing.T) {
	have := ta.Tsi([]float64{1.32, 1.27, 1.42, 1.47, 1.42, 1.45, 1.59}, 3, 2, 2)
	want := [][]float64{{0.3268608414239478, 0.32038834951456274}, {0.5795418491021003, 0.7058823529411765}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Tsi Failed!")
	}
}
func TestBop(t *testing.T) {
	have := ta.Bop([][]float64{{4, 5, 4, 5}, {5, 6, 5, 6}, {6, 8, 5, 6}}, 2)
	want := []float64{1, 0.5}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Bop Failed!")
	}
}
func TestFi(t *testing.T) {
	have := ta.Fi([][]float64{{1.4, 200}, {1.5, 240}, {1.1, 300}, {1.2, 240}, {1.5, 400}}, 4)
	want := []float64{12.00000000000001}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Fi Failed!")
	}
}
func TestAsi(t *testing.T) {
	have := ta.Asi([][]float64{{7, 6, 4}, {9, 7, 5}, {9, 8, 6}})
	want := []float64{0, -12.5}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Asi Failed!")
	}
}
func TestAo(t *testing.T) {
	have := ta.Ao([][]float64{{6, 5}, {8, 6}, {7, 4}, {6, 5}, {7, 6}, {9, 8}}, 2, 5)
	want := []float64{0.0, 0.9000000000000004}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Ao Failed!")
	}
}
func TestPr(t *testing.T) {
	have := ta.Pr([]float64{2, 1, 3, 1, 2}, 4)
	want := []float64{-100, -50}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Pr Failed!")
	}
}
func TestDon(t *testing.T) {
	have := ta.Don([][]float64{{6, 2}, {5, 2}, {5, 3}, {6, 3}, {7, 4}, {6, 3}}, 5)
	want := [][]float64{{7, 4.5, 2}, {7, 4.5, 2}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Don Failed!")
	}
}
func TestIchimoku(t *testing.T) {
	have := ta.Ichimoku([][]float64{{6, 3, 2}, {5, 4, 2}, {5, 4, 3}, {6, 4, 3}, {7, 6, 4}, {6, 5, 3}, {7, 6, 5}, {7, 5, 3}, {8, 6, 5}, {9, 7, 6}, {8, 7, 6}, {7, 5, 5}, {6, 5, 4}, {6, 5, 3}, {6, 3, 2}, {5, 4, 2}}, 2, 4, 6, 4)
	want := [][]float64{{7, 6, 10.5, 6, 5}, {7.5, 6, 7.5, 5.5, 6}, {6.5, 7, 8, 5, 5}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Ichimoku Failed!")
	}
}
func TestStoch(t *testing.T) {
	have := ta.Stoch([][]float64{{3, 2, 1}, {2, 2, 1}, {4, 3, 1}, {2, 2, 1}}, 2, 1, 1)
	want := [][]float64{{66.66666666666667, 66.66666666666667}, {33.333333333333336, 33.333333333333336}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Stoch Failed!")
	}
}
func TestAtr(t *testing.T) {
	have := ta.Atr([][]float64{{3, 2, 1}, {2, 2, 1}, {4, 3, 1}, {2, 2, 1}}, 3)
	want := []float64{2.0, 1.6666666666666667, 2.111111111111111, 1.7407407407407407}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Atr Failed!")
	}
}
func TestKama(t *testing.T) {
	have := ta.Kama([]float64{8, 7, 8, 9, 7, 9}, 2, 4, 8)
	want := []float64{8, 8.64, 8.377600000000001, 8.377600000000001}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Kama Failed!")
	}
}
func TestBands(t *testing.T) {
	have := ta.Bands([]float64{1, 2, 3, 4, 5, 6}, 5, 2)
	want := [][]float64{{5.82842712474619, 3.0, 0.1715728752538097}, {6.82842712474619, 4.0, 1.1715728752538097}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Bands Failed!")
	}
}
func TestBandwidth(t *testing.T) {
	have := ta.Bandwidth([]float64{1, 2, 3, 4, 5, 6}, 5, 2)
	want := []float64{1.8856180831641265, 1.414213562373095}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Bandwidth Failed!")
	}
}
func TestKeltner(t *testing.T) {
	have := ta.Keltner([][]float64{{3, 2, 1}, {2, 2, 1}, {4, 3, 1}, {2, 2, 1}, {3, 3, 1}}, 5, 1)
	want := [][]float64{{3.932266666666667, 2.066666666666667, 0.20106666666666695}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Keltner Failed!")
	}
}
func TestVariance(t *testing.T) {
	have := ta.Variance([]float64{6, 7, 2, 3, 5, 8, 6, 2}, 7)
	want := []float64{3.918367346938776, 5.061224489795919}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Variance Failed!")
	}
}
func TestPercentile(t *testing.T) {
	have := ta.Percentile([][]float64{{6, 4, 7}, {5, 3, 6}, {7, 5, 8}}, 0.5)
	want := []float64{6, 4, 7}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Percentile Failed!")
	}
}
func TestAlligator(t *testing.T) {
	have := ta.Alligator([]float64{8, 7, 8, 9, 7, 8, 9, 6, 7, 8, 6, 8, 10, 8, 7, 9, 8, 7, 9, 6, 7, 9}, 13, 8, 5, 8, 5, 3)
	want := [][]float64{{7.217569412835686, 6.985078985569999, 6.456171046541722}, {7.171597633136094, 7.119368115440011, 6.719144767291392}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Alligator Failed!")
	}
}
func TestGator(t *testing.T) {
	have := ta.Gator([]float64{8, 7, 8, 9, 7, 8, 9, 6, 7, 8, 6, 8, 10, 8, 7, 9, 8, 7, 9, 6, 7, 9}, 13, 8, 5, 8, 5, 3)
	want := [][]float64{{0.23249042726568714, -0.5289079390282767}, {0.05222951769608297, -0.4002233481486188}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Gator Failed!")
	}
}
func TestFractals(t *testing.T) {
	have := ta.Fractals([][]float64{{7, 6}, {8, 6}, {9, 6}, {8, 5}, {7, 4}, {6, 3}, {7, 4}, {8, 5}})
	want := [][]bool{{false, false}, {false, false}, {true, false}, {false, false}, {false, false}, {false, true}, {false, false}, {false, false}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Fractals Failed!")
	}
}
func TestChaikinOsc(t *testing.T) {
	have := ta.ChaikinOsc([][]float64{{2, 3, 4, 6}, {5, 5, 5, 4}, {5, 4, 3, 7}, {4, 3, 3, 4}, {6, 5, 4, 6}, {7, 4, 3, 6}}, 2, 4)
	want := []float64{-1.6666666666666665, -0.28888888888888886, -0.7362962962962962}
	if !assert.Equal(t, want, have) {
		t.Fatalf("ChaikinOsc Failed!")
	}
}
func TestEnvelope(t *testing.T) {
	have := ta.Envelope([]float64{6, 7, 8, 7, 6, 7, 8, 7, 8, 7, 8, 7, 8}, 11, 0.05)
	want := [][]float64{{7.540909090909091, 7.181818181818182, 6.822727272727272}, {7.636363636363637, 7.2727272727272725, 6.909090909090908}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Envelope Failed!")
	}
}
func TestRecentHigh(t *testing.T) {
	have := ta.RecentHigh([]float64{4, 5, 6, 7, 8, 9, 8, 7, 8, 9, 10, 3, 2, 1}, 3)
	want := ta.RecentHighLow(10, 10)
	if !assert.Equal(t, want, have) {
		t.Fatalf("Recent High Failed!")
	}
}
func TestRecentLow(t *testing.T) {
	have := ta.RecentLow([]float64{1, 4, 5, 6, 4, 3, 2, 3, 4, 3, 5, 7, 8, 8, 5}, 4)
	want := ta.RecentHighLow(6, 2)
	if !assert.Equal(t, want, have) {
		t.Fatalf("Recent Low Failed!")
	}
}
func TestFib(t *testing.T) {
	have := ta.Fib(1, 2)
	want := []float64{1, 1.236, 1.3820000000000001, 1.5, 1.6179999999999999, 1.786, 2, 2.6180000000000003, 3.618, 4.618, 5.236}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Fib Failed!")
	}
}
func TestAc(t *testing.T) {
	have := ta.Ac([][]float64{{6, 5}, {8, 6}, {7, 4}, {6, 5}, {7, 6}, {9, 8}}, 2, 4)
	want := []float64{0.125, 0.5625}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Ac Failed!")
	}
}
func TestSupport(t *testing.T) {
	hl := ta.RecentLow([]float64{4, 3, 2, 5, 7, 6, 5, 4, 7, 8, 5, 4, 6, 7, 5}, 20)
	have := ta.Support([]float64{4, 3, 2, 5, 7, 6, 5, 4, 7, 8, 5, 4, 6, 7, 5}, hl)
	want := ta.Line(2, 0.2222222222222222, 2)
	if !assert.Equal(t, want, have) {
		t.Fatalf("Support Failed!")
	}
}
func TestResistance(t *testing.T) {
	hl := ta.RecentHigh([]float64{5, 7, 5, 5, 4, 6, 5, 4, 6, 5, 4, 3, 2, 4, 3, 2, 1}, 20)
	have := ta.Resistance([]float64{5, 7, 5, 5, 4, 6, 5, 4, 6, 5, 4, 3, 2, 4, 3, 2, 1}, hl)
	want := ta.Line(1, -0.14285714285714285, 7)
	if !assert.Equal(t, want, have) {
		t.Fatalf("Resistance Failed!")
	}
	have2 := ta.LineCalc(4, have)
	want2 := 6.428571428571429
	if !assert.Equal(t, want2, have2) {
		t.Fatalf("LineCalc Failed!")
	}
}
func TestEr(t *testing.T) {
	have := ta.Er([]float64{0.02, -0.01, 0.03, 0.05, -0.03});
	want := 0.011934565489708282;
	if !assert.Equal(t, want, have) {
		t.Fatalf("Er Failed!")
	}
}
func TestFisher(t *testing.T) {
	have := ta.Fisher([]float64{8,6,8,9,7,8,9,8,7,8,6,7}, 9);
	want := [][]float64{[]float64{-0.3181527308248248, -0.1104469157900972}, []float64{-0.4491605923903109, -0.3181527308248248}, []float64{-0.6158050492824234, -0.4491605923903109}};
	if !assert.Equal(t, want, have) {
		t.Fatalf("Fisher Failed!")
	}
}
func TestAr(t *testing.T) {
	have := ta.Ar([]float64{0.02, -0.01, 0.03, 0.05, -0.03}, 3);
	want := []float64{0.03667479679633267, -0.053301281310417566};
	if !assert.Equal(t, want, have) {
		t.Fatalf("Ar Failed!");
	}
}
func TestAvgwin(t *testing.T) {
	have := ta.Avgwin([]float64{0.01,0.02,-0.01,-0.03,-0.015,0.005});
	want := 0.011666666666666665;
	if !assert.Equal(t, want, have) {
		t.Fatalf("Avgwin Failed!");
	}
}
func TestAvgloss(t *testing.T) {
	have := ta.Avgloss([]float64{0.01,0.02,-0.01,-0.03,-0.015,0.005});
	want := -0.018333333333333333;
	if !assert.Equal(t, want, have) {
		t.Fatalf("Avgloss Failed!");
	}
}
func TestSe(t *testing.T) {
	have := ta.Se([]float64{34,54,45,43,57,38,49}, 10);
	want := 2.4243661069253055;
	if !assert.Equal(t, want, have) {
		t.Fatalf("Se Failed!");
	}
}
func TestKelly(t *testing.T) {
	have := ta.Kelly([]float64{0.01,0.02,-0.01,-0.03,-0.015,0.045,0.005});
	want := 0.14434748152632182;
	if !assert.Equal(t, want, have) {
		t.Fatalf("Kelly Failed!");
	}
}
func TestZscore(t *testing.T) {
	have := ta.Zscore([]float64{34,54,45,43,57,38,49}, 5);
	want := []float64{1.2664106627730554, -1.3314928442246727, 0.4078462733398033}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Zscore Failed!");
	}
}
func TestNormalizePair(t *testing.T) {
	have := ta.NormalizePair([]float64{10,12,11,13}, []float64{100,130,100,140});
	want := [][]float64{{55,55},{66,71.5},{60.5,54.99999999999999},{71.5,76.99999999999999}};
	if !assert.Equal(t, want, have) {
		t.Fatalf("NormalizePair Failed!");
	}
}
func TestNormalizeFrom(t *testing.T) {
	have := ta.NormalizeFrom([]float64{8,12,10,11}, 100);
	want := []float64{100,150,125,137.5};
	if !assert.Equal(t, want, have) {
		t.Fatalf("NormalizeFrom Failed!");
	}
}
func TestLog(t *testing.T) {
	have := ta.Log([]float64{5, 14, 18, 28, 68, 103});
	want := []float64{1.6094379124341003, 2.639057329615259, 2.8903717578961645, 3.332204510175204, 4.219507705176107, 4.634728988229636};
	if !assert.Equal(t, want, have) {
		t.Fatalf("Log Failed!");
	}
}
func TestExp(t *testing.T) {
	have := ta.Exp([]float64{1.6, 2.63, 2.89, 3.33, 4.22, 4.63});
	want := []float64{4.953032424395115, 13.873769902129904, 17.993309601550315, 27.938341703236507, 68.03348428941966, 102.51406411049346};
	if !assert.Equal(t, want, have) {
		t.Fatalf("Exp Failed!");
	}
}
func TestHalfTrend(t *testing.T) {
	have := ta.HalfTrend([][]float64{{100,97,90},{101,98,94},{103,96,92},{106,100,95},{110,101,100},{112,110,105},{110,100,90},{103,100,97},{95,90,85},{94,80,80},{90,82,81},{85,80,70}}, 6, 3, 2);
	want := [][]interface{}{[]interface{}{110.13773148148148,100.0,89.86226851851852,"long"},[]interface{}{115.77674897119341, 105.0, 94.22325102880659, "long"},[]interface{}{111.32021604938272, 100.0, 88.67978395061728, "long"},[]interface{}{116.10018004115227, 105.0, 93.89981995884773, "long"},[]interface{}{111.2485853909465, 100.0, 88.7514146090535, "long"},[]interface{}{114.76787551440329, 105.0, 95.23212448559671, "long"},[]interface {}{114.99356995884773,100.0,85.00643004115227,"long"}};
	if !assert.Equal(t, want, have) {
		t.Fatalf("HalfTrend Failed!");
	}
}
func TestNcdf(t *testing.T) {
	have := ta.Ncdf(13, 10, 2);
	want := 0.9331737996110652;
	if !assert.Equal(t, want, have) {
		t.Fatalf("Ncdf Failed!");
	}
}
func TestZigZag(t *testing.T) {
	have := ta.ZigZag([][]float64{[]float64{10,9},[]float64{12,10},[]float64{14,12},[]float64{15,13},[]float64{16,15},[]float64{11,10},[]float64{18,15}}, 0.25);
	want := []float64{9, 10.75, 12.5, 14.25, 16, 10, 18};
	if !assert.Equal(t, want, have) {
		t.Fatalf("ZigZag Failed!");
	}
}
func TestPsar(t *testing.T) {
	have := ta.Psar([][]float64{[]float64{82.15,81.29},[]float64{81.89,80.64},[]float64{83.03,81.31},[]float64{83.30,82.65},[]float64{83.85,83.07},[]float64{83.90,83.11},[]float64{83.33,82.49},[]float64{84.30,82.3},[]float64{84.84,84.15},[]float64{85,84.11},[]float64{75.9,74.03},[]float64{76.58,75.39},[]float64{76.98,75.76},[]float64{78,77.17},[]float64{70.87,70.01}}, 0.02, 0.2);
	want := []float64{81.29,82.15,80.64,80.64,80.7464,80.932616,81.17000672,81.3884061824,81.67956556416,82.0588176964608,85,85,84.7806,84.565588,84.35487624000001};
	if !assert.Equal(t, want, have) {
		t.Fatalf("Psar Failed!");
	}
}
func TestFibbands(t *testing.T) {
	have := ta.Fibbands([][]float64{{1, 59}, {1.1, 82}, {1.21, 27}, {1.42, 73}, {1.32, 42}}, 4, 3);
	want := [][]float64{{1.6526058894448858,1.542197040614731,1.4738932612537028,1.4186888368386255,1.363484412423548,1.29518063306252,1.184771784232365,1.0743629354022102,1.0060591560411822,0.9508547316261047,0.8956503072110273,0.8273465278499992,0.7169376790198443},{1.6177775811703725,1.5330576077284503,1.4806460987347188,1.4382861120137578,1.3959261252927968,1.3435146162990652,1.258794642857143,1.1740746694152209,1.1216631604214893,1.0793031737005283,1.0369431869795673,0.9845316779858357,0.8998117045439136}};
	if !assert.Equal(t, want, have) {
		t.Fatalf("Fibbands Failed!");
	}
}
func TestMartingale(t *testing.T) {
	have := ta.Martingale([]float64{-1,-1,1,1,-1,-1}, 5, 200, 2);
	var want float64 = 20;
	if !assert.Equal(t, want, have) {
		t.Fatalf("Martingale Failed!");
	}
}
func TestAntimartingale(t *testing.T) {
	have := ta.Antimartingale([]float64{1,1,-1,-1,1,1}, 5, 200, 2);
	var want float64 = 20;
	if !assert.Equal(t, want, have) {
		t.Fatalf("Antimartingale Failed!");
	}
}
func TestPermutations(t *testing.T) {
	have := ta.Permutations([]float64{10,10,10});
	var want float64 = 1000;
	if !assert.Equal(t, want, have) {
		t.Fatalf("Permutations Failed!");
	}
}
func TestMse(t *testing.T) {
	have := ta.Mse([]float64{7,8,7,8,6,9}, []float64{6,8,8,9,6,8});
	want := 0.6666666666666666;
	if !assert.Equal(t, want, have) {
		t.Fatalf("MSE Failed!");
	}
}
func TestSum(t *testing.T) {
	have := ta.Sum([]float64{2,2,4,2});
	want := float64(10);
	if !assert.Equal(t, want, have) {
		t.Fatalf("Sum Failed!");
	}
}
func TestCum(t *testing.T) {
	have := ta.Cum([]float64{3,5,7,5,10}, 4);
	want := []float64{20,27};
	if !assert.Equal(t, want, have) {
		t.Fatalf("Cum Failed!");
	}
}
func TestSupertrend(t *testing.T) {
	have := ta.Supertrend([][]float64{{3,2,1},{2,2,1},{4,3,1},{2,2,1}}, 3, 0.5);
	want := [][]float64{{3.5555555555555554, 1.4444444444444444},{2.3703703703703702, 0.6296296296296297}};
	if !assert.Equal(t, want, have) {
		t.Fatalf("Supertrend Failed!");
	}
}
func TestCwma(t *testing.T) {
	have := ta.Cwma([]float64{69,68,66,70,68,69}, []float64{1,2,3,5,8});
	want := []float64{68.26315789473684, 68.52631578947368};
	if !assert.Equal(t, want, have) {
		t.Fatalf("Cwma Failed!");
	}
}
func TestVwwma(t *testing.T) {
	have := ta.Vwwma([][]float64{{1,59},{1.1,82},{1.21,27},{1.42,73},{1.32,42}}, 4);
	want := []float64{1.2618288590604028, 1.3160229445506693};
	if !assert.Equal(t, want, have) {
		t.Fatalf("Vwwma Failed!");
	}
}
func TestElderray(t *testing.T) {
	have := ta.Elderray([]float64{6,5,4,7,8,9,6,8}, 7);
	want := [][]float64{{2.571428571428571,-2.428571428571429},{2.2857142857142856,-2.7142857142857144}};
	if !assert.Equal(t, want, have) {
		t.Fatalf("Elderray Failed!");
	}
}
func TestHv(t *testing.T) {
	have := ta.Hv([]float64{7,6,5,7,8,9,7,6,5}, 8);
	want := []float64{0.6420403501307599,0.6823595190987881};
	if !assert.Equal(t, want, have) {
		t.Fatalf("Hv Failed!");
	}
}
func TestPvalue(t *testing.T) {
	have := ta.Pvalue(2.0, 4);
	want := 0.061018363939899845;
	if !assert.Equal(t, want, have) {
		t.Fatalf("Pvalue Failed!");
	}
}
func TestRvi(t *testing.T) {
	have := ta.Rvi([][]float64{{4,6,3,3}, {3,5,2,2}, {2,5,2,4}, {4,6,4,5}, {5,7,4,4}, {4,6,3,4}, {4,7,3,5}, {5,7,5,6}, {6,8,6,6}, {6,9,5,6}, {6,8,6,7}, {7,9,5,6},{6,7,4,5},{5,6,5,6},{6,8,5,5},{5,7,2,6}}, 8);
	want := []float64{0.29878048780487804,0.21951219512195122,0.1589403973509934,0.16083916083916086,0.09859154929577463,0.05109489051094891}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Rvi Failed!");
	}
}
func TestRviSignal(t *testing.T) {
	have := ta.Rvi_signal([]float64{0.29878048780487804,0.21951219512195122,0.1589403973509934,0.16083916083916086,0.09859154929577463,0.05109489051094891});
	want := []float64{0.2027541389316547,0.1596104767996724,0.12148278468863555};
	if !assert.Equal(t, want, have) {
		t.Fatalf("Rvi Signal Failed!")
	}
}
