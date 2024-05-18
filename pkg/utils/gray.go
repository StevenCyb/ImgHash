package utils

import (
	"image"
	"image/color"
)

// ToGrayscale converts an image to grayscale.
func ToGrayscale(img image.Image) *image.Gray {
	bounds := img.Bounds()
	gray := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldPixel := img.At(x, y)
			grayPixel := color.GrayModel.Convert(oldPixel)
			gray.Set(x, y, grayPixel)
		}
	}

	return gray
}
