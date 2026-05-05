# ta.go

Technical-analysis indicators, oscillators, and statistics for Go.
Pure-Go, dependency-free, with first-class performance — see
[BENCHMARKS](bench/BENCHMARKS.md).

```bash
go get github.com/Bitvested/ta.go@latest
```

## Quickstart

ta.go is split into category subpackages. Import the ones you need:

```go
package main

import (
	"fmt"

	indicators "github.com/Bitvested/ta.go/src/indicators"
	ma         "github.com/Bitvested/ta.go/src/moving_averages"
)

func main() {
	prices := []float64{
		44.34, 44.09, 44.15, 43.61, 44.33, 44.83, 45.10, 45.42,
		45.84, 46.08, 45.89, 46.03, 45.61, 46.28, 46.28,
	}

	fmt.Println("SMA(14):", ma.Sma(prices, 14))
	fmt.Println("RSI(14):", indicators.Rsi(prices, 14)) // Wilder smoothing
}
```

## Packages

| Import path                                          | What's inside |
|------------------------------------------------------|---------------|
| `…/src/moving_averages` (`movingaverages`)           | Sma, Ema, Wma, Smma, Hull, Kama, Lsma, Cwma, Vwma, Vwwma, Wsma, Pwma, Hwma, Dema, Tema, T3, Trima, Vidya, Zlema |
| `…/src/indicators`                                   | Rsi, Wrsi, Macd, MacdSignal, MacdBars, AroonUp, AroonDown, Mfi, Roc, Cop, Kst, Obv, Vwap, Mom, Tsi, Bop, Fi, Asi, Alligator, Pr, Stoch, Ichimoku, Atr, Fib, Fractals, FractalsPrice, HalfTrend, ZigZag, Psar, Supertrend, Elderray, Hv, Rvi, Rvi\_signal, Rsi\_divergence, Divergence, Adl, Apo, Cci, Cmf, Cross, Dpo, Pdm, Mdm, Pdi, Mdi, Dx, Adx, Adxr, Emv, Kdj, Mass, Natr, Nvi, Pvi, Ppo, StochRsi, Trix, Ulcer, Vortex |
| `…/src/oscillators`                                  | Gator, MomOsc, ChaikinOsc, AroonOsc, Ao, Ac, Fisher, Kvo, Ult |
| `…/src/bands`                                        | Bands, Bandwidth, Keltner, Don, Fibbands, Envelope |
| `…/src/statistics`                                   | Median, Mad, Aad, Ssd, Std, StdSeries, Variance, Cor, Covariance, Dif, Drawdown, Normsinv, Ncdf, Sim, Percentile, Kmeans, Er, Ar, Winratio, Avgwin, Avgloss, Se, Kelly, Zscore, Normalize, NormalizePair, NormalizeFrom, Denormalize, Standardize, Log, Exp, Martingale, Antimartingale, Sum, Permutations, Mse, Cum, RecentHigh, RecentLow, Pvalue, ExpectedTrails, LrSlope, LrIntercept, LrAngle, Tsf, ReturnPositive, ReturnNegative |
| `…/src/chart_types` (`charttypes`)                   | Ha (Heikin-Ashi), Ren (Renko) |
| `…/src/experimental`                                 | Support, Resistance, Line, LineCalc, DivergenceState |
| `…/src/random`                                       | Prng, Pick, Range, Float, Order |
| `…/src/misc`                                         | Fibnumbers, TimesUp, TimesDown |
| `…/multi`                                            | Sim, SimSeeded — goroutine Monte Carlo |

All paths are prefixed with `github.com/Bitvested/ta.go`. The root
package only exports the `Version` constant — there is no monolithic
re-export.

## Bar conventions

Indicators that take OHLCV-style bars use a fixed column ordering:

