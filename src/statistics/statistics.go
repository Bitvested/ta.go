// Package statistics contains statistical primitives, distribution helpers,
// position-sizing utilities, and other math support for ta.go indicators.
package statistics

import (
	"math"
	"math/rand"
	"sort"

	ma "github.com/Bitvested/ta.go/src/moving_averages"
)

// RecentHighLow is the value returned by RecentHigh and RecentLow. Index is
// the position of the extreme inside the input series; Value is the price.
type RecentHighLow struct {
	Index int
	Value float64
}

// Median is the rolling median. v2.0.0: even-window length now returns the
// average of the two middle sorted values (was lower-middle).
func Median(data []float64, l int) []float64 {
	n := len(data)
	if n < l {
		return nil
	}
	out := make([]float64, 0, n-l+1)
	for i := l; i <= n; i++ {
		pl := append([]float64(nil), data[i-l:i]...)
		sort.Float64s(pl)
		half := l / 2
		if l%2 == 1 {
			out = append(out, pl[half])
		} else {
			out = append(out, (pl[half-1]+pl[half])/2)
		}
	}
	return out
}

func Normalize(data []float64, m float64) []float64 {
	var normal []float64
	tmp := append([]float64(nil), data...)
	sort.Float64s(tmp)
	var max float64 = tmp[len(tmp)-1] * (1 + m)
	var min float64 = tmp[0] * (1 - m)
	for i := 0; i < len(data); i++ {
		normal = append(normal, 1-(max-data[i])/(max-min))
	}
	return normal
}

func Denormalize(norm []float64, data []float64, m float64) []float64 {
	var denormal []float64
	tmp := append([]float64(nil), data...)
	sort.Float64s(tmp)
	var max float64 = tmp[len(tmp)-1] * (1 + m)
	var min float64 = tmp[0] * (1 - m)
	for i := 0; i < len(norm); i++ {
		denormal = append(denormal, min+norm[i]*(max-min))
	}
	return denormal
}

func Standardize(data []float64) []float64 {
	var mean []float64 = ma.Sma(data, len(data))
	var std float64 = Std(data, len(data))
	var res []float64
	for i := 0; i < len(data); i++ {
		res = append(res, (data[i]-mean[0])/std)
	}
	return res
}

func Mad(data []float64, l int) []float64 {
	var med []float64
	for i := l; i <= len(data); i++ {
		tmp := append([]float64(nil), data[i-l:i]...)
		m := Median(tmp, len(tmp))
		var adev []float64
		for x := 0; x < len(tmp); x++ {
			adev = append(adev, math.Abs(tmp[x]-m[0]))
		}
		adev = Median(adev, len(adev))
		med = append(med, adev[0])
	}
	return med
}

func Aad(data []float64, l int) []float64 {
	var med []float64
	for i := l; i <= len(data); i++ {
		tmp := append([]float64(nil), data[i-l:i]...)
		m := ma.Sma(tmp, len(tmp))
		var sum float64 = 0
		for x := 0; x < len(tmp); x++ {
			sum += math.Abs(float64(tmp[x]) - float64(m[0]))
		}
		med = append(med, sum/float64(l))
	}
	return med
}

func Ssd(data []float64, l int) []float64 {
	var sd []float64
	for i := l; i <= len(data); i++ {
		tmp := append([]float64(nil), data[i-l:i]...)
		mean := ma.Sma(tmp, l)
		var sum float64
		for x := 0; x < len(tmp); x++ {
			sum += math.Pow(tmp[x]-mean[0], 2)
		}
		sd = append(sd, math.Pow(sum, 0.5))
	}
	return sd
}

// Std is the population standard deviation over the trailing `l` values.
// v2.0.0: no longer mutates the caller's slice, and computes the mean
// inline with a running sum (avoids the round-trip through ma.Sma).
func Std(data []float64, l int) float64 {
	n := len(data)
	start := 0
	if l < n {
		start = n - l
	}
	var sum float64
	for i := start; i < n; i++ {
		sum += data[i]
	}
	mean := sum / float64(l)
	var sq float64
	for i := start; i < n; i++ {
		d := data[i] - mean
		sq += d * d
	}
	return math.Sqrt(sq / float64(l))
}

