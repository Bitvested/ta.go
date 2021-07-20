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
    fmt.Println(i)
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
  fmt.Println(wmaa);
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
func main() {
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
}