| Indicator family                                                             | Bar layout         |
|------------------------------------------------------------------------------|--------------------|
| `Atr`, `Cci`, `Pdi`/`Mdi`/`Dx`/`Adx`/`Adxr`, `Pdm`/`Mdm`, `Vortex`, `Ult`, `Stoch`, `Kdj`, `Ichimoku`, `Halftrend`, `Psar`, `Supertrend`, `ZigZag`, `Don`, `Keltner` | `[H, C, L]` |
| `Adl`, `Cmf`, `ChaikinOsc`, `Kvo`                                            | `[H, C, L, V]`     |
| `Mfi`                                                                        | `[H, L, C, V]`     |
| `Emv`                                                                        | `[H, L, V]`        |
| `Nvi`, `Pvi`, `Vwma`, `Vwwma`, `Vwap`                                        | `[close, volume]`  |
| `Rvi`                                                                        | `[O, H, L, C]`     |

Restructure your input slices to match before calling.

## Upgrading

If you're moving from a pre-v2.0.0 release, the most likely breakage
points:

| Area | What changed |
|------|--------------|
| Imports | The root `ta` package no longer exports indicators — switch to the category subpackages above. |
| `indicators.Rsi` | Now uses Wilder smoothing (TA-Lib convention). The legacy simple-average behaviour is preserved in `Wrsi`. |
| `indicators.Roc` | Returns rate of change in **percent** (×100). `Cop` and `Kst` cascade. |
| `indicators.Mfi` | Input is `[][]float64` of `[H, L, C, V]` bars. |
| `indicators.Atr`, `oscillators.Fisher`, `movingaverages.Smma` | Output length is `N − l + 1`; TR uses absolute values. `Keltner`, `Supertrend`, `Halftrend` cascade. |
| `statistics.Median` | Even-length windows return the average of the two middle values. `Mad` cascades. |
| `indicators.Macd` | The base `Macd` returns just the MACD line. Signal and histogram split into `MacdSignal` and `MacdBars`. |
| `indicators.Cross` | Returns `[]CrossSignal{Index int; Cross bool}`. Inputs are no longer mutated. |
| NaN policy | `Rsi`, `Roc`, `Dif`, `Mom`, `MomOsc`, `Mfi`, `Fisher`, `Pr`, `Stoch` return `math.NaN()` on flat / divide-by-zero edges instead of silent coercion. |
| No mutation | `Std`, `Cross`, `Divergence`, `Percentile`, and `experimental.Support`/`Resistance` no longer mutate caller inputs. |

Full Phase-by-Phase breakdown in [`CHANGELOG.md`](CHANGELOG.md).

## Performance

ta.go ties or beats `markcheno/go-talib` and `cinar/indicator` at 100 k
bars on every benched indicator. Headline numbers (lower is better,
ns/op):

| Indicator | ta.go   | go-talib | cinar   |
|-----------|--------:|---------:|--------:|
| SMA       |  141 µs |   183 µs |  221 µs |
| EMA       |  301 µs |   299 µs |  502 µs |
| RSI       |  544 µs |   531 µs | 1912 µs |
| MACD      |  936 µs |  1108 µs | 1602 µs |
| Bollinger |  501 µs |   625 µs |  669 µs |
| ATR       |  513 µs |   974 µs |  762 µs |

Full table and reproduction instructions in
[`bench/BENCHMARKS.md`](bench/BENCHMARKS.md).

## Versioning

The user-facing version is **`2.0.0`**, exported from the root
package as `ta.Version`. The Git tag is **`v0.20.0`** — the `v0.x`
prefix lets the existing `github.com/Bitvested/ta.go` import path
keep working without Go's `/v2` module-path requirement. To pin the
2.0.0 line specifically:

```bash
go get github.com/Bitvested/ta.go@v0.20.0
```

```go
import ta "github.com/Bitvested/ta.go"
fmt.Println(ta.Version) // "2.0.0"
```

## Contributing

Pull requests are welcome. For major changes, please open an issue
first to discuss what you would like to change. Please make sure to
update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)

## 👨‍💻 Maintained By

**ta.go** is an open-source project created and actively maintained by [**Nino Kroesen**](https://ninokroesen.com).

This library was built with performance in mind for high-frequency data and algorithmic trading systems. If you are interested in seeing these indicators in action within a live quantitative trading environment, check out [**Bitvested**](https://bitvested.com).

---
