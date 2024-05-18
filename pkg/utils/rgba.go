package utils

import (
	"image"
	"image/draw"
)

// ToRGBA converts an image to RGBA.
func ToRGBA(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)

	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	return rgba
}