// Variance is the rolling population variance over windows of length `l`.
// v2.0.0 perf: running Σx and Σx² (variance = Σx²/l − mean²) with a
// pre-allocated output and a v < 0 clamp guarding against fp cancellation.
func Variance(data []float64, l int) []float64 {
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
	out := make([]float64, n-l+1)
	mean := sumX / lf
	v := sumX2/lf - mean*mean
	if v < 0 {
		v = 0
	}
	out[0] = v
	for i := l; i < n; i++ {
		incoming, outgoing := data[i], data[i-l]
		sumX += incoming - outgoing
		sumX2 += incoming*incoming - outgoing*outgoing
		mean = sumX / lf
		v = sumX2/lf - mean*mean
		if v < 0 {
			v = 0
		}
		out[i-l+1] = v
	}
	return out
}

func Cor(data1 []float64, data2 []float64) float64 {
	var d1avg []float64 = ma.Sma(data1, len(data1))
	var d2avg []float64 = ma.Sma(data2, len(data2))
	var sumavg float64
	var sx float64
	var sy float64
	for i := 0; i < len(data1); i++ {
		x := data1[i] - d1avg[0]
		y := data2[i] - d2avg[0]
		sumavg += (x * y)
		sx += math.Pow(x, 2)
		sy += math.Pow(y, 2)
	}
	n := float64(len(data1) - 1)
	sx /= n
	sy /= n
	sx = math.Pow(sx, 0.5)
	sy = math.Pow(sy, 0.5)
	return (sumavg / (n * sx * sy))
}

// Dif is the relative change (n - o) / o. v2.0.0: returns NaN when o is 0
// instead of producing ±Infinity.
func Dif(n float64, o float64) float64 {
	if o == 0 {
		return math.NaN()
	}
	return (n - o) / o
}

func Drawdown(data []float64) float64 {
	max := data[0]
	min := data[0]
	big := 0.0
	for y := 1; y < len(data); y++ {
		if data[y] > max {
			if min != 0 {
				diff := Dif(min, max)
				if diff < big {
					big = diff
				}
				min = data[y]
			}
			max = data[y]
		}
		if data[y] < min {
			min = data[y]
		}
	}
	diff := Dif(min, max)
	if diff < big {
		big = diff
	}
	return big
}

func Normsinv(p float64) float64 {
	a1 := -39.6968302866538
	a2 := 220.946098424521
	a3 := -275.928510446969
	a4 := 138.357751867269
	a5 := -30.6647980661472
	a6 := 2.50662827745924
	b1 := -54.4760987982241
	b2 := 161.585836858041
	b3 := -155.698979859887
	b4 := 66.8013118877197
	b5 := -13.2806815528857
	c1 := -7.78489400243029e-03
	c2 := -0.322396458041136
	c3 := -2.40075827716184
	c4 := -2.54973253934373
	c5 := 4.37466414146497
	c6 := 2.93816398269878
	d1 := 7.78469570904146e-03
	d2 := 0.32246712907004
	d3 := 2.445134137143
	d4 := 3.75440866190742
	p_low := 0.02425
	p_high := 1 - p_low
	if p < 0 || p > 1 {
		return 0
	}
	if p < p_low {
		q := math.Pow(-2*math.Log(p), 0.5)
		return (((((c1*q+c2)*q+c3)*q+c4)*q+c5)*q + c6) / ((((d1*q+d2)*q+d3)*q+d4)*q + 1)
	}
	if p <= p_high {
		q := p - 0.5
		r := q * q
		return (((((a1*r+a2)*r+a3)*r+a4)*r+a5)*r + a6) * q / (((((b1*r+b2)*r+b3)*r+b4)*r+b5)*r + 1)
	}
	q := math.Pow(-2*math.Log(1-p), 0.5)
	return -(((((c1*q+c2)*q+c3)*q+c4)*q+c5)*q + c6) / ((((d1*q+d2)*q+d3)*q+d4)*q + 1)
}

func Ncdf(x float64, mean float64, std float64) float64 {
	x = (x - mean) / std
	t := 1 / (1 + .2315419*math.Abs(x))
	d := .3989423 * math.Exp(-x*x/2)
	p := d * t * (.3193815 + t*(-.3565638+t*(1.781478+t*(-1.821256+t*1.330274))))
	if x > 0 {
		return 1 - p
	}
	return p
}

