# Technical Analysis (ta.go)

ta.go is a go module for dealing with financial technical analysis.

## Installation

```bash
go get github.com/bitvested/ta.go
```

## Usage
```go
import(
  "github.com/bitvested/ta.go"
)
```

## Examples
- [Simple Moving Average](#sma)
- [Smoothed Moving Average](#smma)
- [Weighted Moving Average](#wma)
- [Wilder's Smoothing Moving Average](#wsma)
- [Parabolic Weighted Moving Average](#pwma)
- [Hyperbolic Weighted Moving Average](#hwma)
- [Hull Moving Average](#hull)
- [Kaufman Adaptive Moving Average](#kama)
- [Volume Weighted Moving Average](#vwma)
- [Exponential Moving Average](#ema)
- [Least Squares Moving Average](#lsma)
- [Moving Average Convergence / Divergence](#macd)
- [Relative Strength Index](#rsi)
- [Wilder's Relative Strength Index](#wrsi)
- [True Strength Index](#tsi)
- [Balance Of Power](#bop)
- [Force Index](#fi)
- [Accumulative Swing Index](#asi)
- [Awesome Oscillator](#ao)
- [Williams %R](#pr)
- [Stochastics](#stoch)
- [Variance](#variance)
- [Standard Deviation](#std)
- [Inverse Normal Distribution](#normsinv)
- [Monte Carlo Simulation](#sim)
- [Percentile](#perc)
- [Correlation](#cor)
- [Percentage Difference](#dif)
- [Drawdown](#drawdown)
- [Median](#median)
- [K-means Clustering](#kmeans)
- [Normalize](#norm)
- [Denormalize](#dnorm)
- [Median Absolute Deviation](#mad)
- [Average Absolute Deviation](#aad)
- [Sum Squared Differences](#ssd)
- [Bollinger Bands](#bands)
- [Bollinger Bandwidth](#bandwidth)
- [Keltner Channels](#kelt)
- [Donchian Channels](#don)
- [Ichimoku Cloud](#ichi)
- [Average True Range](#atr)
- [Aroon Up](#aroon-up)
- [Aroon Down](#aroon-down)
- [Aroon Oscillator](#aroon-osc)
- [Money Flow Index](#mfi)
- [Rate Of Change](#roc)
- [Coppock Curve](#cop)
- [Know Sure Thing](#kst)
- [On-Balance Volume](#obv)
- [Volume-Weighted Average Price](#vwap)
- [Chande Momentum Oscillator](#mom_osc)
- [Momentum](#mom)
- [Heikin Ashi](#ha)
- [Renko](#ren)
#### <a name="sma"></a>Simple Moving Average (SMA)
```go
data := []float64{1, 2, 3, 4, 5, 6, 10};
length := 6;
ta.Sma(data, length);
// output []float64
// [3.5, 5]
```
#### <a name="smma"></a>Smoothed Moving Average (SMMA)
```go
data := []float64{1, 2, 3, 4, 5, 6, 10};
length := 5;
ta.Smma(data, length);
// output []float64
// [3.4, 4.92]
```
#### <a name="wma"></a>Weighted Moving Average (WMA)
```go
data := []float64{69, 68, 66, 70, 68};
length := 4;
ta.Wma(data, length);
// output []float64
// [68.3, 68.2]
```
#### <a name="wsma"></a>Wilder's Smoothing Moving Average
```go
data := []float64{1, 2, 3, 4, 5, 6, 10};
length := 6;
ta.Wsma(data, length);
// output []float64
// [3.5, 4.58]
```
#### <a name="pwma"></a>Parabolic Weighted Moving Average
```go
data := []float64{17, 26, 23, 29, 20};
length := 4;
ta.Pwma(data, length);
// output []float64
// [24.09, 25.18]
```
#### <a name="hwma"></a>Hyperbolic Weighted Moving Average
```go
data := []float64{54, 51, 86, 42, 47};
length := 4;
ta.Hwma(data, length);
// output []float64
// [56.2, 55.0]
```
#### <a name="hull"></a>Hull Moving Average
```go
data := []float64{6, 7, 5, 6, 7, 4, 5, 7};
length := 6;
ta.Hull(data, length);
// output []float64
// [4.76, 5.48]
```
#### <a name="kama"></a>Kaufman Adaptive Moving Average (KAMA)
```go
data := []float64{8, 7, 8, 9, 7, 9};
length1 := 2;
length2 := 4;
length3 := 8;
ta.Kama(data, length1, length2, length3);
// output []float64
// [8, 8.64, 8.57, 8.57]
```
#### <a name="vwma"></a>Volume Weighted Moving Average (VWMA)
```go
data := [][]float64{{1, 59}, {1.1, 82}, {1.21, 27}, {1.42, 73}, {1.32, 42}}; // {price, volume (quantity)}
length := 4;
ta.Vwma(data, length);
// output []float64
// [1.185, 1.259]
```
#### <a name="ema"></a>Exponential Moving Average (EMA)
```go
data := []float64{1, 2, 3, 4, 5, 6, 10};
length := 6;
ta.Ema(data, length);
// output []float64
// [3.5, 5.357]
```
#### <a name="lsma"></a>Least Squares Moving Average (LSMA)
```go
data := []float64{5, 6, 6, 3, 4, 6, 7};
length := 6;
ta.Lsma(data, length);
// output []float64
// [4.714, 5.761]
```
#### <a name="macd"></a>Moving Average Convergence / Divergence (MACD)
```go
data := []float64{1, 2, 3, 4, 5, 6, 14};
length1 := 3;
length2 := 6;
ta.Macd(data, length1, length2);
// output []float64
// [1.5, 3]
```
#### <a name="rsi"></a>Relative Strength Index (RSI)
```go
data := []float64{1, 2, 3, 4, 5, 6, 7, 5};
length := 6;
ta.Rsi(data, length);
// output []float64
// [100.0, 100.0, 66.667]
```
#### <a name="wrsi"></a>Wilder's Relative Strength Index
```go
data := []float64{1, 2, 3, 4, 5, 6, 7, 5, 6};
length := 6;
ta.Wrsi(data, length);
// output []float64
// [100, 71.43, 75.61]
```
#### <a name="tsi"></a>True Strength Index (TSI)
```go
data := []float64{1.32, 1.27, 1.42, 1.47, 1.42, 1.45, 1.59};
longlength := 3;
shortlength := 2;
signallength := 2;
ta.Tsi(data, longlength, shortlength, signallength);
// output [][]float64
// [[0.327, 0.320], [0.579, 0.706]]
// [strength line, signal line]
```
#### <a name="bop"></a>Balance Of Power
```go
data := [][]float64{{4, 5, 4, 5}, {5, 6, 5, 6}, {6, 8, 5, 6}}; // [open, high, low, close]
length := 2;
ta.Bop(data, length);
// output []float64
// [1, 0.5]
```
#### <a name="fi"></a>Force Index
```go
data := [][]float64{{1.4, 200}, {1.5, 240}, {1.1, 300}, {1.2, 240}, {1.5, 400}}; // [close, volume]
length := 4;
ta.Fi(data, length);
// output []float64
// [0.0075]
```
#### <a name="asi"></a>Accumulative Swing Index
```go
data := [][]float64{{7, 6, 4}, {9, 7, 5}, {9, 8, 6}}; // [high, close, low]
ta.Asi(data);
// output []float64
// [0, -12.5]
```
#### <a name="ao"></a>Awesome Oscillator
```go
data := [][]float64{{6, 5}, {8, 6}, {7, 4}, {6, 5}, {7, 6}, {9, 8}}; // [high, low]
shortlength := 2;
longlength := 5;
ta.Ao(data, shortlength, longlength);
// output []float64
// [0, 0.9]
```
#### <a name="pr"></a>Williams %R
```go
data := []float64{2, 1, 3, 1, 2};
length := 4;
ta.Pr(data, length);
// output []float64
// [-0, -100, -50]
```
#### <a name="stoch"></a>Stochastics
```go
data := [][]float64{{3,2,1}, {2,2,1}, {4,3,1}, {2,2,1}}; // [high, close, low]
length := 2;
smoothd := 1;
smoothk := 1;
ta.Stoch(data, length, smoothd, smoothk);
// output []float64
// [[66.667, 66.667], [33.336, 33.336]]
// [kline, dline]
```
#### <a name="variance"></a>Variance
```go
data := []float64{6, 7, 2, 3, 5, 8, 6, 2};
length := 7;
ta.Variance(data, length);
// output []float64
// [3.918, 5.061]
```
#### <a name="std"></a>Standard Deviation
```go
data := []float64{1, 2, 3};
length := 3;
ta.Std(data, length);
// output float64
// 0.81649658092773
```
#### <a name="normsinv"></a>Inverse Normal Distribution
```go
data := 0.4732;
ta.Normsinv(data);
// output float64
// -0.06722824471054376
```
#### <a name="sim"></a>Monte Carlo Simulation
```go
data := []float64{6, 4, 7};
length := 2;
simulations := 2;
ta.Sim(data, length, simulations);
// output [][]float64
// [6, 4, 7, 8, 5, 6, 5.96, 5.7]
```
#### <a name="perc"></a>Percentile
```go
data := [][]float64{{6, 4, 7}, {5, 3, 6}, {7, 5, 8}};
percentile := 0.5;
ta.Percentile(data, percentile);
// output []float64
// [5, 3, 6]
```
#### <a name="cor"></a>Correlation
```go
data1 := []float64{1, 2, 3, 4, 5, 2};
data2 := []float64{1, 3, 2, 4, 6, 3};
ta.Cor(data1, data2);
// output float64
// 0.8808929232684737
```
#### <a name="dif"></a>Percentage Difference
```go
newval := 0.75;
oldval := 0.5;
ta.Dif(newval, oldval);
// output float64
// 0.5
```
#### <a name="drawdown"></a>Drawdown
```go
data := []float64{1, 2, 3, 4, 2, 3};
ta.Drawdown(data);
// output float64
// -0.5
```
#### <a name="median"></a>Median
```go
data := []float64{4, 6, 3, 1, 2, 5};
length := 4;
ta.Median(data, length);
// output []float64
// [3, 2, 2]
```
#### <a name="kmeans"></a>K-means Clustering
```go
data := []float64{2, 3, 4, 5, 3, 5, 7, 8, 6, 8, 6, 4, 2, 6};
length := 4;
ta.Kmeans(data, length);
// output [][]float64
// [[ 4, 5, 5, 4 ], [ 7, 6, 6, 6 ], [ 8, 8 ], [ 2, 3, 3, 2 ]]
```
#### <a name="norm"></a>Normalize
```go
data := []float64{5,4,9,4};
margin := 0.1;
ta.Normalize(data, margin);
// output []float64
// [0.22, 0.06, 0.86, 0.06]
```
#### <a name="dnorm"></a>Denormalize
```go
data := []float64{5,4,9,4}; // original data || [highest, lowest]
norm := float64{0.22, 0.06, 0.86, 0.06, 0.44};
margin := 0.1;
ta.Denormalize(data, norm, margin);
// output []float64
// [5 ,4, 9, 4, 6.4]
```
#### <a name="mad"></a>Median Absolute Deviation
```go
data := []float64{3, 7, 5, 4, 3, 8, 9};
length := 6;
ta.Mad(data, length);
// output []float64
// [1, 2]
```
#### <a name="aad"></a>Average Absolute Deviation
```go
data := []float64{4, 6, 8, 6, 8, 9, 10, 11};
length := 7;
ta.Aad(data, length);
// output []float64
// [1.673, 1.468]
```
#### <a name="ssd"></a>Sum Squared Differences
```go
data := []float64{7, 6, 5, 7, 9, 8, 3, 5, 4};
length := 7;
ta.Ssd(data, length);
// output []float64
// [4.87, 4.986, 5.372]
```
#### <a name="bands"></a>Bollinger Bands
```go
data := []float64{1, 2, 3, 4, 5, 6};
length := 5;
deviations := 2;
ta.Bands(data, length, deviations);
// output []float64
// [[5.828, 3, 0.172], [6.828, 4, 1.172]]
// [upper band, middle band, lower band]
```
#### <a name="bandwidth"></a>Bollinger Bandwidth
```go
data := []float64{1, 2, 3, 4, 5, 6};
length := 5;
deviations := 2;
ta.Bandwidth(data, length, deviations);
// output []float64
// [1.886, 1.344]
```
#### <a name="kelt"></a>Keltner Channels
```go
data := [][]float64{{3,2,1}, {2,2,1}, {4,3,1}, {2,2,1}, {3,3,1}}; // [high, close, low]
length := 5;
deviations := 1;
ta.Keltner(data, length, deviations);
// output [][]float64
// [[3.93, 2.06, 0.20]]
// [upper band, middle band, lower band]
```
#### <a name="don"></a>Donchian Channels
```go
data := [][]float64{{6, 2}, {5, 2}, {5, 3}, {6, 3}, {7, 4}, {6, 3}}; // [high, low]
length := 5;
ta.Don(data, length);
// output []float64
// [[7, 4.5, 2], [7, 4.5, 2]]
// [upper band, base line, lower band]
```
#### <a name="ichi"></a>Ichimoku Cloud
```go
data := [][]float64{{6, 3, 2}, {5, 4, 2}, {5, 4, 3}, {6, 4, 3}, {7, 6, 4}, {6, 5, 3}}; // [high, close, low]
length1 := 9;
length2 := 26;
length3 := 52;
displacement := 26;
ta.Ichimoku(data, length1, length2, length3, displacement);
// output []float64
// [conversion line, base line, leading span A, leading span B, lagging span]
```
#### <a name="atr"></a>Average True Range (ATR)
```go
data := [][]float64{{3,2,1}, {2,2,1}, {4,3,1}, {2,2,1}}; // [high, close, low]
length := 3;
ta.Atr(data, length);
// output []float64
// [2, 1.667, 2.111, 1.741]
```
#### <a name="aroon-up"></a>Aroon Up
```go
data := []float64{5, 4, 5, 2};
length := 3;
ta.AroonUp(data, length);
// output []float64
// [100, 50]
```
#### <a name="aroon-down"></a>Aroon Down
```go
data := []float64{2, 5, 4, 5};
length := 3;
ta.AroonDown(data, length);
// output []float64
// [0, 50]
```
#### <a name="aroon-osc"></a>Aroon Oscillator
```go
data := []float64{2, 5, 4, 5};
length := 3;
ta.AroonOsc(data, length);
// output []float64
// [50, 50]
```
#### <a name="mfi"></a>Money Flow Index
```go
data := [][]float64{{19, 13}, {14, 38}, {21, 25}, {32, 17}}; // [buy volume, sell volume]
length := 3;
ta.Mfi(data, length);
// output []float64
// [41.54, 45.58]
```
#### <a name="roc"></a>Rate Of Change
```go
data := []float64{1, 2, 3, 4};
length := 3;
ta.Roc(data, length);
// output []float64
// [2, 1]
```
#### <a name="cop"></a>Coppock Curve
```go
data := []float64{3, 4, 5, 3, 4, 5, 6, 4, 7, 5, 4, 7, 5};
length1 := 4;
length2 := 6;
length3 := 5;
ta.Cop(data, length1, length2, length3);
// output []float64
// [0.376, 0.237]
```
#### <a name="kst"></a>Know Sure Thing
```go
data := []float64{8, 6, 7, 6, 8, 9, 7, 5, 6, 7, 6, 8, 6, 7, 6, 8, 9, 9, 8, 6, 4, 6, 5, 6, 7, 8, 9};
// roc sma #1
r1 := 5;
s1 := 5;
// roc sma #2
r2 := 7;
s2 := 5;
// roc sma #3
r3 := 10;
s3 := 5;
// roc sma #4
r4 := 15;
s4 := 7;
// signal line
sig := 4;
ta.Kst(data, r1, r2, r3, r4, s1, s2, s3, s4, sig);
// output []float64
// [[-0.68, -0.52], [-0.29, -0.58], [0.35, -0.36]]
// [kst line, signal line]
```
#### <a name="obv"></a>On-Balance Volume
```go
data := [][]float64{{25200, 10}, {30000, 10.15}, {25600, 10.17}, {32000, 10.13}}; // [asset volume, close price]
ta.Obv(data);
// output []float64
// [0, 30000, 55600, 23600]
```
#### <a name="vwap"></a>Volume-Weighted Average Price
```go
data := [][]float64{{127.21, 89329}, {127.17, 16137}, {127.16, 23945}}; // [average price, volume (quantity)]
length := 2;
ta.Vwap(data, length);
// output []float64
// [127.204, 127.164]
```
#### <a name="mom_osc"></a>Chande Momentum Oscillator
```go
data := []float64{1, 1.2, 1.3, 1.3, 1.2, 1.4};
length := 4;
ta.MomOsc(data, length);
// output []float64
// [0, 3.85]
```
#### <a name="mom"></a>Momentum
```go
data := []float64{1, 1.1, 1.2, 1.24, 1.34};
length := 4;
percentage := false;
ta.Mom(data, length, percentage);
// output []float64
// [0.24, 0.24]
```
#### <a name="ha"></a>Heikin Ashi
```go
data := [][]float64{{3, 4, 2, 3}, {3, 6, 3, 5}, {5, 5, 2, 3}}; // [open, high, low, close]
ta.Ha(data);
// output []float64
// [open, high, low, close]
// first 7-10 candles are unreliable
```
#### <a name="ren"></a>Renko
```go
data := [][]float64{{8, 6}, {9, 7}, {9, 8}}; // [high, low]
bricksize := 3;
ta.Ren(data, bricksize);
// output []float64
// [open, high, low, close]
```
## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
