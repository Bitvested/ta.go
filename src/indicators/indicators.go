// Package indicators contains trend, momentum, volatility, and trading
// indicators built on top of the ta.go moving-average and statistics packages.
package indicators

import (
	"math"
	"sort"

	ma "github.com/Bitvested/ta.go/src/moving_averages"
	stats "github.com/Bitvested/ta.go/src/statistics"
)

// --- internal helpers ---------------------------------------------------------

func maxf(v []float64) float64 {
	sort.Float64s(v)
	return v[len(v)-1]
}

func minf(v []float64) float64 {
	sort.Float64s(v)
	return v[0]
}

func intMax(v []int) int {
	sort.Ints(v)
	return v[len(v)-1]
}

func intMin(v []int) int {
	sort.Ints(v)
	return v[0]
}

func reverseFloats(data []float64) []float64 {
	var ret []float64
	for i := len(data) - 1; i >= 0; i-- {
		ret = append(ret, data[i])
	}
	return ret
}

func indexOf(data []float64, index float64) int {
	var ret int = -1
	for i := len(data) - 1; i >= 0; i-- {
		if data[i] == index {
			ret = i
			break
		}
	}
	return ret
}

func map2d(data [][]float64, key int) []float64 {
	var ret []float64
	for i := 0; i < len(data); i++ {
		ret = append(ret, data[i][key])
	}
	return ret
}

// --- exported indicators ------------------------------------------------------

// Rsi is the Wilder-smoothed Relative Strength Index (TA-Lib convention).
// Returns NaN when both gain and loss are zero (flat window) and 100 when
// loss is zero with positive gain.
//
// v2.0.0 perf: pre-allocated output, hoisted 1/l divisor, single delta
// branch per bar.
func Rsi(data []float64, l int) []float64 {
	n := len(data)
	if n <= l {
		return nil
	}
	var gain, loss float64
	for i := 1; i <= l; i++ {
		d := data[i] - data[i-1]
		if d > 0 {
			gain += d
		} else {
			loss -= d
		}
	}
	lf := float64(l)
	avgG := gain / lf
	avgL := loss / lf
	out := make([]float64, n-l)
	out[0] = rsiVal(avgG, avgL)
	for i := l + 1; i < n; i++ {
		d := data[i] - data[i-1]
		if d > 0 {
			avgG += (d - avgG) / lf
			avgL += (0 - avgL) / lf
		} else {
			avgG += (0 - avgG) / lf
			avgL += (-d - avgL) / lf
		}
		out[i-l] = rsiVal(avgG, avgL)
	}
	return out
}

func rsiVal(g, l float64) float64 {
	if g == 0 && l == 0 {
		return math.NaN()
	}
	if l == 0 {
		return 100
	}
	return 100 - 100/(1+g/l)
}

func Wrsi(data []float64, l int) []float64 {
	var arrsi []float64
	var u []float64
	var d []float64
	for i := 1; i < len(data); i++ {
		if data[i]-data[i-1] < 0 {
			d = append(d, math.Abs(data[i]-data[i-1]))
			u = append(u, 0)
		} else {
			d = append(d, 0)
			u = append(u, data[i]-data[i-1])
		}
	}
	d = ma.Wsma(d, l)
	u = ma.Wsma(u, l)
	for i := 0; i < len(d); i++ {
		arrsi = append(arrsi, 100-100/(1+(u[i]/d[i])))
	}
	return arrsi
}

func Macd(data []float64, l1 int, l2 int) []float64 {
	if l1 > l2 {
		l1, l2 = l2, l1
	}
	var ema []float64 = ma.Ema(data, l1)
	var emb []float64 = ma.Ema(data, l2)
	var macd []float64
	ema = ema[l2-l1:]
	for i := 0; i < len(emb); i++ {
		macd = append(macd, ema[i]-emb[i])
	}
	return macd
}

func AroonUp(data []float64, l int) []float64 {
	var aroon []float64
	for i := l; i <= len(data); i++ {
		pl := append([]float64(nil), data[i-l:i]...)
		hl := append([]float64(nil), data[i-l:i]...)
		sort.Float64s(hl)
		pl = reverseFloats(pl)
		index := indexOf(pl, hl[len(hl)-1])
		aroon = append(aroon, (100 * (float64(index)) / (float64(l) - 1)))
	}
	return aroon
}

func AroonDown(data []float64, l int) []float64 {
	var aroon []float64
	for i := l; i <= len(data); i++ {
		pl := append([]float64(nil), data[i-l:i]...)
		hl := append([]float64(nil), data[i-l:i]...)
		sort.Float64s(hl)
		index := indexOf(pl, hl[0])
		aroon = append(aroon, (100 * (float64(index)) / (float64(l) - 1)))
	}
	return aroon
}

// Mfi is the Money Flow Index. Input is OHLCV-style: each bar is
// [high, low, close, volume]. Output length = len(data) - length.
// Returns NaN when both flows are zero, 100 when only negative flow is zero.
func Mfi(data [][]float64, l int) []float64 {
	n := len(data)
	if n <= l {
		return nil
	}
	pos := make([]float64, n)
	neg := make([]float64, n)
	prevTp := (data[0][0] + data[0][1] + data[0][2]) / 3
	for i := 1; i < n; i++ {
		tp := (data[i][0] + data[i][1] + data[i][2]) / 3
		flow := tp * data[i][3]
		switch {
		case tp > prevTp:
			pos[i] = flow
		case tp < prevTp:
			neg[i] = flow
		}
		prevTp = tp
	}
	var sumPos, sumNeg float64
	for i := 1; i <= l; i++ {
		sumPos += pos[i]
		sumNeg += neg[i]
	}
	out := []float64{mfiVal(sumPos, sumNeg)}
	for i := l + 1; i < n; i++ {
		sumPos += pos[i] - pos[i-l]
		sumNeg += neg[i] - neg[i-l]
		out = append(out, mfiVal(sumPos, sumNeg))
	}
	return out
}

func mfiVal(p, n float64) float64 {
	if p == 0 && n == 0 {
		return math.NaN()
	}
	if n == 0 {
		return 100
	}
	return 100 * p / (p + n)
}

// Roc returns rate-of-change as a percentage (TA-Lib convention).
// Returns NaN when the reference value is zero.
func Roc(data []float64, l int) []float64 {
	n := len(data)
	if n <= l {
		return nil
	}
	out := make([]float64, 0, n-l)
	for i := l; i <= n; i++ {
		ref := data[i-l]
		if ref == 0 {
			out = append(out, math.NaN())
		} else {
			out = append(out, (data[i-1]-ref)/ref*100)
		}
	}
	return out
}