func Sim(data []float64, l int, sims int) [][]float64 {
	var sd [][]float64
	for i := 0; i < sims; i++ {
		projected := append([]float64(nil), data...)
		for x := 0; x < l; x++ {
			var change []float64
			for y := 1; y < len(projected); y++ {
				df := Dif(projected[y], projected[y-1])
				change = append(change, df)
			}
			mean := ma.Sma(change, len(change))
			std := Std(change, len(change))
			random := Normsinv(rand.Float64())
			projected = append(projected, projected[len(projected)-1]*math.Exp((mean[0]-(std*std)/2)+std*random))
		}
		sd = append(sd, projected)
	}
	return sd
}

// Percentile returns, for each column of a 2D series, the value at the
// `perc` percentile (0..1). v2.0.0: no longer mutates the input rows —
// each column is copied to a fresh slice before sorting.
func Percentile(data [][]float64, perc float64) []float64 {
	if len(data) == 0 {
		return nil
	}
	cols := len(data[0])
	out := make([]float64, 0, cols)
	for i := 0; i < cols; i++ {
		col := make([]float64, len(data))
		for r := 0; r < len(data); r++ {
			col[r] = data[r][i]
		}
		sort.Float64s(col)
		idx := int(math.Round(float64(len(col)-1) * perc))
		out = append(out, col[idx])
	}
	return out
}

// Kmeans runs 1-D k-means clustering. v2.0.0 fix: actually iterates to
// convergence (the previous implementation reset means once and stopped).
// Each iteration: snapshot centers, assign points to nearest center,
// recompute centers as the mean of each cluster, loop until no center moves.
func Kmeans(data []float64, clusters int) [][]float64 {
	init := int(math.Round(float64(len(data)) / float64(clusters+1)))
	centers := make([]float64, clusters)
	for i := 0; i < clusters; i++ {
		centers[i] = data[init*(i+1)]
	}
	means := make([][]float64, clusters)
	changed := true
	for changed {
		changed = false
		for i := range means {
			means[i] = means[i][:0]
		}
		for x := 0; x < len(data); x++ {
			best := math.Abs(centers[0] - data[x])
			n := 0
			for y := 1; y < clusters; y++ {
				r := math.Abs(centers[y] - data[x])
				if r <= best {
					best = r
					n = y
				}
			}
			means[n] = append(means[n], data[x])
		}
		for x := 0; x < clusters; x++ {
			if len(means[x]) == 0 {
				continue
			}
			var sum float64
			for _, v := range means[x] {
				sum += v
			}
			m := sum / float64(len(means[x]))
			if m != centers[x] {
				centers[x] = m
				changed = true
			}
		}
	}
	return means
}

func Er(data []float64) float64 {
	var wins []float64
	var losses []float64
	wp := 1.0
	lp := 1.0
	for i := 0; i < len(data); i++ {
		if data[i] >= 0 {
			wins = append(wins, 1+data[i])
		} else {
			losses = append(losses, 1+data[i])
		}
	}
	var datlen float64 = float64(len(data))
	var loslen float64 = float64(len(losses))
	var winlen float64 = float64(len(wins))
	win := winlen / datlen
	loss := loslen / datlen
	for i := 0; i < int(loslen); i++ {
		lp = lp * losses[i]
	}
	for i := 0; i < int(winlen); i++ {
		wp = wp * wins[i]
	}
	var fw float64 = (1.0 / winlen)
	var fl float64 = (1.0 / loslen)
	return ((((math.Pow(wp, fw)) - 1) * 100) * win + (((math.Pow(lp, fl)) - 1) * 100) * loss) / 100
}

func Ar(data []float64, l int) []float64 {
	var out []float64
	for i := l; i < len(data); i++ {
		exp := Er(data[i-l : i])
		out = append(out, data[i]-exp)
	}
	return out
}

func Winratio(data []float64) float64 {
	var wins int = 0
	var losses int = 0
	for i := 0; i < len(data); i++ {
		if data[i] >= 0 {
			wins++
		} else {
			losses++
		}
	}
	return float64(wins) / float64(losses+wins)
}

func WinAverage(data []float64) float64 {
	var wins []float64
	for i := 0; i < len(data); i++ {
		if data[i] >= 0 {
			wins = append(wins, data[i])
		}
	}
	avg := ma.Sma(wins, len(wins))
	return avg[0]
}

func LossAverage(data []float64) float64 {
	var losses []float64
	for i := 0; i < len(data); i++ {
		if data[i] < 0 {
			losses = append(losses, data[i])
		}
	}
	avg := ma.Sma(losses, len(losses))
	return avg[0]
}

