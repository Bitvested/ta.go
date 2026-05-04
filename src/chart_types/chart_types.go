// Package charttypes contains alternative price-bar representations such as
// Heikin-Ashi and Renko.
package charttypes

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func maxf(v []float64) float64 {
	sort.Float64s(v)
	return v[len(v)-1]
}

func minf(v []float64) float64 {
	sort.Float64s(v)
	return v[0]
}

func decimalplaces(n float64) int {
	decimalPlaces := fmt.Sprintf("%f", n-math.Floor(n))
	decimalPlaces = strings.Replace(decimalPlaces, "0.", "", -1)
	decimalPlaces = strings.TrimRight(decimalPlaces, "0")
	return len(decimalPlaces)
}

func Ha(data [][]float64) [][]float64 {
	var ha [][]float64 = [][]float64{{(data[0][0] + data[0][3]) / 2, data[0][1], data[0][2], (data[0][0] + data[0][1] + data[0][2] + data[0][3]) / 4}}
	for i := 1; i < len(data); i++ {
		ha = append(ha, []float64{
			(ha[len(ha)-1][0] + ha[len(ha)-1][3]) / 2,
			maxf([]float64{ha[len(ha)-1][0], ha[len(ha)-1][3], data[i][1]}),
			minf([]float64{ha[len(ha)-1][0], ha[len(ha)-1][3], data[i][2]}),
			(data[i][0] + data[i][1] + data[i][2] + data[i][3]) / 4,
		})
	}
	return ha
}

func Ren(data [][]float64, bs float64) [][]float64 {
	var re [][]float64
	decimals := float64(decimalplaces(bs))
	bh := math.Ceil(data[0][0]/bs*math.Pow(10, decimals)) / math.Pow(10, decimals) * bs
	bl := bh - bs
	for i := 1; i < len(data); i++ {
		if data[i][0] > bh+bs {
			for ok := true; ok == true; ok = (data[i][0] > bh+bs) {
				re = append(re, []float64{bh, bh + bs, bh, bh + bs})
				bh += bs
				bl += bs
			}
		}
		if data[i][1] < bl-bs {
			for ok := true; ok == true; ok = (data[i][1] < bl-bs) {
				re = append(re, []float64{bl, bl, bl - bs, bl - bs})
				bh -= bs
				bl -= bs
			}
		}
	}
	return re
}