func Cop(data []float64, l1 int, l2 int, l3 int) []float64 {
	max := math.Max(float64(l1), float64(l2))
	var co []float64
	for i := int(max) + l3; i < len(data); i++ {
		r1 := append([]float64(nil), data[i-(int(max)+l3):i]...)
		r2 := r1[:]
		var tmp []float64
		r1 = Roc(r1, l1)
		r2 = Roc(r2, l2)
		r1 = r1[len(r1)-len(r2):]
		r2 = r2[len(r2)-len(r1):]
		for a := 0; a < len(r1); a++ {
			tmp = append(tmp, r1[a]+r2[a])
		}
		tmp = ma.Wma(tmp, l3)
		co = append(co, tmp[len(tmp)-1])
	}
	return co
}

func Kst(data []float64, r1 int, r2 int, r3 int, r4 int, s1 int, s2 int, s3 int, s4 int, sig int) [][]float64 {
	var ks []float64
	var fs [][]float64
	ms := (intMax([]int{r1, r2, r3, r4}) + intMax([]int{s4, s3, s2, s1}))
	for i := ms; i <= len(data); i++ {
		pl := append([]float64(nil), data[i-ms:i]...)
		rcma1 := Roc(pl, r1)
		rcma2 := Roc(pl, r2)
		rcma3 := Roc(pl, r3)
		rcma4 := Roc(pl, r4)
		rcma1 = ma.Sma(rcma1, s1)
		rcma2 = ma.Sma(rcma2, s2)
		rcma3 = ma.Sma(rcma3, s3)
		rcma4 = ma.Sma(rcma4, s4)
		ks = append(ks, rcma1[len(rcma1)-1]+rcma2[len(rcma2)-1]+rcma3[len(rcma3)-1]+rcma4[len(rcma4)-1])
	}
	sl := ma.Sma(ks, sig)
	ks = ks[len(ks)-len(sl):]
	for i := 0; i < len(sl); i++ {
		fs = append(fs, []float64{ks[i], sl[i]})
	}
	return fs
}

func Obv(data [][]float64) []float64 {
	var obv []float64
	obv = append(obv, 0)
	for i := 1; i < len(data); i++ {
		if data[i][1] > data[i-1][1] {
			obv = append(obv, obv[len(obv)-1]+data[i][0])
		}
		if data[i][1] < data[i-1][1] {
			obv = append(obv, obv[len(obv)-1]-data[i][0])
		}
		if data[i][1] == data[i-1][1] {
			obv = append(obv, obv[len(obv)-1])
		}
	}
	return obv
}

func Vwap(data [][]float64, l int) []float64 {
	var vwap []float64
	var weighted [][]float64
	for i := 0; i < len(data); i++ {
		weighted = append(weighted, []float64{data[i][0] * data[i][1], data[i][1]})
	}
	for i := l; i <= len(weighted); i++ {
		pl := append([][]float64(nil), weighted[i-l:i]...)
		var totalv float64
		var totalp float64
		for o := 0; o < len(pl); o++ {
			totalv += pl[o][1]
			totalp += pl[o][0]
		}
		vwap = append(vwap, totalp/totalv)
	}
	return vwap
}

// Mom is momentum. With p=false returns simple difference; with p=true
// returns percent (×100) and yields NaN when the reference value is zero.
func Mom(data []float64, l int, p bool) []float64 {
	n := len(data)
	if n < l {
		return nil
	}
	out := make([]float64, 0, n-l+1)
	for i := l - 1; i < n; i++ {
		ref := data[i-(l-1)]
		if p {
			if ref == 0 {
				out = append(out, math.NaN())
			} else {
				out = append(out, data[i]/ref*100)
			}
		} else {
			out = append(out, data[i]-ref)
		}
	}
	return out
}

func Tsi(data []float64, llen int, slen int, sig int) [][]float64 {
	var mom []float64
	var abs []float64
	var ts []float64
	var tsi [][]float64
	for i := 1; i < len(data); i++ {
		mom = append(mom, data[i]-data[i-1])
		abs = append(abs, math.Abs(data[i]-data[i-1]))
	}
	sma1 := ma.Ema(mom, llen)
	sma2 := ma.Ema(abs, llen)
	ema1 := ma.Ema(sma1, slen)
	ema2 := ma.Ema(sma2, slen)
	for i := 0; i < len(ema1); i++ {
		ts = append(ts, ema1[i]/ema2[i])
	}
	tma := ma.Ema(ts, sig)
	ts = ts[len(ts)-len(tma):]
	for i := 0; i < len(tma); i++ {
		tsi = append(tsi, []float64{tma[i], ts[i]})
	}
	return tsi
}

func Bop(data [][]float64, l int) []float64 {
	var bo []float64
	for i := 0; i < len(data); i++ {
		bo = append(bo, (data[i][3]-data[i][0])/(data[i][1]-data[i][2]))
	}
	bo = ma.Sma(bo, l)
	return bo
}

func Fi(data [][]float64, l int) []float64 {
	var ff []float64
	var pl []float64
	for i := 1; i < len(data); i++ {
		pl = append(pl, (data[i][0]-data[i-1][0])*data[i][1])
		if len(pl) >= l {
			vfi := ma.Ema(pl, l)
			ff = append(ff, vfi[len(vfi)-1])
			pl = pl[1:]
		}
	}
	return ff
}

func Asi(data [][]float64) []float64 {
	var a []float64
	for i := 1; i < len(data); i++ {
		c := data[i][1]
		cy := data[i-1][1]
		h := data[i][0]
		hy := data[i-1][0]
		l := data[i][2]
		ly := data[i-1][2]
		var k float64
		if hy-c > ly-c {
			k = hy - c
		} else {
			k = ly - c
		}
		o := data[i][0]
		oy := data[i-1][0]
		var r float64
		t := maxf([]float64{data[i][0], data[i-1][0]}) - minf([]float64{data[i][2], data[i-1][2]})
		if h-cy > l-cy && h-cy > h-l {
			r = h - cy - (l-cy)/2 + (cy-oy)/4
		}
		if l-cy > h-cy && l-cy > h-l {
			r = l - cy - (h-cy)/2 + (cy-oy)/4
		}
		if h-l > h-cy && h-l > l-cy {
			r = h - l + (cy-oy)/4
		}
		a = append(a, 50*((cy-c+(cy-oy)/2+(c-o)/2)/r)*k/t)
	}
	return a
}

func Alligator(data []float64, jl int, tl int, ll int, js int, ts int, ls int) [][]float64 {
	var ret [][]float64
	var jaw []float64 = ma.Smma(data, jl)
	var teeth []float64 = ma.Smma(data, tl)
	var lips []float64 = ma.Smma(data, ll)
	teeth = teeth[len(teeth)-len(jaw):]
	lips = lips[len(lips)-len(jaw):]
	for i := len(jaw) - 1; i >= 7; i-- {
		ret = append(ret, []float64{jaw[i-(js-1)], teeth[i-(ts-1)], lips[i-(ls-1)]})
	}
	return ret
}

