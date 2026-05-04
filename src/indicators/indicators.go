// Package indicators contains trend, momentum, volatility, and trading
// indicators built on top of the ta.go moving-average and statistics packages.
package indicators

import (
	"math"
	"sort"

	ma "github.com/Bitvested/ta.go/src/moving_averages"
	stats "github.com/Bitvested/ta.go/src/statistics"
)

// --- internal helpers ---------------------------------------------------------

func maxf(v []float64) float64 {
	sort.Float64s(v)
	return v[len(v)-1]
}

func minf(v []float64) float64 {
	sort.Float64s(v)
	return v[0]
}

func intMax(v []int) int {
	sort.Ints(v)
	return v[len(v)-1]
}

func intMin(v []int) int {
	sort.Ints(v)
	return v[0]
}

func reverseFloats(data []float64) []float64 {
	var ret []float64
	for i := len(data) - 1; i >= 0; i-- {
		ret = append(ret, data[i])
	}
	return ret
}

func indexOf(data []float64, index float64) int {
	var ret int = -1
	for i := len(data) - 1; i >= 0; i-- {
		if data[i] == index {
			ret = i
			break
		}
	}
	return ret
}

func map2d(data [][]float64, key int) []float64 {
	var ret []float64
	for i := 0; i < len(data); i++ {
		ret = append(ret, data[i][key])
	}
	return ret
}

// --- exported indicators ------------------------------------------------------

func Rsi(data []float64, l int) []float64 {
	var arrsi []float64
	for i := l; i <= len(data); i++ {
		pl := append([]float64(nil), data[i-l:i]...)
		var loss float64
		var gain float64
		for q := 1; q < len(pl); q++ {
			if pl[q]-pl[q-1] < 0 {
				loss += math.Abs(pl[q] - pl[q-1])
			} else {
				gain += pl[q] - pl[q-1]
			}
		}
		var rsi float64 = 100 - 100/(1+((gain/float64(l))/(loss/float64(l))))
		arrsi = append(arrsi, rsi)
	}
	return arrsi
}

func Wrsi(data []float64, l int) []float64 {
	var arrsi []float64
	var u []float64
	var d []float64
	for i := 1; i < len(data); i++ {
		if data[i]-data[i-1] < 0 {
			d = append(d, math.Abs(data[i]-data[i-1]))
			u = append(u, 0)
		} else {
			d = append(d, 0)
			u = append(u, data[i]-data[i-1])
		}
	}
	d = ma.Wsma(d, l)
	u = ma.Wsma(u, l)
	for i := 0; i < len(d); i++ {
		arrsi = append(arrsi, 100-100/(1+(u[i]/d[i])))
	}
	return arrsi
}

func Macd(data []float64, l1 int, l2 int) []float64 {
	if l1 > l2 {
		l1, l2 = l2, l1
	}
	var ema []float64 = ma.Ema(data, l1)
	var emb []float64 = ma.Ema(data, l2)
	var macd []float64
	ema = ema[l2-l1:]
	for i := 0; i < len(emb); i++ {
		macd = append(macd, ema[i]-emb[i])
	}
	return macd
}

func AroonUp(data []float64, l int) []float64 {
	var aroon []float64
	for i := l; i <= len(data); i++ {
		pl := append([]float64(nil), data[i-l:i]...)
		hl := append([]float64(nil), data[i-l:i]...)
		sort.Float64s(hl)
		pl = reverseFloats(pl)
		index := indexOf(pl, hl[len(hl)-1])
		aroon = append(aroon, (100 * (float64(index)) / (float64(l) - 1)))
	}
	return aroon
}

func AroonDown(data []float64, l int) []float64 {
	var aroon []float64
	for i := l; i <= len(data); i++ {
		pl := append([]float64(nil), data[i-l:i]...)
		hl := append([]float64(nil), data[i-l:i]...)
		sort.Float64s(hl)
		index := indexOf(pl, hl[0])
		aroon = append(aroon, (100 * (float64(index)) / (float64(l) - 1)))
	}
	return aroon
}

func Mfi(data [][]float64, l int) []float64 {
	var mfi []float64
	var n []float64
	var p []float64
	for i := 0; i < len(data); i++ {
		n = append(n, data[i][1])
		p = append(p, data[i][0])
	}
	for i := l; i <= len(data); i++ {
		var pos float64
		var neg float64
		for q := i - l; q < i; q++ {
			pos += p[q]
			neg += n[q]
		}
		mfi = append(mfi, 100-100/(1+pos/neg))
	}
	return mfi
}

