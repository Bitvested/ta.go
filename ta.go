package ta
import (
  "fmt"
  "sort"
  "math"
  "math/rand"
  "strings"
)
func Median(data []float64, l int) []float64  {
  var med []float64;
  for i := l; i <= len(data); i++ {
    pl := append([]float64(nil), data[i-l:i]...);
    sort.Float64s(pl);
    med = append(med, pl[len(pl)/2-1]);
  }
  return med;
}
func Normalize(data []float64, m float64) []float64 {
  var normal []float64;
  tmp := append([]float64(nil), data...);
  sort.Float64s(tmp);
  var max float64 = tmp[len(tmp)-1]*(1+m);
  var min float64 = tmp[0]*(1-m);
  for i := 0; i < len(data); i++ {
    normal = append(normal, 1-(max-data[i])/(max-min));
  }
  return normal;
}
func Denormalize(norm []float64, data []float64, m float64) []float64 {
  var denormal []float64;
  tmp := append([]float64(nil), data...);
  sort.Float64s(tmp);
  var max float64 = tmp[len(tmp)-1]*(1+m);
  var min float64 = tmp[0]*(1-m);
  for i := 0; i < len(norm); i++ {
    denormal = append(denormal, min+norm[i]*(max-min));
  }
  return denormal;
}
func Standardize(data []float64) []float64 {
  var mean []float64 = Sma(data, len(data));
  var std float64 = Std(data, len(data));
  var res []float64;
  for i := 0; i < len(data); i++ {
    res = append(res, (data[i]-mean[0])/std);
  }
  return res;
}
func Mad(data []float64, l int) []float64 {
  var med []float64;
  for i := l; i <= len(data); i++ {
    tmp := append([]float64(nil), data[i-l:i]...);
    m := Median(tmp, len(tmp));
    var adev []float64;
    for x := 0; x < len(tmp); x++ { adev = append(adev, math.Abs(tmp[x]-m[0])); }
    adev = Median(adev, len(adev));
    med = append(med, adev[0]);
  }
  return med;
}
func Aad(data []float64, l int) []float64 {
  var med []float64;
  for i := l; i <= len(data); i++ {
    tmp := append([]float64(nil), data[i-l:i]...);
    m := Sma(tmp, len(tmp));
    var sum float64 = 0;
    for x := 0; x < len(tmp); x++ { sum += math.Abs(float64(tmp[x])-float64(m[0])); }
    med = append(med, sum/float64(l));
  }
  return med;
}
func Sma(data []float64, l int) []float64 {
  var sma []float64;
  for i := l; i <= len(data); i++ {
    var avg float64;
    tmp := append([]float64(nil), data[i-l:i]...);
    for x := 0; x < len(tmp); x++ {avg += tmp[x]}
    sma = append(sma, avg/float64(l));
  }
  return sma;
}
func Ssd(data []float64, l int) []float64 {
  var sd []float64;
  for i := l; i <= len(data); i++ {
    tmp := append([]float64(nil), data[i-l:i]...);
    mean := Sma(tmp, l);
    var sum float64;
    for x := 0; x < len(tmp); x++ {
      sum += math.Pow(tmp[x]-mean[0], 2);
    }
    sd = append(sd, math.Pow(sum, 0.5));
  }
  return sd;
}
func Rsi(data []float64, l int) []float64 {
  var arrsi []float64;
  for i := l; i < len(data); i++ {
    pl := append([]float64(nil), data[i-l:i+1]...); var loss float64; var gain float64;
    for q := 1; q < len(pl); q++ {
      if pl[q]-pl[q-1] < 0 {
        loss += math.Abs(pl[q]-pl[q-1]);
      } else {
        gain += pl[q]-pl[q-1];
      }
    }
    var rsi float64 = 100-100/(1+((gain/float64(l))/(loss/float64(l))));
    arrsi = append(arrsi, rsi);
  }
  return arrsi;
}
func Wsma(data []float64, l int) []float64 {
  var wsm []float64;
  var weight float64 = 1/float64(l);
  for i := l; i <= len(data); i++ {
    if len(wsm) > 0 {
      wsm = append(wsm, (data[i-1]-wsm[len(wsm)-1])*weight+wsm[len(wsm)-1]);
      continue;
    }
    pl := append([]float64(nil), data[i-l:i]...);
    var average float64;
    for q := 0; q < len(pl); q++ { average += pl[q]; }
    wsm = append(wsm, average/float64(l));
  }
  return wsm;
}
func Wrsi(data []float64, l int) []float64 {
  var arrsi []float64; var u []float64; var d []float64;
  for i := 1; i < len(data); i++ {
    if(data[i]-data[i-1] < 0) {
      d = append(d, math.Abs(data[i]-data[i-1]));
      u = append(u, 0);
    } else {
      d = append(d, 0);
      u = append(u, data[i]-data[i-1]);
    }
  }
  d = Wsma(d, l); u = Wsma(u, l);
  for i := 0; i < len(d); i++ {
    arrsi = append(arrsi, 100-100/(1+(u[i]/d[i])));
  }
  return arrsi;
}
func Ema(data []float64, l int) []float64 {
  var ema []float64; var weight float64 = 2 / (float64(l)+1);
  for i := l; i <= len(data); i++ {
    if len(ema) > 0 {
      ema = append(ema, (data[i-1] - ema[len(ema)-1]) * weight + ema[len(ema)-1]);
      continue;
    }
    pl := append([]float64(nil), data[i-l:i]...);
    var average float64;
    for q := 0; q < len(pl); q++ { average += pl[q]; }
    ema = append(ema, average/float64(l));
  }
  return ema;
}
func Smma(data []float64, l int) []float64 {
  var smma []float64;
  for i := l; i <= len(data); i++ {
    pl := append([]float64(nil), data[i-l:i]...);
    var average float64;
    for q := 0; q < len(pl); q++ {
      average += pl[q];
    }
    if len(smma) <= 0 {
      smma = append(smma, average/float64(l));
    } else {
      smma = append(smma, (average-smma[len(smma)-1])/float64(l));
    }
  }
  return smma[1:];
}
func Wma(data []float64, l int) []float64 {
  var weight int; var wma []float64;
  for i := 1; i <= l; i++ { weight += i; }
  for i := l; i <= len(data); i++ {
    pl := append([]float64(nil), data[i-l:i]...);
    var average float64;
    for q := 0; q < l; q++ { average += (pl[q] * float64(q+1) / float64(weight)) }
    wma = append(wma, average);
  }
  return wma;
}
func Pwma(data []float64, l int) []float64 {
  var weight int; var wmaa []float64; var b int = l; var weights []float64;
  for i := int(float64(l) / 2); i >= 1; i-- {
    if(i % 1 != 0) {
      weight += (i * b);
    } else {
      weights = append(weights, float64(i*b));
      weight += (i*b*2);
    }
    weights = append([]float64{float64(i*b)}, weights...);
    b--;
  }
  for i := l; i <= len(data); i++ {
    var average float64;
    pl := append([]float64(nil), data[i-l:i]...);
    for a := l-1; a >= 0; a-- {
      average += (pl[a]*weights[a]/float64(weight));
    }
    wmaa = append(wmaa, average);
  }
  return wmaa
}
func Hwma(data []float64, l int) []float64 {
  var weight int; var wmaa []float64; var b int = l; var weights []float64;
  for i := 1; i <= int(l / 2); i++ {
    if(i % 1 != 0) {
      weight += (i*b);
    } else {
      weights = append(weights, float64(i*b));
      weight += (i*b*2);
    }
    weights = append([]float64{float64(i*b)}, weights...);
    b--;
  }
  for i := l; i <= len(data); i++ {
    var average float64; pl := append([]float64(nil), data[i-l:i]...);
    for a := l - 1; a >= 0; a-- {
      average += (pl[a] * weights[a] / float64(weight));
    }
    wmaa = append(wmaa, average);
  }
  return wmaa;
}
func Vwma(data [][]float64, l int) []float64 {
  var vwma []float64; var weighted [][]float64;
  for i := 0; i < len(data); i++ { weighted = append(weighted, []float64{data[i][0]*data[i][1], data[i][1]}); }
  for i := l; i <= len(data); i++ {
    pl := append([][]float64(nil), weighted[i-l:i]...);
    var totalv float64; var totalp float64;
    for o := 0; o < len(pl); o++ {
      totalv += pl[o][1];
      totalp += pl[o][0];
    }
    vwma = append(vwma, totalp/totalv);
  }
  return vwma;
}
func Lsma(data []float64, l int) []float64 {
  var lr []float64;
  for i := l; i <= len(data); i++ {
    pl := append([]float64(nil), data[i-l:i]...); var sum_x float64 = 0; var sum_y float64 = 0; var sum_xy float64 = 0; var sum_xx float64 = 0; var sum_yy float64 = 0;
    for q := 1; q <= l; q++ {
      sum_x += float64(q);
      sum_y += pl[q-1];
      sum_xy += (pl[q-1]*float64(q));
      sum_xx += (float64(q)*float64(q));
      sum_yy += (pl[q-1]*pl[q-1]);
    }
    m := ((sum_xy-sum_x*sum_y/float64(l))/(sum_xx-sum_x*sum_x/float64(l)));
    b := sum_y/float64(l)-m*sum_x/float64(l);
    lr = append(lr, m*float64(l)+b);
  }
  return lr;
}
func Hull(data []float64, l int) []float64 {
  var pl []float64; var hma []float64; var ewma []float64 = Wma(data, l); var sqn int = int(math.Pow(float64(l), 0.5));
  var first []float64 = Wma(data, int(l/2));
  first = first[len(first)-len(ewma):];
  for i := 0; i < len(ewma); i++ {
    pl = append(pl, (first[i]*2)-ewma[i]);
    if len(pl) >= sqn {
      var h []float64 = Wma(pl, sqn);
      hma = append(hma, h[0]);
      pl = pl[1:];
    }
  }
  return hma;
}
func Macd(data []float64, l1 int, l2 int) []float64 {
  if(l1 > l2) { l1, l2 = l2, l1; }
  var ema []float64 = Ema(data, l1); var emb []float64 = Ema(data, l2);
  var macd []float64;
  ema = ema[l2-l1:];
  for i := 0; i < len(emb); i++ {
    macd = append(macd, ema[i]-emb[i]);
  }
  return macd;
}
func Std(data []float64, l int) float64 {
  var mean []float64 = Sma(data, l);
  var average float64;
  for i := len(data)-l; i < len(data); i++ {
    average += math.Pow(data[i]-mean[len(mean)-1], 2);
  }
  var std float64 = math.Pow(average/float64(l), 0.5);
  return std;
}
func Normsinv(p float64) float64 {
  a1 := -39.6968302866538; a2 := 220.946098424521; a3 := -275.928510446969; a4 := 138.357751867269; a5 := -30.6647980661472; a6 := 2.50662827745924;
  b1 := -54.4760987982241; b2 := 161.585836858041; b3 := -155.698979859887; b4 := 66.8013118877197; b5 := -13.2806815528857; c1 := -7.78489400243029E-03;
  c2 := -0.322396458041136; c3 := -2.40075827716184; c4 := -2.54973253934373; c5 := 4.37466414146497; c6 := 2.93816398269878; d1 := 7.78469570904146E-03;
  d2 := 0.32246712907004; d3 := 2.445134137143; d4 := 3.75440866190742; p_low := 0.02425; p_high := 1 - p_low;
  if p < 0 || p > 1 {
    return 0
  }
  if p < p_low {
    q := math.Pow(-2*math.Log(p), 0.5);
    return (((((c1*q+c2)*q+c3)*q+c4)*q+c5)*q+c6)/((((d1*q+d2)*q+d3)*q+d4)*q+1);
  }
  if p <= p_high {
    q := p-0.5;
    r := q*q;
    return (((((a1*r+a2)*r+a3)*r+a4)*r+a5)*r+a6)*q/(((((b1*r+b2)*r+b3)*r+b4)*r+b5)*r+1);
  }
  q := math.Pow(-2*math.Log(1-p), 0.5);
  return -(((((c1*q+c2)*q+c3)*q+c4)*q+c5)*q+c6)/((((d1*q+d2)*q+d3)*q+d4)*q+1);
}
func Cor(data1 []float64, data2 []float64) float64 {
  var d1avg []float64 = Sma(data1, len(data1)); var d2avg []float64 = Sma(data2, len(data2));
  var sumavg float64; var sx float64; var sy float64;
  for i := 0; i < len(data1); i++ {
    x := data1[i]-d1avg[0]; y := data2[i]-d2avg[0];
    sumavg += (x*y); sx += math.Pow(x, 2); sy += math.Pow(y, 2);
  }
  n := float64(len(data1)-1);
  sx /= n; sy /= n; sx = math.Pow(sx, 0.5); sy = math.Pow(sy, 0.5);
  return (sumavg/(n*sx*sy));
}
func Dif(n float64, o float64) float64 {
  return (n-o)/o;
}
func Drawdown(data []float64) float64 {
  max := data[0]; min := data[0]; big := 0.0;
  for y := 1; y < len(data); y++ {
    if data[y] > max {
      if min != 0 {
        diff := Dif(min, max);
        if diff < big { big = diff; }
        min = data[y];
      }
      max = data[y];
    }
    if data[y] < min { min = data[y]; }
  }
  diff := Dif(min, max);
  if diff < big {
    big = diff
  }
  return big;
}
func reverseFloats(data []float64) []float64 {
  var ret []float64;
  for i := len(data)-1; i >= 0; i-- {
    ret = append(ret, data[i])
  }
  return ret;
}
func indexOf(data []float64, index float64) int {
  var ret int;
  for i := len(data)-1; i >= 0; i-- {
    if data[i] == index {
      ret = i;
      break;
    }
  }
  return ret;
}
func AroonUp(data []float64, l int) []float64 {
  var aroon []float64;
  for i := l; i <= len(data); i++ {
    pl := append([]float64(nil), data[i-l:i]...);
    hl := append([]float64(nil), data[i-l:i]...);
    sort.Float64s(hl);
    pl = reverseFloats(pl);
    index := indexOf(pl, hl[len(hl)-1]);
    aroon = append(aroon, (100 * (float64(index))/(float64(l)-1)));
  }
  return aroon;
}
func AroonDown(data []float64, l int) []float64 {
  var aroon []float64;
  for i := l; i <= len(data); i++ {
    pl := append([]float64(nil), data[i-l:i]...);
    hl := append([]float64(nil), data[i-l:i]...);
    sort.Float64s(hl);
    //pl = reverseFloats(pl);
    index := indexOf(pl, hl[0]);
    aroon = append(aroon, (100 * (float64(index))/(float64(l)-1)));
  }
  return aroon;
}
func AroonOsc(data []float64, l int) []float64 {
  var aroon []float64;
  u := AroonUp(data, l); d := AroonDown(data, l);
  for i := 0; i < len(u); i++ { aroon = append(aroon, u[i]-d[i]) }
  return aroon;
}
func Mfi(data [][]float64, l int) []float64 {
  var mfi []float64; var n []float64; var p []float64;
  for i := 0; i < len(data); i++ {
    n = append(n, data[i][1]);
    p = append(p, data[i][0]);
  }
  for i := l; i <= len(data); i++ {
    var pos float64; var neg float64;
    for q := i-l; q < i; q++ {
      pos += p[q];
      neg += n[q];
    }
    mfi = append(mfi, 100 - 100 / (1 + pos / neg));
  }
  return mfi;
}
func Roc(data []float64, l int) []float64 {
  var roc []float64;
  for i := l; i <= len(data); i++ {
    roc = append(roc, (data[i-1]-data[i-l])/data[i-l]);
  }
  return roc;
}
func Cop(data []float64, l1 int, l2 int, l3 int) []float64 {
  max := math.Max(float64(l1), float64(l2)); var co []float64;
  for i := int(max) + l3; i < len(data); i++ {
    r1 := append([]float64(nil), data[i-(int(max)+l3):i]...); r2 := r1[:]; var tmp []float64;
    r1 = Roc(r1, l1); r2 = Roc(r2, l2);
    r1 = r1[len(r1)-len(r2):]; r2 = r2[len(r2)-len(r1):];
    for a := 0; a < len(r1); a++ {
      tmp = append(tmp, r1[a]+r2[a]);
    }
    tmp = Wma(tmp, l3);
    co = append(co, tmp[len(tmp)-1]);
  }
  return co;
}
func max(v []int) int {
  sort.Ints(v);
  return v[len(v)-1];
}
func min(v []int) int {
  sort.Ints(v);
  return v[0];
}
func maxf(v []float64) float64 {
  sort.Float64s(v);
  return v[len(v)-1];
}
func minf(v []float64) float64 {
  sort.Float64s(v);
  return v[0];
}
func Kst(data []float64, r1 int, r2 int, r3 int, r4 int, s1 int, s2 int, s3 int, s4 int, sig int) [][]float64 {
  var ks []float64; var fs [][]float64; ms := (max([]int{r1,r2,r3,r4}) + max([]int{s4, s3, s2, s1}));
  for i := ms; i <= len(data); i++ {
    pl := append([]float64(nil), data[i-ms:i]...);
    rcma1 := Roc(pl, r1);
    rcma2 := Roc(pl, r2);
    rcma3 := Roc(pl, r3);
    rcma4 := Roc(pl, r4);
    rcma1 = Sma(rcma1, s1);
    rcma2 = Sma(rcma2, s2);
    rcma3 = Sma(rcma3, s3);
    rcma4 = Sma(rcma4, s4);
    ks = append(ks, rcma1[len(rcma1)-1] + rcma2[len(rcma2)-1] + rcma3[len(rcma3)-1] +rcma4[len(rcma4)-1]);
  }
  sl := Sma(ks, sig);
  ks = ks[len(ks)-len(sl):];
  for i := 0; i < len(sl); i++ {
    fs = append(fs, []float64{ks[i], sl[i]});
  }
  return fs;
}
func Obv(data [][]float64) []float64 {
  var obv []float64;
  obv = append(obv, 0);
  for i := 1; i < len(data); i++ {
    if data[i][1] > data[i-1][1] {
      obv = append(obv, obv[len(obv)-1]+data[i][0]);
    }
    if data[i][1] < data[i-1][1] {
      obv = append(obv, obv[len(obv)-1]-data[i][0]);
    }
    if(data[i][1] == data[i-1][1]) {
      obv = append(obv, obv[len(obv)-1]);
    }
  }
  return obv;
}
func Vwap(data [][]float64, l int) []float64 {
  var vwap []float64; var weighted [][]float64;
  for i := 0; i < len(data); i++ { weighted = append(weighted, []float64{data[i][0]*data[i][1], data[i][1]}); }
  for i := l; i <= len(weighted); i++ {
    pl := append([][]float64(nil), weighted[i-l:i]...);
    var totalv float64; var totalp float64;
    for o := 0; o < len(pl); o++ {
      totalv += pl[o][1];
      totalp += pl[o][0];
    }
    vwap = append(vwap, totalp/totalv);
  }
  return vwap;
}
func Mom(data []float64, l int, p bool) []float64 {
  var mom []float64;
  for i := l-1; i < len(data); i++ {
    if(p) {
      mom = append(mom, data[i]/data[i-(l-1)]*100);
    } else {
      mom = append(mom, data[i]-data[i-(l-1)]);
    }
  }
  return mom;
}
func MomOsc(data []float64, l int) []float64 {
  var osc []float64;
  for i := l; i <= len(data); i++ {
    var sumh float64; var suml float64;
    for a := 1; a < l; a++ {
      if(data[i-l+(a-1)] < data[i-l+a]) {
        sumh += data[i-l+a];
      } else {
        suml += data[i-l+a];
      }
    }
    osc = append(osc, (sumh-suml)/(sumh+suml)*100);
  }
  return osc;
}
func Ha(data [][]float64) [][]float64 {
  var ha [][]float64 = [][]float64{{(data[0][0]+data[0][3])/2, data[0][1], data[0][2], (data[0][0]+data[0][1]+data[0][2]+data[0][3])/4}};
  for i := 1; i < len(data); i++ {
    ha = append(ha, []float64{(ha[len(ha)-1][0]+ha[len(ha)-1][3])/2, maxf([]float64{ha[len(ha)-1][0], ha[len(ha)-1][3], data[i][1]}), minf([]float64{ha[len(ha)-1][0], ha[len(ha)-1][3], data[i][2]}), (data[i][0]+data[i][1]+data[i][2]+data[i][3])/4});
  }
  return ha;
}
func decimalplaces(n float64) int {
  decimalPlaces := fmt.Sprintf("%f", n-math.Floor(n)); // produces 0.xxxx0000
  decimalPlaces = strings.Replace(decimalPlaces, "0.", "", -1); // remove 0.
  decimalPlaces = strings.TrimRight(decimalPlaces, "0"); // remove trailing 0s
  return len(decimalPlaces);
}
func Ren(data [][]float64, bs float64) [][]float64 {
  var re [][]float64;
  decimals := float64(decimalplaces(bs));
  bh := math.Ceil(data[0][0]/bs*math.Pow(10, decimals))/math.Pow(10, decimals)*bs; bl:=bh-bs;
  for i := 1; i < len(data); i++ {
    if data[i][0] > bh + bs {
      for ok := true; ok == true; ok = (data[i][0]>bh+bs) {
        re = append(re, []float64{bh,bh+bs,bh,bh+bs});
        bh+=bs;
        bl+=bs;
      }
    }
    if data[i][1] < bl - bs {
      for ok := true; ok == true; ok = (data[i][1]<bl-bs) {
        re = append(re, []float64{bl,bl,bl-bs,bl-bs});
        bh-=bs;
        bl-=bs;
      }
    }
  }
  return re;
}
func Tsi(data []float64, llen int, slen int, sig int) [][]float64 {
  var mom []float64; var abs []float64; var ts []float64; var tsi [][]float64;
  for i := 1; i < len(data); i++ {
    mom = append(mom, data[i]-data[i-1]);
    abs = append(abs, math.Abs(data[i]-data[i-1]));
  }
  sma1 := Ema(mom, llen);
  sma2 := Ema(abs, llen);
  ema1 := Ema(sma1, slen);
  ema2 := Ema(sma2, slen);
  for i := 0; i < len(ema1); i++ {
    ts = append(ts, ema1[i] / ema2[i]);
  }
  tma := Ema(ts, sig);
  ts = ts[len(ts)-len(tma):];
  for i := 0; i < len(tma); i++ {
    tsi = append(tsi, []float64{tma[i], ts[i]});
  }
  return tsi;
}
func Bop(data [][]float64, l int) []float64 {
  var bo []float64;
  for i := 0; i < len(data); i++ {
    bo = append(bo, (data[i][3]-data[i][0])/(data[i][1]-data[i][2]));
  }
  bo = Sma(bo, l);
  return bo;
}
func Fi(data [][]float64, l int) []float64 {
  var ff []float64; var pl []float64;
  for i := 1; i < len(data); i++ {
    pl = append(pl, (data[i][0]-data[i-1][0])*data[i][1]);
    if len(pl) >= l {
      vfi := Ema(pl, l);
      ff = append(ff, vfi[len(vfi)-1]);
      pl = pl[1:];
    }
  }
  return ff;
}
func Asi(data [][]float64) []float64 {
  var a []float64;
  for i := 1; i < len(data); i++ {
    c := data[i][1]; cy := data[i-1][1]; h := data[i][0]; hy := data[i-1][0];
    l := data[i][2]; ly := data[i-1][2]; var k float64; if(hy-c>ly-c) {k=hy-c} else {k=ly-c};
    o := data[i][0]; oy := data[i-1][0]; var r float64; t := maxf([]float64{data[i][0], data[i-1][0]}) - minf([]float64{data[i][2], data[i-1][2]});
    if(h-cy>l-cy&&h-cy>h-l) {r = h-cy-(l-cy)/2+(cy-oy)/4;}
    if(l-cy>h-cy&&l-cy>h-l) {r = l-cy-(h-cy)/2+(cy-oy)/4;}
    if(h-l>h-cy&&h-l>l-cy) {r = h-l+(cy-oy)/4}
    a = append(a, 50 * ((cy-c+(cy-oy)/2+(c-o)/2)/r)*k/t);
  }
  return a;
}
func Ao(data [][]float64, l1 int, l2 int) []float64 {
  var processed []float64; var a []float64;
  for i := 0; i < len(data); i++ {processed = append(processed, (data[i][0]+data[i][1])/2);}
  for i := l2; i <= len(processed); i++ {
    pl := processed[i-l2:i];
    f := Sma(pl, l1);
    s := Sma(pl, l2);
    a = append(a, f[len(f)-1]-s[len(s)-1]);
  }
  return a;
}
func Pr(data []float64, l int) []float64 {
  var n []float64;
  for i := l; i <= len(data); i++ {
    pl := append([]float64(nil), data[i-l:i]...); highd := maxf(pl); lowd := minf(pl);
    n = append(n, (highd-data[i-1])/(highd-lowd)*-100);
  }
  return n;
}
func Don(data [][]float64, l int) [][]float64 {
  var channel [][]float64;
  for i := l; i <= len(data); i++ {
    pl := data[i-l:i]; var highs []float64; var lows []float64;
    for h := 0; h < len(pl); h++ {
      highs = append(highs, pl[h][0]);
      lows = append(lows, pl[h][1]);
    }
    max := maxf(highs);
    min := minf(lows);
    channel = append(channel, []float64{max, (max+min)/2, min});
  }
  return channel;
}
func Ichimoku(data [][]float64, l1 int, l2 int, l3 int, l4 int) [][]float64 {
  var pl [][]float64; var cloud [][]float64; var place [][]float64;
  for i := 0; i < len(data); i++ {
    pl = append(pl, data[i]);
    if len(pl) >= l3 {
      var highs []float64; var lows []float64;
      for a := 0; a < len(pl); a++ {
        highs = append(highs, pl[a][0]);
        lows = append(lows, pl[a][2]);
      }
      tsen := (maxf(highs[len(highs)-l1:]) + minf(lows[len(lows)-l1:])) / 2;
      ksen := (maxf(highs[len(highs)-l2:]) + minf(lows[len(lows)-l2:])) / 2;
      senka := data[i][1] + ksen;
      senkb := (maxf(highs[len(highs)-l3:]) + minf(lows[len(lows)-l2:])) / 2;
      chik := data[i][1];
      place = append(place, []float64{tsen, ksen, senka, senkb, chik});
      pl = pl[1:];
    }
  }
  for i := l4; i < len(place)-l4; i++ {
    cloud = append(cloud, []float64{place[i][0], place[i][1], place[i+l4][2], place[i+(l4)][3], place[i-l4][4]});
  }
  return cloud;
}
func Stoch(data [][]float64, l int, sd int, sk int) [][]float64 {
  var stoch [][]float64; var high []float64; var low []float64; var ka []float64;
  if l < sd { l, sd = sd, l }
  if sk > sd { sk, sd = sd, sk }
  for i := 0; i < len(data); i++ {
    high = append(high, data[i][0]);
    low = append(low, data[i][2]);
    if len(high) >= l {
      highd := maxf(high);
      lowd := minf(low);
      k := 100 * (data[i][1]-lowd)/(highd-lowd);
      ka = append(ka, k);
    }
    if(sk > 0 && len(ka) > sk) {
      smoothedk := Sma(ka, sk);
      ka = append(ka, smoothedk[len(smoothedk)-1]);
    }
    if len(ka) - sk >= sd {
      d := Sma(ka, sd);
      stoch = append(stoch, []float64{ka[len(ka)-1], d[len(d)-1]});
      high = high[1:];
      low = low[1:];
      ka = ka[1:];
    }
  }
  return stoch;
}
func Atr(data [][]float64, l int) []float64 {
  var atr = []float64{data[0][0]-data[0][2]};
  for i := 1; i < len(data); i++ {
    t0 := maxf([]float64{data[i][0]-data[i-1][1], data[i][2]-data[i-1][1], data[i][0]-data[i][2]});
    atr = append(atr, (atr[len(atr)-1]*(float64(l)-1)+t0)/float64(l));
  }
  return atr;
}
func Kama(data []float64, l1 int, l2 int, l3 int) []float64 {
  ka := Sma(data, l1); ka = []float64{ka[len(ka)-1]};
  for i := l1+1; i < len(data); i++ {
    var vola float64; change := math.Abs(data[i]-data[i-l1]);
    for a := 1; a < l1; a++ {
      vola += math.Abs(data[i-a]-data[i-a-1]);
    }
    sc := math.Pow(change/vola*(2/(float64(l2)+1)-2/(float64(l3)+1)+2/(float64(l3)+1)), 2);
    ka = append(ka, ka[len(ka)-1]+sc*(data[i]-ka[len(ka)-1]));
  }
  return ka;
}
func Bands(data []float64, l int, dev float64) [][]float64 {
  var pl []float64; var deviation []float64; var boll [][]float64;
      sma := Sma(data, l);
  for i := 0; i < len(data); i++ {
    pl = append(pl, data[i]);
    if len(pl) >= l {
      devi := Std(pl, l);
      deviation = append(deviation, devi);
      pl = pl[1:];
    }
  }
  for i := 0; i < len(sma); i++ {
    boll = append(boll, []float64{sma[i]+deviation[i]*dev, sma[i], sma[i]-deviation[i]*dev});
  }
  return boll;
}
func Bandwidth(data []float64, l int, dev float64) []float64 {
  var boll []float64;
  band := Bands(data, l, dev);
  for i := 0; i < len(band); i++ {
    boll = append(boll, (band[i][0]-band[i][2])/band[i][1]);
  }
  return boll;
}
func Keltner(data [][]float64, l int, dev float64) [][]float64 {
  var closing []float64; atr := Atr(data, l); var kelt [][]float64;
  for i := 0; i < len(data); i++ {
    closing = append(closing, (data[i][0]+data[i][1]+data[i][2])/3);
  }
  kma := Sma(closing, l);
  atr = atr[l-1:];
  for i := 0; i < len(kma); i++ {
    kelt = append(kelt, []float64{kma[i]+atr[i]*dev, kma[i], kma[i]-atr[i]*dev});
  }
  return kelt;
}
func Variance(data []float64, l int) []float64 {
  var va []float64;
  for i := l; i <= len(data); i++ {
    tmp := data[i-l:i]; mean := Sma(tmp, l); var sum float64;
    for x := 0; x < len(tmp); x++ {
      sum += math.Pow(tmp[x]-mean[len(mean)-1], 2);
    }
    va = append(va, sum/float64(l));
  }
  return va;
}
func Sim(data []float64, l int, sims int) [][]float64 {
  var sd [][]float64;
  for i := 0; i < sims; i++ {
    projected := append([]float64(nil), data...);
    for x := 0; x < l; x++ {
      var change []float64;
      for y := 1; y < len(projected); y++ {
        df := Dif(projected[y],projected[y-1]);
        change = append(change, df);
      }
      mean := Sma(change, len(change));
      std := Std(change, len(change));
      random := Normsinv(rand.Float64());
      projected = append(projected, projected[len(projected)-1]*math.Exp((mean[0]-(std*std)/2)+std*random));
    }
    sd = append(sd, projected);
  }
  return sd;
}
func Percentile(data [][]float64, perc float64) []float64 {
  var ret []float64;
  for i := 0; i < len(data[0]); i++ {
    var tmp []float64;
    for x := 0; x < len(data); x++ {
      tmp = append(tmp, data[x][i]);
    }
    sort.Float64s(tmp);
    ret = append(ret, tmp[int(float64(len(tmp)-1)*perc)]);
  }
  return ret;
}
func Kmeans(data []float64, clusters int) [][]float64 {
  var means [][]float64; var n int; var centers []float64; var old []float64; changed := false; init := int(math.Round(float64(len(data))/float64(clusters+1)));
  for i := 0; i < clusters; i++ { centers = append(centers, data[init*(1+i)]) }
  for ok := true; ok == true; ok = changed {
    for i := 0; i < clusters; i++ { means = append(means, []float64{}) }
    for x := 0; x < len(data); x++ {
      var oldrange float64 = -1;
      for y := 0; y < clusters; y++ {
        r := math.Abs(centers[y]-data[x]);
        if oldrange == -1 {
          oldrange = r;
          n = y;
        } else if r <= oldrange {
          oldrange = r;
          n = y;
        }
      }
      means[n] = append(means[n], data[x]);
    }
    old = centers;
    for x := 0; x < clusters; x++ {
      var sum float64;
      for y := 0; y < len(means[x]); y++ {
        sum += means[x][y];
      }
      m := sum / float64(len(means[n]));
      centers[x] = m;
    }
    for x := 0; x < clusters; x++ { if centers[x] != old[x] { break; } }
  }
  return means;
}
func Alligator(data []float64, jl int, tl int, ll int, js int, ts int, ls int) [][]float64 {
  var ret [][]float64; var jaw []float64 = Smma(data, jl); var teeth []float64 = Smma(data, tl); var lips []float64 = Smma(data, ll);
  teeth = teeth[len(teeth)-len(jaw):]; lips = lips[len(lips)-len(jaw):];
  for i := len(jaw)-1; i >= 7; i-- {
    ret = append(ret, []float64{jaw[i-(js-1)], teeth[i-(ts-1)], lips[i-(ls-1)]});
  }
  return ret;
}
func Gator(data []float64, jl int, tl int, ll int, js int, ts int, ls int) [][]float64 {
  var ret [][]float64; var jaw []float64 = Smma(data, jl); var teeth []float64 = Smma(data, tl); var lips []float64 = Smma(data, ll);
  teeth = teeth[len(teeth)-len(jaw):]; lips = lips[len(lips)-len(jaw):];
  for i := len(jaw)-1; i >= (js-1); i-- {
    ret = append(ret, []float64{jaw[i-(js-1)]-teeth[i-(ts-1)], -(math.Abs(teeth[i-(ts-1)]-lips[i-(ls-1)]))})
  }
  return ret;
}
func Fractals(data [][]float64) [][]bool {
  var fractals = [][]bool{{false,false},{false,false}};
  for i := 2; i < len(data)-2; i++ {
    var up bool; var down bool;
    if(data[i-2][0] < data[i][0] && data[i-1][0] < data[i][0] && data[i][0] > data[i+1][0] && data[i][0] > data[i+2][0]) { up = true; } else { up = false; }
    if(data[i-2][1] > data[i][1] && data[i-1][1] > data[i][1] && data[i][1] < data[i+1][1] && data[i][1] < data[i+2][1]) { down = true; } else { down = false; }
    fractals = append(fractals, []bool{up, down});
  }
  fractals = append(fractals, []bool{false, false});
  fractals = append(fractals, []bool{false, false});
  return fractals;
}
func ChaikinOsc(data [][]float64, ema1 int, ema2 int) []float64 {
  var cha []float64; var adl []float64;
  for i := 0; i < len(data); i++ {
    var mfm = ((data[i][1]-data[i][2])-(data[i][0]-data[i][1]))/(data[i][0]-data[i][2]);
    if(math.IsNaN(mfm)) {
      adl = append(adl, 0);
    } else {
      adl = append(adl, mfm*data[i][3]);
    }
  }
  ef := Ema(adl, ema1); es := Ema(adl, ema2);
  if(len(ef) > len(es)) { ef = ef[len(ef)-len(es):] } else { es = es[len(es)-len(ef):] }
  for i := 0; i < len(ef); i++ { cha = append(cha, ef[i]-es[i]); }
  return cha;
}
func Envelope(data []float64, l int, p float64) [][]float64 {
  var enve [][]float64;
  for i := l; i < len(data); i++ {
    sm := Sma(data[i-l:i], l);
    enve = append(enve, []float64{sm[0]*(1+p),sm[0],sm[0]*(1-p)});
  }
  return enve;
}
