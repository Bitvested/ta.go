# Changelog

All notable changes to ta.go.

## v2.0.0 — git tag `v0.20.0`

Major refactor. ~52 new functions, formula and convention fixes on
existing ones, no-mutation guarantees, a goroutine Monte Carlo
subpackage, and performance rewrites that make ta.go competitive with
or faster than `markcheno/go-talib` and `cinar/indicator` at 100 k
bars.

`Version = "2.0.0"` is exported from the root `ta` package; the Git
tag itself is `v0.20.0` so the existing `github.com/Bitvested/ta.go`
import path keeps working without Go's `/v2` module-path rule.

### Layout

- The single `ta.go` file is gone. Implementations now live under
  category subpackages:
  - `src/moving_averages` (`movingaverages`)
  - `src/indicators`
  - `src/oscillators`
  - `src/bands`
  - `src/statistics`
  - `src/chart_types` (`charttypes`)
  - `src/experimental`
  - `src/random` (new package)
  - `src/misc`
- New top-level `multi/` subpackage for goroutine Monte Carlo.
- New top-level `bench/` Go module (separate `go.mod`) for
  cross-library comparative benchmarks — does not pollute the
  consumer `go.sum`.
- New `internal/testutil` with `tu.Slice`, `tu.Slice2D`,
  `tu.Scalar` — float-tolerant comparators (rel/abs tolerance
  1e-12, NaN-aware).

### Formula fixes

- **`indicators.Psar`** — downtrend extreme now tracks the bar's
  low (`data[i][1]`); previous code mistakenly tracked the high.
- **`statistics.Kmeans`** — value-copies `old`, resets `means` per
  loop, recomputes by cluster index, and resets `changed`.
- **`indicators.Ichimoku`** — Senkou Span B's lows window now uses
  `length3`; the per-window max/min loops were replaced with
  rolling-extremum scans.
- **`movingaverages.Smma`** — Wilder recurrence with SMA seed of
  the first `N` values; output length is `N − l + 1`.
- **`indicators.Alligator` / `oscillators.Gator`** — fixtures
  regenerated against the canonical Williams formulas (cascades
  from the Smma fix).
- **`oscillators.MomOsc`** — Chande CMO formula on first
  differences (positive vs absolute negative deltas).
- **`indicators.Fractals`** — returns `[][]bool`; the broken
  price-mode shape was replaced with a separate
  `indicators.FractalsPrice` that pads two leading and two
  trailing `[-1, -1]` rows.
- **`indicators.ZigZag`** — fixed the state-transition tracker
  (the scalar branch was resetting the wrong extreme). Both
  scalar and OHLC branches share consistent semantics.

### Convention changes (breaking)

- **`indicators.Rsi`** — Wilder smoothing (TA-Lib convention) is
  the new default. The legacy simple-average behaviour is exposed
  as `Wrsi` for back-compat callers.
- **`indicators.Roc`** — returns percent (×100). `Cop` and `Kst`
  fixtures regenerate as a linear cascade.
- **`indicators.Mfi`** — input is `[][]float64` of `[H, L, C, V]`
  bars (was pre-computed positive/negative flow).
- **`statistics.Median`** — even windows return the average of the
  two middle values (was lower-median). `Mad` cascades.
- **`indicators.Atr`** — output length is `N − l + 1` (was `N`); TR
  uses absolute values for `|H − Cprev|` and `|L − Cprev|`. `Keltner`,
  `Supertrend`, `Halftrend` lose their compensating splice.
- **`oscillators.Fisher`** — output length is `N − l + 1` (drop
  trailing slice).

### No-mutation guarantees

The following functions no longer mutate caller inputs:
`statistics.Std`, `indicators.Cross` (now also returns
`[]CrossSignal{Index, Cross}`), `indicators.Divergence`,
`statistics.Percentile`, `experimental.Support` / `Resistance`.

### NaN handling standardised

