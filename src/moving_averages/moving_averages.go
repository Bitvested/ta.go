// Package movingaverages contains all moving-average implementations
// used across the ta.go technical-analysis library.
package movingaverages

import "math"

func Sma(data []float64, l int) []float64 {
	var sma []float64
	for i := l; i <= len(data); i++ {
		var avg float64
		tmp := append([]float64(nil), data[i-l:i]...)
		for x := 0; x < len(tmp); x++ {
			avg += tmp[x]
		}
		sma = append(sma, avg/float64(l))
	}
	return sma
}

func Ema(data []float64, l int) []float64 {
	var ema []float64
	var weight float64 = 2 / (float64(l) + 1)
	for i := l; i <= len(data); i++ {
		if len(ema) > 0 {
			ema = append(ema, (data[i-1]-ema[len(ema)-1])*weight+ema[len(ema)-1])
			continue
		}
		pl := append([]float64(nil), data[i-l:i]...)
		var average float64
		for q := 0; q < len(pl); q++ {
			average += pl[q]
		}
		ema = append(ema, average/float64(l))
	}
	return ema
}

// Smma is the Wilder smoothed moving average. Seed = SMA of the first
// `l` values; subsequent values use SMMA[i] = SMMA[i-1] + (data[i] - SMMA[i-1]) / l.
// Output length is len(data) - l + 1, matching sibling MA functions.
func Smma(data []float64, l int) []float64 {
	n := len(data)
	if n < l {
		return nil
	}
	var seed float64
	for i := 0; i < l; i++ {
		seed += data[i]
	}
	prev := seed / float64(l)
	out := []float64{prev}
	for i := l; i < n; i++ {
		prev = prev + (data[i]-prev)/float64(l)
		out = append(out, prev)
	}
	return out
}

func Wma(data []float64, l int) []float64 {
	var weight int
	var wma []float64
	for i := 1; i <= l; i++ {
		weight += i
	}
	for i := l; i <= len(data); i++ {
		pl := append([]float64(nil), data[i-l:i]...)
		var average float64
		for q := 0; q < l; q++ {
			average += (pl[q] * float64(q+1) / float64(weight))
		}
		wma = append(wma, average)
	}
	return wma
}

func Pwma(data []float64, l int) []float64 {
	var weight int
	var wmaa []float64
	var b int = l
	var weights []float64
	for i := int(float64(l) / 2); i >= 1; i-- {
		if i%1 != 0 {
			weight += (i * b)
		} else {
			weights = append(weights, float64(i*b))
			weight += (i * b * 2)
		}
		weights = append([]float64{float64(i * b)}, weights...)
		b--
	}
	for i := l; i <= len(data); i++ {
		var average float64
		pl := append([]float64(nil), data[i-l:i]...)
		for a := l - 1; a >= 0; a-- {
			average += (pl[a] * weights[a] / float64(weight))
		}
		wmaa = append(wmaa, average)
	}
	return wmaa
}

func Hwma(data []float64, l int) []float64 {
	var weight int
	var wmaa []float64
	var b int = l
	var weights []float64
	for i := 1; i <= int(l/2); i++ {
		if i%1 != 0 {
			weight += (i * b)
		} else {
			weights = append(weights, float64(i*b))
			weight += (i * b * 2)
		}
		weights = append([]float64{float64(i * b)}, weights...)
		b--
	}
	for i := l; i <= len(data); i++ {
		var average float64
		pl := append([]float64(nil), data[i-l:i]...)
		for a := l - 1; a >= 0; a-- {
			average += (pl[a] * weights[a] / float64(weight))
		}
		wmaa = append(wmaa, average)
	}
	return wmaa
}

func Vwma(data [][]float64, l int) []float64 {
	var vwma []float64
	var weighted [][]float64
	for i := 0; i < len(data); i++ {
		weighted = append(weighted, []float64{data[i][0] * data[i][1], data[i][1]})
	}
	for i := l; i <= len(data); i++ {
		pl := append([][]float64(nil), weighted[i-l:i]...)
		var totalv float64
		var totalp float64
		for o := 0; o < len(pl); o++ {
			totalv += pl[o][1]
			totalp += pl[o][0]
		}
		vwma = append(vwma, totalp/totalv)
	}
	return vwma
}

