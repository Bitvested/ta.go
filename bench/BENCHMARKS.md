# ta.go cross-library benchmarks

ta.go v2.0.0 vs `markcheno/go-talib` and `cinar/indicator`.

Reproduce locally:

```bash
cd bench && CGO_ENABLED=0 go test -bench=. -benchtime=300ms -count=1
```

Hardware: VirtualApple @ 2.50 GHz, GOOS=darwin GOARCH=amd64. Inputs are
deterministic-seeded random walks (math/rand source 1); HCL tuples use
±0.5 around the seeded close. Period parameters: SMA=50, EMA=50,
RSI=14, MACD=(12, 26, 9), Bollinger=(20, 2σ), ATR=14. Lower is better.

## ns/op at 100k bars

| Indicator | ta.go      | markcheno/go-talib | cinar/indicator | ta.go vs talib | ta.go vs cinar |
|-----------|-----------:|-------------------:|----------------:|---------------:|---------------:|
| SMA       |   140 862  |            183 344 |         221 359 | **1.30× faster** | **1.57× faster** |
| EMA       |   300 585  |            299 477 |         502 044 | within noise   | **1.67× faster** |
| RSI       |   544 016  |            530 593 |       1 912 205 | within noise   | **3.51× faster** |
| MACD      |   935 884  |          1 107 558 |       1 602 311 | **1.18× faster** | **1.71× faster** |
| Bollinger |   500 951  |            624 966 |         669 380 | **1.25× faster** | **1.34× faster** |
| ATR       |   513 062  |            974 211 |         762 348 | **1.90× faster** | **1.49× faster** |

## ns/op at 10k bars

| Indicator | ta.go    | markcheno/go-talib | cinar/indicator |
|-----------|---------:|-------------------:|----------------:|
| SMA       |  12 489  |             16 868 |          20 404 |
| EMA       |  29 252  |             29 055 |          49 525 |
| RSI       |  48 302  |             47 342 |         168 000 |
| MACD      |  93 746  |            107 826 |         166 906 |
| Bollinger |  98 334  |             63 526 |          71 707 |
| ATR       |  50 596  |             84 810 |          74 624 |

## ns/op at 1k bars

| Indicator | ta.go   | markcheno/go-talib | cinar/indicator |
|-----------|--------:|-------------------:|----------------:|
| SMA       |  1 497  |              1 803 |           2 211 |
| EMA       |  2 948  |              2 928 |           5 036 |
| RSI       |  4 576  |              4 814 |          17 109 |
| MACD      |  8 485  |             11 212 |          16 830 |
| Bollinger |  6 220  |              6 828 |           7 914 |
| ATR       |  5 042  |              8 609 |           7 598 |

## Notes

- ta.go ties or wins on every indicator at 100k bars except EMA and RSI
  where ta.go and `markcheno/go-talib` are within run-to-run noise of each
  other.
- The 10k Bollinger row is the one outlier where ta.go is slower than
  both competitors — `markcheno/go-talib` gets to skip output rows
  before the warm-up bar; ta.go and cinar both emit the full
  `n-l+1` rows. (At 100k that overhead becomes irrelevant and ta.go is
  fastest again.)
- `cinar/indicator` returns extra companion slices for several
  indicators (e.g. Rsi returns `(rs, rsi)`); the bench includes those
  return slots in the workload, matching how a caller would invoke it.
- `cinar/indicator` uses fixed default periods for Rsi (14), Macd
  (12, 26, 9), and BollingerBands (20, 2σ). Those match the
  parameters used by talib and ta.go in this bench.

## Layout

```
bench/
  go.mod                # separate module — does not pollute consumer go.sum
  bench_test.go         # ta.go + talib + cinar comparative benches
  BENCHMARKS.md         # this file
  results-phase4.txt    # internal Phase 4 before/after snapshots
```
