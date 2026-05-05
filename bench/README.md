# bench/

Separate Go module containing cross-library performance comparisons
for ta.go.

This module is **not** part of the main `github.com/Bitvested/ta.go`
module — its `go.mod` declares an independent module path so the
external benchmark dependencies (`github.com/markcheno/go-talib`,
`github.com/cinar/indicator`) never leak into a consumer's `go.sum`.

The local `replace` directive in `bench/go.mod` points the bench
module at `..`, so it always benches the in-tree ta.go.

## Run

```bash
cd bench
CGO_ENABLED=0 go test -bench=. -benchtime=300ms -count=3
```

`CGO_ENABLED=0` is needed only on this developer's macOS install
(Go 1.19.1 has a dyld load issue here); on other platforms it can
be omitted.

## Files

- `bench_test.go` — comparative benches for SMA, EMA, RSI, MACD,
  Bollinger Bands, and ATR at 1k / 10k / 100k bars across
  ta.go, markcheno/go-talib, and cinar/indicator.
- `BENCHMARKS.md` — published comparison table generated from a
  representative run.
- `results-phase4.txt` — Phase 4 before/after snapshots from the
  internal optimisation pass; not a cross-library comparison.