// Pr is Williams %R. Returns NaN when the window range is zero (flat data).
func Pr(data []float64, l int) []float64 {
	n := len(data)
	if n < l {
		return nil
	}
	out := make([]float64, 0, n-l+1)
	for i := l; i <= n; i++ {
		hi := math.Inf(-1)
		lo := math.Inf(1)
		for j := i - l; j < i; j++ {
			if data[j] > hi {
				hi = data[j]
			}
			if data[j] < lo {
				lo = data[j]
			}
		}
		rng := hi - lo
		if rng == 0 {
			out = append(out, math.NaN())
		} else {
			out = append(out, (hi-data[i-1])/rng*-100)
		}
	}
	return out
}

// Stoch is the Stochastic oscillator (smoothK, smoothD). v2.0.0 uses a
// monotonic-deque high/low tracker plus running-sum smoothing for the K and
// D series; output is [%K, %D] pairs aligned at the tail. NaN is emitted
// when the high/low range is zero.
func Stoch(data [][]float64, l int, sd int, sk int) [][]float64 {
	if l < sd {
		l, sd = sd, l
	}
	if sk > sd {
		sk, sd = sd, sk
	}
	n := len(data)
	if n < l {
		return nil
	}
	qHi := make([]int, 0, n)
	qLo := make([]int, 0, n)
	rawBuf := make([]float64, sk)
	smBuf := make([]float64, sd)
	var rawSum, smSum float64
	rawCount, smCount := 0, 0
	rawIdx, smIdx := 0, 0
	out := make([][]float64, 0)
	for i := 0; i < n; i++ {
		hi, lo := data[i][0], data[i][2]
		for len(qHi) > 0 && data[qHi[len(qHi)-1]][0] <= hi {
			qHi = qHi[:len(qHi)-1]
		}
		qHi = append(qHi, i)
		if qHi[0] <= i-l {
			qHi = qHi[1:]
		}
		for len(qLo) > 0 && data[qLo[len(qLo)-1]][2] >= lo {
			qLo = qLo[:len(qLo)-1]
		}
		qLo = append(qLo, i)
		if qLo[0] <= i-l {
			qLo = qLo[1:]
		}
		if i < l-1 {
			continue
		}
		highd := data[qHi[0]][0]
		lowd := data[qLo[0]][2]
		rng := highd - lowd
		var rawK float64
		if rng == 0 {
			rawK = math.NaN()
		} else {
			rawK = 100 * (data[i][1] - lowd) / rng
		}
		if rawCount < sk {
			rawBuf[rawIdx] = rawK
			rawSum += rawK
			rawCount++
		} else {
			rawSum += rawK - rawBuf[rawIdx]
			rawBuf[rawIdx] = rawK
		}
		rawIdx = (rawIdx + 1) % sk
		if rawCount < sk {
			continue
		}
		smK := rawSum / float64(sk)
		if smCount < sd {
			smBuf[smIdx] = smK
			smSum += smK
			smCount++
		} else {
			smSum += smK - smBuf[smIdx]
			smBuf[smIdx] = smK
		}
		smIdx = (smIdx + 1) % sd
		if smCount < sd {
			continue
		}
		out = append(out, []float64{smK, smSum / float64(sd)})
	}
	return out
}

// Ichimoku Cloud. v2.0.0 fixes the Senkou Span B bug (lows window now uses
// length3, matching the highs window) and replaces the per-window max/min
// loop with a rolling-extremum pass.
func Ichimoku(data [][]float64, l1 int, l2 int, l3 int, l4 int) [][]float64 {
	n := len(data)
	if n < l3 {
		return nil
	}
	place := make([][]float64, 0, n-l3+1)
	for i := l3 - 1; i < n; i++ {
		h1, l1lo := math.Inf(-1), math.Inf(1)
		for j := i - l1 + 1; j <= i; j++ {
			if data[j][0] > h1 {
				h1 = data[j][0]
			}
			if data[j][2] < l1lo {
				l1lo = data[j][2]
			}
		}
		h2, l2lo := math.Inf(-1), math.Inf(1)
		for j := i - l2 + 1; j <= i; j++ {
			if data[j][0] > h2 {
				h2 = data[j][0]
			}
			if data[j][2] < l2lo {
				l2lo = data[j][2]
			}
		}
		h3, l3lo := math.Inf(-1), math.Inf(1)
		for j := i - l3 + 1; j <= i; j++ {
			if data[j][0] > h3 {
				h3 = data[j][0]
			}
			if data[j][2] < l3lo {
				l3lo = data[j][2]
			}
		}
		tsen := (h1 + l1lo) / 2
		ksen := (h2 + l2lo) / 2
		senka := data[i][1] + ksen
		senkb := (h3 + l3lo) / 2
		chik := data[i][1]
		place = append(place, []float64{tsen, ksen, senka, senkb, chik})
	}
	cloud := make([][]float64, 0)
	for i := l4; i < len(place)-l4; i++ {
		cloud = append(cloud, []float64{place[i][0], place[i][1], place[i+l4][2], place[i+l4][3], place[i-l4][4]})
	}
	return cloud
}

// Atr is the Wilder Average True Range. Bars are [high, close, low].
// v2.0.0: TR uses absolute values (|H-Cprev|, |L-Cprev|) and output length
// is len(data) - length + 1, matching sibling indicators.
func Atr(data [][]float64, l int) []float64 {
	n := len(data)
	if n < l {
		return nil
	}
	prev := data[0][0] - data[0][2]
	out := make([]float64, 0, n-l+1)
	if l == 1 {
		out = append(out, prev)
	}
	for i := 1; i < n; i++ {
		cprev := data[i-1][1]
		tr := data[i][0] - data[i][2]
		if v := math.Abs(data[i][0] - cprev); v > tr {
			tr = v
		}
		if v := math.Abs(data[i][2] - cprev); v > tr {
			tr = v
		}
		prev = (prev*float64(l-1) + tr) / float64(l)
		if i >= l-1 {
			out = append(out, prev)
		}
	}
	return out
}

func Fib(start float64, end float64) []float64 {
	return []float64{start, (end-start)*.236 + start, (end-start)*.382 + start, (end-start)*.5 + start, (end-start)*.618 + start, (end-start)*.786 + start, end, (end-start)*1.618 + start, (end-start)*2.618 + start, (end-start)*3.618 + start, (end-start)*4.236 + start}
}

