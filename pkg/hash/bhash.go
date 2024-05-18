package hash

import (
	"errors"
	"image"
	"image/color"
	"imghash/pkg/utils"
)

var ErrImageDimensionsNotDivisibleByBlockSize = errors.New("image dimensions are not divisible by block size")

// Block Hashing (bHash).
func BHash(img image.Image) (uint64, error) {
	if img == nil {
		return 0, ErrEmptyImage
	}

	// Normally this algorithm does not resize the image
	// but for performance reasons this could be done anyway.
	// With 128x128 it has nice block size of 8x8.
	img = utils.ResizeRGBInterArea(img, 128, 128)

	// Create a slice to store the block differences
	diffMap := make([]bool, 64)
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			// Calculate the bounds of the current block
			blockBounds := image.Rect(x*8, y*8, (x+1)*8, (y+1)*8)

			// Calculate the average intensity of the current block
			averageIntensity := calculateBlockAverageIntensity(img, blockBounds)

			// Compare the intensity of the current block with its neighboring blocks
			if x < 8-1 {
				rightBlockBounds := image.Rect((x+1)*8, y*8, (x+2)*8, (y+1)*8)
				rightBlockAverageIntensity := calculateBlockAverageIntensity(img, rightBlockBounds)
				diffMap[y*8+x] = averageIntensity > rightBlockAverageIntensity
			}
			if y < 8-1 {
				bottomBlockBounds := image.Rect(x*8, (y+1)*8, (x+1)*8, (y+2)*8)
				bottomBlockAverageIntensity := calculateBlockAverageIntensity(img, bottomBlockBounds)
				diffMap[y*8+x+8] = averageIntensity > bottomBlockAverageIntensity
			}
		}
	}

	// Transform the block differences into a 64-bit integer hash
	var hash uint64
	for i := 0; i < 64; i++ {
		if diffMap[i] {
			hash |= 1 << uint(63-i)
		}
	}

	return hash, nil
}

func calculateBlockAverageIntensity(img image.Image, bounds image.Rectangle) uint8 {
	var totalIntensity uint64
	var pixelCount uint64

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			totalIntensity += uint64(gray.Y)
			pixelCount++
		}
	}

	return uint8(totalIntensity / pixelCount)
}