func Avgwin(data []float64) float64 {
	var wins []float64
	for i := 0; i < len(data); i++ {
		if data[i] >= 0 {
			wins = append(wins, data[i])
		}
	}
	avg := ma.Sma(wins, len(wins))
	return avg[0]
}

func Avgloss(data []float64) float64 {
	var loss []float64
	for i := 0; i < len(data); i++ {
		if data[i] < 0 {
			loss = append(loss, data[i])
		}
	}
	avg := ma.Sma(loss, len(loss))
	return avg[0]
}

func Se(data []float64, l int) float64 {
	stdv := Std(data, len(data))
	return stdv / math.Pow(float64(l), 0.5)
}

func Kelly(data []float64) float64 {
	exp := Er(data) + 1
	winr := Winratio(data)
	return winr - (1-winr)/exp
}

func Zscore(data []float64, l int) []float64 {
	var out []float64
	for i := l; i <= len(data); i++ {
		pl := data[i-l : i]
		mean := ma.Sma(pl, l)
		stdv := Std(pl, l)
		out = append(out, (data[i-1]-mean[0])/stdv)
	}
	return out
}

func NormalizePair(data1 []float64, data2 []float64) [][]float64 {
	f := (data1[0] + data2[0]) / 2
	var ret [][]float64
	ret = append(ret, []float64{f, f})
	for i := 1; i < len(data1); i++ {
		ret = append(ret, []float64{ret[len(ret)-1][0] * ((data1[i]-data1[i-1])/data1[i-1] + 1), ret[len(ret)-1][1] * ((data2[i]-data2[i-1])/data2[i-1] + 1)})
	}
	return ret
}

func NormalizeFrom(data []float64, value float64) []float64 {
	var ret = []float64{value}
	for i := 1; i < len(data); i++ {
		ret = append(ret, ret[len(ret)-1]*((data[i]-data[i-1])/data[i-1]+1))
	}
	return ret
}

func Log(data []float64) []float64 {
	var ret []float64
	for i := 0; i < len(data); i++ {
		ret = append(ret, math.Log(data[i]))
	}
	return ret
}

func Exp(data []float64) []float64 {
	var ret []float64
	for i := 0; i < len(data); i++ {
		ret = append(ret, math.Exp(data[i]))
	}
	return ret
}

func Martingale(data []float64, bet float64, max float64, multiplier float64) float64 {
	current := bet
	for i := 0; i < len(data); i++ {
		if data[i] < 0 {
			current *= multiplier
		} else if data[i] > 0 {
			current = bet
		}
	}
	return current
}

func Antimartingale(data []float64, bet float64, max float64, multiplier float64) float64 {
	current := bet
	for i := 0; i < len(data); i++ {
		if data[i] > 0 {
			current *= multiplier
		} else if data[i] < 0 {
			current = bet
		}
	}
	return current
}

func Sum(data []float64) float64 {
	var sum float64 = 0
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	return sum
}

func Permutations(data []float64) float64 {
	var total float64 = 1
	for i := 0; i < len(data); i++ {
		total *= data[i]
	}
	return total
}

func Mse(data1 []float64, data2 []float64) float64 {
	var err float64 = 0
	for i := 0; i < len(data1); i++ {
		err += math.Pow(data2[i]-data1[i], 2)
	}
	return err / float64(len(data1))
}

func Cum(data []float64, l int) []float64 {
	var res []float64
	for i := l; i <= len(data); i++ {
		res = append(res, Sum(data[i-l:i]))
	}
	return res
}

func RecentHigh(data []float64, lb int) RecentHighLow {
	xback := lb
	hindex := 0
	highest := data[len(data)-1]
	for i := len(data) - 2; i >= 0; i-- {
		if data[i] >= highest && xback > 0 {
			highest = data[i]
			hindex = i
			xback = lb
		} else {
			xback--
		}
		if xback <= 0 {
			break
		}
	}
	return RecentHighLow{Index: hindex, Value: highest}
}

func RecentLow(data []float64, lb int) RecentHighLow {
	xback := lb
	lindex := 0
	lowest := data[len(data)-1]
	for i := len(data) - 2; i >= 0; i-- {
		if data[i] <= lowest && xback > 0 {
			lowest = data[i]
			lindex = i
			xback = lb
		} else {
			xback--
		}
		if xback <= 0 {
			break
		}
	}
	return RecentHighLow{Index: lindex, Value: lowest}
}