func Roc(data []float64, l int) []float64 {
	var roc []float64
	for i := l; i <= len(data); i++ {
		roc = append(roc, (data[i-1]-data[i-l])/data[i-l])
	}
	return roc
}

func Cop(data []float64, l1 int, l2 int, l3 int) []float64 {
	max := math.Max(float64(l1), float64(l2))
	var co []float64
	for i := int(max) + l3; i < len(data); i++ {
		r1 := append([]float64(nil), data[i-(int(max)+l3):i]...)
		r2 := r1[:]
		var tmp []float64
		r1 = Roc(r1, l1)
		r2 = Roc(r2, l2)
		r1 = r1[len(r1)-len(r2):]
		r2 = r2[len(r2)-len(r1):]
		for a := 0; a < len(r1); a++ {
			tmp = append(tmp, r1[a]+r2[a])
		}
		tmp = ma.Wma(tmp, l3)
		co = append(co, tmp[len(tmp)-1])
	}
	return co
}

func Kst(data []float64, r1 int, r2 int, r3 int, r4 int, s1 int, s2 int, s3 int, s4 int, sig int) [][]float64 {
	var ks []float64
	var fs [][]float64
	ms := (intMax([]int{r1, r2, r3, r4}) + intMax([]int{s4, s3, s2, s1}))
	for i := ms; i <= len(data); i++ {
		pl := append([]float64(nil), data[i-ms:i]...)
		rcma1 := Roc(pl, r1)
		rcma2 := Roc(pl, r2)
		rcma3 := Roc(pl, r3)
		rcma4 := Roc(pl, r4)
		rcma1 = ma.Sma(rcma1, s1)
		rcma2 = ma.Sma(rcma2, s2)
		rcma3 = ma.Sma(rcma3, s3)
		rcma4 = ma.Sma(rcma4, s4)
		ks = append(ks, rcma1[len(rcma1)-1]+rcma2[len(rcma2)-1]+rcma3[len(rcma3)-1]+rcma4[len(rcma4)-1])
	}
	sl := ma.Sma(ks, sig)
	ks = ks[len(ks)-len(sl):]
	for i := 0; i < len(sl); i++ {
		fs = append(fs, []float64{ks[i], sl[i]})
	}
	return fs
}

func Obv(data [][]float64) []float64 {
	var obv []float64
	obv = append(obv, 0)
	for i := 1; i < len(data); i++ {
		if data[i][1] > data[i-1][1] {
			obv = append(obv, obv[len(obv)-1]+data[i][0])
		}
		if data[i][1] < data[i-1][1] {
			obv = append(obv, obv[len(obv)-1]-data[i][0])
		}
		if data[i][1] == data[i-1][1] {
			obv = append(obv, obv[len(obv)-1])
		}
	}
	return obv
}

func Vwap(data [][]float64, l int) []float64 {
	var vwap []float64
	var weighted [][]float64
	for i := 0; i < len(data); i++ {
		weighted = append(weighted, []float64{data[i][0] * data[i][1], data[i][1]})
	}
	for i := l; i <= len(weighted); i++ {
		pl := append([][]float64(nil), weighted[i-l:i]...)
		var totalv float64
		var totalp float64
		for o := 0; o < len(pl); o++ {
			totalv += pl[o][1]
			totalp += pl[o][0]
		}
		vwap = append(vwap, totalp/totalv)
	}
	return vwap
}

func Mom(data []float64, l int, p bool) []float64 {
	var mom []float64
	for i := l - 1; i < len(data); i++ {
		if p {
			mom = append(mom, data[i]/data[i-(l-1)]*100)
		} else {
			mom = append(mom, data[i]-data[i-(l-1)])
		}
	}
	return mom
}

func Tsi(data []float64, llen int, slen int, sig int) [][]float64 {
	var mom []float64
	var abs []float64
	var ts []float64
	var tsi [][]float64
	for i := 1; i < len(data); i++ {
		mom = append(mom, data[i]-data[i-1])
		abs = append(abs, math.Abs(data[i]-data[i-1]))
	}
	sma1 := ma.Ema(mom, llen)
	sma2 := ma.Ema(abs, llen)
	ema1 := ma.Ema(sma1, slen)
	ema2 := ma.Ema(sma2, slen)
	for i := 0; i < len(ema1); i++ {
		ts = append(ts, ema1[i]/ema2[i])
	}
	tma := ma.Ema(ts, sig)
	ts = ts[len(ts)-len(tma):]
	for i := 0; i < len(tma); i++ {
		tsi = append(tsi, []float64{tma[i], ts[i]})
	}
	return tsi
}

