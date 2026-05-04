package misc_test

import (
	"testing"

	misc "github.com/Bitvested/ta.go/src/misc"
	"github.com/stretchr/testify/assert"
)

func TestTimesUp(t *testing.T) {
	have := misc.TimesUp([]float64{5, 6, 7, 8, 7, 6, 5}, 3)
	want := []float64{1, 0, 0, 0}
	if !assert.Equal(t, want, have) {
		t.Fatalf("TimesUp Failed!")
	}
}

func TestTimesDown(t *testing.T) {
	have := misc.TimesDown([]float64{5, 6, 7, 8, 7, 6, 5}, 3)
	want := []float64{0, 0, 0, 1}
	if !assert.Equal(t, want, have) {
		t.Fatalf("TimesDown Failed!")
	}
}
