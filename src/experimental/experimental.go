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

func maxf(v []float64) float64 {
	sort.Float64s(v)
	return v[len(v)-1]
}

func minf(v []float64) float64 {
	sort.Float64s(v)
	return v[0]
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
			lowform = minf(d[:])
			hl.Value = d[hl.Index]
			findex = true
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
			highform = maxf(d[:])
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