func Bop(data [][]float64, l int) []float64 {
	var bo []float64
	for i := 0; i < len(data); i++ {
		bo = append(bo, (data[i][3]-data[i][0])/(data[i][1]-data[i][2]))
	}
	bo = ma.Sma(bo, l)
	return bo
}

func Fi(data [][]float64, l int) []float64 {
	var ff []float64
	var pl []float64
	for i := 1; i < len(data); i++ {
		pl = append(pl, (data[i][0]-data[i-1][0])*data[i][1])
		if len(pl) >= l {
			vfi := ma.Ema(pl, l)
			ff = append(ff, vfi[len(vfi)-1])
			pl = pl[1:]
		}
	}
	return ff
}

func Asi(data [][]float64) []float64 {
	var a []float64
	for i := 1; i < len(data); i++ {
		c := data[i][1]
		cy := data[i-1][1]
		h := data[i][0]
		hy := data[i-1][0]
		l := data[i][2]
		ly := data[i-1][2]
		var k float64
		if hy-c > ly-c {
			k = hy - c
		} else {
			k = ly - c
		}
		o := data[i][0]
		oy := data[i-1][0]
		var r float64
		t := maxf([]float64{data[i][0], data[i-1][0]}) - minf([]float64{data[i][2], data[i-1][2]})
		if h-cy > l-cy && h-cy > h-l {
			r = h - cy - (l-cy)/2 + (cy-oy)/4
		}
		if l-cy > h-cy && l-cy > h-l {
			r = l - cy - (h-cy)/2 + (cy-oy)/4
		}
		if h-l > h-cy && h-l > l-cy {
			r = h - l + (cy-oy)/4
		}
		a = append(a, 50*((cy-c+(cy-oy)/2+(c-o)/2)/r)*k/t)
	}
	return a
}

func Alligator(data []float64, jl int, tl int, ll int, js int, ts int, ls int) [][]float64 {
	var ret [][]float64
	var jaw []float64 = ma.Smma(data, jl)
	var teeth []float64 = ma.Smma(data, tl)
	var lips []float64 = ma.Smma(data, ll)
	teeth = teeth[len(teeth)-len(jaw):]
	lips = lips[len(lips)-len(jaw):]
	for i := len(jaw) - 1; i >= 7; i-- {
		ret = append(ret, []float64{jaw[i-(js-1)], teeth[i-(ts-1)], lips[i-(ls-1)]})
	}
	return ret
}

func Pr(data []float64, l int) []float64 {
	var n []float64
	for i := l; i <= len(data); i++ {
		pl := append([]float64(nil), data[i-l:i]...)
		highd := maxf(pl)
		lowd := minf(pl)
		n = append(n, (highd-data[i-1])/(highd-lowd)*-100)
	}
	return n
}

func Stoch(data [][]float64, l int, sd int, sk int) [][]float64 {
	var stoch [][]float64
	var high []float64
	var low []float64
	var ka []float64
	if l < sd {
		l, sd = sd, l
	}
	if sk > sd {
		sk, sd = sd, sk
	}
	for i := 0; i < len(data); i++ {
		high = append(high, data[i][0])
		low = append(low, data[i][2])
		if len(high) >= l {
			highd := maxf(high)
			lowd := minf(low)
			k := 100 * (data[i][1] - lowd) / (highd - lowd)
			ka = append(ka, k)
		}
		if sk > 0 && len(ka) > sk {
			smoothedk := ma.Sma(ka, sk)
			ka = append(ka, smoothedk[len(smoothedk)-1])
		}
		if len(ka)-sk >= sd {
			d := ma.Sma(ka, sd)
			stoch = append(stoch, []float64{ka[len(ka)-1], d[len(d)-1]})
			high = high[1:]
			low = low[1:]
			ka = ka[1:]
		}
	}
	return stoch
}

