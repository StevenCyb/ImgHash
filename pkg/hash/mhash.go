package hash

import (
	"image"
	"imghash/pkg/utils"
	"sort"
)

// Median Hashing (mHash).
func MHash(img image.Image) (uint64, error) {
	if img == nil {
		return 0, ErrEmptyImage
	}

	// Resize the image
	img = utils.ResizeRGBInterArea(img, 8, 8)

	// Convert the image to grayscale
	grayImg := utils.ToGrayscale(img)

	// Find the median
	pixels := make([]float64, 64)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			pixels[i*8+j] = float64(grayImg.GrayAt(i, j).Y)
		}
	}
	sort.Float64s(pixels)
	median := pixels[31]

	// Convert each pixel to a bit based on whether it's above or below the average
	bitmap := make([]bool, 64)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			bitmap[i*8+j] = float64(grayImg.GrayAt(i, j).Y) > median
		}
	}

	// Transform the 64 bitmap into a 64 integer.
	var hash uint64
	for i := 0; i < 64; i++ {
		if bitmap[i] {
			hash |= 1 << uint(63-i)
		}
	}

	return hash, nil
}
