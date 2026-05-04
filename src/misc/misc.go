// Package misc contains small utilities and constants that don't fit a
// dedicated category (Fibonacci sequence, monotonicity counters).
package misc

var Fibnumbers = []float64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181}

func TimesUp(data []float64, l int) []float64 {
	var out []float64
	for i := l; i < len(data); i++ {
		var up float64 = 1
		for x := i - l + 1; x <= i; x++ {
			if data[x-1] > data[x] {
				up = 0
				break
			}
		}
		out = append(out, up)
	}
	return out
}

func TimesDown(data []float64, l int) []float64 {
	var out []float64
	for i := l; i < len(data); i++ {
		var dn float64 = 1
		for x := i - l + 1; x <= i; x++ {
			if data[x-1] < data[x] {
				dn = 0
				break
			}
		}
		out = append(out, dn)
	}
	return out
}
