package hash

import (
	"image"
	"imghash/pkg/utils"
)

// Average Hashing (aHash).
func AHash(img image.Image) (uint64, error) {
	if img == nil {
		return 0, ErrEmptyImage
	}

	// Resize the image
	img = utils.ResizeRGBInterArea(img, 8, 8)

	// Convert the image to grayscale
	grayImg := utils.ToGrayscale(img)

	// Compute the average color
	averageColor := float64(0)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			averageColor += float64(grayImg.GrayAt(i, j).Y)
		}
	}
	averageColor /= float64(64)

	// Convert each pixel to a bit based on whether it's above or below the average
	bitmap := make([]bool, 64)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			bitmap[i*8+j] = float64(grayImg.GrayAt(i, j).Y) > averageColor
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
