package utils

import (
	"image"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ResizeRGBInterArea_ScaleDown(t *testing.T) {
	t.Parallel()

	input := image.NewRGBA(image.Rect(0, 0, 4, 4))
	input.Set(0, 0, color.RGBA{255, 255, 255, 255})
	input.Set(0, 1, color.RGBA{255, 255, 255, 255})
	input.Set(0, 2, color.RGBA{0, 255, 0, 255})
	input.Set(0, 3, color.RGBA{0, 255, 0, 255})
	input.Set(1, 0, color.RGBA{255, 255, 255, 255})
	input.Set(1, 1, color.RGBA{255, 255, 255, 255})
	input.Set(1, 2, color.RGBA{0, 255, 0, 255})
	input.Set(1, 3, color.RGBA{0, 255, 0, 255})
	input.Set(2, 0, color.RGBA{255, 0, 0, 255})
	input.Set(2, 1, color.RGBA{255, 0, 0, 255})
	input.Set(2, 2, color.RGBA{0, 0, 0, 255})
	input.Set(2, 3, color.RGBA{0, 0, 0, 255})
	input.Set(3, 0, color.RGBA{255, 0, 0, 255})
	input.Set(3, 1, color.RGBA{255, 0, 0, 255})
	input.Set(3, 2, color.RGBA{0, 0, 0, 255})
	input.Set(3, 3, color.RGBA{0, 0, 0, 255})

	expect := image.NewRGBA(image.Rect(0, 0, 2, 2))
	expect.Set(0, 0, color.RGBA{255, 255, 255, 255})
	expect.Set(0, 1, color.RGBA{0, 255, 0, 255})
	expect.Set(1, 0, color.RGBA{255, 0, 0, 255})
	expect.Set(1, 1, color.RGBA{0, 0, 0, 255})

	resized := ResizeRGBInterArea(input, 2, 2)
	assert.Equal(t, expect.Pix, resized.(*image.RGBA).Pix)
}

func Test_ResizeRGBLinear_ScaleUp(t *testing.T) {
	t.Parallel()

	input := image.NewRGBA(image.Rect(0, 0, 2, 2))
	input.Set(0, 0, color.RGBA{255, 255, 255, 255})
	input.Set(0, 1, color.RGBA{0, 255, 0, 255})
	input.Set(1, 0, color.RGBA{255, 0, 0, 255})
	input.Set(1, 1, color.RGBA{0, 0, 0, 255})

	expect := image.NewRGBA(image.Rect(0, 0, 4, 4))
	expect.Set(0, 0, color.RGBA{255, 255, 255, 255})
	expect.Set(0, 1, color.RGBA{255, 255, 255, 255})
	expect.Set(0, 2, color.RGBA{0, 255, 0, 255})
	expect.Set(0, 3, color.RGBA{0, 255, 0, 255})
	expect.Set(1, 0, color.RGBA{255, 255, 255, 255})
	expect.Set(1, 1, color.RGBA{255, 255, 255, 255})
	expect.Set(1, 2, color.RGBA{0, 255, 0, 255})
	expect.Set(1, 3, color.RGBA{0, 255, 0, 255})
	expect.Set(2, 0, color.RGBA{255, 0, 0, 255})
	expect.Set(2, 1, color.RGBA{255, 0, 0, 255})
	expect.Set(2, 2, color.RGBA{0, 0, 0, 255})
	expect.Set(2, 3, color.RGBA{0, 0, 0, 255})
	expect.Set(3, 0, color.RGBA{255, 0, 0, 255})
	expect.Set(3, 1, color.RGBA{255, 0, 0, 255})
	expect.Set(3, 2, color.RGBA{0, 0, 0, 255})
	expect.Set(3, 3, color.RGBA{0, 0, 0, 255})

	resized := ResizeRGBInterArea(input, 4, 4)

	assert.Equal(t, expect.Pix, resized.(*image.RGBA).Pix)
}

func Test_ResizeRGBInterAreaWithRatio_ScaleDown(t *testing.T) {
	t.Parallel()

	input := image.NewRGBA(image.Rect(0, 0, 4, 4))
	input.Set(0, 0, color.RGBA{255, 255, 255, 255})
	input.Set(0, 1, color.RGBA{255, 255, 255, 255})
	input.Set(0, 2, color.RGBA{0, 255, 0, 255})
	input.Set(0, 3, color.RGBA{0, 255, 0, 255})
	input.Set(1, 0, color.RGBA{255, 255, 255, 255})
	input.Set(1, 1, color.RGBA{255, 255, 255, 255})
	input.Set(1, 2, color.RGBA{0, 255, 0, 255})
	input.Set(1, 3, color.RGBA{0, 255, 0, 255})
	input.Set(2, 0, color.RGBA{255, 0, 0, 255})
	input.Set(2, 1, color.RGBA{255, 0, 0, 255})
	input.Set(2, 2, color.RGBA{0, 0, 0, 255})
	input.Set(2, 3, color.RGBA{0, 0, 0, 255})
	input.Set(3, 0, color.RGBA{255, 0, 0, 255})
	input.Set(3, 1, color.RGBA{255, 0, 0, 255})
	input.Set(3, 2, color.RGBA{0, 0, 0, 255})
	input.Set(3, 3, color.RGBA{0, 0, 0, 255})

	expect := image.NewRGBA(image.Rect(0, 0, 2, 2))
	expect.Set(0, 0, color.RGBA{255, 255, 255, 255})
	expect.Set(0, 1, color.RGBA{0, 255, 0, 255})
	expect.Set(1, 0, color.RGBA{255, 0, 0, 255})
	expect.Set(1, 1, color.RGBA{0, 0, 0, 255})

	resized := ResizeRGBInterAreaWithRatio(input, 2)
	assert.Equal(t, expect.Pix, resized.(*image.RGBA).Pix)
}

func Test_ResizeRGBInterAreaWithRatio_ScaleUp(t *testing.T) {
	t.Parallel()

	input := image.NewRGBA(image.Rect(0, 0, 2, 2))
	input.Set(0, 0, color.RGBA{255, 255, 255, 255})
	input.Set(0, 1, color.RGBA{0, 255, 0, 255})
	input.Set(1, 0, color.RGBA{255, 0, 0, 255})
	input.Set(1, 1, color.RGBA{0, 0, 0, 255})

	expect := image.NewRGBA(image.Rect(0, 0, 4, 4))
	expect.Set(0, 0, color.RGBA{255, 255, 255, 255})
	expect.Set(0, 1, color.RGBA{255, 255, 255, 255})
	expect.Set(0, 2, color.RGBA{0, 255, 0, 255})
	expect.Set(0, 3, color.RGBA{0, 255, 0, 255})
	expect.Set(1, 0, color.RGBA{255, 255, 255, 255})
	expect.Set(1, 1, color.RGBA{255, 255, 255, 255})
	expect.Set(1, 2, color.RGBA{0, 255, 0, 255})
	expect.Set(1, 3, color.RGBA{0, 255, 0, 255})
	expect.Set(2, 0, color.RGBA{255, 0, 0, 255})
	expect.Set(2, 1, color.RGBA{255, 0, 0, 255})
	expect.Set(2, 2, color.RGBA{0, 0, 0, 255})
	expect.Set(2, 3, color.RGBA{0, 0, 0, 255})
	expect.Set(3, 0, color.RGBA{255, 0, 0, 255})
	expect.Set(3, 1, color.RGBA{255, 0, 0, 255})
	expect.Set(3, 2, color.RGBA{0, 0, 0, 255})
	expect.Set(3, 3, color.RGBA{0, 0, 0, 255})

	resized := ResizeRGBInterAreaWithRatio(input, 4)

	assert.Equal(t, expect.Pix, resized.(*image.RGBA).Pix)
}
