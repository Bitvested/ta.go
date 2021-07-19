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
func main() {
  Median([]float64{4,6,3,1,2,5}, 4);
  Normalize([]float64{5,4,9,4}, 0.1);
  Denormalize([]float64{5,4,9,4}, []float64{0.2222222222222222, 0.06349206349206349, 0.8571428571428571, 0.06349206349206349, 0.4444444444444444}, 0.1);
  Mad([]float64{3, 7, 5, 4, 3, 8, 9}, 6);
  Aad([]float64{4, 6, 8, 6, 8, 9, 10, 11}, 7);
  Sma([]float64{1, 2, 3, 4, 5, 6, 10}, 6);
}
