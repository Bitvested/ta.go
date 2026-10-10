package indicators_test

import (
	"encoding/csv"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	ind "github.com/Bitvested/ta.go/src/indicators"
)

// Parity test against a TradingView chart-data export.
//
// Everything else in the suite is derived from the Pine source pasted in the ta.py
// issue "Williams Alligator not matching the one in trading view"; this is the only
// check that compares against numbers TradingView itself plotted, and therefore the
// only one that can actually establish parity.
//
// To capture the fixture: put the Williams Alligator on a chart (defaults), scroll
// fully left so every loaded bar is in memory, then Chart menu -> "Download chart
// data" -> CSV. The export carries a column per study in the Data Window, so
// Jaw/Teeth/Lips come along with the OHLC. Save it as
// src/indicators/testdata/tradingview_alligator.csv. Trimming rows off the end is
// fine; trimming rows off the front is not.
//
// The window matters. The SMMA is recursive, so its value at any bar depends on
// every bar before it: if the export starts partway through the loaded history, the
// leading rows cannot match and no tolerance helps. This test reports that case as
// a window problem rather than as a value mismatch.

const (
	// TradingView writes a fixed number of decimals, so this is the export's print
	// precision, not algorithmic slack (the library's own tolerance is 1e-12).
	tvTolerance = 1e-6
	jawFirstBar = 13 - 1 + 8 // jaw length - 1 + jaw shift
)

func loadTradingViewExport(t *testing.T) ([]float64, [][]string) {
	t.Helper()

	path := filepath.Join("testdata", "tradingview_alligator.csv")
	handle, err := os.Open(path)
	if err != nil {
		t.Skipf("no TradingView export at %s — see this file's comment", path)
	}
	defer handle.Close()

	records, err := csv.NewReader(handle).ReadAll()
	if err != nil {
		t.Fatalf("%s: %v", path, err)
	}
	if len(records) < 2 {
		t.Fatalf("%s has no data rows", path)
	}

	header := records[0]
	find := func(names ...string) int {
		for i, h := range header {
			h = strings.ToLower(strings.TrimSpace(strings.TrimPrefix(h, "\xef\xbb\xbf")))
			for _, name := range names {
				if h == name {
					return i
				}
			}
		}
		return -1
	}
	highIdx, lowIdx := find("high"), find("low")
	plotIdx := []int{find("jaw"), find("teeth"), find("lips")}
	if highIdx < 0 || lowIdx < 0 || plotIdx[0] < 0 || plotIdx[1] < 0 || plotIdx[2] < 0 {
		t.Fatalf("%s has columns %v, so it is missing either the OHLC or the "+
			"Jaw/Teeth/Lips study columns. Re-export with the Alligator on the chart; "+
			"the export only carries studies that are in the Data Window.", path, header)
	}

	body := records[1:]
	hl2 := make([]float64, len(body))
	plots := make([][]string, len(plotIdx))
	for col := range plots {
		plots[col] = make([]string, len(body))
	}
	for i, record := range body {
		high, errHigh := strconv.ParseFloat(strings.TrimSpace(record[highIdx]), 64)
		low, errLow := strconv.ParseFloat(strings.TrimSpace(record[lowIdx]), 64)
		if errHigh != nil || errLow != nil {
			t.Fatalf("%s row %d: bad OHLC value (%v, %v)", path, i, errHigh, errLow)
		}
		hl2[i] = (high + low) / 2
		for col, idx := range plotIdx {
			plots[col][i] = strings.TrimSpace(record[idx])
		}
	}
	return hl2, plots
}

func TestTradingViewExportBeginsWhereComputationBegan(t *testing.T) {
	_, plots := loadTradingViewExport(t)

	firstFilled := -1
	for i, text := range plots[0] {
		if text != "" {
			firstFilled = i
			break
		}
	}
	if firstFilled < 0 {
		t.Fatal("the export has no Jaw column values at all")
	}
	if firstFilled != jawFirstBar {
		t.Fatalf("jaw first appears at row %d, expected %d: the export window does "+
			"not match the loaded history, so the SMMA seeds differ", firstFilled, jawFirstBar)
	}
}

func TestTradingViewJawTeethLipsMatch(t *testing.T) {
	hl2, plots := loadTradingViewExport(t)

	have := ind.Alligator(hl2, 13, 8, 5, 8, 5, 3)
	if len(have) != len(hl2) {
		t.Fatalf("got %d rows for %d bars", len(have), len(hl2))
	}

	compared := 0
	for col, name := range []string{"jaw", "teeth", "lips"} {
		for i, text := range plots[col] {
			if text == "" {
				continue
			}
			expected, err := strconv.ParseFloat(text, 64)
			if err != nil {
				t.Fatalf("%s row %d: unparseable value %q", name, i, text)
			}
			actual := have[i][col]
			if math.IsNaN(actual) {
				t.Errorf("%s[%d] is NaN here but TradingView plotted %v", name, i, expected)
				continue
			}
			if math.Abs(actual-expected) > tvTolerance {
				t.Errorf("%s[%d]: ta.go %v, TradingView %v", name, i, actual, expected)
			}
			compared++
		}
	}
	if compared == 0 {
		t.Fatal("the export had no data points to compare")
	}
}