func Ichimoku(data [][]float64, l1 int, l2 int, l3 int, l4 int) [][]float64 {
	var pl [][]float64
	var cloud [][]float64
	var place [][]float64
	for i := 0; i < len(data); i++ {
		pl = append(pl, data[i])
		if len(pl) >= l3 {
			var highs []float64
			var lows []float64
			for a := 0; a < len(pl); a++ {
				highs = append(highs, pl[a][0])
				lows = append(lows, pl[a][2])
			}
			tsen := (maxf(highs[len(highs)-l1:]) + minf(lows[len(lows)-l1:])) / 2
			ksen := (maxf(highs[len(highs)-l2:]) + minf(lows[len(lows)-l2:])) / 2
			senka := data[i][1] + ksen
			senkb := (maxf(highs[len(highs)-l3:]) + minf(lows[len(lows)-l2:])) / 2
			chik := data[i][1]
			place = append(place, []float64{tsen, ksen, senka, senkb, chik})
			pl = pl[1:]
		}
	}
	for i := l4; i < len(place)-l4; i++ {
		cloud = append(cloud, []float64{place[i][0], place[i][1], place[i+l4][2], place[i+(l4)][3], place[i-l4][4]})
	}
	return cloud
}

func Atr(data [][]float64, l int) []float64 {
	var atr = []float64{data[0][0] - data[0][2]}
	for i := 1; i < len(data); i++ {
		t0 := maxf([]float64{data[i][0] - data[i-1][1], data[i][2] - data[i-1][1], data[i][0] - data[i][2]})
		atr = append(atr, (atr[len(atr)-1]*(float64(l)-1)+t0)/float64(l))
	}
	return atr
}

func Fib(start float64, end float64) []float64 {
	return []float64{start, (end-start)*.236 + start, (end-start)*.382 + start, (end-start)*.5 + start, (end-start)*.618 + start, (end-start)*.786 + start, end, (end-start)*1.618 + start, (end-start)*2.618 + start, (end-start)*3.618 + start, (end-start)*4.236 + start}
}

func Fractals(data [][]float64) [][]bool {
	var fractals = [][]bool{{false, false}, {false, false}}
	for i := 2; i < len(data)-2; i++ {
		var up bool
		var down bool
		if data[i-2][0] < data[i][0] && data[i-1][0] < data[i][0] && data[i][0] > data[i+1][0] && data[i][0] > data[i+2][0] {
			up = true
		} else {
			up = false
		}
		if data[i-2][1] > data[i][1] && data[i-1][1] > data[i][1] && data[i][1] < data[i+1][1] && data[i][1] < data[i+2][1] {
			down = true
		} else {
			down = false
		}
		fractals = append(fractals, []bool{up, down})
	}
	fractals = append(fractals, []bool{false, false})
	fractals = append(fractals, []bool{false, false})
	return fractals
}

func HalfTrend(data [][]float64, atrlen int, amplitude int, deviation float64) [][]interface{} {
	out := make([][]interface{}, 0)
	nexttrend := []int{0}
	trend := []int{0}
	var up []float64
	var down []float64
	var direction string
	for i := atrlen; i <= len(data); i++ {
		pl := data[i-atrlen : i]
		maxlow := pl[len(pl)-2][2]
		minhigh := pl[len(pl)-2][0]
		atr2 := Atr(pl, atrlen)
		atr2 = []float64{atr2[len(atr2)-1] / 2}
		dev := deviation * atr2[len(atr2)-1]
		highprice := maxf([]float64{pl[len(pl)-1][0], pl[len(pl)-2][0]})
		lowprice := minf([]float64{pl[len(pl)-1][2], pl[len(pl)-2][2]})
		highs := map2d(pl[len(pl)-amplitude:len(pl)], 0)
		lows := map2d(pl[len(pl)-amplitude:len(pl)], 2)
		highma := ma.Sma(highs, amplitude)
		lowma := ma.Sma(lows, amplitude)
		var atrHigh float64
		var atrLow float64
		if nexttrend[len(nexttrend)-1] == 1 {
			maxlow = maxf([]float64{lowprice, maxlow})
			if highma[0] < maxlow && pl[len(pl)-1][1] < pl[len(pl)-2][2] {
				trend = append(trend, 1)
				nexttrend = append(nexttrend, 0)
				minhigh = pl[len(pl)-2][0]
			}
		} else {
			minhigh = minf([]float64{highprice, minhigh})
			if lowma[0] > minhigh && pl[len(pl)-1][1] < pl[len(pl)-2][0] {
				trend = append(trend, 0)
				nexttrend = append(nexttrend, 1)
				maxlow = lowprice
			}
		}
		if trend[len(trend)-1] == 0 {
			if len(trend) > 1 && trend[len(trend)-2] != 0 {
				if len(down) < 2 {
					up = append(up, down[len(down)-1])
				} else {
					up = append(up, down[len(down)-2])
				}
			} else {
				if len(up) < 2 {
					up = append(up, maxlow)
				} else {
					up = append(up, maxf([]float64{up[len(up)-2], maxlow}))
				}
			}
			direction = "long"
			atrHigh = up[len(up)-1] + dev
			atrLow = up[len(up)-1] - dev
		} else {
			if len(trend) > 2 && trend[len(trend)-2] != 1 {
				if len(up) < 2 {
					down = append(down, up[len(up)-1])
				} else {
					down = append(down, up[len(up)-2])
				}
			} else {
				if len(down) < 2 {
					down = append(down, minhigh)
				} else {
					down = append(down, minf([]float64{minhigh, down[len(down)-2]}))
				}
			}
			direction = "long"
			atrHigh = down[len(down)-1] + dev
			atrLow = down[len(down)-1] - dev
		}
		if trend[len(trend)-1] == 0 {
			out = append(out, []interface{}{atrHigh, up[len(up)-1], atrLow, direction})
		} else {
			out = append(out, []interface{}{atrHigh, down[len(down)-1], atrLow})
		}
	}
	return out
}

