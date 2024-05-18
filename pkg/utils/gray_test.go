package utils

import (
	"image"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToGrayscale(t *testing.T) {
	t.Parallel()

	img := image.NewRGBA(image.Rect(0, 0, 2, 2))
	img.Set(0, 0, color.RGBA{255, 0, 0, 255})
	img.Set(0, 1, color.RGBA{255, 255, 255, 255})
	img.Set(1, 0, color.RGBA{0, 0, 0, 255})
	img.Set(1, 1, color.RGBA{0, 0, 255, 255})

	expected := image.NewGray(image.Rect(0, 0, 2, 2))
	expected.Set(0, 0, color.Gray{76})
	expected.Set(0, 1, color.Gray{255})
	expected.Set(1, 0, color.Gray{0})
	expected.Set(1, 1, color.Gray{29})

	gray := ToGrayscale(img)
	assert.Equal(t, expected.Pix, gray.Pix)
}
