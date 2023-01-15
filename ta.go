package ta

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
)

func Median(data []float64, l int) []float64 {
	var med []float64
	for i := l; i <= len(data); i++ {
		pl := append([]float64(nil), data[i-l:i]...)
		sort.Float64s(pl)
		med = append(med, pl[len(pl)/2-1])
	}
	return med
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
	var mean []float64 = Sma(data, len(data))
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
		m := Sma(tmp, len(tmp))
		var sum float64 = 0
		for x := 0; x < len(tmp); x++ {
			sum += math.Abs(float64(tmp[x]) - float64(m[0]))
		}
		med = append(med, sum/float64(l))
	}
	return med
}
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
func Ssd(data []float64, l int) []float64 {
	var sd []float64
	for i := l; i <= len(data); i++ {
		tmp := append([]float64(nil), data[i-l:i]...)
		mean := Sma(tmp, l)
		var sum float64
		for x := 0; x < len(tmp); x++ {
			sum += math.Pow(tmp[x]-mean[0], 2)
		}
		sd = append(sd, math.Pow(sum, 0.5))
	}
	return sd
}
func Er(data []float64) float64 {
	var wins []float64; var losses []float64; wp := 1.0; lp := 1.0;
	for i := 0; i < len(data); i++ {
		if(data[i] >= 0) {
			wins = append(wins, 1+data[i]);
		} else {
			losses = append(losses, 1+data[i]);
		}
	}
	var datlen float64 = float64(len(data));
	var loslen float64 = float64(len(losses));
	var winlen float64 = float64(len(wins));
	win := winlen / datlen;
	loss := loslen / datlen;
	for i := 0; i < int(loslen); i++ { lp = lp * losses[i];	}
	for i := 0; i < int(winlen); i++ { wp = wp * wins[i]; }
	var fw float64 = (1.0 / winlen); var fl float64 = (1.0 / loslen);
	return ((((math.Pow(wp, fw))-1)*100) * win + (((math.Pow(lp, fl))-1)*100) * loss) / 100;
}
func Winratio(data []float64) float64 {
	var wins int = 0; var losses int = 0;
	for i := 0; i < len(data); i++ {
		if data[i] >= 0 { wins++ } else { losses++ }
	}
	return float64(wins) / float64(losses + wins);
}
func WinAverage(data []float64) float64 {
	var wins []float64;
	for i := 0; i < len(data); i++ {
		if data[i] >= 0 { wins = append(wins, data[i]) }
	}
	avg := Sma(wins, len(wins));
	return avg[0];
}
func LossAverage(data []float64) float64 {
	var losses []float64;
	for i := 0; i < len(data); i++ {
		if data[i] < 0 { losses = append(losses, data[i]) }
	}
	avg := Sma(losses, len(losses));
	return avg[0];
}
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
	d = Wsma(d, l)
	u = Wsma(u, l)
	for i := 0; i < len(d); i++ {
		arrsi = append(arrsi, 100-100/(1+(u[i]/d[i])))
	}
	return arrsi
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
func Smma(data []float64, l int) []float64 {
	var smma []float64
	for i := l; i <= len(data); i++ {
		pl := append([]float64(nil), data[i-l:i]...)
		var average float64
		for q := 0; q < len(pl); q++ {
			average += pl[q]
		}
		if len(smma) <= 0 {
			smma = append(smma, average/float64(l))
		} else {
			smma = append(smma, (average-smma[len(smma)-1])/float64(l))
		}
	}
	return smma[1:]
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
func Macd(data []float64, l1 int, l2 int) []float64 {
	if l1 > l2 {
		l1, l2 = l2, l1
	}
	var ema []float64 = Ema(data, l1)
	var emb []float64 = Ema(data, l2)
	var macd []float64
	ema = ema[l2-l1:]
	for i := 0; i < len(emb); i++ {
		macd = append(macd, ema[i]-emb[i])
	}
	return macd
}
func Std(data []float64, l int) float64 {
	var mean []float64 = Sma(data, l)
	var average float64
	for i := len(data) - l; i < len(data); i++ {
		average += math.Pow(data[i]-mean[len(mean)-1], 2)
	}
	var std float64 = math.Pow(average/float64(l), 0.5)
	return std
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
func Cor(data1 []float64, data2 []float64) float64 {
	var d1avg []float64 = Sma(data1, len(data1))
	var d2avg []float64 = Sma(data2, len(data2))
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
func Dif(n float64, o float64) float64 {
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
func reverseFloats(data []float64) []float64 {
	var ret []float64
	for i := len(data) - 1; i >= 0; i-- {
		ret = append(ret, data[i])
	}
	return ret
}
func indexOf(data []float64, index float64) int {
	var ret int = -1;
	for i := len(data) - 1; i >= 0; i-- {
		if data[i] == index {
			ret = i
			break
		}
	}
	return ret
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
func AroonOsc(data []float64, l int) []float64 {
	var aroon []float64
	u := AroonUp(data, l)
	d := AroonDown(data, l)
	for i := 0; i < len(u); i++ {
		aroon = append(aroon, u[i]-d[i])
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
		tmp = Wma(tmp, l3)
		co = append(co, tmp[len(tmp)-1])
	}
	return co
}
func max(v []int) int {
	sort.Ints(v)
	return v[len(v)-1]
}
func min(v []int) int {
	sort.Ints(v)
	return v[0]
}
func maxf(v []float64) float64 {
	sort.Float64s(v)
	return v[len(v)-1]
}
func minf(v []float64) float64 {
	sort.Float64s(v)
	return v[0]
}
func Kst(data []float64, r1 int, r2 int, r3 int, r4 int, s1 int, s2 int, s3 int, s4 int, sig int) [][]float64 {
	var ks []float64
	var fs [][]float64
	ms := (max([]int{r1, r2, r3, r4}) + max([]int{s4, s3, s2, s1}))
	for i := ms; i <= len(data); i++ {
		pl := append([]float64(nil), data[i-ms:i]...)
		rcma1 := Roc(pl, r1)
		rcma2 := Roc(pl, r2)
		rcma3 := Roc(pl, r3)
		rcma4 := Roc(pl, r4)
		rcma1 = Sma(rcma1, s1)
		rcma2 = Sma(rcma2, s2)
		rcma3 = Sma(rcma3, s3)
		rcma4 = Sma(rcma4, s4)
		ks = append(ks, rcma1[len(rcma1)-1]+rcma2[len(rcma2)-1]+rcma3[len(rcma3)-1]+rcma4[len(rcma4)-1])
	}
	sl := Sma(ks, sig)
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
func MomOsc(data []float64, l int) []float64 {
	var osc []float64
	for i := l; i <= len(data); i++ {
		var sumh float64
		var suml float64
		for a := 1; a < l; a++ {
			if data[i-l+(a-1)] < data[i-l+a] {
				sumh += data[i-l+a]
			} else {
				suml += data[i-l+a]
			}
		}
		osc = append(osc, (sumh-suml)/(sumh+suml)*100)
	}
	return osc
}
func Ha(data [][]float64) [][]float64 {
	var ha [][]float64 = [][]float64{{(data[0][0] + data[0][3]) / 2, data[0][1], data[0][2], (data[0][0] + data[0][1] + data[0][2] + data[0][3]) / 4}}
	for i := 1; i < len(data); i++ {
		ha = append(ha, []float64{(ha[len(ha)-1][0] + ha[len(ha)-1][3]) / 2, maxf([]float64{ha[len(ha)-1][0], ha[len(ha)-1][3], data[i][1]}), minf([]float64{ha[len(ha)-1][0], ha[len(ha)-1][3], data[i][2]}), (data[i][0] + data[i][1] + data[i][2] + data[i][3]) / 4})
	}
	return ha
}
func decimalplaces(n float64) int {
	decimalPlaces := fmt.Sprintf("%f", n-math.Floor(n))
	decimalPlaces = strings.Replace(decimalPlaces, "0.", "", -1)
	decimalPlaces = strings.TrimRight(decimalPlaces, "0")
	return len(decimalPlaces)
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
func Tsi(data []float64, llen int, slen int, sig int) [][]float64 {
	var mom []float64
	var abs []float64
	var ts []float64
	var tsi [][]float64
	for i := 1; i < len(data); i++ {
		mom = append(mom, data[i]-data[i-1])
		abs = append(abs, math.Abs(data[i]-data[i-1]))
	}
	sma1 := Ema(mom, llen)
	sma2 := Ema(abs, llen)
	ema1 := Ema(sma1, slen)
	ema2 := Ema(sma2, slen)
	for i := 0; i < len(ema1); i++ {
		ts = append(ts, ema1[i]/ema2[i])
	}
	tma := Ema(ts, sig)
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
	bo = Sma(bo, l)
	return bo
}
func Fi(data [][]float64, l int) []float64 {
	var ff []float64
	var pl []float64
	for i := 1; i < len(data); i++ {
		pl = append(pl, (data[i][0]-data[i-1][0])*data[i][1])
		if len(pl) >= l {
			vfi := Ema(pl, l)
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
func Ao(data [][]float64, l1 int, l2 int) []float64 {
	var processed []float64
	var a []float64
	for i := 0; i < len(data); i++ {
		processed = append(processed, (data[i][0]+data[i][1])/2)
	}
	for i := l2; i <= len(processed); i++ {
		pl := processed[i-l2 : i]
		f := Sma(pl, l1)
		s := Sma(pl, l2)
		a = append(a, f[len(f)-1]-s[len(s)-1])
	}
	return a
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
func Don(data [][]float64, l int) [][]float64 {
	var channel [][]float64
	for i := l; i <= len(data); i++ {
		pl := data[i-l : i]
		var highs []float64
		var lows []float64
		for h := 0; h < len(pl); h++ { // todo: optimize
			highs = append(highs, pl[h][0])
			lows = append(lows, pl[h][1])
		}
		max := maxf(highs)
		min := minf(lows)
		channel = append(channel, []float64{max, (max + min) / 2, min})
	}
	return channel
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
			smoothedk := Sma(ka, sk)
			ka = append(ka, smoothedk[len(smoothedk)-1])
		}
		if len(ka)-sk >= sd {
			d := Sma(ka, sd)
			stoch = append(stoch, []float64{ka[len(ka)-1], d[len(d)-1]})
			high = high[1:]
			low = low[1:]
			ka = ka[1:]
		}
	}
	return stoch
}
func Atr(data [][]float64, l int) []float64 {
	var atr = []float64{data[0][0] - data[0][2]}
	for i := 1; i < len(data); i++ {
		t0 := maxf([]float64{data[i][0] - data[i-1][1], data[i][2] - data[i-1][1], data[i][0] - data[i][2]})
		atr = append(atr, (atr[len(atr)-1]*(float64(l)-1)+t0)/float64(l))
	}
	return atr
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
func Bands(data []float64, l int, dev float64) [][]float64 {
	var pl []float64
	var deviation []float64
	var boll [][]float64
	sma := Sma(data, l)
	for i := 0; i < len(data); i++ {
		pl = append(pl, data[i])
		if len(pl) >= l {
			devi := Std(pl, l)
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
func Keltner(data [][]float64, l int, dev float64) [][]float64 {
	var closing []float64
	atr := Atr(data, l)
	var kelt [][]float64
	for i := 0; i < len(data); i++ {
		closing = append(closing, (data[i][0]+data[i][1]+data[i][2])/3)
	}
	kma := Sma(closing, l)
	atr = atr[l-1:]
	for i := 0; i < len(kma); i++ {
		kelt = append(kelt, []float64{kma[i] + atr[i]*dev, kma[i], kma[i] - atr[i]*dev})
	}
	return kelt
}
func Variance(data []float64, l int) []float64 {
	var va []float64
	for i := l; i <= len(data); i++ {
		tmp := data[i-l : i]
		mean := Sma(tmp, l)
		var sum float64
		for x := 0; x < len(tmp); x++ {
			sum += math.Pow(tmp[x]-mean[len(mean)-1], 2)
		}
		va = append(va, sum/float64(l))
	}
	return va
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
			mean := Sma(change, len(change))
			std := Std(change, len(change))
			random := Normsinv(rand.Float64())
			projected = append(projected, projected[len(projected)-1]*math.Exp((mean[0]-(std*std)/2)+std*random))
		}
		sd = append(sd, projected)
	}
	return sd
}
func Percentile(data [][]float64, perc float64) []float64 {
	var ret []float64
	for i := 0; i < len(data[0]); i++ {
		var tmp []float64
		for x := 0; x < len(data); x++ {
			tmp = append(tmp, data[x][i])
		}
		sort.Float64s(tmp)
		ret = append(ret, tmp[int(float64(len(tmp)-1)*perc)])
	}
	return ret
}
func Kmeans(data []float64, clusters int) [][]float64 {
	var means [][]float64
	var n int
	var centers []float64
	var old []float64
	changed := false
	init := int(math.Round(float64(len(data)) / float64(clusters+1)))
	for i := 0; i < clusters; i++ {
		centers = append(centers, data[init*(1+i)])
	}
	for ok := true; ok == true; ok = changed {
		for i := 0; i < clusters; i++ {
			means = append(means, []float64{})
		}
		for x := 0; x < len(data); x++ {
			var oldrange float64 = -1
			for y := 0; y < clusters; y++ {
				r := math.Abs(centers[y] - data[x])
				if oldrange == -1 {
					oldrange = r
					n = y
				} else if r <= oldrange {
					oldrange = r
					n = y
				}
			}
			means[n] = append(means[n], data[x])
		}
		old = centers
		for x := 0; x < clusters; x++ {
			var sum float64
			for y := 0; y < len(means[x]); y++ {
				sum += means[x][y]
			}
			m := sum / float64(len(means[n]))
			centers[x] = m
		}
		for x := 0; x < clusters; x++ {
			if centers[x] != old[x] {
				break
			}
		}
	}
	return means
}
func Alligator(data []float64, jl int, tl int, ll int, js int, ts int, ls int) [][]float64 {
	var ret [][]float64
	var jaw []float64 = Smma(data, jl)
	var teeth []float64 = Smma(data, tl)
	var lips []float64 = Smma(data, ll)
	teeth = teeth[len(teeth)-len(jaw):]
	lips = lips[len(lips)-len(jaw):]
	for i := len(jaw) - 1; i >= 7; i-- {
		ret = append(ret, []float64{jaw[i-(js-1)], teeth[i-(ts-1)], lips[i-(ls-1)]})
	}
	return ret
}
func Gator(data []float64, jl int, tl int, ll int, js int, ts int, ls int) [][]float64 {
	var ret [][]float64
	var jaw []float64 = Smma(data, jl)
	var teeth []float64 = Smma(data, tl)
	var lips []float64 = Smma(data, ll)
	teeth = teeth[len(teeth)-len(jaw):]
	lips = lips[len(lips)-len(jaw):]
	for i := len(jaw) - 1; i >= (js - 1); i-- {
		ret = append(ret, []float64{jaw[i-(js-1)] - teeth[i-(ts-1)], -(math.Abs(teeth[i-(ts-1)] - lips[i-(ls-1)]))})
	}
	return ret
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
	ef := Ema(adl, ema1)
	es := Ema(adl, ema2)
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
func Envelope(data []float64, l int, p float64) [][]float64 {
	var enve [][]float64
	for i := l; i < len(data); i++ {
		sm := Sma(data[i-l:i], l)
		enve = append(enve, []float64{sm[0] * (1 + p), sm[0], sm[0] * (1 - p)})
	}
	return enve
}

type recentHighLow struct {
	Index int
	Value float64
}

func RecentHighLow(index int, value float64) recentHighLow {
	return recentHighLow{Index: index, Value: value}
}
func RecentHigh(data []float64, lb int) recentHighLow {
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
	return recentHighLow{Index: hindex, Value: highest}
}
func RecentLow(data []float64, lb int) recentHighLow {
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
	return recentHighLow{Index: lindex, Value: lowest}
}
func Fib(start float64, end float64) []float64 {
	return []float64{start, (end-start)*.236 + start, (end-start)*.382 + start, (end-start)*.5 + start, (end-start)*.618 + start, (end-start)*.786 + start, end, (end-start)*1.618 + start, (end-start)*2.618 + start, (end-start)*3.618 + start, (end-start)*4.236 + start}
}
func Ac(data [][]float64, len1 int, len2 int) []float64 {
	a := Ao(data, len1, len2)
	sm := Sma(a, len1)
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

type line struct {
	Slope float64
	Value float64
	Index int
}

func Line(index int, slope float64, value float64) line {
	return line{Slope: slope, Value: value, Index: index}
}
func LineCalc(pos int, calc line) float64 {
	return float64(pos)*calc.Slope + calc.Value
}
func Support(d []float64, hl recentHighLow) line {
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
		return line{Slope: 0, Value: hl.Value, Index: hl.Index}
	} else {
		return line{Slope: lowform, Value: hl.Value, Index: hl.Index}
	}
}
func Resistance(d []float64, hl recentHighLow) line {
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
		return line{Slope: 0, Value: hl.Value, Index: hl.Index}
	} else {
		return line{Slope: -highform, Value: hl.Value, Index: hl.Index}
	}
}
func Fisher(data []float64, l int) [][]float64 {
	var out [][]float64; var fish, v1 float64 = 0, 0;
	for i := l; i <= len(data); i++ {
		pl := data[i-l:i]; pf := fish;
		mn := minf(pl);
		v1 = .33*2*((data[i-1]-mn)/maxf(pl)-.5)+.67*v1;
		if v1 > 0.99 { v1 = 0.999 }
		if v1 < -0.99 { v1 = -0.999 }
		fish = 0.5 * math.Log((1+v1)/(1-v1)) + 0.5 * pf;
		out = append(out, []float64{fish, pf});
	}
	return out[1:];
}
func Ar(data []float64, l int) []float64 {
	var out []float64;
	for i := l; i < len(data); i++ {
		exp := Er(data[i-l:i]);
		out = append(out, data[i]-exp);
	}
	return out;
}
func Avgwin(data []float64) float64 {
	var wins []float64;
	for i := 0; i < len(data); i++ {
		if data[i] >= 0 { wins = append(wins, data[i]) }
	}
	avg := Sma(wins, len(wins));
	return avg[0];
}
func Avgloss(data []float64) float64 {
	var loss []float64;
	for i := 0; i < len(data); i++ {
		if data[i] < 0 { loss = append(loss, data[i]) }
	}
	avg := Sma(loss, len(loss));
	return avg[0];
}
func Se(data []float64, l int) float64 {
	stdv := Std(data, len(data));
	return stdv / math.Pow(float64(l), 0.5);
}
func Kelly(data []float64) float64 {
	exp := Er(data) + 1;
	winr := Winratio(data);
	return winr - (1-winr) / exp;
}
func Zscore(data []float64, l int) []float64 {
	var out []float64;
	for i := l; i <= len(data); i++ {
		pl := data[i-l:i];
		mean := Sma(pl, l);
		stdv := Std(pl, l);
		out = append(out, (data[i-1]-mean[0])/stdv);
	}
	return out;
}
func NormalizePair(data1 []float64, data2 []float64) [][]float64 {
	f := (data1[0] + data2[0]) / 2;
	var ret [][]float64;
	ret = append(ret, []float64{f,f});
	for i := 1; i < len(data1); i++ {
		ret = append(ret, []float64{ret[len(ret)-1][0]*((data1[i]-data1[i-1])/data1[i-1]+1),ret[len(ret)-1][1]*((data2[i]-data2[i-1])/data2[i-1]+1)});
	}
	return ret;
}
func NormalizeFrom(data []float64, value float64) []float64 {
	var ret = []float64{value};
	for i := 1; i < len(data); i++ {
		ret = append(ret, ret[len(ret)-1]*((data[i]-data[i-1])/data[i-1]+1));
	}
	return ret;
}
func Log(data []float64) []float64 {
	var ret []float64;
	for i := 0; i < len(data); i++ {
		ret = append(ret, math.Log(data[i]));
	}
	return ret;
}
func Exp(data []float64) []float64 {
	var ret []float64;
	for i := 0; i < len(data); i++ {
		ret = append(ret, math.Exp(data[i]));
	}
	return ret;
}
func map2d(data [][]float64, key int) []float64 {
	var ret []float64;
	for i := 0; i < len(data); i++ {
		ret = append(ret, data[i][key]);
	}
	return ret;
}
func HalfTrend(data [][]float64, atrlen int, amplitude int, deviation float64) [][]interface{} {
	out := make([][]interface{}, 0); nexttrend := []int{0}; trend := []int{0};
	var up []float64; var down []float64; var direction string;
	for i := atrlen; i <= len(data); i++ {
		pl := data[i-atrlen:i];
		maxlow := pl[len(pl)-2][2];
		minhigh := pl[len(pl)-2][0];
		atr2 := Atr(pl, atrlen);
		atr2 = []float64{atr2[len(atr2)-1] / 2};
		dev := deviation * atr2[len(atr2)-1];
		highprice := maxf([]float64{pl[len(pl)-1][0],pl[len(pl)-2][0]});
		lowprice := minf([]float64{pl[len(pl)-1][2],pl[len(pl)-2][2]});
		highs := map2d(pl[len(pl)-amplitude:len(pl)], 0);
		lows := map2d(pl[len(pl)-amplitude:len(pl)], 2);
		highma := Sma(highs, amplitude);
		lowma := Sma(lows, amplitude);
		var atrHigh float64;
		var atrLow float64;
		if nexttrend[len(nexttrend)-1] == 1 {
			maxlow = maxf([]float64{lowprice, maxlow});
			if highma[0] < maxlow && pl[len(pl)-1][1] < pl[len(pl)-2][2] {
				trend = append(trend, 1);
				nexttrend = append(nexttrend, 0);
				minhigh = pl[len(pl)-2][0]
			}
		} else {
			minhigh = minf([]float64{highprice, minhigh});
			if lowma[0] > minhigh && pl[len(pl)-1][1] < pl[len(pl)-2][0] {
				trend = append(trend, 0);
				nexttrend = append(nexttrend, 1);
				maxlow = lowprice;
			}
		}
		if trend[len(trend)-1] == 0 {
			if len(trend) > 1 && trend[len(trend)-2] != 0 {
				if len(down) < 2 {
					up = append(up, down[len(down)-1]);
				} else {
					up = append(up, down[len(down)-2]);
				}
			} else {
				if len(up) < 2 {
					up = append(up, maxlow);
				} else {
					up = append(up, maxf([]float64{up[len(up)-2], maxlow}));
				}
			}
			direction = "long";
			atrHigh = up[len(up)-1] + dev;
			atrLow = up[len(up)-1] - dev;
		} else {
			if len(trend) > 2 && trend[len(trend)-2] != 1 {
				if len(up) < 2 {
					down = append(down, up[len(up)-1]);
				} else {
					down = append(down, up[len(up)-2]);
				}
			} else {
				if len(down) < 2 {
					down = append(down, minhigh);
				} else {
					down = append(down, minf([]float64{minhigh, down[len(down)-2]}))
				}
			}
			direction = "long";
			atrHigh = down[len(down)-1] + dev;
			atrLow = down[len(down)-1] - dev;
		}
		if trend[len(trend)-1] == 0 {
			out = append(out, []interface{}{atrHigh, up[len(up)-1], atrLow, direction});
		} else {
			out = append(out, []interface{}{atrHigh, down[len(down)-1], atrLow});
		}
	}
	return out;
}
func Ncdf(x float64, mean float64, std float64) float64 {
	x = (x - mean) / std;
	t := 1 / (1+.2315419*math.Abs(x));
	d := .3989423*math.Exp(-x*x/2);
	p := d*t*(.3193815 + t * (-.3565638+t*(1.781478+t*(-1.821256+t*1.330274))));
	if x > 0 {
		return 1-p;
	} else {
		return p;
	}
}
func ZigZag(data [][]float64, perc float64) []float64 {
	var indexes []recentHighLow; min := math.Inf(1); max := math.Inf(-1);
	lmin := false; lmax := false;
	for i := 0; i < len(data); i++ {
		if lmin {
			if indexes[len(indexes)-1].Value >= data[i][1] {
				indexes[len(indexes)-1].Value = data[i][1];
				indexes[len(indexes)-1].Index = i;
			}
			if min >= data[i][1] { min = data[i][1]; }
			hdif := (data[i][0]-min)/min;
			if hdif > perc {
				indexes = append(indexes, RecentHighLow(i, data[i][0]));
				lmax = true;
				lmin = false;
				min = math.Inf(1);
			}
		} else if lmax {
			if indexes[len(indexes)-1].Value <= data[i][0] {
				indexes[len(indexes)-1].Value = data[i][0];
				indexes[len(indexes)-1].Index = i;
			}
			if max <= data[i][1] { max = data[i][1]; }
			ldif := (max-data[i][1])/data[i][1];
			if ldif > perc {
				indexes = append(indexes, RecentHighLow(i, data[i][1]));
				lmin = true;
				lmax = false;
				max = math.Inf(-1);
			}
		} else {
			if min >= data[i][1] {min = data[i][1];}
			if max <= data[i][0] {max = data[i][0];}
			hdif := (data[i][0]-min)/min;
			ldif := (max-data[i][1])/max;
			if ldif > perc && hdif < perc {
				lmin = true;
				indexes = append(indexes, RecentHighLow(0, data[0][0]));
				indexes = append(indexes, RecentHighLow(i, data[i][1]));
			} else if hdif > perc && ldif < perc {
				lmax = true;
				indexes = append(indexes, RecentHighLow(0, data[0][1]));
				indexes = append(indexes, RecentHighLow(i, data[i][0]));
			} else {
				if ldif > hdif {
					lmin = true;
					indexes = append(indexes, RecentHighLow(0, data[0][0]));
					indexes = append(indexes, RecentHighLow(i, data[i][0]));
				} else {
					lmax = true;
					indexes = append(indexes, RecentHighLow(0, data[0][1]));
					indexes = append(indexes, RecentHighLow(i, data[i][0]));
				}
			}
		}
	}
	final := []float64{indexes[0].Value};
	for i := 1; i < len(indexes); i++ {
		len := indexes[i].Index - indexes[i-1].Index;
		delta := (indexes[i].Value - indexes[i-1].Value) / float64(len);
		for z := 1; z <= len; z++ { final = append(final, float64(z)*delta+indexes[i-1].Value); }
	}
	return final;
}
func Psar(data [][]float64, step float64, max float64) []float64 {
	furthest := data[0]; up := true; accel := step; prev := data[0];
	sar := data[0][1]; extreme := data[0][0]; final := []float64{sar};
	for i := 1; i < len(data); i++ {
		sar = sar + accel * (extreme - sar);
		if up {
			sar = minf([]float64{sar, furthest[1], prev[1]});
			if data[i][0] > extreme {
				extreme = data[i][0];
				accel = minf([]float64{accel+step, max});
			}
		} else {
			sar = maxf([]float64{sar, furthest[0], prev[0]});
			if data[i][1] < extreme {
				extreme = data[i][0];
				accel = minf([]float64{accel + step, max});
			}
		}
		if (up && data[i][1] < sar) || (!up && data[i][0] > sar) {
			accel = step;
			sar = extreme;
			up = !up;
			if !up {
				extreme = data[i][1];
			} else {
				extreme = data[i][0];
			}
		}
		furthest = prev;
		prev = data[i];
		final = append(final, sar);
	}
	return final;
}
func Fibbands(data [][]float64, length int, deviations float64) [][]float64 {
	var pl []float64; var deviation []float64; ma := Vwma(data, length); var boll [][]float64;
	for i := 0; i < len(data); i++ {
		pl = append(pl, data[i][0]);
		if len(pl) >= length {
			devi := Std(pl, length);
			deviation = append(deviation, devi * deviations);
			pl = pl[1:];
		}
	}
	for i := 0; i < len(ma); i++ {
		upper1 := ma[i] + (0.236 * deviation[i]);
		upper2 := ma[i] + 0.382 * deviation[i];
		upper3 := ma[i] + 0.5 * deviation[i];
		upper4 := ma[i] + 0.618 * deviation[i];
		upper5 := ma[i] + 0.764 * deviation[i];
		upper6 := ma[i] + deviation[i];
		lower1 := ma[i] - 0.236 * deviation[i];
		lower2 := ma[i] - 0.382 * deviation[i];
		lower3 := ma[i] - 0.5 * deviation[i];
		lower4 := ma[i] - 0.618 * deviation[i];
		lower5 := ma[i] - 0.764 * deviation[i];
		lower6 := ma[i] - deviation[i];
		boll = append(boll, []float64{upper6, upper5, upper4, upper3, upper2, upper1, ma[i], lower1, lower2, lower3, lower4, lower5, lower6});
	}
	return boll;
}
func Martingale(data []float64, bet float64, max float64, multiplier float64) float64 {
	current := bet;
	for i := 0; i < len(data); i++ {
		if data[i] < 0 {
			current *= multiplier;
		} else if data[i] > 0 {
			current = bet;
		}
	}
	return current;
}
func Antimartingale(data []float64, bet float64, max float64, multiplier float64) float64 {
	current := bet;
	for i := 0; i < len(data); i++ {
		if data[i] > 0 {
			current *= multiplier;
		} else if data[i] < 0 {
			current = bet;
		}
	}
	return current;
}
func Sum(data []float64) float64 {
	var sum float64 = 0;
	for i := 0; i < len(data); i++ {
		sum += data[i];
	}
	return sum;
}
func Supertrend(data [][]float64, l int, multiplier float64) [][]float64 {
	atr := Atr(data, l);
	var trend [][]float64;
	for i := l-1; i < len(data); i++ {
		trend = append(trend, []float64{(data[i][0] + data[i][2]) / 2 + multiplier * atr[i], (data[i][0] + data[i][2]) / 2 - multiplier * atr[i]});
	}
	return trend;
}
func Cwma(data []float64, weights []float64) []float64 {
	var ma []float64;
	for i := len(weights); i <= len(data); i++ {
		pl := data[i-len(weights):i]; var sum float64 = 0; var weight float64 = 0;
		for q := 0; q < len(weights); q++ {
			sum += pl[q] * weights[q];
			weight += weights[q];
		}
		ma = append(ma, sum / float64(weight));
	}
	return ma;
}
var Fibnumbers []float64 = []float64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181};
func Permutations(data []float64) float64 {
	var total float64 = 1;
	for i := 0; i < len(data); i++ {
		total *= data[i];
	}
	return total;
}
func Mse(data1 []float64, data2 []float64) float64 {
	var err float64 = 0;
	for i := 0; i < len(data1); i++ {
		err += math.Pow(data2[i]-data1[i], 2);
	}
	return err / float64(len(data1));
}
func Cum(data []float64, l int) []float64 {
	var res []float64;
	for i := l; i <= len(data); i++ {
		res = append(res, Sum(data[i-l:i]));
	}
	return res;
}
func Vwwma(data [][]float64, l int) []float64 {
	var vwma []float64; var weight float64 = 0;
	var weighted [][]float64;
	for i := 0; i < len(data); i++ {
		weighted = append(weighted, []float64{data[i][0] * data[i][1], data[i][1]});
	}
	for i := 1; i <= l; i++ { weight += float64(i); }
	for i := l; i <= len(data); i++ {
		pl := append([][]float64(nil), weighted[i-l:i]...);
		var totalv float64;
		var totalp float64;
		for o := 0; o < len(pl); o++ {
			totalv += pl[o][1] * float64(o+1) / weight;
			totalp += pl[o][0] * float64(o+1) / weight;
		}
		vwma = append(vwma, totalp/totalv);
	}
	return vwma
}
func Elderray(data []float64, l int) [][]float64 {
	var eld [][]float64;
	for i := l; i <= len(data); i++ {
		pl := append([]float64{}, data[i-l:i]...);
		sort.Float64s(pl);
		em := Ema(data[i-l:i], l);
		eld = append(eld, []float64{pl[len(pl)-1]-em[0],pl[0]-em[0]});
	}
	return eld;
}
func Hv(data []float64, l int) []float64 {
	var hv []float64;
	for i := l; i <= len(data); i++ {
		ss := Ssd(data[i-l:i], l);
		vari := ss[0] / float64(l);
		hv = append(hv, math.Sqrt(vari));
	}
	return hv;
}
var t_table [][]float64 = [][]float64{
	{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
	{1,.816,.765,.741,.727,.718,.711,.706,.703,.7,.697,.695,.694,.692,.691,.69,.689,.688,.688,.687,.686,.686,.685,.685,.684},
	{1.376,1.061,.978,.941,.92,.906,.896,.889,.883,.879,.876,.873,.87,.868,.866,.865,.863,.862,.861,.86,.859,.858,.858,.857,.856},
	{1.963,1.386,1.25,1.19,1.156,1.134,1.119,1.108,1.1,1.093,1.088,1.088,1.079,1.076,1.074,1.071,1.069,1.067,1.066,1.064,1.063,1.061,1.06,1.059,1.058},
	{3.078,1.886,1.638,1.533,1.476,1.44,1.415,1.397,1.383,1.372,1.363,1.356,1.35,1.345,1.341,1.337,1.333,1.33,1.328,1.325,1.323,1.321,1.319,1.318,1.316},
	{6.314,2.92,2.353,2.132,2.015,1.943,1.895,1.86,1.833,1.812,1.796,1.782,1.771,1.761,1.753,1.746,1.74,1.734,1.729,1.725,1.721,1.717,1.714,1.711,1.708},
	{12.71,4.303,3.182,2.776,2.571,2.447,2.365,2.306,2.262,2.228,2.201,2.179,2.16,2.145,2.131,2.12,2.11,2.101,2.093,2.086,2.08,2.074,2.069,2.064,2.06},
	{31.82,6.965,4.541,3.747,3.365,3.143,2.998,2.896,2.821,2.764,2.718,2.681,2.65,2.624,2.602,2.583,2.567,2.552,2.539,2.528,2.518,2.508,2.5,2.492,2.485},
	{63.66,9.925,5.841,4.604,4.032,3.707,3.499,3.355,3.25,3.169,3.106,3.055,3.012,2.977,2.947,2.921,2.898,2.878,2.861,2.845,2.831,2.819,2.807,2.797,2.787},
	{318.13,22.327,10.215,7.173,5.893,5.208,4.785,4.501,4.297,4.144,4.025,3.93,3.852,3.787,3.733,3.686,3.646,3.61,3.579,3.552,3.527,3.505,3.485,3.467,3.45},
	{636.62,31.599,12.924,8.61,6.869,5.959,5.408,5.041,4.781,4.587,4.437,4.318,4.221,4.14,4.073,4.015,3.965,3.922,3.883,3.85,3.819,3.792,3.768,3.745,3.725},
}
var v_table []float64 = []float64{.5,.25,.2,.15,.1,.05,.025,.01,.005,.001,.0005};
func Pvalue(t float64, df int) float64 {
	if df > 25 || df < 1 { return 1; } // todo err
	for x := 0; x < len(t_table)-1; x++ {
		if(t >= t_table[x][df-1] && t <= t_table[x+1][df-1]) {
			return v_table[x+1] + (v_table[x+1]-v_table[x]) * (t_table[x+1][df-1] - t) / (t_table[x][df-1] - t_table[x+1][df-1]);
		}
	}
	return 0.0005;
}
func Rvi(data [][]float64, ln int) []float64 {
	var num []float64; var dnom []float64; var rv []float64;
	for i := 3; i < len(data); i++ {
		num = append(num, (data[i][3]-data[i][0]+2*(data[i][3]-data[i-1][0])+2*(data[i][3]-data[i-2][0])+(data[i][3]-data[i-3][0]))/6);
		dnom = append(dnom, (data[i][1]-data[i][2]+2*(data[i][1]-data[i-1][2])+2*(data[i][1]-data[i-2][2])+(data[i][1]-data[i-3][2]))/6);
		if len(num) >= ln {
			sn := Sma(num, ln);
			dn := Sma(dnom, ln);
			rv = append(rv, sn[0]/dn[0]);
			num = num[1:];
			dnom = dnom[1:];
		}
	}
	return rv;
}
func Rvi_signal(rv []float64) []float64 {
	var sig []float64;
	for i := 3; i < len(rv); i++ {
		sig = append(sig,(rv[i]+2*rv[i-1]+2*rv[i-2]+rv[i-3])/6);
	}
	return sig;
}
func Rsi_divergence(data []float64, ln int, rs func([]float64, int) []float64) []float64 {
	rd := rs(data, ln); var out []float64;
	data = Mom(data[ln-1:len(data)], 2, false);
	for i := 0; i < len(data); i++ {
		if (data[i] > 0 && rd[i] < 0) || (data[i] < 0 && rd[i] > 0) {
			out = append(out, 1);
		} else {
			out = append(out, 0);
		}
	}
	return out;
}
func Divergence(data1 []float64, data2 []float64) []float64 {
	if len(data1) > len(data2) { data1 = data1[len(data1)-len(data2):]; }
	if len(data2) > len(data1) { data2 = data2[len(data2)-len(data1):]; }
	var out []float64;
	for i := 1; i < len(data1); i++ {
		if (data1[i] > data1[i-1] && data2[i] < data2[i-1]) || (data1[i] < data1[i-1] && data2[i] > data2[i-1]) {
			out = append(out, 1);
		} else {
			out = append(out, 0);
		}
	}
	return out;
}
func Times_up(data []float64, l int) []float64 {
	var out []float64;
	for i := l; i < len(data); i++ {
		var up float64 = 1;
		for x := i-l+1; x <= i; x++ {
			if data[x-1] > data[x] {
				up = 0;
				break;
			}
		}
		out = append(out, up);
	}
	return out;
}