func ZigZag(data [][]float64, perc float64) []float64 {
	var indexes []stats.RecentHighLow
	min := math.Inf(1)
	max := math.Inf(-1)
	lmin := false
	lmax := false
	for i := 0; i < len(data); i++ {
		if lmin {
			if indexes[len(indexes)-1].Value >= data[i][1] {
				indexes[len(indexes)-1].Value = data[i][1]
				indexes[len(indexes)-1].Index = i
			}
			if min >= data[i][1] {
				min = data[i][1]
			}
			hdif := (data[i][0] - min) / min
			if hdif > perc {
				indexes = append(indexes, stats.RecentHighLow{Index: i, Value: data[i][0]})
				lmax = true
				lmin = false
				min = math.Inf(1)
			}
		} else if lmax {
			if indexes[len(indexes)-1].Value <= data[i][0] {
				indexes[len(indexes)-1].Value = data[i][0]
				indexes[len(indexes)-1].Index = i
			}
			if max <= data[i][1] {
				max = data[i][1]
			}
			ldif := (max - data[i][1]) / data[i][1]
			if ldif > perc {
				indexes = append(indexes, stats.RecentHighLow{Index: i, Value: data[i][1]})
				lmin = true
				lmax = false
				max = math.Inf(-1)
			}
		} else {
			if min >= data[i][1] {
				min = data[i][1]
			}
			if max <= data[i][0] {
				max = data[i][0]
			}
			hdif := (data[i][0] - min) / min
			ldif := (max - data[i][1]) / max
			if ldif > perc && hdif < perc {
				lmin = true
				indexes = append(indexes, stats.RecentHighLow{Index: 0, Value: data[0][0]})
				indexes = append(indexes, stats.RecentHighLow{Index: i, Value: data[i][1]})
			} else if hdif > perc && ldif < perc {
				lmax = true
				indexes = append(indexes, stats.RecentHighLow{Index: 0, Value: data[0][1]})
				indexes = append(indexes, stats.RecentHighLow{Index: i, Value: data[i][0]})
			} else {
				if ldif > hdif {
					lmin = true
					indexes = append(indexes, stats.RecentHighLow{Index: 0, Value: data[0][0]})
					indexes = append(indexes, stats.RecentHighLow{Index: i, Value: data[i][0]})
				} else {
					lmax = true
					indexes = append(indexes, stats.RecentHighLow{Index: 0, Value: data[0][1]})
					indexes = append(indexes, stats.RecentHighLow{Index: i, Value: data[i][0]})
				}
			}
		}
	}
	final := []float64{indexes[0].Value}
	for i := 1; i < len(indexes); i++ {
		length := indexes[i].Index - indexes[i-1].Index
		delta := (indexes[i].Value - indexes[i-1].Value) / float64(length)
		for z := 1; z <= length; z++ {
			final = append(final, float64(z)*delta+indexes[i-1].Value)
		}
	}
	return final
}

