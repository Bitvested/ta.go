package charttypes_test

import (
	"testing"

	ct "github.com/Bitvested/ta.go/src/chart_types"
	"github.com/stretchr/testify/assert"
)

func TestHa(t *testing.T) {
	have := ct.Ha([][]float64{{3, 4, 2, 3}, {3, 6, 3, 5}, {5, 5, 2, 3}})
	want := [][]float64{{3, 4, 2, 3}, {3, 6, 3, 4.25}, {3.625, 5, 2, 3.75}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Ha Failed!")
	}
}

func TestRen(t *testing.T) {
	have := ct.Ren([][]float64{{8, 6}, {9, 7}, {9, 8}, {13, 10}}, 2)
	want := [][]float64{{8, 10, 8, 10}, {10, 12, 10, 12}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Ren Failed!")
	}
}
