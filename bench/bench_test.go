// Package bench runs comparative benchmarks between ta.go and other
// public Go technical-analysis libraries (markcheno/go-talib,
// cinar/indicator).
//
// This module is intentionally separate from the main ta.go module so
// the third-party deps don't pollute the consumer go.sum.
package bench

import (
	"math/rand"
	"testing"

	tabands "github.com/Bitvested/ta.go/src/bands"
	taind "github.com/Bitvested/ta.go/src/indicators"
	tama "github.com/Bitvested/ta.go/src/moving_averages"

	cinar "github.com/cinar/indicator"
	talib "github.com/markcheno/go-talib"
)

// --- shared input data ------------------------------------------------------

func priceSeries(n int) []float64 {
	r := rand.New(rand.NewSource(1))
	out := make([]float64, n)
	out[0] = 100
	for i := 1; i < n; i++ {
		out[i] = out[i-1] * (1 + (r.Float64()-0.5)*0.02)
	}
	return out
}

// hcl returns high/close/low slices in ta.go's [H,C,L] tuple layout
// plus parallel high/low/close slices for the talib/cinar APIs.
func hcl(n int) (taTuples [][]float64, h, l, c []float64) {
	r := rand.New(rand.NewSource(2))
	taTuples = make([][]float64, n)
	h = make([]float64, n)
	l = make([]float64, n)
	c = make([]float64, n)
	closing := 100.0
	for i := 0; i < n; i++ {
		closing *= 1 + (r.Float64()-0.5)*0.02
		h[i] = closing + 0.5
		c[i] = closing
		l[i] = closing - 0.5
		taTuples[i] = []float64{h[i], c[i], l[i]}
	}
	return
}

var sizes = []struct {
	name string
	n    int
}{{"1k", 1000}, {"10k", 10000}, {"100k", 100000}}

// --- SMA --------------------------------------------------------------------

func BenchmarkSmaTaGo(b *testing.B) {
	for _, s := range sizes {
		d := priceSeries(s.n)
		b.Run(s.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = tama.Sma(d, 50)
			}
		})
	}
}

func BenchmarkSmaTalib(b *testing.B) {
	for _, s := range sizes {
		d := priceSeries(s.n)
		b.Run(s.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = talib.Sma(d, 50)
			}
		})
	}
}

func BenchmarkSmaCinar(b *testing.B) {
	for _, s := range sizes {
		d := priceSeries(s.n)
		b.Run(s.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = cinar.Sma(50, d)
			}
		})
	}
}

// --- EMA --------------------------------------------------------------------

func BenchmarkEmaTaGo(b *testing.B) {
	for _, s := range sizes {
		d := priceSeries(s.n)
		b.Run(s.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = tama.Ema(d, 50)
			}
		})
	}
}

func BenchmarkEmaTalib(b *testing.B) {
	for _, s := range sizes {
		d := priceSeries(s.n)
		b.Run(s.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = talib.Ema(d, 50)
			}
		})
	}
}

func BenchmarkEmaCinar(b *testing.B) {
	for _, s := range sizes {
		d := priceSeries(s.n)
		b.Run(s.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = cinar.Ema(50, d)
			}
		})
	}
}

// --- RSI --------------------------------------------------------------------

func BenchmarkRsiTaGo(b *testing.B) {
	for _, s := range sizes {
		d := priceSeries(s.n)
		b.Run(s.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = taind.Rsi(d, 14)
			}
		})
	}
}

func BenchmarkRsiTalib(b *testing.B) {
	for _, s := range sizes {
		d := priceSeries(s.n)
		b.Run(s.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = talib.Rsi(d, 14)
			}
		})
	}
}

// cinar.Rsi uses a fixed period of 14, matches ta.go default; returns
// (rs, rsi).
func BenchmarkRsiCinar(b *testing.B) {
	for _, s := range sizes {
		d := priceSeries(s.n)
		b.Run(s.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = cinar.Rsi(d)
			}
		})
	}
}

// --- MACD -------------------------------------------------------------------

func BenchmarkMacdTaGo(b *testing.B) {
	for _, s := range sizes {
		d := priceSeries(s.n)
		b.Run(s.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = taind.Macd(d, 12, 26)
			}
		})
	}
}

func BenchmarkMacdTalib(b *testing.B) {
	for _, s := range sizes {
		d := priceSeries(s.n)
		b.Run(s.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _, _ = talib.Macd(d, 12, 26, 9)
			}
		})
	}
}

func BenchmarkMacdCinar(b *testing.B) {
	for _, s := range sizes {
		d := priceSeries(s.n)
		b.Run(s.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = cinar.Macd(d)
			}
		})
	}
}

// --- Bollinger Bands --------------------------------------------------------

func BenchmarkBandsTaGo(b *testing.B) {
	for _, s := range sizes {
		d := priceSeries(s.n)
		b.Run(s.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = tabands.Bands(d, 20, 2)
			}
		})
	}
}

func BenchmarkBandsTalib(b *testing.B) {
	for _, s := range sizes {
		d := priceSeries(s.n)
		b.Run(s.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _, _ = talib.BBands(d, 20, 2, 2, talib.SMA)
			}
		})
	}
}

func BenchmarkBandsCinar(b *testing.B) {
	for _, s := range sizes {
		d := priceSeries(s.n)
		b.Run(s.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _, _ = cinar.BollingerBands(d)
			}
		})
	}
}

// --- ATR --------------------------------------------------------------------

func BenchmarkAtrTaGo(b *testing.B) {
	for _, s := range sizes {
		hcls, _, _, _ := hcl(s.n)
		b.Run(s.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = taind.Atr(hcls, 14)
			}
		})
	}
}

func BenchmarkAtrTalib(b *testing.B) {
	for _, s := range sizes {
		_, h, l, c := hcl(s.n)
		b.Run(s.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = talib.Atr(h, l, c, 14)
			}
		})
	}
}

func BenchmarkAtrCinar(b *testing.B) {
	for _, s := range sizes {
		_, h, l, c := hcl(s.n)
		b.Run(s.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = cinar.Atr(14, h, l, c)
			}
		})
	}
}
