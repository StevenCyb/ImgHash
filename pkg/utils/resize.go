package utils

import (
	"image"
	"image/color"
)

// ResizeRGBInterArea resizes an image using the nearest neighbor algorithm.
func ResizeRGBInterArea(img image.Image, width, height int) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))

	// Calculate scaling factors
	xScale := float64(bounds.Dx()) / float64(width)
	yScale := float64(bounds.Dy()) / float64(height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Calculate the corresponding rectangle in the source image
			x0 := int(float64(x) * xScale)
			y0 := int(float64(y) * yScale)
			x1 := int(float64(x+1) * xScale)
			y1 := int(float64(y+1) * yScale)

			// Average pixel values in the source rectangle
			var totalR, totalG, totalB uint32
			count := 0
			for j := y0; j < y1; j++ {
				for i := x0; i < x1; i++ {
					r, g, b, _ := img.At(i, j).RGBA()
					totalR += r
					totalG += g
					totalB += b
					count++
				}
			}

			// Ensure count is never zero
			if count == 0 {
				count = 1
				if x0 < bounds.Max.X && y0 < bounds.Max.Y {
					r, g, b, _ := img.At(x0, y0).RGBA()
					totalR = r
					totalG = g
					totalB = b
				} else {
					// Default to black if the source pixel is out of bounds
					totalR = 0
					totalG = 0
					totalB = 0
				}
			}

			avgR := totalR / uint32(count)
			avgG := totalG / uint32(count)
			avgB := totalB / uint32(count)

			// Set the averaged pixel value in the new image
			newImg.Set(x, y, color.RGBA{uint8(avgR >> 8), uint8(avgG >> 8), uint8(avgB >> 8), 255})
		}
	}

	return newImg
}