// Fractals returns boolean up/down fractal markers for each bar (Williams
// fractals on [high, low] tuples). The first two and last two bars are
// padded with [false, false]; the price-mode variant is FractalsPrice.
func Fractals(data [][]float64) [][]bool {
	out := [][]bool{{false, false}, {false, false}}
	for i := 2; i < len(data)-2; i++ {
		isUp := data[i-2][0] < data[i][0] &&
			data[i-1][0] < data[i][0] &&
			data[i][0] > data[i+1][0] &&
			data[i][0] > data[i+2][0]
		isDown := data[i-2][1] > data[i][1] &&
			data[i-1][1] > data[i][1] &&
			data[i][1] < data[i+1][1] &&
			data[i][1] < data[i+2][1]
		out = append(out, []bool{isUp, isDown})
	}
	out = append(out, []bool{false, false}, []bool{false, false})
	return out
}

// FractalsPrice is the price-mode variant of Fractals: each bar emits
// [highPriceIfUp, lowPriceIfDown] with -1 sentinel for "no fractal".
// Two leading and two trailing pad rows of [-1, -1] keep the output the
// same length as data.
func FractalsPrice(data [][]float64) [][]float64 {
	pad := []float64{-1, -1}
	out := [][]float64{pad, pad}
	for i := 2; i < len(data)-2; i++ {
		isUp := data[i-2][0] < data[i][0] &&
			data[i-1][0] < data[i][0] &&
			data[i][0] > data[i+1][0] &&
			data[i][0] > data[i+2][0]
		isDown := data[i-2][1] > data[i][1] &&
			data[i-1][1] > data[i][1] &&
			data[i][1] < data[i+1][1] &&
			data[i][1] < data[i+2][1]
		hi := -1.0
		lo := -1.0
		if isUp {
			hi = data[i][0]
		}
		if isDown {
			lo = data[i][1]
		}
		out = append(out, []float64{hi, lo})
	}
	out = append(out, pad, pad)
	return out
}

// HalfTrend mirrors the v2.0.0 JS implementation: 4-tuple output
// [atrHigh, mid, atrLow, "long"|"short"]; loop is `i < len(data)` (one
// fewer iteration than the previous Go version); direction is set
// correctly per branch (was always "long" in the old code).
func HalfTrend(data [][]float64, atrlen int, amplitude int, deviation float64) [][]interface{} {
	out := make([][]interface{}, 0)
	nexttrend := []int{0}
	trend := []int{0}
	up := []float64{0}
	down := []float64{0}
	for i := atrlen; i < len(data); i++ {
		pl := data[i-atrlen : i]
		maxlow := pl[len(pl)-2][2]
		minhigh := pl[len(pl)-2][0]
		atrSeq := Atr(pl, atrlen)
		atr2 := atrSeq[len(atrSeq)-1] / 2
		dev := deviation * atr2

		// max(high) / min(low) over the window
		highprice := math.Inf(-1)
		lowprice := math.Inf(1)
		for _, b := range pl {
			if b[0] > highprice {
				highprice = b[0]
			}
			if b[2] < lowprice {
				lowprice = b[2]
			}
		}
		highs := map2d(pl[len(pl)-amplitude:], 0)
		lows := map2d(pl[len(pl)-amplitude:], 2)
		highma := ma.Sma(highs, amplitude)
		lowma := ma.Sma(lows, amplitude)

		if nexttrend[len(nexttrend)-1] == 1 {
			if lowprice > maxlow {
				maxlow = lowprice
			}
			if highma[0] < maxlow && pl[len(pl)-1][1] < pl[len(pl)-2][2] {
				trend = append(trend, 1)
				nexttrend = append(nexttrend, 0)
				minhigh = pl[len(pl)-2][0]
			}
		} else {
			if highprice < minhigh {
				minhigh = highprice
			}
			if lowma[0] > minhigh && pl[len(pl)-1][1] < pl[len(pl)-2][0] {
				trend = append(trend, 0)
				nexttrend = append(nexttrend, 1)
				maxlow = lowprice
			}
		}

		var direction string
		if trend[len(trend)-1] == 0 {
			if len(trend) > 1 && trend[len(trend)-2] != 0 {
				if len(down) < 2 {
					up = append(up, down[len(down)-1])
				} else {
					up = append(up, down[len(down)-2])
				}
			} else {
				if len(up) < 2 {
					up = append(up, maxlow)
				} else if up[len(up)-2] > maxlow {
					up = append(up, up[len(up)-2])
				} else {
					up = append(up, maxlow)
				}
			}
			direction = "long"
		} else {
			if len(trend) > 1 && trend[len(trend)-2] != 1 {
				if len(up) < 2 {
					down = append(down, up[len(up)-1])
				} else {
					down = append(down, up[len(up)-2])
				}
			} else {
				if len(down) < 2 {
					down = append(down, minhigh)
				} else if down[len(down)-2] < minhigh {
					down = append(down, down[len(down)-2])
				} else {
					down = append(down, minhigh)
				}
			}
			direction = "short"
		}

		var anchor float64
		if direction == "long" {
			anchor = up[len(up)-1]
		} else {
			anchor = down[len(down)-1]
		}
		atrHigh := anchor + dev
		atrLow := anchor - dev

		var mid float64
		if trend[len(trend)-1] == 0 {
			mid = up[len(up)-1]
		} else {
			mid = down[len(down)-1]
		}
		out = append(out, []interface{}{atrHigh, mid, atrLow, direction})
	}
	return out
}

