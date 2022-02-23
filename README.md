# Technical Analysis (ta.go)

ta.go is a go module for dealing with financial technical analysis.

## Installation

```bash
go install github.com/Bitvested/ta.go@latest
```

## Usage
```go
import(
  ta "github.com/Bitvested/ta.go"
)
```

## Examples
#### Moving Averages
- [Simple Moving Average](#sma)
- [Smoothed Moving Average](#smma)
- [Weighted Moving Average](#wma)
- [Exponential Moving Average](#ema)
- [Hull Moving Average](#hull)
- [Least Squares Moving Average](#lsma)
- [Volume Weighted Moving Average](#vwma)
- [Wilder's Smoothing Moving Average](#wsma)
- [Parabolic Weighted Moving Average](#pwma)
- [Hyperbolic Weighted Moving Average](#hwma)
- [Kaufman Adaptive Moving Average](#kama)
#### Indicators
- [Moving Average Convergence / Divergence](#macd)
- [Relative Strength Index](#rsi)
- [Wilder's Relative Strength Index](#wrsi)
- [True Strength Index](#tsi)
- [Balance Of Power](#bop)
- [Force Index](#fi)
- [Accumulative Swing Index](#asi)
- [Alligator Indicator](#alli)
- [Williams %R](#pr)
- [Stochastics](#stoch)
- [Fibonacci Retracement](#fib)
- [Bollinger Bandwidth](#bandwidth)
- [Ichimoku Cloud](#ichi)
- [Average True Range](#atr)
- [Aroon Up](#aroon-up)
- [Aroon Down](#aroon-down)
- [Money Flow Index](#mfi)
- [Rate Of Change](#roc)
- [Coppock Curve](#cop)
- [Know Sure Thing](#kst)
- [On-Balance Volume](#obv)
- [Volume-Weighted Average Price](#vwap)
- [Fractals](#fractals)
- [Momentum](#mom)
- [HalfTrend](#half)
#### Oscillators
- [Alligator Oscillator](#gator)
- [Chande Momentum Oscillator](#mom_osc)
- [Chaikin Oscillator](#chaikin_osc)
- [Aroon Oscillator](#aroon-osc)
- [Awesome Oscillator](#ao)
- [Accelerator Oscillator](#ac)
- [Fisher Transform](#fish)
#### Bands
- [Bollinger Bands](#bands)
- [Keltner Channels](#kelt)
- [Donchian Channels](#don)
- [Envelope](#env)
#### Statistics
- [Standard Deviation](#std)
- [Variance](#variance)
- [Inverse Normal Distribution](#normsinv)
- [Monte Carlo Simulation](#sim)
- [Percentile](#perc)
- [Correlation](#cor)
- [Percentage Difference](#dif)
- [Expected Return](#er)
- [Abnormal Return](#ar)
- [Kelly Criterion](#kelly)
- [Winratio](#winratio)
- [Average Win](#avgwin)
- [Average Loss](#avgloss)
- [Drawdown](#drawdown)
- [Median](#median)
- [Recent High](#rh)
- [Recent Low](#rl)
- [Median Absolute Deviation](#mad)
- [Average Absolute Deviation](#aad)
- [Standard Error](#stderr)
- [Sum Squared Differences](#ssd)
- [Logarithm](#log)
- [Exponent](#exp)
- [Normalize](#norm)
- [Denormalize](#dnorm)
- [Normalize Pair](#normp)
- [Normalize From](#normf)
- [Standardize](#standard)
- [Z-Score](#zscore)
- [K-means Clustering](#kmeans)
#### Chart Types
- [Heikin Ashi](#ha)
- [Renko](#ren)
#### Experimental
- [Support Line](#sup)
- [Resistance Line](#res)
#### <a name="sma"></a>Simple Moving Average (SMA)
```go
data := []float64{1, 2, 3, 4, 5, 6, 10};
length := 6;
ta.Sma(data, length);
// output []float64
// {3.5, 5}
```
#### <a name="smma"></a>Smoothed Moving Average (SMMA)
```go
data := []float64{1, 2, 3, 4, 5, 6, 10};
length := 5;
ta.Smma(data, length);
// output []float64
// {3.4, 4.92}
```
#### <a name="wma"></a>Weighted Moving Average (WMA)
```go
data := []float64{69, 68, 66, 70, 68};
length := 4;
ta.Wma(data, length);
// output []float64
// {68.3, 68.2}
```
#### <a name="ema"></a>Exponential Moving Average (EMA)
```go
data := []float64{1, 2, 3, 4, 5, 6, 10};
length := 6;
ta.Ema(data, length);
// output []float64
// {3.5, 5.357}
```
#### <a name="hull"></a>Hull Moving Average
```go
data := []float64{6, 7, 5, 6, 7, 4, 5, 7};
length := 6;
ta.Hull(data, length);
// output []float64
// {4.76, 5.48}
```
#### <a name="lsma"></a>Least Squares Moving Average (LSMA)
```go
data := []float64{5, 6, 6, 3, 4, 6, 7};
length := 6;
ta.Lsma(data, length);
// output []float64
// {4.714, 5.761}
```
#### <a name="vwma"></a>Volume Weighted Moving Average (VWMA)
```go
data := [][]float64{{1, 59}, {1.1, 82}, {1.21, 27}, {1.42, 73}, {1.32, 42}}; // {price, volume (quantity)}
length := 4;
ta.Vwma(data, length);
// output []float64
// {1.185, 1.259}
```
#### <a name="wsma"></a>Wilder's Smoothing Moving Average
```go
data := []float64{1, 2, 3, 4, 5, 6, 10};
length := 6;
ta.Wsma(data, length);
// output []float64
// {3.5, 4.58}
```
#### <a name="pwma"></a>Parabolic Weighted Moving Average
```go
data := []float64{17, 26, 23, 29, 20};
length := 4;
ta.Pwma(data, length);
// output []float64
// {24.09, 25.18}
```
#### <a name="hwma"></a>Hyperbolic Weighted Moving Average
```go
data := []float64{54, 51, 86, 42, 47};
length := 4;
ta.Hwma(data, length);
// output []float64
// {56.2, 55.0}
```
#### <a name="kama"></a>Kaufman Adaptive Moving Average (KAMA)
```go
data := []float64{8, 7, 8, 9, 7, 9};
length1 := 2;
length2 := 4;
length3 := 8;
ta.Kama(data, length1, length2, length3);
// output []float64
// {8, 8.64, 8.57, 8.57}
```
#### <a name="macd"></a>Moving Average Convergence / Divergence (MACD)
```go
data := []float64{1, 2, 3, 4, 5, 6, 14};
length1 := 3;
length2 := 6;
ta.Macd(data, length1, length2);
// output []float64
// {1.5, 3}
```
#### <a name="rsi"></a>Relative Strength Index (RSI)
```go
data := []float64{1, 2, 3, 4, 5, 6, 7, 5};
length := 6;
ta.Rsi(data, length);
// output []float64
// {100.0, 100.0, 66.667}
```
#### <a name="wrsi"></a>Wilder's Relative Strength Index
```go
data := []float64{1, 2, 3, 4, 5, 6, 7, 5, 6};
length := 6;
ta.Wrsi(data, length);
// output []float64
// {100, 71.43, 75.61}
```
#### <a name="tsi"></a>True Strength Index (TSI)
```go
data := []float64{1.32, 1.27, 1.42, 1.47, 1.42, 1.45, 1.59};
longlength := 3;
shortlength := 2;
signallength := 2;
ta.Tsi(data, longlength, shortlength, signallength);
// output [][]float64
// {{0.327, 0.320}, {0.579, 0.706}}
// {strength line, signal line}
```
#### <a name="bop"></a>Balance Of Power
```go
data := [][]float64{{4, 5, 4, 5}, {5, 6, 5, 6}, {6, 8, 5, 6}}; // {open, high, low, close}
length := 2;
ta.Bop(data, length);
// output []float64
// {1, 0.5}
```
#### <a name="fi"></a>Force Index
```go
data := [][]float64{{1.4, 200}, {1.5, 240}, {1.1, 300}, {1.2, 240}, {1.5, 400}}; // {close, volume}
length := 4;
ta.Fi(data, length);
// output []float64
// {0.0075}
```
#### <a name="asi"></a>Accumulative Swing Index
```go
data := [][]float64{{7, 6, 4}, {9, 7, 5}, {9, 8, 6}}; // {high, close, low}
ta.Asi(data);
// output []float64
// {0, -12.5}
```
#### <a name="alli"></a>Alligator Indicator
```go
data := []float64{8,7,8,9,7,8,9,6,7,8,6,8,10,8,7,9,8,7,9,6,7,9};
jawlength := 13;
teethlength := 8;
lipslength := 5;
jawshift := 8;
teethshift := 5;
lipshift := 3;
ta.Alligator(data, jawlength, teethlength, liplength, jawshift, teethshift, lipshift);
// output [][]float64
// {{jaw, teeth, lips}}
```
#### <a name="pr"></a>Williams %R
```go
data := []float64{2, 1, 3, 1, 2};
length := 4;
ta.Pr(data, length);
// output []float64
// {-0, -100, -50}
```
#### <a name="stoch"></a>Stochastics
```go
data := [][]float64{{3,2,1}, {2,2,1}, {4,3,1}, {2,2,1}}; // {high, close, low}
length := 2;
smoothd := 1;
smoothk := 1;
ta.Stoch(data, length, smoothd, smoothk);
// output []float64
// {{66.667, 66.667}, {33.336, 33.336}}
// {{kline, dline}}
```
#### <a name="fib"></a>Fibonacci Retracement
```go
start := 1;
end := 2;
ta.Fib(start, end);
// output []float64
// {1, 1.236, 1.382, 1.5, 1.618, 1.786, 2, 2.618, 3.618, 4.618, 5.236}
```
#### <a name="bandwidth"></a>Bollinger Bandwidth
```go
data := []float64{1, 2, 3, 4, 5, 6};
length := 5;
deviations := 2;
ta.Bandwidth(data, length, deviations);
// output []float64
// {1.886, 1.344}
```
#### <a name="ichi"></a>Ichimoku Cloud
```go
data := [][]float64{{6, 3, 2}, {5, 4, 2}, {5, 4, 3}, {6, 4, 3}, {7, 6, 4}, {6, 5, 3}}; // {high, close, low}
length1 := 9;
length2 := 26;
length3 := 52;
displacement := 26;
ta.Ichimoku(data, length1, length2, length3, displacement);
// output []float64
// {conversion line, base line, leading span A, leading span B, lagging span}
```
#### <a name="atr"></a>Average True Range (ATR)
```go
data := [][]float64{{3,2,1}, {2,2,1}, {4,3,1}, {2,2,1}}; // {high, close, low}
length := 3;
ta.Atr(data, length);
// output []float64
// {2, 1.667, 2.111, 1.741}
```
#### <a name="aroon-up"></a>Aroon Up
```go
data := []float64{5, 4, 5, 2};
length := 3;
ta.AroonUp(data, length);
// output []float64
// {100, 50}
```
#### <a name="aroon-down"></a>Aroon Down
```go
data := []float64{2, 5, 4, 5};
length := 3;
ta.AroonDown(data, length);
// output []float64
// {0, 50}
```
#### <a name="mfi"></a>Money Flow Index
```go
data := [][]float64{{19, 13}, {14, 38}, {21, 25}, {32, 17}}; // {buy volume, sell volume}
length := 3;
ta.Mfi(data, length);
// output []float64
// {41.54, 45.58}
```
#### <a name="roc"></a>Rate Of Change
```go
data := []float64{1, 2, 3, 4};
length := 3;
ta.Roc(data, length);
// output []float64
// {2, 1}
```
#### <a name="cop"></a>Coppock Curve
```go
data := []float64{3, 4, 5, 3, 4, 5, 6, 4, 7, 5, 4, 7, 5};
length1 := 4;
length2 := 6;
length3 := 5;
ta.Cop(data, length1, length2, length3);
// output []float64
// {0.376, 0.237}
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
// {{-0.68, -0.52}, {-0.29, -0.58}, {0.35, -0.36}}
// {kst line, signal line}
```
#### <a name="obv"></a>On-Balance Volume
```go
data := [][]float64{{25200, 10}, {30000, 10.15}, {25600, 10.17}, {32000, 10.13}}; // {asset volume, close price}
ta.Obv(data);
// output []float64
// {0, 30000, 55600, 23600}
```
#### <a name="vwap"></a>Volume-Weighted Average Price
```go
data := [][]float64{{127.21, 89329}, {127.17, 16137}, {127.16, 23945}}; // {average price, volume (quantity)}
length := 2;
ta.Vwap(data, length);
// output []float64
// {127.204, 127.164}
```
#### <a name="fractals"></a>Fractals
```go
data := [][]float64{{7,6},{8,6},{9,6},{8,5},{7,4},{6,3},{7,4},{8,5}};
ta.Fractals(data);
// output [][]float64  (same length as input)
// {{false, false},{false,false},{true,false},{false,false},{false,false},{false,true},{false,false},{false,false}}
// {upper fractal, lower fractal}
```
#### <a name="mom"></a>Momentum
```go
data := []float64{1, 1.1, 1.2, 1.24, 1.34};
length := 4;
percentage := false;
ta.Mom(data, length, percentage);
// output []float64
// {0.24, 0.24}
```
#### <a id="half"></a>HalfTrend
```go
// experimental (untested) function (may change in the future), ported from:
// https://www.tradingview.com/script/U1SJ8ubc-HalfTrend/
// data = [high, close, low]
data := [][]float64{{100,97,90},{101,98,94},{103,96,92},{106,100,95},{110,101,100},{112,110,105},{110,100,90},{103,100,97},{95,90,85},{94,80,80},{90,82,81},{85,80,70}};
atrlen := 6;
amplitude := 3;
deviation := 2;
ta.HalfTrend(data, atrlen, amplitude, deviation);
// output [][]interface{}
// {high, halftrend, low, signal}
```
### Oscillators
#### <a id="gator"></a>Alligator Oscillator
```go
data := []float64{8,7,8,9,7,8,9,6,7,8,6,8,10,8,7,9,8,7,9,6,7,9}
jawlength := 13;
teethlength := 8;
liplength := 5;
jawshift := 8;
teethshift := 5;
lipshift := 3;
ta.Gator(data, jawlength, teethlength, liplength, jawshift, teethshift, lipshift);
// output [][]float64
// {{upper gator, lower gator}}
```
#### <a name="mom_osc"></a>Chande Momentum Oscillator
```go
data := []float64{1, 1.2, 1.3, 1.3, 1.2, 1.4};
length := 4;
ta.MomOsc(data, length);
// output []float64
// {0, 3.85}
```
#### <a name="chaikin_osc"></a>Chaikin Oscillator
```go
data := [][]float64{{2,3,4,6},{5,5,5,4},{5,4,3,7},{4,3,3,4},{6,5,4,6},{7,4,3,6}}
length1 := 2;
length2 := 4;
ta.ChaikinOsc(data, length1, length2);
// output []float64
// {-1.667, -0.289, -0.736}
```
#### <a name="aroon-osc"></a>Aroon Oscillator
```go
data := []float64{2, 5, 4, 5};
length := 3;
ta.AroonOsc(data, length);
// output []float64
// {50, 50}
```
#### <a name="ao"></a>Awesome Oscillator
```go
data := [][]float64{{6, 5}, {8, 6}, {7, 4}, {6, 5}, {7, 6}, {9, 8}}; // {high, low}
shortlength := 2;
longlength := 5;
ta.Ao(data, shortlength, longlength);
// output []float64
// {0, 0.9}
```
#### <a name="ac"></a>Accelerator Oscillator
```go
data := [][]float64{{6, 5}, {8, 6}, {7, 4}, {6, 5}, {7, 6}, {9, 8}};
shortlength := 2;
longlength := 4;
ta.Ac(data, shortlength, longlength);
// output []float64
// {-5.875, -6.125, -6.5}
```
#### <a name="fish"></a>Fisher Transform
```go
data := []float64{8,6,8,9,7,8,9,8,7,8,6,7};
length := 9;
ta.Fisher(data, length);
// output [][]float64
// [[-0.318, -0.11], [-0.449, -0.318], [-0.616, -0.449]] // [fisher, trigger]
```
### Bands
#### <a name="bands"></a>Bollinger Bands
```go
data := []float64{1, 2, 3, 4, 5, 6};
length := 5;
deviations := 2;
ta.Bands(data, length, deviations);
// output []float64
// {{5.828, 3, 0.172}, {6.828, 4, 1.172}}
// {upper band, middle band, lower band}
```
#### <a name="kelt"></a>Keltner Channels
```go
data := [][]float64{{3,2,1}, {2,2,1}, {4,3,1}, {2,2,1}, {3,3,1}}; // {high, close, low}
length := 5;
deviations := 1;
ta.Keltner(data, length, deviations);
// output [][]float64
// {{3.93, 2.06, 0.20}}
// {upper band, middle band, lower band}
```
#### <a name="don"></a>Donchian Channels
```go
data := [][]float64{{6, 2}, {5, 2}, {5, 3}, {6, 3}, {7, 4}, {6, 3}}; // {high, low}
length := 5;
ta.Don(data, length);
// output []float64
// {{7, 4.5, 2}, {7, 4.5, 2}}
// {upper band, base line, lower band}
```
#### <a name="env"></a>Envelope
```go
data := []float64{6,7,8,7,6,7,8,7,8,7,8,7,8};
length := 11;
percentage := 0.05;
ta.Envelope(data, length, percentage);
// output [][]float64
// {{7.541, 7.182, 6.823}, {7.636, 7.273, 6.909}}
// {upper band, base line, lower band}
```
### Statistics
#### <a name="std"></a>Standard Deviation
```go
data := []float64{1, 2, 3};
length := 3;
ta.Std(data, length);
// output float64
// 0.81649658092773
```
#### <a name="variance"></a>Variance
```go
data := []float64{6, 7, 2, 3, 5, 8, 6, 2};
length := 7;
ta.Variance(data, length);
// output []float64
// {3.918, 5.061}
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
// {5, 3, 6}
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
#### <a name="er"></a>Expected Return
```go
data := []float64{0.02, -0.01, 0.03, 0.05, -0.03}; // historical return data
ta.Er(data);
// output float64
// 0.0119
```
#### <a name="ar"></a>Abnormal Return
```go
data := []float64{0.02, -0.01, 0.03, 0.05, -0.03}; // historical return data
length := 3;
ta.Ar(data, length);
// output float64
// [0.037, -0.053]
```
#### <a name="kelly"></a>Kelly Criterion
```go
data := []float64{0.01, 0.02, -0.01, -0.03, -0.015, 0.045, 0.005};
ta.Kelly(data);
// output float64
// 0.1443
```
#### <a name="winratio"></a>Winratio
```go
data := []float64{0.01, 0.02, -0.01, -0.03, -0.015, 0.005};
ta.Winratio(data);
// output float64
// 0.5
```
#### <a name="avgwin"></a>Average Win
```go
data := []float64{0.01, 0.02, -0.01, -0.03, -0.015, 0.005};
ta.Avgwin(data);
// output float64
// 0.012
```
#### <a name="avgloss"></a>Average Loss
```go
data := []float64{0.01, 0.02, -0.01, -0.03, -0.015, 0.005};
ta.Avgloss(data);
// output float64
// -0.018
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
// {3, 2, 2}
```
#### <a name="rh"></a>Recent High
```go
data := []float64{4,5,6,7,8,9,8,7,8,9,10,3,2,1};
lookback := 3;
ta.RecentHigh(data, lookback);
// output (struct)
// {Index: int, Value: float64}
```
#### <a name="rh"></a>Recent High
```go
data := []float64{1,4,5,6,4,3,2,3,4,3,5,7,8,8,5};
lookback := 3;
ta.RecentLow(data, lookback);
// output (struct)
// {Index: int, Value: float64}
```
#### <a name="mad"></a>Median Absolute Deviation
```go
data := []float64{3, 7, 5, 4, 3, 8, 9};
length := 6;
ta.Mad(data, length);
// output []float64
// {1, 2}
```
#### <a name="aad"></a>Average Absolute Deviation
```go
data := []float64{4, 6, 8, 6, 8, 9, 10, 11};
length := 7;
ta.Aad(data, length);
// output []float64
// {1.673, 1.468}
```
#### <a name="stderr"></a>Standard Error
```go
data := []float64{34, 54, 45, 43, 57, 38, 49};
size := 10;
ta.Se(data, size);
// output float64
// 2.424
```
#### <a name="ssd"></a>Sum Squared Differences
```go
data := []float64{7, 6, 5, 7, 9, 8, 3, 5, 4};
length := 7;
ta.Ssd(data, length);
// output []float64
// {4.87, 4.986, 5.372}
```
#### <a id="log"></a>Logarithm
```go
data := []float64{5, 14, 18, 28, 68, 103};
ta.Log(data);
// output (array)
// [1.61, 2.64, 2.89, 3.33, 4.22, 4.63]
```
#### <a id="exp"></a>Exponent
```go
data := []float64{1.6, 2.63, 2.89, 3.33, 4.22, 4.63};
ta.Exp(data);
// output (array)
// [4.95, 13.87, 17.99, 27.94, 68.03, 102.51]
```
#### <a name="norm"></a>Normalize
```go
data := []float64{5,4,9,4};
margin := 0.1;
ta.Normalize(data, margin);
// output []float64
// {0.22, 0.06, 0.86, 0.06}
```
#### <a name="dnorm"></a>Denormalize
```go
data := []float64{5,4,9,4}; // original data || {highest, lowest}
norm := float64{0.22, 0.06, 0.86, 0.06, 0.44};
margin := 0.1;
ta.Denormalize(data, norm, margin);
// output []float64
// {5 ,4, 9, 4, 6.4}
```
#### <a name="normp"></a>Normalize Pair
```go
pair1 := []float64{10,12,11,13};
pair2 := []float64{100,130,100,140};
ta.NormalizePair(pair1, pair2);
// output [][]float64
// [[55, 55], [66, 71.5], [60.5, 54.99], [71.5, 76.99]]
```
#### <a name="normf"></a>Normalize From
```go
data := []float64{8,12,10,11};
baseline := 100;
// output []float64
// [100, 150, 125, 137.5]
```
#### <a name="standard"></a>Standardize
```go
data := []float64{6,4,6,8,6};
ta.Standardize(data);
// output []float64
// {0, -1.581, 0, 1.581, 0}
```
#### <a name="zscore"></a>Z-Score
```go
data := []float64{34,54,45,43,57,38,49};
length := 5;
ta.Zscore(data, length);
// output []float64
// [1.266, -1.331, 0.408]
```
#### <a name="kmeans"></a>K-means Clustering
```go
data := []float64{2, 3, 4, 5, 3, 5, 7, 8, 6, 8, 6, 4, 2, 6};
length := 4;
ta.Kmeans(data, length);
// output [][]float64
// {{ 4, 5, 5, 4 }, { 7, 6, 6, 6 }, { 8, 8 }, { 2, 3, 3, 2 }}
```
### Chart Types
#### <a name="ha"></a>Heikin Ashi
```go
data := [][]float64{{3, 4, 2, 3}, {3, 6, 3, 5}, {5, 5, 2, 3}}; // {open, high, low, close}
ta.Ha(data);
// output []float64
// {open, high, low, close}
// first 7-10 candles are unreliable
```
#### <a name="ren"></a>Renko
```go
data := [][]float64{{8, 6}, {9, 7}, {9, 8}}; // {high, low}
bricksize := 3;
ta.Ren(data, bricksize);
// output []float64
// {open, high, low, close}
```
### Experimental Functions
#### <a id="sup"></a>Support Line
```go
data := []float64{4,3,2,5,7,6,5,4,7,8,5,4,6,7,5};
start := ta.RecentLow(data, 20);
support := ta.Support(data, start);
// output (struct) {Slope, Value, Index}
// support.Slope = delta y per x
// support.Lowest = lowest (start) value at x = 0
// support.Index = (start) index of the lowest value
// to get the line at the current candle / chart period
current := ta.LineCalc(len(data)-support.Index, support);
```
#### <a id="res"></a>Resistance Line
```go
data := []float64{5,7,5,5,4,6,5,4,6,5,4,3,2,4,3,2,1};
start := ta.RecentHigh(data, 20);
resistance := ta.Resistance(data, start);
// output (struct) {Slope, Value, Index}
// resistance.Slope = delta y per x
// resistance.Highest = highest (start) value
// resistance.Index = (start) index of highest value
// to get the line at the current candle / chart period
current := ta.LineCalc(len(data)-resistance.Index, resistance);
```
## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
