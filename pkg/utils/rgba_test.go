package utils

import (
	"image"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToRGBA(t *testing.T) {
	t.Parallel()

	img := image.NewGray(image.Rect(0, 0, 2, 2))
	img.Set(0, 0, color.Gray{76})
	img.Set(0, 1, color.Gray{255})
	img.Set(1, 0, color.Gray{0})
	img.Set(1, 1, color.Gray{29})

	expected := image.NewRGBA(image.Rect(0, 0, 2, 2))
	expected.Set(0, 0, color.RGBA{76, 76, 76, 255})
	expected.Set(0, 1, color.RGBA{255, 255, 255, 255})
	expected.Set(1, 0, color.RGBA{0, 0, 0, 255})
	expected.Set(1, 1, color.RGBA{29, 29, 29, 255})

	rgba := ToRGBA(img)
	assert.Equal(t, expected.Pix, rgba.Pix)
}
