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

func ChaikinOsc(data [][]float64, ema1 int, ema2 int) []float64 {
	var cha []float64
	var adl []float64
	for i := 0; i < len(data); i++ {
		var mfm = ((data[i][1] - data[i][2]) - (data[i][0] - data[i][1])) / (data[i][0] - data[i][2])
		if math.IsNaN(mfm) {
			adl = append(adl, 0)
		} else {
			adl = append(adl, mfm*data[i][3])
		}
	}
	ef := ma.Ema(adl, ema1)
	es := ma.Ema(adl, ema2)
	if len(ef) > len(es) {
		ef = ef[len(ef)-len(es):]
	} else {
		es = es[len(es)-len(ef):]
	}
	for i := 0; i < len(ef); i++ {
		cha = append(cha, ef[i]-es[i])
	}
	return cha
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
func Fisher(data []float64, l int) [][]float64 {
	n := len(data)
	if n < l {
		return nil
	}
	out := make([][]float64, 0, n-l+1)
	var fish, v1 float64
	for i := l; i <= n; i++ {
		mn := math.Inf(1)
		mx := math.Inf(-1)
		for j := i - l; j < i; j++ {
			if data[j] < mn {
				mn = data[j]
			}
			if data[j] > mx {
				mx = data[j]
			}
		}
		rng := mx - mn
		var norm float64
		if rng == 0 {
			norm = 0
		} else {
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
		out = append(out, []float64{fish, pf})
	}
	return out
}
