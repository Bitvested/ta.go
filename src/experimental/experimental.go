// Package experimental contains prototype indicators that have not yet
// stabilized into a permanent place in the API.
package experimental

import (
	"sort"

	stats "github.com/Bitvested/ta.go/src/statistics"
)

// Line represents a linear-trend descriptor used by Support and Resistance.
// Index/Slope/Value describe the line at the anchor index.
type Line struct {
	Slope float64
	Value float64
	Index int
}

// LineCalc evaluates a Line at position pos.
func LineCalc(pos int, calc Line) float64 {
	return float64(pos)*calc.Slope + calc.Value
}

// maxf and minf are non-mutating: the input slice is not sorted in place.
// The caller (Support / Resistance) holds slices that flow back into
// DivergenceState's window views, so mutation here would corrupt
// subsequent bars.
func maxf(v []float64) float64 {
	cp := append([]float64(nil), v...)
	sort.Float64s(cp)
	return cp[len(cp)-1]
}

func minf(v []float64) float64 {
	cp := append([]float64(nil), v...)
	sort.Float64s(cp)
	return cp[0]
}

func Support(d []float64, hl stats.RecentHighLow) Line {
	var index2 int
	var findex bool = false
	lowform := hl.Value
	for do := true; do; do = !findex {
		for i := hl.Index; i < len(d); i++ {
			newlow := (hl.Value - d[i]) / float64(hl.Index-i)
			if newlow < lowform {
				lowform = newlow
				index2 = i
			}
		}
		if hl.Index+1 == index2 && index2 != len(d)-1 {
			hl.Index = index2
			lowform = minf(d)
			hl.Value = d[hl.Index]
			findex = false
		} else {
			findex = true
		}
		if hl.Index == len(d)-1 {
			findex = true
		}
	}
	if index2 == len(d)-1 || hl.Index == len(d)-1 {
		return Line{Slope: 0, Value: hl.Value, Index: hl.Index}
	}
	return Line{Slope: lowform, Value: hl.Value, Index: hl.Index}
}

func Resistance(d []float64, hl stats.RecentHighLow) Line {
	var index2 int
	var findex bool = false
	highform := hl.Value
	for do := true; do; do = !findex {
		for i := hl.Index; i < len(d); i++ {
			newhigh := (d[i] - hl.Value) / float64(hl.Index-i)
			if newhigh < highform {
				highform = newhigh
				index2 = i
			}
		}
		if hl.Index+1 == index2 && index2 != len(d)-1 {
			hl.Index = index2
			highform = maxf(d)
			hl.Value = d[hl.Index]
			findex = false
		} else {
			findex = true
		}
		if hl.Index == len(d)-1 {
			findex = true
		}
	}
	if index2 == len(d)-1 || hl.Index == len(d)-1 {
		return Line{Slope: 0, Value: hl.Value, Index: hl.Index}
	}
	return Line{Slope: -highform, Value: hl.Value, Index: hl.Index}
}

// DivergenceState classifies the current bar's relation between two
// price series (e.g., price vs an oscillator) as convergence,
// divergence, or one of six labelled divergence regimes (regular,
// hidden, or exaggerated × bullish/bearish).
//
// length is the rolling window size, lb the look-back used by recent
// high/low extraction, smoother is the bar offset used to detect
// directional disagreement (default 1 in ta.js), and the thresholds
// gate the regime labels.
//
// v2.0.0: inputs are not mutated — when one series is longer the
// trim is via a non-mutating slice expression.
func DivergenceState(data1, data2 []float64, length, lb, smoother int, thresholdEx, thresholdNm float64) [][]string {
	if len(data1) > len(data2) {
		data1 = data1[len(data1)-len(data2):]
	}
	if len(data2) > len(data1) {
		data2 = data2[len(data2)-len(data1):]
	}
	var out [][]string
	for i := length; i < len(data1); i++ {
		pl1 := data1[i-length : i+1]
		s1 := Support(pl1, stats.RecentLow(pl1, lb))
		s1Delta := s1.Slope / s1.Value
		r1 := Resistance(pl1, stats.RecentHigh(pl1, lb))
		// Go's Resistance.Slope is the visual line slope (= −highform);
		// ta.js' .slope field is highform itself, so flip back here.
		r1Delta := -r1.Slope / r1.Value

		pl2 := data2[i-length : i+1]
		s2 := Support(pl2, stats.RecentLow(pl2, lb))
		s2Delta := s2.Slope / s2.Value
		r2 := Resistance(pl2, stats.RecentHigh(pl2, lb))
		r2Delta := -r2.Slope / r2.Value

		disagree := (data1[i] > data1[i-smoother] && data2[i] < data2[i-smoother]) ||
			(data1[i] < data1[i-smoother] && data2[i] > data2[i-smoother])
		if !disagree {
			out = append(out, []string{"convergence"})
			continue
		}
		var obj []string
		if r1Delta < -thresholdEx && r2Delta > -thresholdNm {
			obj = append(obj, "exaggerated_bearish")
		}
		if s1Delta < thresholdNm && s2Delta > thresholdEx {
			obj = append(obj, "exaggerated_bullish")
		}
		if r1Delta < -thresholdNm && r2Delta < thresholdNm {
			obj = append(obj, "hidden_bearish")
		}
		if s1Delta > thresholdNm && s2Delta < -thresholdNm {
			obj = append(obj, "hidden_bullish")
		}
		if r1Delta > thresholdNm && r2Delta < -thresholdNm {
			obj = append(obj, "regular_bearish")
		}
		if s1Delta < -thresholdNm && s2Delta > thresholdNm {
			obj = append(obj, "regular_bullish")
		}
		if len(obj) == 0 {
			obj = append(obj, "divergence")
		}
		out = append(out, obj)
	}
	return out
}
