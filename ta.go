package main
import (
  "fmt"
  "sort"
  "math"
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
func AroonUp(data []float64, l int) []float64 {
  var aroon []float64;
  for i := l; i <= len(data); i++ {
    pl := append([]float64(nil), data[i-l:i]...);
    hl := append([]float64(nil), data[i-l:i]...);
    sort.Float64s(hl);
    index := sort.SearchFloat64s(pl, hl[len(hl)-1]);
    aroon = append(aroon, (100 * (float64(l)-float64(index))/float64(l)));
  }
  return aroon;
}
func AroonDown(data []float64, l int) []float64 {
  var aroon []float64;
  for i := l; i <= len(data); i++ {
    pl := append([]float64(nil), data[i-l:i]...);
    hl := append([]float64(nil), data[i-l:i]...);
    sort.Float64s(hl);
    index := sort.SearchFloat64s(pl, hl[0]);
    aroon = append(aroon, (100 * (float64(l)-float64(index))/float64(l)));
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
  fmt.Println(mom);
  return mom;
}
func main() {
  Mom([]float64{1, 1.1, 1.2, 1.24, 1.34}, 4, false);
  Vwap([][]float64{{127.21, 89329}, {127.17, 16137}, {127.16, 23945}}, 2);
  Obv([][]float64{{25200, 10}, {30000, 10.15}, {25600, 10.17}, {32000, 10.13}});
  Kst([]float64{8, 6, 7, 6, 8, 9, 7, 5, 6, 7, 6, 8, 6, 7, 6, 8, 9, 9, 8, 6, 4, 6, 5, 6, 7, 8, 9}, 5, 7, 10, 15, 5, 5, 5, 7, 4);
  Cop([]float64{3, 4, 5, 3, 4, 5, 6, 4, 7, 5, 4, 7, 5}, 4, 6, 5);
  Roc([]float64{1, 2, 3, 4}, 3);
  Mfi([][]float64{{19, 13}, {14, 38}, {21, 25}, {32, 17}}, 3);
  AroonOsc([]float64{2, 5, 4, 5}, 3);
  AroonDown([]float64{2, 5, 4, 5}, 3);
  AroonUp([]float64{5, 4, 5, 2}, 3);
  Drawdown([]float64{1,2,3,4,2,3});
  Dif(0.75, 0.5);
  Median([]float64{4,6,3,1,2,5}, 4);
  Normalize([]float64{5,4,9,4}, 0.1);
  Denormalize([]float64{5,4,9,4}, []float64{0.2222222222222222, 0.06349206349206349, 0.8571428571428571, 0.06349206349206349, 0.4444444444444444}, 0.1);
  Mad([]float64{3, 7, 5, 4, 3, 8, 9}, 6);
  Aad([]float64{4, 6, 8, 6, 8, 9, 10, 11}, 7);
  Sma([]float64{1, 2, 3, 4, 5, 6, 10}, 6);
  Ssd([]float64{7, 6, 5, 7, 9, 8, 3, 5, 4}, 7);
  Rsi([]float64{1, 2, 3, 4, 5, 6, 7, 5}, 6);
  Wsma([]float64{1, 2, 3, 4, 5, 6, 10}, 6);
  Wrsi([]float64{1, 2, 3, 4, 5, 6, 7, 5, 6}, 6);
  Ema([]float64{1, 2, 3, 4, 5, 6, 10}, 6);
  Smma([]float64{1, 2, 3, 4, 5, 6, 10}, 5);
  Wma([]float64{69, 68, 66, 70, 68}, 4);
  Pwma([]float64{17, 26, 23, 29, 20}, 4);
  Hwma([]float64{54, 51, 86, 42, 47}, 4);
  Vwma([][]float64{{1, 59}, {1.1, 82}, {1.21, 27}, {1.42, 73}, {1.32, 42}}, 4);
  Lsma([]float64{5, 6, 6, 3, 4, 6, 7}, 6);
  Hull([]float64{6, 7, 5, 6, 7, 4, 5, 7}, 6);
  Macd([]float64{1, 2, 3, 4, 5, 6, 14}, 3, 6);
  Std([]float64{1,2,3}, 3);
  Normsinv(0.4732);
  Cor([]float64{1, 2, 3, 4, 5, 2}, []float64{1, 3, 2, 4, 6, 3});
}
