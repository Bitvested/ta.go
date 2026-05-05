// Module ta.go/bench — separate from the main ta.go module so external
// benchmark dependencies (markcheno/go-talib, cinar/indicator) don't
// leak into the main go.sum.
module github.com/Bitvested/ta.go/bench

go 1.19

replace github.com/Bitvested/ta.go => ../

require (
	github.com/Bitvested/ta.go v0.0.0-00010101000000-000000000000
	github.com/cinar/indicator v1.3.0
	github.com/markcheno/go-talib v0.0.0-20250114000313-ec55a20c902f
)