func Vwwma(data [][]float64, l int) []float64 {
	var vwma []float64
	var weight float64 = 0
	var weighted [][]float64
	for i := 0; i < len(data); i++ {
		weighted = append(weighted, []float64{data[i][0] * data[i][1], data[i][1]})
	}
	for i := 1; i <= l; i++ {
		weight += float64(i)
	}
	for i := l; i <= len(data); i++ {
		pl := append([][]float64(nil), weighted[i-l:i]...)
		var totalv float64
		var totalp float64
		for o := 0; o < len(pl); o++ {
			totalv += pl[o][1] * float64(o+1) / weight
			totalp += pl[o][0] * float64(o+1) / weight
		}
		vwma = append(vwma, totalp/totalv)
	}
	return vwma
}

func Wsma(data []float64, l int) []float64 {
	var wsm []float64
	var weight float64 = 1 / float64(l)
	for i := l; i <= len(data); i++ {
		if len(wsm) > 0 {
			wsm = append(wsm, (data[i-1]-wsm[len(wsm)-1])*weight+wsm[len(wsm)-1])
			continue
		}
		pl := append([]float64(nil), data[i-l:i]...)
		var average float64
		for q := 0; q < len(pl); q++ {
			average += pl[q]
		}
		wsm = append(wsm, average/float64(l))
	}
	return wsm
}

func Lsma(data []float64, l int) []float64 {
	var lr []float64
	for i := l; i <= len(data); i++ {
		pl := append([]float64(nil), data[i-l:i]...)
		var sum_x float64 = 0
		var sum_y float64 = 0
		var sum_xy float64 = 0
		var sum_xx float64 = 0
		var sum_yy float64 = 0
		for q := 1; q <= l; q++ {
			sum_x += float64(q)
			sum_y += pl[q-1]
			sum_xy += (pl[q-1] * float64(q))
			sum_xx += (float64(q) * float64(q))
			sum_yy += (pl[q-1] * pl[q-1])
		}
		m := ((sum_xy - sum_x*sum_y/float64(l)) / (sum_xx - sum_x*sum_x/float64(l)))
		b := sum_y/float64(l) - m*sum_x/float64(l)
		lr = append(lr, m*float64(l)+b)
		_ = sum_yy
	}
	return lr
}

func Hull(data []float64, l int) []float64 {
	var pl []float64
	var hma []float64
	var ewma []float64 = Wma(data, l)
	var sqn int = int(math.Pow(float64(l), 0.5))
	var first []float64 = Wma(data, int(l/2))
	first = first[len(first)-len(ewma):]
	for i := 0; i < len(ewma); i++ {
		pl = append(pl, (first[i]*2)-ewma[i])
		if len(pl) >= sqn {
			var h []float64 = Wma(pl, sqn)
			hma = append(hma, h[0])
			pl = pl[1:]
		}
	}
	return hma
}

func Kama(data []float64, l1 int, l2 int, l3 int) []float64 {
	ka := Sma(data, l1)
	ka = []float64{ka[len(ka)-1]}
	for i := l1 + 1; i < len(data); i++ {
		var vola float64
		change := math.Abs(data[i] - data[i-l1])
		for a := 1; a < l1; a++ {
			vola += math.Abs(data[i-a] - data[i-a-1])
		}
		sc := math.Pow(change/vola*(2/(float64(l2)+1)-2/(float64(l3)+1)+2/(float64(l3)+1)), 2)
		ka = append(ka, ka[len(ka)-1]+sc*(data[i]-ka[len(ka)-1]))
	}
	return ka
}

func Cwma(data []float64, weights []float64) []float64 {
	var ma []float64
	for i := len(weights); i <= len(data); i++ {
		pl := data[i-len(weights) : i]
		var sum float64 = 0
		var weight float64 = 0
		for q := 0; q < len(weights); q++ {
			sum += pl[q] * weights[q]
			weight += weights[q]
		}
		ma = append(ma, sum/float64(weight))
	}
	return ma
}