func ZigZag(data [][]float64, perc float64) []float64 {
	var indexes []stats.RecentHighLow
	min := math.Inf(1)
	max := math.Inf(-1)
	lmin := false
	lmax := false
	for i := 0; i < len(data); i++ {
		if lmin {
			if indexes[len(indexes)-1].Value >= data[i][1] {
				indexes[len(indexes)-1].Value = data[i][1]
				indexes[len(indexes)-1].Index = i
			}
			if min >= data[i][1] {
				min = data[i][1]
			}
			hdif := (data[i][0] - min) / min
			if hdif > perc {
				indexes = append(indexes, stats.RecentHighLow{Index: i, Value: data[i][0]})
				lmax = true
				lmin = false
				min = math.Inf(1)
			}
		} else if lmax {
			if indexes[len(indexes)-1].Value <= data[i][0] {
				indexes[len(indexes)-1].Value = data[i][0]
				indexes[len(indexes)-1].Index = i
			}
			if max <= data[i][1] {
				max = data[i][1]
			}
			ldif := (max - data[i][1]) / data[i][1]
			if ldif > perc {
				indexes = append(indexes, stats.RecentHighLow{Index: i, Value: data[i][1]})
				lmin = true
				lmax = false
				max = math.Inf(-1)
			}
		} else {
			if min >= data[i][1] {
				min = data[i][1]
			}
			if max <= data[i][0] {
				max = data[i][0]
			}
			hdif := (data[i][0] - min) / min
			ldif := (max - data[i][1]) / max
			if ldif > perc && hdif < perc {
				lmin = true
				indexes = append(indexes, stats.RecentHighLow{Index: 0, Value: data[0][0]})
				indexes = append(indexes, stats.RecentHighLow{Index: i, Value: data[i][1]})
			} else if hdif > perc && ldif < perc {
				lmax = true
				indexes = append(indexes, stats.RecentHighLow{Index: 0, Value: data[0][1]})
				indexes = append(indexes, stats.RecentHighLow{Index: i, Value: data[i][0]})
			} else {
				if ldif > hdif {
					lmin = true
					indexes = append(indexes, stats.RecentHighLow{Index: 0, Value: data[0][0]})
					indexes = append(indexes, stats.RecentHighLow{Index: i, Value: data[i][0]})
				} else {
					lmax = true
					indexes = append(indexes, stats.RecentHighLow{Index: 0, Value: data[0][1]})
					indexes = append(indexes, stats.RecentHighLow{Index: i, Value: data[i][0]})
				}
			}
		}
	}
	final := []float64{indexes[0].Value}
	for i := 1; i < len(indexes); i++ {
		length := indexes[i].Index - indexes[i-1].Index
		delta := (indexes[i].Value - indexes[i-1].Value) / float64(length)
		for z := 1; z <= length; z++ {
			final = append(final, float64(z)*delta+indexes[i-1].Value)
		}
	}
	return final
}

func Psar(data [][]float64, step float64, max float64) []float64 {
	furthest := data[0]
	up := true
	accel := step
	prev := data[0]
	sar := data[0][1]
	extreme := data[0][0]
	final := []float64{sar}
	for i := 1; i < len(data); i++ {
		sar = sar + accel*(extreme-sar)
		if up {
			sar = minf([]float64{sar, furthest[1], prev[1]})
			if data[i][0] > extreme {
				extreme = data[i][0]
				accel = minf([]float64{accel + step, max})
			}
		} else {
			sar = maxf([]float64{sar, furthest[0], prev[0]})
			if data[i][1] < extreme {
				extreme = data[i][1] // v2.0.0 fix: was data[i][0] (high) — must be data[i][1] (low) in downtrend
				accel = minf([]float64{accel + step, max})
			}
		}
		if (up && data[i][1] < sar) || (!up && data[i][0] > sar) {
			accel = step
			sar = extreme
			up = !up
			if !up {
				extreme = data[i][1]
			} else {
				extreme = data[i][0]
			}
		}
		furthest = prev
		prev = data[i]
		final = append(final, sar)
	}
	return final
}

// Supertrend bands: midprice ± multiplier * ATR. v2.0.0 inherits Atr's
// new output shape (N - L + 1); the indexing offset is adjusted accordingly.
func Supertrend(data [][]float64, l int, multiplier float64) [][]float64 {
	atr := Atr(data, l)
	if len(atr) == 0 {
		return nil
	}
	trend := make([][]float64, 0, len(atr))
	for i := l - 1; i < len(data); i++ {
		mid := (data[i][0] + data[i][2]) / 2
		a := atr[i-(l-1)]
		trend = append(trend, []float64{mid + multiplier*a, mid - multiplier*a})
	}
	return trend
}

func Elderray(data []float64, l int) [][]float64 {
	var eld [][]float64
	for i := l; i <= len(data); i++ {
		pl := append([]float64{}, data[i-l:i]...)
		sort.Float64s(pl)
		em := ma.Ema(data[i-l:i], l)
		eld = append(eld, []float64{pl[len(pl)-1] - em[0], pl[0] - em[0]})
	}
	return eld
}

func Hv(data []float64, l int) []float64 {
	var hv []float64
	for i := l; i <= len(data); i++ {
		ss := stats.Ssd(data[i-l:i], l)
		vari := ss[0] / float64(l)
		hv = append(hv, math.Sqrt(vari))
	}
	return hv
}

func Rvi(data [][]float64, ln int) []float64 {
	var num []float64
	var dnom []float64
	var rv []float64
	for i := 3; i < len(data); i++ {
		num = append(num, (data[i][3]-data[i][0]+2*(data[i][3]-data[i-1][0])+2*(data[i][3]-data[i-2][0])+(data[i][3]-data[i-3][0]))/6)
		dnom = append(dnom, (data[i][1]-data[i][2]+2*(data[i][1]-data[i-1][2])+2*(data[i][1]-data[i-2][2])+(data[i][1]-data[i-3][2]))/6)
		if len(num) >= ln {
			sn := ma.Sma(num, ln)
			dn := ma.Sma(dnom, ln)
			rv = append(rv, sn[0]/dn[0])
			num = num[1:]
			dnom = dnom[1:]
		}
	}
	return rv
}

func Rvi_signal(rv []float64) []float64 {
	var sig []float64
	for i := 3; i < len(rv); i++ {
		sig = append(sig, (rv[i]+2*rv[i-1]+2*rv[i-2]+rv[i-3])/6)
	}
	return sig
}

func Rsi_divergence(data []float64, ln int, rs func([]float64, int) []float64) []float64 {
	rd := rs(data, ln)
	var out []float64
	data = Mom(data[ln-1:len(data)], 2, false)
	for i := 0; i < len(data); i++ {
		if (data[i] > 0 && rd[i] < 0) || (data[i] < 0 && rd[i] > 0) {
			out = append(out, 1)
		} else {
			out = append(out, 0)
		}
	}
	return out
}

// CrossSignal records a single cross between two series — Index is the
// bar at which the cross occurs and Cross is true for an upward cross
// (d1 rising above d2), false for a downward cross.
type CrossSignal struct {
	Index int
	Cross bool
}

// Cross reports the points where d1 crosses d2. v2.0.0: input series
// are not mutated; d1 is non-mutatingly trimmed to d2's length when
// longer.
func Cross(d1, d2 []float64) []CrossSignal {
	if len(d1) > len(d2) {
		d1 = d1[len(d1)-len(d2):]
	}
	if len(d1) == 0 {
		return nil
	}
	crossed := d1[0] > d2[0]
	out := make([]CrossSignal, 0)
	for i := 0; i < len(d1); i++ {
		if d1[i] < d2[i] && crossed {
			out = append(out, CrossSignal{Index: i, Cross: false})
			crossed = false
		}
		if d1[i] > d2[i] && !crossed {
			out = append(out, CrossSignal{Index: i, Cross: true})
			crossed = true
		}
	}
	return out
}

