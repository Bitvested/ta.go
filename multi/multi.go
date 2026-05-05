// Package multi provides goroutine-parallel Monte Carlo projections
// built on top of the ta.go statistics primitives.
//
// Given a price series, Sim extends it forward by `length` bars across
// `sims` independent geometric Brownian motion (GBM) runs and returns
// the data prefixed onto the per-bar `perc`-rank quantile of the
// ensemble. SimSeeded is the deterministic variant — drawing all
// random samples from a single seeded PRNG before fan-out, so the
// result depends only on the seed and inputs (not on goroutine
// scheduling).
package multi

import (
	"math"
	"math/rand"
	"runtime"
	"sort"
	"sync"

	ma "github.com/Bitvested/ta.go/src/moving_averages"
	stats "github.com/Bitvested/ta.go/src/statistics"
)

// Sim runs `sims` GBM Monte Carlo extensions of `data` by `length`
// bars and returns `data` followed by the per-bar `perc`-rank quantile
// of the ensemble (perc in (0, 1]). If perc is ≤ 0 or > 1 the median
// (0.5) is used. Random draws come from math/rand's package source —
// for reproducible results use SimSeeded.
func Sim(data []float64, length, sims int, perc float64) []float64 {
	return runSim(data, length, sims, perc, nil)
}

// SimSeeded is the deterministic variant of Sim: all random draws are
// produced sequentially from a single seeded source before fan-out, so
// the output depends only on (data, length, sims, perc, seed).
func SimSeeded(data []float64, length, sims int, perc float64, seed int64) []float64 {
	return runSim(data, length, sims, perc, rand.New(rand.NewSource(seed)))
}

func runSim(data []float64, length, sims int, perc float64, src *rand.Rand) []float64 {
	if length <= 0 || sims <= 0 || len(data) < 2 {
		return nil
	}
	if perc <= 0 || perc > 1 {
		perc = 0.5
	}
	// Pre-generate random draws in a single goroutine for determinism.
	draws := make([][]float64, sims)
	for i := 0; i < sims; i++ {
		draws[i] = make([]float64, length)
		for j := 0; j < length; j++ {
			if src != nil {
				draws[i][j] = src.Float64()
			} else {
				draws[i][j] = rand.Float64()
			}
		}
	}
	// Fan out workers.
	workers := runtime.NumCPU()
	if workers > sims {
		workers = sims
	}
	results := make([][]float64, sims)
	jobs := make(chan int, sims)
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx := range jobs {
				results[idx] = simulate(data, length, draws[idx])
			}
		}()
	}
	for i := 0; i < sims; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	// Build the percentile projection: data + per-future-bar quantile.
	n := len(data)
	out := make([]float64, 0, n+length)
	out = append(out, data...)
	col := make([]float64, sims)
	for i := n; i < n+length; i++ {
		for s := 0; s < sims; s++ {
			col[s] = results[s][i]
		}
		sort.Float64s(col)
		idx := int(math.Round(float64(sims-1) * perc))
		if idx < 0 {
			idx = 0
		}
		if idx >= sims {
			idx = sims - 1
		}
		out = append(out, col[idx])
	}
	return out
}

// simulate runs one GBM extension of `data` by `length` bars, drawing
// the per-bar normal samples from the supplied uniform `randoms` slice
// via Normsinv. Returns the original data followed by the projected
// tail.
func simulate(data []float64, length int, randoms []float64) []float64 {
	projected := make([]float64, len(data), len(data)+length)
	copy(projected, data)
	for x := 0; x < length; x++ {
		n := len(projected)
		change := make([]float64, n-1)
		for y := 1; y < n; y++ {
			change[y-1] = stats.Dif(projected[y], projected[y-1])
		}
		mean := ma.Sma(change, len(change))
		std := stats.Std(change, len(change))
		random := stats.Normsinv(randoms[x])
		projected = append(projected, projected[n-1]*math.Exp((mean[0]-(std*std)/2)+std*random))
	}
	return projected
}
