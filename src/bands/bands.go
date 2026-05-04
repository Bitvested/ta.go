// Package bands contains envelope/channel indicators that wrap a moving
// average with a deviation band (Bollinger, Keltner, Donchian, etc.).
package bands

import (
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

func Bands(data []float64, l int, dev float64) [][]float64 {
	var pl []float64
	var deviation []float64
	var boll [][]float64
	sma := ma.Sma(data, l)
	for i := 0; i < len(data); i++ {
		pl = append(pl, data[i])
		if len(pl) >= l {
			devi := stats.Std(pl, l)
			deviation = append(deviation, devi)
			pl = pl[1:]
		}
	}
	for i := 0; i < len(sma); i++ {
		boll = append(boll, []float64{sma[i] + deviation[i]*dev, sma[i], sma[i] - deviation[i]*dev})
	}
	return boll
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

func Don(data [][]float64, l int) [][]float64 {
	var channel [][]float64
	for i := l; i <= len(data); i++ {
		pl := data[i-l : i]
		var highs []float64
		var lows []float64
		for h := 0; h < len(pl); h++ {
			highs = append(highs, pl[h][0])
			lows = append(lows, pl[h][1])
		}
		max := maxf(highs)
		min := minf(lows)
		channel = append(channel, []float64{max, (max + min) / 2, min})
	}
	return channel
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
