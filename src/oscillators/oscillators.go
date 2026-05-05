// Package oscillators contains oscillator indicators that derive from the
// underlying moving-average and indicator primitives in ta.go.
package oscillators

import (
	"math"
	"sort"

	ind "github.com/Bitvested/ta.go/src/indicators"
	ma "github.com/Bitvested/ta.go/src/moving_averages"
)

func maxf(v []float64) float64 {
	sort.Float64s(v)
	return v[len(v)-1]
}

func minf(v []float64) float64 {
	sort.Float64s(v)
	return v[0]
}

func Gator(data []float64, jl int, tl int, ll int, js int, ts int, ls int) [][]float64 {
	var ret [][]float64
	var jaw []float64 = ma.Smma(data, jl)
	var teeth []float64 = ma.Smma(data, tl)
	var lips []float64 = ma.Smma(data, ll)
	teeth = teeth[len(teeth)-len(jaw):]
	lips = lips[len(lips)-len(jaw):]
	for i := len(jaw) - 1; i >= (js - 1); i-- {
		ret = append(ret, []float64{jaw[i-(js-1)] - teeth[i-(ts-1)], -(math.Abs(teeth[i-(ts-1)] - lips[i-(ls-1)]))})
	}
	return ret
}

// MomOsc is the Chande Momentum Oscillator (CMO). v2.0.0 corrects the
// formula: sum the price *changes* (not the prices themselves) for up-days
// vs down-days. NaN is emitted when both sums are zero (flat window).
func MomOsc(data []float64, l int) []float64 {
	n := len(data)
	if n <= l {
		return nil
	}
	out := make([]float64, 0, n-l)
	for i := l; i < n; i++ {
		var sumh, suml float64
		for j := i - l + 1; j <= i; j++ {
			d := data[j] - data[j-1]
			if d > 0 {
				sumh += d
			} else if d < 0 {
				suml -= d
			}
		}
		denom := sumh + suml
		if denom == 0 {
			out = append(out, math.NaN())
		} else {
			out = append(out, (sumh-suml)/denom*100)
		}
	}
	return out
}

// ChaikinOsc is the Chaikin Oscillator. v2.0.0 removes the dangerous
// "NaN → 0" coercion: bars with zero range now propagate NaN through ADL
// and out, instead of silently zeroing the contribution. Bars are
// [high, close, low, volume].
func ChaikinOsc(data [][]float64, ema1 int, ema2 int) []float64 {
	n := len(data)
	adl := make([]float64, n)
	for i := 0; i < n; i++ {
		rng := data[i][0] - data[i][2]
		var mfm float64
		if rng == 0 {
			mfm = math.NaN()
		} else {
			mfm = ((data[i][1] - data[i][2]) - (data[i][0] - data[i][1])) / rng
		}
		adl[i] = mfm * data[i][3]
	}
	ef := ma.Ema(adl, ema1)
	es := ma.Ema(adl, ema2)
	if len(ef) > len(es) {
		ef = ef[len(ef)-len(es):]
	} else if len(es) > len(ef) {
		es = es[len(es)-len(ef):]
	}
	out := make([]float64, len(ef))
	for i := range ef {
		out[i] = ef[i] - es[i]
	}
	return out
}

func AroonOsc(data []float64, l int) []float64 {
	var aroon []float64
	u := ind.AroonUp(data, l)
	d := ind.AroonDown(data, l)
	for i := 0; i < len(u); i++ {
		aroon = append(aroon, u[i]-d[i])
	}
	return aroon
}

func Ao(data [][]float64, l1 int, l2 int) []float64 {
	var processed []float64
	var a []float64
	for i := 0; i < len(data); i++ {
		processed = append(processed, (data[i][0]+data[i][1])/2)
	}
	for i := l2; i <= len(processed); i++ {
		pl := processed[i-l2 : i]
		f := ma.Sma(pl, l1)
		s := ma.Sma(pl, l2)
		a = append(a, f[len(f)-1]-s[len(s)-1])
	}
	return a
}

func Ac(data [][]float64, len1 int, len2 int) []float64 {
	a := Ao(data, len1, len2)
	sm := ma.Sma(a, len1)
	var acr []float64
	if len(a) > len(sm) {
		a = a[len(a)-len(sm):]
	} else {
		sm = sm[len(sm)-len(a):]
	}
	for i := 0; i < len(a); i++ {
		acr = append(acr, a[i]-sm[i])
	}
	return acr
}