// MacdSignal is the EMA of the MACD line over `lsig` bars.
func MacdSignal(data []float64, l1, l2, lsig int) []float64 {
	return ma.Ema(Macd(data, l1, l2), lsig)
}

// MacdBars is the histogram (MACD − MACD signal) component of MACD.
func MacdBars(data []float64, l1, l2, lsig int) []float64 {
	macd := Macd(data, l1, l2)
	mas := ma.Ema(macd, lsig)
	off := len(macd) - len(mas)
	out := make([]float64, len(mas))
	for i := 0; i < len(mas); i++ {
		out[i] = macd[i+off] - mas[i]
	}
	return out
}

// Adl is the Chaikin Accumulation/Distribution Line. Bars are
// [H, C, L, V]; output is the running cumulative money-flow volume.
func Adl(data [][]float64) []float64 {
	n := len(data)
	if n == 0 {
		return nil
	}
	out := make([]float64, n)
	var acc float64
	for i := 0; i < n; i++ {
		h, c, l, v := data[i][0], data[i][1], data[i][2], data[i][3]
		r := h - l
		var mfm float64
		if r != 0 {
			mfm = ((c - l) - (h - c)) / r
		}
		acc += mfm * v
		out[i] = acc
	}
	return out
}

// Apo is the Absolute Price Oscillator — equivalent to MACD with the
// same length parameters.
func Apo(data []float64, l1, l2 int) []float64 {
	return Macd(data, l1, l2)
}

// Cci is the Commodity Channel Index. Bars are [H, C, L]. Returns
// NaN for windows where the mean absolute deviation is zero.
func Cci(data [][]float64, l int) []float64 {
	n := len(data)
	if n < l {
		return nil
	}
	tp := make([]float64, n)
	for i := 0; i < n; i++ {
		tp[i] = (data[i][0] + data[i][1] + data[i][2]) / 3
	}
	out := make([]float64, n-l+1)
	for i := l; i <= n; i++ {
		var sum float64
		for j := i - l; j < i; j++ {
			sum += tp[j]
		}
		smaTp := sum / float64(l)
		var mad float64
		for j := i - l; j < i; j++ {
			mad += math.Abs(tp[j] - smaTp)
		}
		mad /= float64(l)
		if mad == 0 {
			out[i-l] = math.NaN()
		} else {
			out[i-l] = (tp[i-1] - smaTp) / (0.015 * mad)
		}
	}
	return out
}

// Cmf is the Chaikin Money Flow over `l` bars. Bars are [H, C, L, V].
func Cmf(data [][]float64, l int) []float64 {
	n := len(data)
	if n < l {
		return nil
	}
	mfv := make([]float64, n)
	for i := 0; i < n; i++ {
		h, c, lo, v := data[i][0], data[i][1], data[i][2], data[i][3]
		r := h - lo
		if r != 0 {
			mfv[i] = (((c - lo) - (h - c)) / r) * v
		}
	}
	var sM, sV float64
	for i := 0; i < l; i++ {
		sM += mfv[i]
		sV += data[i][3]
	}
	out := make([]float64, n-l+1)
	if sV == 0 {
		out[0] = math.NaN()
	} else {
		out[0] = sM / sV
	}
	for i := l; i < n; i++ {
		sM += mfv[i] - mfv[i-l]
		sV += data[i][3] - data[i-l][3]
		if sV == 0 {
			out[i-l+1] = math.NaN()
		} else {
			out[i-l+1] = sM / sV
		}
	}
	return out
}

// Dpo is the Detrended Price Oscillator: data − backshifted SMA.
func Dpo(data []float64, l int) []float64 {
	sma := ma.Sma(data, l)
	if len(sma) == 0 {
		return nil
	}
	off := l - 2 - l/2
	out := make([]float64, len(sma))
	for i := 0; i < len(sma); i++ {
		out[i] = data[i+off] - sma[i]
	}
	return out
}

// Pdm is Wilder's smoothed positive directional movement (+DM).
// Bars are [H, C, L].
func Pdm(data [][]float64, l int) []float64 {
	n := len(data)
	if n < l+1 {
		return nil
	}
	var sum float64
	for i := 1; i <= l; i++ {
		up := data[i][0] - data[i-1][0]
		dn := data[i-1][2] - data[i][2]
		if up > dn && up > 0 {
			sum += up
		}
	}
	prev := sum / float64(l)
	out := make([]float64, n-l)
	out[0] = prev
	for i := l + 1; i < n; i++ {
		up := data[i][0] - data[i-1][0]
		dn := data[i-1][2] - data[i][2]
		var dm float64
		if up > dn && up > 0 {
			dm = up
		}
		prev = (prev*float64(l-1) + dm) / float64(l)
		out[i-l] = prev
	}
	return out
}

// Mdm is Wilder's smoothed negative directional movement (−DM).
func Mdm(data [][]float64, l int) []float64 {
	n := len(data)
	if n < l+1 {
		return nil
	}
	var sum float64
	for i := 1; i <= l; i++ {
		up := data[i][0] - data[i-1][0]
		dn := data[i-1][2] - data[i][2]
		if dn > up && dn > 0 {
			sum += dn
		}
	}
	prev := sum / float64(l)
	out := make([]float64, n-l)
	out[0] = prev
	for i := l + 1; i < n; i++ {
		up := data[i][0] - data[i-1][0]
		dn := data[i-1][2] - data[i][2]
		var dm float64
		if dn > up && dn > 0 {
			dm = dn
		}
		prev = (prev*float64(l-1) + dm) / float64(l)
		out[i-l] = prev
	}
	return out
}

// Pdi is the Wilder positive directional indicator: 100 × +DM / TR.
func Pdi(data [][]float64, l int) []float64 {
	n := len(data)
	if n < l+1 {
		return nil
	}
	var sumDm, sumTr float64
	for i := 1; i <= l; i++ {
		h, lo, cprev := data[i][0], data[i][2], data[i-1][1]
		up := data[i][0] - data[i-1][0]
		dn := data[i-1][2] - data[i][2]
		if up > dn && up > 0 {
			sumDm += up
		}
		tr := h - lo
		if v := math.Abs(h - cprev); v > tr {
			tr = v
		}
		if v := math.Abs(lo - cprev); v > tr {
			tr = v
		}
		sumTr += tr
	}
	dm := sumDm / float64(l)
	tr := sumTr / float64(l)
	out := make([]float64, n-l)
	if tr == 0 {
		out[0] = math.NaN()
	} else {
		out[0] = 100 * dm / tr
	}
	for i := l + 1; i < n; i++ {
		h, lo, cprev := data[i][0], data[i][2], data[i-1][1]
		up := data[i][0] - data[i-1][0]
		dn := data[i-1][2] - data[i][2]
		var curDm float64
		if up > dn && up > 0 {
			curDm = up
		}
		curTr := h - lo
		if v := math.Abs(h - cprev); v > curTr {
			curTr = v
		}
		if v := math.Abs(lo - cprev); v > curTr {
			curTr = v
		}
		dm = (dm*float64(l-1) + curDm) / float64(l)
		tr = (tr*float64(l-1) + curTr) / float64(l)
		if tr == 0 {
			out[i-l] = math.NaN()
		} else {
			out[i-l] = 100 * dm / tr
		}
	}
	return out
}

