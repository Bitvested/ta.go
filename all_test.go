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
	want := []float64{100, 71.42857142857143}
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
func Fisher(t *testing.T) {
	have := ta.Fisher([]float64{8,6,8,9,7,8,9,8,7,8,6,7}, 9);
	want := [][]float64{{-0.20692076425551026, 0.11044691579009712},{-0.3930108381942109, -0.20692076425551026}};
	if !assert.Equal(t, want, have) {
		t.Fatalf("Fisher Failed!")
	}
}
func Ar(t *testing.T) {
	have := ta.Ar([]float64{0.02, -0.01, 0.03, 0.05, -0.03}, 3);
	want := []float64{0.03667479679633267, -0.053301281310417566};
	if !assert.Equal(t, want, have) {
		t.Fatalf("Ar Failed!");
	}
}