func Psar(data [][]float64, step float64, max float64) []float64 {
	furthest := data[0]
	up := true
	accel := step
	prev := data[0]
	sar := data[0][1]
	extreme := data[0][0]
	final := []float64{sar}
	for i := 1; i < len(data); i++ {
		sar = sar + accel*(extreme-sar)
		if up {
			sar = minf([]float64{sar, furthest[1], prev[1]})
			if data[i][0] > extreme {
				extreme = data[i][0]
				accel = minf([]float64{accel + step, max})
			}
		} else {
			sar = maxf([]float64{sar, furthest[0], prev[0]})
			if data[i][1] < extreme {
				extreme = data[i][0]
				accel = minf([]float64{accel + step, max})
			}
		}
		if (up && data[i][1] < sar) || (!up && data[i][0] > sar) {
			accel = step
			sar = extreme
			up = !up
			if !up {
				extreme = data[i][1]
			} else {
				extreme = data[i][0]
			}
		}
		furthest = prev
		prev = data[i]
		final = append(final, sar)
	}
	return final
}

func Supertrend(data [][]float64, l int, multiplier float64) [][]float64 {
	atr := Atr(data, l)
	var trend [][]float64
	for i := l - 1; i < len(data); i++ {
		trend = append(trend, []float64{(data[i][0]+data[i][2])/2 + multiplier*atr[i], (data[i][0]+data[i][2])/2 - multiplier*atr[i]})
	}
	return trend
}

func Elderray(data []float64, l int) [][]float64 {
	var eld [][]float64
	for i := l; i <= len(data); i++ {
		pl := append([]float64{}, data[i-l:i]...)
		sort.Float64s(pl)
		em := ma.Ema(data[i-l:i], l)
		eld = append(eld, []float64{pl[len(pl)-1] - em[0], pl[0] - em[0]})
	}
	return eld
}

func Hv(data []float64, l int) []float64 {
	var hv []float64
	for i := l; i <= len(data); i++ {
		ss := stats.Ssd(data[i-l:i], l)
		vari := ss[0] / float64(l)
		hv = append(hv, math.Sqrt(vari))
	}
	return hv
}

func Rvi(data [][]float64, ln int) []float64 {
	var num []float64
	var dnom []float64
	var rv []float64
	for i := 3; i < len(data); i++ {
		num = append(num, (data[i][3]-data[i][0]+2*(data[i][3]-data[i-1][0])+2*(data[i][3]-data[i-2][0])+(data[i][3]-data[i-3][0]))/6)
		dnom = append(dnom, (data[i][1]-data[i][2]+2*(data[i][1]-data[i-1][2])+2*(data[i][1]-data[i-2][2])+(data[i][1]-data[i-3][2]))/6)
		if len(num) >= ln {
			sn := ma.Sma(num, ln)
			dn := ma.Sma(dnom, ln)
			rv = append(rv, sn[0]/dn[0])
			num = num[1:]
			dnom = dnom[1:]
		}
	}
	return rv
}

func Rvi_signal(rv []float64) []float64 {
	var sig []float64
	for i := 3; i < len(rv); i++ {
		sig = append(sig, (rv[i]+2*rv[i-1]+2*rv[i-2]+rv[i-3])/6)
	}
	return sig
}

func Rsi_divergence(data []float64, ln int, rs func([]float64, int) []float64) []float64 {
	rd := rs(data, ln)
	var out []float64
	data = Mom(data[ln-1:len(data)], 2, false)
	for i := 0; i < len(data); i++ {
		if (data[i] > 0 && rd[i] < 0) || (data[i] < 0 && rd[i] > 0) {
			out = append(out, 1)
		} else {
			out = append(out, 0)
		}
	}
	return out
}

func Divergence(data1 []float64, data2 []float64) []float64 {
	if len(data1) > len(data2) {
		data1 = data1[len(data1)-len(data2):]
	}
	if len(data2) > len(data1) {
		data2 = data2[len(data2)-len(data1):]
	}
	var out []float64
	for i := 1; i < len(data1); i++ {
		if (data1[i] > data1[i-1] && data2[i] < data2[i-1]) || (data1[i] < data1[i-1] && data2[i] > data2[i-1]) {
			out = append(out, 1)
		} else {
			out = append(out, 0)
		}
	}
	return out
}