// Mdi is the Wilder negative directional indicator: 100 × −DM / TR.
func Mdi(data [][]float64, l int) []float64 {
	n := len(data)
	if n < l+1 {
		return nil
	}
	var sumDm, sumTr float64
	for i := 1; i <= l; i++ {
		h, lo, cprev := data[i][0], data[i][2], data[i-1][1]
		up := data[i][0] - data[i-1][0]
		dn := data[i-1][2] - data[i][2]
		if dn > up && dn > 0 {
			sumDm += dn
		}
		tr := h - lo
		if v := math.Abs(h - cprev); v > tr {
			tr = v
		}
		if v := math.Abs(lo - cprev); v > tr {
			tr = v
		}
		sumTr += tr
	}
	dm := sumDm / float64(l)
	tr := sumTr / float64(l)
	out := make([]float64, n-l)
	if tr == 0 {
		out[0] = math.NaN()
	} else {
		out[0] = 100 * dm / tr
	}
	for i := l + 1; i < n; i++ {
		h, lo, cprev := data[i][0], data[i][2], data[i-1][1]
		up := data[i][0] - data[i-1][0]
		dn := data[i-1][2] - data[i][2]
		var curDm float64
		if dn > up && dn > 0 {
			curDm = dn
		}
		curTr := h - lo
		if v := math.Abs(h - cprev); v > curTr {
			curTr = v
		}
		if v := math.Abs(lo - cprev); v > curTr {
			curTr = v
		}
		dm = (dm*float64(l-1) + curDm) / float64(l)
		tr = (tr*float64(l-1) + curTr) / float64(l)
		if tr == 0 {
			out[i-l] = math.NaN()
		} else {
			out[i-l] = 100 * dm / tr
		}
	}
	return out
}

// Dx is the Wilder Directional Movement Index: 100 × |+DI − −DI| / (+DI + −DI).
func Dx(data [][]float64, l int) []float64 {
	p := Pdi(data, l)
	m := Mdi(data, l)
	out := make([]float64, len(p))
	for i := 0; i < len(p); i++ {
		sum := p[i] + m[i]
		if sum == 0 {
			out[i] = math.NaN()
		} else {
			out[i] = 100 * math.Abs(p[i]-m[i]) / sum
		}
	}
	return out
}

// Adx is the Wilder Average Directional Index — Wilder smoothing of DX.
func Adx(data [][]float64, l int) []float64 {
	d := Dx(data, l)
	if len(d) < l {
		return nil
	}
	var sum float64
	for i := 0; i < l; i++ {
		sum += d[i]
	}
	prev := sum / float64(l)
	out := make([]float64, len(d)-l+1)
	out[0] = prev
	for i := l; i < len(d); i++ {
		prev = (prev*float64(l-1) + d[i]) / float64(l)
		out[i-l+1] = prev
	}
	return out
}

// Adxr is the Average Directional Movement Index Rating, the average
// of the current ADX and the ADX from `l-1` bars ago.
func Adxr(data [][]float64, l int) []float64 {
	a := Adx(data, l)
	if len(a) < l {
		return nil
	}
	out := make([]float64, len(a)-l+1)
	for i := l - 1; i < len(a); i++ {
		out[i-l+1] = (a[i] + a[i-l+1]) / 2
	}
	return out
}

// Emv is the Ease of Movement indicator. Bars are [H, L, V].
// v2.0.0: returns NaN for bars with zero range or zero volume.
func Emv(data [][]float64) []float64 {
	n := len(data)
	if n < 2 {
		return nil
	}
	out := make([]float64, n-1)
	for i := 1; i < n; i++ {
		h, l, v := data[i][0], data[i][1], data[i][2]
		hp, lp := data[i-1][0], data[i-1][1]
		move := (h+l)/2 - (hp+lp)/2
		r := h - l
		if r == 0 || v == 0 {
			out[i-1] = math.NaN()
		} else {
			out[i-1] = move * r * 10000 / v
		}
	}
	return out
}

// Kdj augments Stoch with the J line (3K − 2D). The K/D parameter
// order matches ta.js' kdj signature (length, smoothK, smoothD).
func Kdj(data [][]float64, l, smoothK, smoothD int) [][]float64 {
	s := Stoch(data, l, smoothK, smoothD)
	out := make([][]float64, len(s))
	for i := 0; i < len(s); i++ {
		k, d := s[i][0], s[i][1]
		out[i] = []float64{k, d, 3*k - 2*d}
	}
	return out
}

// Mass is the Mass Index over `l` bars. Bars are [H, L].
func Mass(data [][]float64, l int) []float64 {
	n := len(data)
	rng := make([]float64, n)
	for i := 0; i < n; i++ {
		rng[i] = data[i][0] - data[i][1]
	}
	e1 := ma.Ema(rng, 9)
	if len(e1) < 9 {
		return nil
	}
	e2 := ma.Ema(e1, 9)
	if len(e2) < l {
		return nil
	}
	off := len(e1) - len(e2)
	ratio := make([]float64, len(e2))
	for i := 0; i < len(e2); i++ {
		if e2[i] == 0 {
			ratio[i] = math.NaN()
		} else {
			ratio[i] = e1[i+off] / e2[i]
		}
	}
	var sum float64
	for i := 0; i < l; i++ {
		sum += ratio[i]
	}
	out := make([]float64, len(ratio)-l+1)
	out[0] = sum
	for i := l; i < len(ratio); i++ {
		sum += ratio[i] - ratio[i-l]
		out[i-l+1] = sum
	}
	return out
}

// Natr is the Normalized ATR — 100 × ATR / close.
func Natr(data [][]float64, l int) []float64 {
	a := Atr(data, l)
	if len(a) == 0 {
		return nil
	}
	out := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		c := data[i+l-1][1]
		if c == 0 {
			out[i] = math.NaN()
		} else {
			out[i] = 100 * a[i] / c
		}
	}
	return out
}

// Nvi is the Negative Volume Index. Bars are [close, volume].
func Nvi(data [][]float64) []float64 {
	n := len(data)
	if n == 0 {
		return nil
	}
	out := make([]float64, n)
	out[0] = 1000
	for i := 1; i < n; i++ {
		cprev := data[i-1][0]
		if data[i][1] < data[i-1][1] && cprev != 0 {
			out[i] = out[i-1] * data[i][0] / cprev
		} else {
			out[i] = out[i-1]
		}
	}
	return out
}

