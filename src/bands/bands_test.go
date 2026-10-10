package bands_test

import (
	"testing"

	tu "github.com/Bitvested/ta.go/internal/testutil"
	bnd "github.com/Bitvested/ta.go/src/bands"
	"github.com/stretchr/testify/assert"
)

func TestBands(t *testing.T) {
	have := bnd.Bands([]float64{1, 2, 3, 4, 5, 6}, 5, 2)
	want := [][]float64{{5.82842712474619, 3.0, 0.1715728752538097}, {6.82842712474619, 4.0, 1.1715728752538097}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Bands Failed!")
	}
}

func TestBandwidth(t *testing.T) {
	have := bnd.Bandwidth([]float64{1, 2, 3, 4, 5, 6}, 5, 2)
	want := []float64{1.8856180831641265, 1.414213562373095}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Bandwidth Failed!")
	}
}

func TestKeltner(t *testing.T) {
	// v2.0.0: cascades from Atr's absolute-TR fix.
	have := bnd.Keltner([][]float64{{3, 2, 1}, {2, 2, 1}, {4, 3, 1}, {2, 2, 1}, {3, 3, 1}}, 5, 1)
	want := [][]float64{{4.092266666666667, 2.066666666666667, 0.04106666666666703}}
	tu.Slice2D(t, have, want, "Keltner")
}

func TestDon(t *testing.T) {
	have := bnd.Don([][]float64{{6, 2}, {5, 2}, {5, 3}, {6, 3}, {7, 4}, {6, 3}}, 5)
	want := [][]float64{{7, 4.5, 2}, {7, 4.5, 2}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Don Failed!")
	}
}

func TestFibbands(t *testing.T) {
	have := bnd.Fibbands([][]float64{{1, 59}, {1.1, 82}, {1.21, 27}, {1.42, 73}, {1.32, 42}}, 4, 3)
	want := [][]float64{{1.6526058894448858, 1.542197040614731, 1.4738932612537028, 1.4186888368386255, 1.363484412423548, 1.29518063306252, 1.184771784232365, 1.0743629354022102, 1.0060591560411822, 0.9508547316261047, 0.8956503072110273, 0.8273465278499992, 0.7169376790198443}, {1.6177775811703725, 1.5330576077284503, 1.4806460987347188, 1.4382861120137578, 1.3959261252927968, 1.3435146162990652, 1.258794642857143, 1.1740746694152209, 1.1216631604214893, 1.0793031737005283, 1.0369431869795673, 0.9845316779858357, 0.8998117045439136}}
	tu.Slice2D(t, have, want, "Fibbands")
}

func TestEnvelope(t *testing.T) {
	have := bnd.Envelope([]float64{6, 7, 8, 7, 6, 7, 8, 7, 8, 7, 8, 7, 8}, 11, 0.05)
	want := [][]float64{{7.540909090909091, 7.181818181818182, 6.822727272727272}, {7.636363636363637, 7.2727272727272725, 6.909090909090908}}
	if !assert.Equal(t, want, have) {
		t.Fatalf("Envelope Failed!")
	}
}
