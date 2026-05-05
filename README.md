# ta.go

A technical-analysis library for Go, mirroring the API and semantics of
the sister [ta.js](https://github.com/Bitvested/ta.js) v2.0.0 release.

```bash
go get github.com/Bitvested/ta.go@latest
```

> **Versioning.** The user-facing version is `2.0.0` (matching ta.js); the
> Go module is tagged in the `v0.x` series so the existing import path
> `github.com/Bitvested/ta.go` keeps working without Go's `/v2`
> module-path requirement. The release tag for the `2.0.0` line is
> [`v0.20.0`](#release-tag); `Version` is exported as a constant from
> the root package.

## Quickstart

ta.go is split into category subpackages — there is no monolithic root
package, so import the categories you need:

```go
package main

import (
	"fmt"

	indicators "github.com/Bitvested/ta.go/src/indicators"
	ma         "github.com/Bitvested/ta.go/src/moving_averages"
)

func main() {
	prices := []float64{44.34, 44.09, 44.15, 43.61, 44.33, 44.83,
		45.10, 45.42, 45.84, 46.08, 45.89, 46.03, 45.61, 46.28, 46.28}

	fmt.Println("SMA(14):", ma.Sma(prices, 14))
	fmt.Println("RSI(14):", indicators.Rsi(prices, 14)) // Wilder smoothing
}
```

## Packages

All implementations live under `src/<category>/`. Each subpackage is
independently importable:

| Import path                                                  | Contents |
|--------------------------------------------------------------|----------|
| `github.com/Bitvested/ta.go/src/moving_averages` (`movingaverages`) | Sma, Ema, Wma, Smma, Hull, Kama, Lsma, Cwma, Vwma, Vwwma, Wsma, Pwma, Hwma, Dema, Tema, T3, Trima, Vidya, Zlema |
| `github.com/Bitvested/ta.go/src/indicators`                  | Rsi, Wrsi, Macd, MacdSignal, MacdBars, AroonUp, AroonDown, Mfi, Roc, Cop, Kst, Obv, Vwap, Mom, Tsi, Bop, Fi, Asi, Alligator, Pr, Stoch, Ichimoku, Atr, Fib, Fractals, FractalsPrice, HalfTrend, ZigZag, Psar, Supertrend, Elderray, Hv, Rvi, Rvi_signal, Rsi_divergence, Divergence, Adl, Apo, Cci, Cmf, Cross, Dpo, Pdm, Mdm, Pdi, Mdi, Dx, Adx, Adxr, Emv, Kdj, Mass, Natr, Nvi, Pvi, Ppo, StochRsi, Trix, Ulcer, Vortex |
| `github.com/Bitvested/ta.go/src/oscillators`                 | Gator, MomOsc, ChaikinOsc, AroonOsc, Ao, Ac, Fisher, Kvo, Ult |
| `github.com/Bitvested/ta.go/src/bands`                       | Bands, Bandwidth, Keltner, Don, Fibbands, Envelope |
| `github.com/Bitvested/ta.go/src/statistics`                  | Median, Mad, Aad, Ssd, Std, StdSeries, Variance, Cor, Covariance, Dif, Drawdown, Normsinv, Ncdf, Sim, Percentile, Kmeans, Er, Ar, Winratio, Avgwin, Avgloss, Se, Kelly, Zscore, Normalize, NormalizePair, NormalizeFrom, Denormalize, Standardize, Log, Exp, Martingale, Antimartingale, Sum, Permutations, Mse, Cum, RecentHigh, RecentLow, Pvalue, ExpectedTrails, LrSlope, LrIntercept, LrAngle, Tsf, ReturnPositive, ReturnNegative |
| `github.com/Bitvested/ta.go/src/chart_types` (`charttypes`)  | Ha (Heikin-Ashi), Ren (Renko) |
| `github.com/Bitvested/ta.go/src/experimental`                | Support, Resistance, Line, LineCalc, DivergenceState |
| `github.com/Bitvested/ta.go/src/random`                      | Prng, Pick, Range, Float, Order |
| `github.com/Bitvested/ta.go/src/misc`                        | Fibnumbers, TimesUp, TimesDown |
| `github.com/Bitvested/ta.go/multi`                           | Sim, SimSeeded — goroutine Monte Carlo |

## Bar input conventions

Indicators that take OHLCV-style bars use a fixed column ordering. These
match ta.js v2.0.0:

| Indicator family                                               | Bar layout         |
|----------------------------------------------------------------|--------------------|
| `Atr`, `Cci`, `Pdi`/`Mdi`/`Dx`/`Adx`/`Adxr`, `Pdm`/`Mdm`, `Vortex`, `Ult`, `Stoch`, `Kdj`, `Ichimoku`, `Halftrend`, `Psar`, `Supertrend`, `ZigZag`, `Don`, `Keltner` | `[H, C, L]` (or `[H, C, L, …]` extended) |
| `Adl`, `Cmf`, `ChaikinOsc`, `Kvo`                              | `[H, C, L, V]`     |
| `Mfi`                                                          | `[H, L, C, V]`     |
| `Emv`                                                          | `[H, L, V]`        |
| `Nvi`, `Pvi`, `Vwma`, `Vwwma`, `Vwap`                          | `[close, volume]`  |
| `Rvi`                                                          | `[O, H, L, C]`     |

If your data is in a different ordering, restructure it before calling.

## Migration notes (pre-v2.0.0 → v2.0.0)

If you are upgrading from a pre-v2.0.0 ta.go release, the most likely
breakage points are:

| Change | What's different |
|--------|------------------|
| **Imports** | The root `ta` package no longer exports indicators — use the category subpackages listed above. The root only holds the `Version` constant. |
| **`indicators.Rsi`** | Now uses Wilder smoothing (TA-Lib convention). The old simple-average behaviour is gone; if you need a tolerant migration, the existing `Wrsi` alias keeps the legacy semantics. |
| **`indicators.Roc`** | Returns rate of change in **percent** (×100), not as a fraction. `Cop`, `Kst` cascades adjusted accordingly. |
| **`indicators.Mfi`** | Input is now `[][]float64` of `[H, L, C, V]` bars. Pre-computed positive/negative flow no longer accepted. |
| **`indicators.Atr` / `oscillators.Fisher` / `movingaverages.Smma`** | Output length is `N − l + 1` (was `N`); TR uses absolute values for `\|H − Cprev\|` and `\|L − Cprev\|`. `Keltner`, `Supertrend`, `Halftrend` cascade from the new shape. |
| **`statistics.Median`** | Even-length windows now return the average of the two middle values (was the lower-median). `Mad` cascades. |
| **`indicators.Macd`** | The single `Macd` returns just the MACD line. The signal line and histogram split into `MacdSignal` and `MacdBars`. |
| **`indicators.Cross`** | Returns `[]CrossSignal{Index int; Cross bool}` (struct slice). Inputs are no longer mutated. |
| **NaN policy** | `Rsi`, `Roc`, `Dif`, `Mom`, `MomOsc`, `Mfi`, `Fisher`, `Pr`, `Stoch` return `math.NaN()` on flat windows / division-by-zero edges (instead of `0` or `1.0` coercions). Callers that previously relied on the silent-coercion behaviour need to handle NaN explicitly. |
| **No-mutation guarantee** | `Std`, `Cross`, `Divergence`, `Percentile`, and `experimental.Support`/`Resistance` no longer mutate caller inputs. |

See [`CHANGELOG.md`](CHANGELOG.md) for the full Phase-by-Phase list.

## JS ↔ Go name map

The mapping from ta.js v2.0.0 names to Go is uniform: snake_case →
PascalCase, dispatched into a category subpackage. A handful of
representative examples:

| ta.js                          | ta.go                                                        |
|--------------------------------|--------------------------------------------------------------|
| `ta.sma(d, 14)`                | `movingaverages.Sma(d, 14)`                                  |
| `ta.ema(d, 26)`                | `movingaverages.Ema(d, 26)`                                  |
| `ta.dema(d, 30)`               | `movingaverages.Dema(d, 30)`                                 |
| `ta.rsi(d, 14)`                | `indicators.Rsi(d, 14)`                                      |
| `ta.macd(d, 12, 26)`           | `indicators.Macd(d, 12, 26)`                                 |
| `ta.macd_signal(d, 12, 26, 9)` | `indicators.MacdSignal(d, 12, 26, 9)`                        |
| `ta.aroon_up(d, 14)`           | `indicators.AroonUp(d, 14)`                                  |
| `ta.bands(d, 20, 2)`           | `bands.Bands(d, 20, 2)`                                      |
| `ta.atr(hcl, 14)`              | `indicators.Atr(hcl, 14)`                                    |
| `ta.std(d, 14)`                | `statistics.Std(d, 14)`                                      |
| `ta.std_series(d, 14)`         | `statistics.StdSeries(d, 14)`                                |
| `ta.lr_slope(d, 14)`           | `statistics.LrSlope(d, 14)`                                  |
| `ta.divergence_state(...)`     | `experimental.DivergenceState(...)`                          |
| `ta.random.prng('seed')`       | `random.Prng("seed")`                                        |
| `ta.multi.sim(d, 50, 1000, 0.5)` | `multi.Sim(d, 50, 1000, 0.5)` (goroutine fan-out)         |

## Performance

ta.go ties or beats `markcheno/go-talib` and `cinar/indicator` at 100 k
bars on every benched indicator. Headline numbers (lower is better,
ns/op):

| Indicator | ta.go | go-talib | cinar |
|-----------|------:|---------:|------:|
| SMA       |  141 µs |   183 µs |  221 µs |
| EMA       |  301 µs |   299 µs |  502 µs |
| RSI       |  544 µs |   531 µs | 1912 µs |
| MACD      |  936 µs |  1108 µs | 1602 µs |
| Bollinger |  501 µs |   625 µs |  669 µs |
| ATR       |  513 µs |   974 µs |  762 µs |

Full table and reproduction instructions live in
[`bench/BENCHMARKS.md`](bench/BENCHMARKS.md). Internal Phase 4
before/after snapshots are in [`bench/results-phase4.txt`](bench/results-phase4.txt).

## Release tag

The `v2.0.0` line is published as Git tag **`v0.20.0`**. The `v0.x`
naming sidesteps Go's `/v2` module-path rule so existing imports keep
working — pin via the tag if you want the v2.0.0 line specifically:

```bash
go get github.com/Bitvested/ta.go@v0.20.0
```

Inside Go code:

```go
import ta "github.com/Bitvested/ta.go"
fmt.Println(ta.Version) // "2.0.0"
```

## License

MIT — see [LICENSE](LICENSE).