// Pvi is the Positive Volume Index. Bars are [close, volume].
func Pvi(data [][]float64) []float64 {
	n := len(data)
	if n == 0 {
		return nil
	}
	out := make([]float64, n)
	out[0] = 1000
	for i := 1; i < n; i++ {
		cprev := data[i-1][0]
		if data[i][1] > data[i-1][1] && cprev != 0 {
			out[i] = out[i-1] * data[i][0] / cprev
		} else {
			out[i] = out[i-1]
		}
	}
	return out
}

// Ppo is the Percentage Price Oscillator: 100 × MACD / EMA(data, l2).
func Ppo(data []float64, l1, l2 int) []float64 {
	if l1 > l2 {
		l1, l2 = l2, l1
	}
	eb := ma.Ema(data, l2)
	m := Macd(data, l1, l2)
	out := make([]float64, len(m))
	for i := 0; i < len(m); i++ {
		if eb[i] == 0 {
			out[i] = math.NaN()
		} else {
			out[i] = 100 * m[i] / eb[i]
		}
	}
	return out
}

// StochRsi is the Stochastic-RSI oscillator: %K and %D of the rolling
// high/low of RSI over `stochLength` bars.
func StochRsi(data []float64, rsiLength, stochLength, smoothK, smoothD int) [][]float64 {
	r := Rsi(data, rsiLength)
	if len(r) < stochLength {
		return nil
	}
	rawK := make([]float64, len(r)-stochLength+1)
	for i := stochLength; i <= len(r); i++ {
		lo, hi := math.Inf(1), math.Inf(-1)
		for j := i - stochLength; j < i; j++ {
			if r[j] > hi {
				hi = r[j]
			}
			if r[j] < lo {
				lo = r[j]
			}
		}
		rng := hi - lo
		if rng == 0 {
			rawK[i-stochLength] = math.NaN()
		} else {
			rawK[i-stochLength] = 100 * (r[i-1] - lo) / rng
		}
	}
	fastK := ma.Sma(rawK, smoothK)
	fastD := ma.Sma(fastK, smoothD)
	if len(fastD) == 0 {
		return nil
	}
	off := len(fastK) - len(fastD)
	out := make([][]float64, len(fastD))
	for i := 0; i < len(fastD); i++ {
		out[i] = []float64{fastK[i+off], fastD[i]}
	}
	return out
}

// Trix is the percent rate-of-change of a triple-smoothed EMA.
func Trix(data []float64, l int) []float64 {
	e1 := ma.Ema(data, l)
	if len(e1) < l {
		return nil
	}
	e2 := ma.Ema(e1, l)
	if len(e2) < l {
		return nil
	}
	e3 := ma.Ema(e2, l)
	if len(e3) < 2 {
		return nil
	}
	out := make([]float64, len(e3)-1)
	for i := 1; i < len(e3); i++ {
		prev := e3[i-1]
		if prev == 0 {
			out[i-1] = math.NaN()
		} else {
			out[i-1] = 100 * (e3[i] - prev) / prev
		}
	}
	return out
}

// Ulcer is the Ulcer Index — RMS percent drawdown from the rolling
// `l`-bar high.
func Ulcer(data []float64, l int) []float64 {
	n := len(data)
	if n < l {
		return nil
	}
	dd := make([]float64, n)
	for i := 0; i < n; i++ {
		hi := math.Inf(-1)
		start := i - l + 1
		if start < 0 {
			start = 0
		}
		for j := start; j <= i; j++ {
			if data[j] > hi {
				hi = data[j]
			}
		}
		if hi == 0 {
			dd[i] = 0
		} else {
			dd[i] = 100 * (data[i] - hi) / hi
		}
	}
	var sum float64
	for i := 0; i < l; i++ {
		sum += dd[i] * dd[i]
	}
	out := make([]float64, n-l+1)
	out[0] = math.Sqrt(sum / float64(l))
	for i := l; i < n; i++ {
		sum += dd[i]*dd[i] - dd[i-l]*dd[i-l]
		out[i-l+1] = math.Sqrt(sum / float64(l))
	}
	return out
}

// Vortex is the Vortex Indicator pair [+VI, −VI]. Bars are [H, C, L].
func Vortex(data [][]float64, l int) [][]float64 {
	n := len(data)
	if n < l+1 {
		return nil
	}
	vmp := make([]float64, n)
	vmn := make([]float64, n)
	tr := make([]float64, n)
	for i := 1; i < n; i++ {
		h, lo, cprev := data[i][0], data[i][2], data[i-1][1]
		hprev, lprev := data[i-1][0], data[i-1][2]
		vmp[i] = math.Abs(h - lprev)
		vmn[i] = math.Abs(lo - hprev)
		t := h - lo
		if v := math.Abs(h - cprev); v > t {
			t = v
		}
		if v := math.Abs(lo - cprev); v > t {
			t = v
		}
		tr[i] = t
	}
	var sP, sN, sT float64
	for i := 1; i <= l; i++ {
		sP += vmp[i]
		sN += vmn[i]
		sT += tr[i]
	}
	out := make([][]float64, n-l)
	if sT == 0 {
		out[0] = []float64{math.NaN(), math.NaN()}
	} else {
		out[0] = []float64{sP / sT, sN / sT}
	}
	for i := l + 1; i < n; i++ {
		sP += vmp[i] - vmp[i-l]
		sN += vmn[i] - vmn[i-l]
		sT += tr[i] - tr[i-l]
		if sT == 0 {
			out[i-l] = []float64{math.NaN(), math.NaN()}
		} else {
			out[i-l] = []float64{sP / sT, sN / sT}
		}
	}
	return out
}

// Divergence reports per-bar divergence (1) or no-divergence (0) between
// two series. v2.0.0 fix: no longer mutates either input — the trim to the
// common tail is done with non-mutating slice expressions.
func Divergence(data1 []float64, data2 []float64) []float64 {
	if len(data1) > len(data2) {
		data1 = data1[len(data1)-len(data2):]
	}
	if len(data2) > len(data1) {
		data2 = data2[len(data2)-len(data1):]
	}
	out := make([]float64, 0, len(data1))
	for i := 1; i < len(data1); i++ {
		up1 := data1[i] > data1[i-1]
		dn1 := data1[i] < data1[i-1]
		up2 := data2[i] > data2[i-1]
		dn2 := data2[i] < data2[i-1]
		if (up1 && dn2) || (dn1 && up2) {
			out = append(out, 1)
		} else {
			out = append(out, 0)
		}
	}
	return out
}