var t_table [][]float64 = [][]float64{
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{1, .816, .765, .741, .727, .718, .711, .706, .703, .7, .697, .695, .694, .692, .691, .69, .689, .688, .688, .687, .686, .686, .685, .685, .684},
	{1.376, 1.061, .978, .941, .92, .906, .896, .889, .883, .879, .876, .873, .87, .868, .866, .865, .863, .862, .861, .86, .859, .858, .858, .857, .856},
	{1.963, 1.386, 1.25, 1.19, 1.156, 1.134, 1.119, 1.108, 1.1, 1.093, 1.088, 1.088, 1.079, 1.076, 1.074, 1.071, 1.069, 1.067, 1.066, 1.064, 1.063, 1.061, 1.06, 1.059, 1.058},
	{3.078, 1.886, 1.638, 1.533, 1.476, 1.44, 1.415, 1.397, 1.383, 1.372, 1.363, 1.356, 1.35, 1.345, 1.341, 1.337, 1.333, 1.33, 1.328, 1.325, 1.323, 1.321, 1.319, 1.318, 1.316},
	{6.314, 2.92, 2.353, 2.132, 2.015, 1.943, 1.895, 1.86, 1.833, 1.812, 1.796, 1.782, 1.771, 1.761, 1.753, 1.746, 1.74, 1.734, 1.729, 1.725, 1.721, 1.717, 1.714, 1.711, 1.708},
	{12.71, 4.303, 3.182, 2.776, 2.571, 2.447, 2.365, 2.306, 2.262, 2.228, 2.201, 2.179, 2.16, 2.145, 2.131, 2.12, 2.11, 2.101, 2.093, 2.086, 2.08, 2.074, 2.069, 2.064, 2.06},
	{31.82, 6.965, 4.541, 3.747, 3.365, 3.143, 2.998, 2.896, 2.821, 2.764, 2.718, 2.681, 2.65, 2.624, 2.602, 2.583, 2.567, 2.552, 2.539, 2.528, 2.518, 2.508, 2.5, 2.492, 2.485},
	{63.66, 9.925, 5.841, 4.604, 4.032, 3.707, 3.499, 3.355, 3.25, 3.169, 3.106, 3.055, 3.012, 2.977, 2.947, 2.921, 2.898, 2.878, 2.861, 2.845, 2.831, 2.819, 2.807, 2.797, 2.787},
	{318.13, 22.327, 10.215, 7.173, 5.893, 5.208, 4.785, 4.501, 4.297, 4.144, 4.025, 3.93, 3.852, 3.787, 3.733, 3.686, 3.646, 3.61, 3.579, 3.552, 3.527, 3.505, 3.485, 3.467, 3.45},
	{636.62, 31.599, 12.924, 8.61, 6.869, 5.959, 5.408, 5.041, 4.781, 4.587, 4.437, 4.318, 4.221, 4.14, 4.073, 4.015, 3.965, 3.922, 3.883, 3.85, 3.819, 3.792, 3.768, 3.745, 3.725},
}
var v_table []float64 = []float64{.5, .25, .2, .15, .1, .05, .025, .01, .005, .001, .0005}

func Pvalue(t float64, df int) float64 {
	if df > 25 || df < 1 {
		return 1
	}
	for x := 0; x < len(t_table)-1; x++ {
		if t >= t_table[x][df-1] && t <= t_table[x+1][df-1] {
			return v_table[x+1] + (v_table[x+1]-v_table[x])*(t_table[x+1][df-1]-t)/(t_table[x][df-1]-t_table[x+1][df-1])
		}
	}
	return 0.0005
}

// Covariance is the rolling sample covariance over windows of length `l`.
// Both inputs must have the same length; returns nil if shorter than `l`.
func Covariance(data1, data2 []float64, l int) []float64 {
	n := len(data1)
	if n < l || len(data2) < l {
		return nil
	}
	out := make([]float64, 0, n-l+1)
	for z := l; z <= n; z++ {
		var xMean, yMean float64
		for i := z - l; i < z; i++ {
			xMean += data1[i]
			yMean += data2[i]
		}
		xMean /= float64(l)
		yMean /= float64(l)
		var sum float64
		for i := z - l; i < z; i++ {
			sum += (data1[i] - xMean) * (data2[i] - yMean)
		}
		out = append(out, sum/float64(l))
	}
	return out
}

