package utils

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForwardDCT(t *testing.T) {
	t.Parallel()

	img := image.Gray{
		Pix:    []uint8{0, 3, 0, 5},
		Stride: 2,
		Rect:   image.Rect(0, 0, 2, 2),
	}

	dct, err := ForwardDCT(img)
	assert.NoError(t, err)
	assert.Len(t, dct, 2)
	assert.Len(t, dct[0], 2)

	expectAround := [][]float64([][]float64{{0.9999, -0.2499}, {-0.9999, 0.2499}})
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			assert.InDelta(t, expectAround[i][j], dct[i][j], 0.0001)
		}
	}
}
