package hash

import (
	"image"
	"imghash/pkg/utils"
)

// Difference Hashing (dHash).
func DHash(img image.Image) (uint64, error) {
	if img == nil {
		return 0, ErrEmptyImage
	}

	// Resize the image
	img = utils.ResizeRGBInterArea(img, 9, 8)

	// Convert the image to grayscale
	grayImg := utils.ToGrayscale(img)

	// From each row, the first 8 pixels are examined serially from left
	// to right and compared to their neighbor to the right
	diffMap := make([]bool, 64)
	for y := 0; y < 8; y++ {
		previous := grayImg.GrayAt(0, y).Y
		for x := 1; x < 9; x++ {
			current := grayImg.GrayAt(x, y).Y
			// Compare with right neighbor
			diffMap[y*8+x-1] = previous > current
			previous = current
		}
	}

	// Transform the 64 bitmap into a 64 integer
	var hash uint64
	for i := 0; i < 64; i++ {
		if diffMap[i] {
			hash |= 1 << uint(63-i)
		}
	}

	return hash, nil
}
