// Package movingaverages contains all moving-average implementations
// used across the ta.go technical-analysis library.
package movingaverages

import "math"

// Sma is the simple moving average over `l` bars. v2.0.0 perf: O(n)
// running-sum recurrence (sum += new − old) and a single pre-allocated
// output slice — no per-window copy.
func Sma(data []float64, l int) []float64 {
	n := len(data)
	if n < l {
		return nil
	}
	out := make([]float64, n-l+1)
	var sum float64
	for i := 0; i < l; i++ {
		sum += data[i]
	}
	lf := float64(l)
	out[0] = sum / lf
	for i := l; i < n; i++ {
		sum += data[i] - data[i-l]
		out[i-l+1] = sum / lf
	}
	return out
}

// Ema is the exponential moving average over `l` bars. Seed is the SMA of
// the first `l` values; subsequent values use a single recurrence step.
//
// v2.0.0 perf: pre-allocated output, seed folded out of the inner loop,
// no per-iteration branch on whether the seed has been emitted yet.
func Ema(data []float64, l int) []float64 {
	n := len(data)
	if n < l {
		return nil
	}
	weight := 2 / (float64(l) + 1)
	out := make([]float64, n-l+1)
	var seed float64
	for i := 0; i < l; i++ {
		seed += data[i]
	}
	prev := seed / float64(l)
	out[0] = prev
	for i := l; i < n; i++ {
		prev = (data[i]-prev)*weight + prev
		out[i-l+1] = prev
	}
	return out
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

// Wma is the weighted moving average over `l` bars (linear weights 1..l).
// v2.0.0 perf: pre-allocated output, no per-window slice copy.
//
// The inner loop keeps the original `data[k] * (q+1) / weight` association
// (per-term divide rather than divide-once at the end) so the result is
// bitwise-identical to the previous implementation — Hull and other
// callers depend on exact fp equality.
func Wma(data []float64, l int) []float64 {
	n := len(data)
	if n < l {
		return nil
	}
	weight := float64(l * (l + 1) / 2)
	out := make([]float64, n-l+1)
	for i := l; i <= n; i++ {
		var avg float64
		base := i - l
		for q := 0; q < l; q++ {
			avg += data[base+q] * float64(q+1) / weight
		}
		out[i-l] = avg
	}
	return out
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
	// Round rather than truncate, matching ta.js's Math.round: int(l/2) shortens the
	// half-length at every odd length, and int(sqrt(l)) flattens the final smoothing
	// window whenever sqrt rounds up (lengths 8, 14, ...). Both put the output on a
	// different bar grid than the other ports.
	var sqn int = int(math.Round(math.Sqrt(float64(l))))
	var first []float64 = Wma(data, int(math.Round(float64(l)/2)))
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

// Kama is Kaufman's Adaptive Moving Average.
//
//	ER   = |data[i] - data[i-l1]| / sum(|data[k] - data[k-1]|, k = i-l1+1 .. i)
//	SC   = (ER * (fast - slow) + slow) ^ 2,  fast = 2/(l2+1), slow = 2/(l3+1)
//	KAMA = KAMA[i-1] + SC * (data[i] - KAMA[i-1])
//
// Seeded with data[l1-1] and emitted from that bar on, so the output is
// n-l1+1 long. Matches the tulind reference. The volatility sum is carried
// across bars instead of re-summed per bar, which keeps the run linear.
func Kama(data []float64, l1 int, l2 int, l3 int) []float64 {
	n := len(data)
	if n < l1 {
		return nil
	}
	var vola float64
	for i := 1; i < l1; i++ {
		vola += math.Abs(data[i] - data[i-1])
	}
	fast, slow := 2/(float64(l2)+1), 2/(float64(l3)+1)
	ka := data[l1-1]
	out := make([]float64, 1, n-l1+1)
	out[0] = ka
	for i := l1; i < n; i++ {
		vola += math.Abs(data[i] - data[i-1])
		if i > l1 {
			vola -= math.Abs(data[i-l1] - data[i-l1-1])
		}
		// A window with no movement leaves the ratio undefined; treat it as a
		// one-sided move, which pins the smoothing constant to `fast`.
		er := 1.0
		if vola != 0 {
			er = math.Abs(data[i]-data[i-l1]) / vola
		}
		sc := er*(fast-slow) + slow
		ka += sc * sc * (data[i] - ka)
		out = append(out, ka)
	}
	return out
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

// Dema is the Double Exponential Moving Average: 2·EMA − EMA(EMA).
// Output length = N − 2·l + 2 (returns nil if N < 2·l − 1).
func Dema(data []float64, l int) []float64 {
	n := len(data)
	if n < 2*l-1 {
		return nil
	}
	alpha := 2.0 / (float64(l) + 1)
	var sum float64
	for i := 0; i < l; i++ {
		sum += data[i]
	}
	e1 := sum / float64(l)
	sumE1 := e1
	for i := l; i < 2*l-1; i++ {
		e1 = (data[i]-e1)*alpha + e1
		sumE1 += e1
	}
	e2 := sumE1 / float64(l)
	out := make([]float64, n-2*l+2)
	out[0] = 2*e1 - e2
	for i := 2*l - 1; i < n; i++ {
		e1 = (data[i]-e1)*alpha + e1
		e2 = (e1-e2)*alpha + e2
		out[i-2*l+2] = 2*e1 - e2
	}
	return out
}

// Tema is the Triple Exponential Moving Average: 3·EMA − 3·EMA² + EMA³.
func Tema(data []float64, l int) []float64 {
	e1 := Ema(data, l)
	if len(e1) < l {
		return nil
	}
	e2 := Ema(e1, l)
	if len(e2) < l {
		return nil
	}
	e3 := Ema(e2, l)
	off1 := len(e1) - len(e3)
	off2 := len(e2) - len(e3)
	out := make([]float64, len(e3))
	for i := 0; i < len(e3); i++ {
		out[i] = 3*e1[i+off1] - 3*e2[i+off2] + e3[i]
	}
	return out
}

// T3 is the Tillson T3 moving average (six cascaded EMAs with vfactor weights).
func T3(data []float64, l int, vfactor float64) []float64 {
	e1 := Ema(data, l)
	if len(e1) < l {
		return nil
	}
	e2 := Ema(e1, l)
	if len(e2) < l {
		return nil
	}
	e3 := Ema(e2, l)
	if len(e3) < l {
		return nil
	}
	e4 := Ema(e3, l)
	if len(e4) < l {
		return nil
	}
	e5 := Ema(e4, l)
	if len(e5) < l {
		return nil
	}
	e6 := Ema(e5, l)
	v2 := vfactor * vfactor
	v3 := v2 * vfactor
	c1 := -v3
	c2 := 3*v2 + 3*v3
	c3 := -6*v2 - 3*vfactor - 3*v3
	c4 := 1 + 3*vfactor + v3 + 3*v2
	o3 := len(e3) - len(e6)
	o4 := len(e4) - len(e6)
	o5 := len(e5) - len(e6)
	out := make([]float64, len(e6))
	for i := 0; i < len(e6); i++ {
		out[i] = c1*e6[i] + c2*e5[i+o5] + c3*e4[i+o4] + c4*e3[i+o3]
	}
	return out
}

// Trima is the triangular moving average — SMA of an SMA, with inner/outer
// window sizes chosen so the combined response is triangular.
func Trima(data []float64, l int) []float64 {
	var nInner, nOuter int
	if l%2 == 0 {
		nInner = l/2 + 1
		nOuter = l / 2
	} else {
		nInner = (l + 1) / 2
		nOuter = nInner
	}
	return Sma(Sma(data, nInner), nOuter)
}

// Vidya is Chande's Variable Index Dynamic Average. The smoothing
// coefficient scales with the ratio of short- to long-window pop std-dev.
func Vidya(data []float64, shortLength, longLength int, alpha float64) []float64 {
	n := len(data)
	if n < longLength {
		return nil
	}
	prev := data[longLength-2]
	out := []float64{prev}
	for i := longLength - 1; i < n; i++ {
		sShort := windowPopStd(data[i-shortLength+1 : i+1])
		sLong := windowPopStd(data[i-longLength+1 : i+1])
		var k float64
		if sLong != 0 {
			k = sShort / sLong
		}
		prev = alpha*k*data[i] + (1-alpha*k)*prev
		out = append(out, prev)
	}
	return out
}

// windowPopStd is the population std-dev of an entire slice. Inlined to
// avoid a circular import with the statistics package (which depends on
// movingaverages).
func windowPopStd(data []float64) float64 {
	n := len(data)
	if n == 0 {
		return 0
	}
	var sum float64
	for _, v := range data {
		sum += v
	}
	mean := sum / float64(n)
	var sq float64
	for _, v := range data {
		d := v - mean
		sq += d * d
	}
	return math.Sqrt(sq / float64(n))
}

// Zlema is the Zero-Lag Exponential Moving Average. It applies de-lagging
// adjustment 2·data[i] − data[i−lag] before EMA smoothing.
func Zlema(data []float64, l int) []float64 {
	n := len(data)
	if n == 0 {
		return nil
	}
	lag := (l - 1) / 2
	alpha := 2.0 / (float64(l) + 1)
	prev := data[0]
	out := make([]float64, n)
	out[0] = prev
	for i := 1; i < n; i++ {
		var adj float64
		if i >= lag {
			adj = 2*data[i] - data[i-lag]
		} else {
			adj = data[i]
		}
		prev = alpha*adj + (1-alpha)*prev
		out[i] = prev
	}
	return out
}
