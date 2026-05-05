// Package bands contains envelope/channel indicators that wrap a moving
// average with a deviation band (Bollinger, Keltner, Donchian, etc.).
package bands

import (
	"math"
	"sort"

	ind "github.com/Bitvested/ta.go/src/indicators"
	ma "github.com/Bitvested/ta.go/src/moving_averages"
	stats "github.com/Bitvested/ta.go/src/statistics"
)

func maxf(v []float64) float64 {
	sort.Float64s(v)
	return v[len(v)-1]
}

func minf(v []float64) float64 {
	sort.Float64s(v)
	return v[0]
}

// Bands is the Bollinger band — center is SMA(data, l) and the upper/lower
// envelopes are ±dev × pop-std. v2.0.0 perf: pre-allocated output and an
// inline running Σx / Σx² recurrence (variance = Σx²/l − mean²) so the
// hot path no longer rebuilds windows or calls statistics.Std per bar.
func Bands(data []float64, l int, dev float64) [][]float64 {
	n := len(data)
	if n < l {
		return nil
	}
	lf := float64(l)
	var sumX, sumX2 float64
	for i := 0; i < l; i++ {
		sumX += data[i]
		sumX2 += data[i] * data[i]
	}
	rows := n - l + 1
	flat := make([]float64, 3*rows)
	out := make([][]float64, rows)
	for i := range out {
		out[i] = flat[i*3 : i*3+3 : i*3+3]
	}
	mean := sumX / lf
	v := sumX2/lf - mean*mean
	if v < 0 {
		v = 0
	}
	std := math.Sqrt(v)
	out[0][0], out[0][1], out[0][2] = mean+std*dev, mean, mean-std*dev
	for i := l; i < n; i++ {
		incoming, outgoing := data[i], data[i-l]
		sumX += incoming - outgoing
		sumX2 += incoming*incoming - outgoing*outgoing
		mean = sumX / lf
		v = sumX2/lf - mean*mean
		if v < 0 {
			v = 0
		}
		std = math.Sqrt(v)
		row := out[i-l+1]
		row[0], row[1], row[2] = mean+std*dev, mean, mean-std*dev
	}
	return out
}

func Bandwidth(data []float64, l int, dev float64) []float64 {
	var boll []float64
	band := Bands(data, l, dev)
	for i := 0; i < len(band); i++ {
		boll = append(boll, (band[i][0]-band[i][2])/band[i][1])
	}
	return boll
}

// Keltner channels: Sma(typical) ± dev * Atr. v2.0.0: inherits Atr's new
// shape (N - L + 1), so the compensating splice is removed; sma and atr are
// already the same length and tail-aligned.
func Keltner(data [][]float64, l int, dev float64) [][]float64 {
	closing := make([]float64, len(data))
	for i := range data {
		closing[i] = (data[i][0] + data[i][1] + data[i][2]) / 3
	}
	atr := ind.Atr(data, l)
	kma := ma.Sma(closing, l)
	if len(atr) != len(kma) {
		// Defensive: keep the shorter tail.
		if len(atr) > len(kma) {
			atr = atr[len(atr)-len(kma):]
		} else {
			kma = kma[len(kma)-len(atr):]
		}
	}
	kelt := make([][]float64, 0, len(kma))
	for i := range kma {
		kelt = append(kelt, []float64{kma[i] + atr[i]*dev, kma[i], kma[i] - atr[i]*dev})
	}
	return kelt
}

// Don is the Donchian channel: rolling max-of-highs and min-of-lows over
// a window of `l` bars. Bar format is [high, low].
//
// v2.0.0 perf: single linear scan per window with no slice copies / sort,
// pre-allocated [][]float64 output backed by one flat slice.
func Don(data [][]float64, l int) [][]float64 {
	n := len(data)
	if n < l {
		return nil
	}
	rows := n - l + 1
	flat := make([]float64, 3*rows)
	out := make([][]float64, rows)
	for i := range out {
		out[i] = flat[i*3 : i*3+3 : i*3+3]
	}
	for i := l; i <= n; i++ {
		hi := math.Inf(-1)
		lo := math.Inf(1)
		for j := i - l; j < i; j++ {
			if data[j][0] > hi {
				hi = data[j][0]
			}
			if data[j][1] < lo {
				lo = data[j][1]
			}
		}
		row := out[i-l]
		row[0], row[1], row[2] = hi, (hi+lo)/2, lo
	}
	return out
}

func Fibbands(data [][]float64, length int, deviations float64) [][]float64 {
	var pl []float64
	var deviation []float64
	mavg := ma.Vwma(data, length)
	var boll [][]float64
	for i := 0; i < len(data); i++ {
		pl = append(pl, data[i][0])
		if len(pl) >= length {
			devi := stats.Std(pl, length)
			deviation = append(deviation, devi*deviations)
			pl = pl[1:]
		}
	}
	for i := 0; i < len(mavg); i++ {
		upper1 := mavg[i] + (0.236 * deviation[i])
		upper2 := mavg[i] + 0.382*deviation[i]
		upper3 := mavg[i] + 0.5*deviation[i]
		upper4 := mavg[i] + 0.618*deviation[i]
		upper5 := mavg[i] + 0.764*deviation[i]
		upper6 := mavg[i] + deviation[i]
		lower1 := mavg[i] - 0.236*deviation[i]
		lower2 := mavg[i] - 0.382*deviation[i]
		lower3 := mavg[i] - 0.5*deviation[i]
		lower4 := mavg[i] - 0.618*deviation[i]
		lower5 := mavg[i] - 0.764*deviation[i]
		lower6 := mavg[i] - deviation[i]
		boll = append(boll, []float64{upper6, upper5, upper4, upper3, upper2, upper1, mavg[i], lower1, lower2, lower3, lower4, lower5, lower6})
	}
	return boll
}

func Envelope(data []float64, l int, p float64) [][]float64 {
	var enve [][]float64
	for i := l; i < len(data); i++ {
		sm := ma.Sma(data[i-l:i], l)
		enve = append(enve, []float64{sm[0] * (1 + p), sm[0], sm[0] * (1 - p)})
	}
	return enve
}