`Rsi`, `Roc`, `Dif`, `Mom` (percent mode), `MomOsc`, `Mfi`,
`Fisher`, `Pr`, `Stoch` now return `math.NaN()` on flat-window /
divide-by-zero edges instead of silent `0`/`1.0` coercions.
`statistics.Kelly` no longer coerces NaN → 1.0;
`indicators.ChaikinOsc` no longer coerces → 0.

### New functions (~52)

#### Moving averages (6)
`Dema`, `Tema`, `T3`, `Trima`, `Vidya`, `Zlema`.

#### Indicators (26)
`Adl`, `Adx`, `Adxr`, `Apo`, `Cci`, `Cmf`, `Cross`, `Dpo`, `Dx`,
`Emv`, `Kdj`, `Mass`, `Mdi`, `Mdm`, `Natr`, `Nvi`, `Pdi`, `Pdm`,
`Ppo`, `Pvi`, `StochRsi`, `Trix`, `Ulcer`, `Vortex`, plus the
`Macd` split: `MacdSignal`, `MacdBars`.

#### Oscillators (2)
`Kvo`, `Ult`.

#### Statistics (9)
`Covariance`, `ExpectedTrails`, `LrSlope`, `LrIntercept`,
`LrAngle`, `Tsf`, `StdSeries`, `ReturnPositive`, `ReturnNegative`.
(`Ssd` already existed.)

#### Random (5, new package)
`Prng` (deterministic mulberry32 closure), `Pick`, `Range`,
`Float`, `Order`. Sampling helpers accept an optional rng with
`math/rand` fallback when nil.

#### Experimental (1)
`DivergenceState`.

While porting `DivergenceState`, two latent bugs in
`experimental.Support` were fixed: the in-package `minf`/`maxf`
helpers sorted the caller's slice in place (now copy first), and
the anchor-advance branch set `findex = true` instead of `false`
(so the outer do-while exited after one iteration). `Resistance`
was already correct on the second point.

### Performance rewrites

Captured per-rewrite in `bench/results-phase4.txt`. Headline
speedups at 100 k bars:

| Function   | Before  | After  | Speedup |
|------------|--------:|-------:|--------:|
| `Sma`      | 30 ms   | 154 µs | ≈195× |
| `Variance` | 110 ms  | 182 µs | ≈600× |
| `Bands`    | 22 ms   | 524 µs | ≈42× |
| `Ema`      | 960 µs  | 320 µs | ≈3× |
| `Wma`      | 800 µs  | 230 µs | ≈3.5× (10 k) |
| `Rsi`      | 935 µs  | 520 µs | ≈1.8× |

Plus drop-per-window-copies/sort in `AroonUp`/`AroonDown`/
`Elderray`/`Don`/`Fisher` (≈13–50× at 10 k).

### Multi-thread Monte Carlo

New `multi/` subpackage:

- `multi.Sim(data, length, sims, perc) []float64` — fans
  `min(sims, runtime.NumCPU())` worker goroutines reading from a
  job channel; returns `data` followed by the per-bar `perc`-rank
  quantile of the ensemble.
- `multi.SimSeeded(data, length, sims, perc, seed) []float64` —
  deterministic variant. All uniform draws are pre-generated in a
  single goroutine before fan-out, so output depends only on
  `(data, length, sims, perc, seed)`.

### Cross-library benchmarks

New `bench/` separate Go module compares ta.go against
`github.com/markcheno/go-talib` and `github.com/cinar/indicator`
for SMA, EMA, RSI, MACD, Bollinger Bands, and ATR at 1 k / 10 k /
100 k bars. ta.go ties or wins at 100 k on every benched
indicator. See `bench/BENCHMARKS.md`.

### Tests

Test count grew from 99/100 (one pre-existing float-drift
failure) to 165/165 passing. Float-tolerant comparator at
`internal/testutil/close.go` underpins fp-sensitive fixtures.