// Fisher transform. v2.0.0 fixes:
//   - Output length normalized to len(data) - l + 1 (drops the trailing slice).
//   - Norm of (price - min) divided by (max - min), not by max alone.
//   - Zero-range guard returns NaN propagation through v1.
// Kvo is the Klinger Volume Oscillator. Bars are [H, C, L, V]. The
// trend flips when sum(H+C+L) crosses; cumulative measurement (CM)
// resets to the prior bar's range and accumulates the current range.
func Kvo(data [][]float64, fast, slow int) []float64 {
	n := len(data)
	if n <= 1 {
		return nil
	}
	aF := 2.0 / (float64(fast) + 1)
	aS := 2.0 / (float64(slow) + 1)
	prevHlc := data[0][0] + data[0][1] + data[0][2]
	trend := -1
	var cm, emaFast, emaSlow float64
	out := make([]float64, n-1)
	for i := 1; i < n; i++ {
		h, c, l, v := data[i][0], data[i][1], data[i][2], data[i][3]
		hlc := h + c + l
		dm := h - l
		if hlc > prevHlc && trend != 1 {
			trend = 1
			cm = data[i-1][0] - data[i-1][2]
		} else if hlc < prevHlc && trend != 0 {
			trend = 0
			cm = data[i-1][0] - data[i-1][2]
		}
		cm += dm
		sign := 1.0
		if trend == 0 {
			sign = -1
		}
		var vf float64
		if cm != 0 {
			vf = v * math.Abs(2*dm/cm-1) * 100 * sign
		}
		if i == 1 {
			emaFast = vf
			emaSlow = vf
		} else {
			emaFast = (vf-emaFast)*aF + emaFast
			emaSlow = (vf-emaSlow)*aS + emaSlow
		}
		out[i-1] = emaFast - emaSlow
		prevHlc = hlc
	}
	return out
}

// Ult is the Williams Ultimate Oscillator over three lookback periods
// `p1 < p2 < p3`. Bars are [H, C, L]. Returns nil when n−1 < max(p).
func Ult(data [][]float64, p1, p2, p3 int) []float64 {
	n := len(data)
	need := p1
	if p2 > need {
		need = p2
	}
	if p3 > need {
		need = p3
	}
	if n-1 < need {
		return nil
	}
	bp := make([]float64, n)
	tr := make([]float64, n)
	for i := 1; i < n; i++ {
		h, c, l := data[i][0], data[i][1], data[i][2]
		cprev := data[i-1][1]
		lo := l
		if cprev < lo {
			lo = cprev
		}
		hi := h
		if cprev > hi {
			hi = cprev
		}
		bp[i] = c - lo
		tr[i] = hi - lo
	}
	var sBp1, sTr1, sBp2, sTr2, sBp3, sTr3 float64
	for i := need - p1 + 1; i <= need; i++ {
		sBp1 += bp[i]
		sTr1 += tr[i]
	}
	for i := need - p2 + 1; i <= need; i++ {
		sBp2 += bp[i]
		sTr2 += tr[i]
	}
	for i := need - p3 + 1; i <= need; i++ {
		sBp3 += bp[i]
		sTr3 += tr[i]
	}
	out := []float64{ultVal(sBp1, sTr1, sBp2, sTr2, sBp3, sTr3)}
	for i := need + 1; i < n; i++ {
		sBp1 += bp[i] - bp[i-p1]
		sTr1 += tr[i] - tr[i-p1]
		sBp2 += bp[i] - bp[i-p2]
		sTr2 += tr[i] - tr[i-p2]
		sBp3 += bp[i] - bp[i-p3]
		sTr3 += tr[i] - tr[i-p3]
		out = append(out, ultVal(sBp1, sTr1, sBp2, sTr2, sBp3, sTr3))
	}
	return out
}

func ultVal(sBp1, sTr1, sBp2, sTr2, sBp3, sTr3 float64) float64 {
	a1 := math.NaN()
	if sTr1 != 0 {
		a1 = sBp1 / sTr1
	}
	a2 := math.NaN()
	if sTr2 != 0 {
		a2 = sBp2 / sTr2
	}
	a3 := math.NaN()
	if sTr3 != 0 {
		a3 = sBp3 / sTr3
	}
	return 100 * (4*a1 + 2*a2 + a3) / 7
}

// Fisher is the Ehlers Fisher Transform. v2.0.0 perf: pre-allocated
// [][]float64 output backed by a single flat slice. The inner max/min
// scan was already single-pass.
func Fisher(data []float64, l int) [][]float64 {
	n := len(data)
	if n < l {
		return nil
	}
	rows := n - l + 1
	flat := make([]float64, 2*rows)
	out := make([][]float64, rows)
	for i := range out {
		out[i] = flat[i*2 : i*2+2 : i*2+2]
	}
	var fish, v1 float64
	for i := l; i <= n; i++ {
		mn := math.Inf(1)
		mx := math.Inf(-1)
		for j := i - l; j < i; j++ {
			v := data[j]
			if v < mn {
				mn = v
			}
			if v > mx {
				mx = v
			}
		}
		rng := mx - mn
		var norm float64
		if rng != 0 {
			norm = (data[i-1]-mn)/rng - 0.5
		}
		v1 = 0.33*2*norm + 0.67*v1
		if v1 > 0.99 {
			v1 = 0.999
		}
		if v1 < -0.99 {
			v1 = -0.999
		}
		pf := fish
		fish = 0.5*math.Log((1+v1)/(1-v1)) + 0.5*pf
		row := out[i-l]
		row[0], row[1] = fish, pf
	}
	return out
}