// ExpectedTrails is the expected number of trials needed to observe
// every outcome of an `n`-sided fair die — the harmonic-sum estimate
// `ceil(n · H_n)` (coupon-collector heuristic).
func ExpectedTrails(n int) int {
	var sum float64
	for i := 1; i <= n; i++ {
		sum += 1 / float64(i)
	}
	return int(math.Ceil(float64(n) * sum))
}

// lrWindows is the rolling linear-regression decomposition over `l` bars
// using x = 0..l-1. Returns one [slope, intercept] pair per window —
// shared internally by LrSlope, LrIntercept, LrAngle, and Tsf.
func lrWindows(data []float64, l int) [][2]float64 {
	n := len(data)
	if n < l {
		return nil
	}
	sumX := float64((l - 1) * l) / 2
	sumX2 := float64((l-1)*l*(2*l-1)) / 6
	denom := float64(l)*sumX2 - sumX*sumX
	out := make([][2]float64, n-l+1)
	for i := l; i <= n; i++ {
		var sumY, sumXY float64
		for q := 0; q < l; q++ {
			y := data[i-l+q]
			sumY += y
			sumXY += y * float64(q)
		}
		var m float64
		if denom == 0 {
			m = math.NaN()
		} else {
			m = (float64(l)*sumXY - sumX*sumY) / denom
		}
		b := (sumY - m*sumX) / float64(l)
		out[i-l] = [2]float64{m, b}
	}
	return out
}

// LrSlope is the slope component of the rolling linear regression of
// `data` over `l` bars (TA-Lib LINEARREG_SLOPE).
func LrSlope(data []float64, l int) []float64 {
	w := lrWindows(data, l)
	out := make([]float64, len(w))
	for i := range w {
		out[i] = w[i][0]
	}
	return out
}

// LrIntercept is the intercept component of the rolling linear regression
// (TA-Lib LINEARREG_INTERCEPT).
func LrIntercept(data []float64, l int) []float64 {
	w := lrWindows(data, l)
	out := make([]float64, len(w))
	for i := range w {
		out[i] = w[i][1]
	}
	return out
}

// LrAngle is `atan(slope)` converted to degrees per window.
func LrAngle(data []float64, l int) []float64 {
	w := lrWindows(data, l)
	const radToDeg = 180 / math.Pi
	out := make([]float64, len(w))
	for i := range w {
		out[i] = math.Atan(w[i][0]) * radToDeg
	}
	return out
}

// Tsf is the Time Series Forecast — the linear-regression projection
// one bar past the end of each window: `slope · l + intercept`.
func Tsf(data []float64, l int) []float64 {
	w := lrWindows(data, l)
	out := make([]float64, len(w))
	for i := range w {
		out[i] = w[i][0]*float64(l) + w[i][1]
	}
	return out
}

// StdSeries is the rolling population std-dev over windows of length `l`,
// computed via running sumX and sumX² (O(1) per step after seeding).
func StdSeries(data []float64, l int) []float64 {
	n := len(data)
	if n < l {
		return nil
	}
	var sumX, sumX2 float64
	for i := 0; i < l; i++ {
		sumX += data[i]
		sumX2 += data[i] * data[i]
	}
	out := make([]float64, 0, n-l+1)
	out = append(out, stdSeriesStep(sumX, sumX2, l))
	for i := l; i < n; i++ {
		incoming, outgoing := data[i], data[i-l]
		sumX += incoming - outgoing
		sumX2 += incoming*incoming - outgoing*outgoing
		out = append(out, stdSeriesStep(sumX, sumX2, l))
	}
	return out
}

func stdSeriesStep(sumX, sumX2 float64, l int) float64 {
	mean := sumX / float64(l)
	v := sumX2/float64(l) - mean*mean
	if v < 0 {
		v = 0
	}
	return math.Sqrt(v)
}

// ReturnPositive returns only the strictly positive entries of `data`,
// preserving order.
func ReturnPositive(data []float64) []float64 {
	out := make([]float64, 0)
	for _, v := range data {
		if v > 0 {
			out = append(out, v)
		}
	}
	return out
}

// ReturnNegative returns only the strictly negative entries of `data`,
// preserving order.
func ReturnNegative(data []float64) []float64 {
	out := make([]float64, 0)
	for _, v := range data {
		if v < 0 {
			out = append(out, v)
		}
	}
	return out
}
