package hash

import (
	"image"
	"imghash/pkg/utils"
)

// Perceptual Hashing (pHash) function based on the DCT algorithm.
func PHash(img image.Image, hashSize HashSize) (uint64, error) {
	if img == nil {
		return 0, ErrEmptyImage
	} else if hashSize != HashSize8 && hashSize != HashSize16 && hashSize != HashSize32 && hashSize != HashSize64 {
		return 0, ErrInvalidHashSize
	}

	size := int(hashSize)

	// Resize the image
	img = utils.ResizeRGBInterArea(img, size, size)

	// Convert the image to grayscale
	grayImg := utils.ToGrayscale(img)

	// Apply the forward DCT
	dct, err := utils.ForwardDCT(*grayImg)
	if err != nil {
		return 0, err
	}

	// Calculate the mean for the top left block (8x8)
	mean := float64(0)
	lowestFrequenciesBlock := make([][]float64, 8)
	for i := 0; i < 8; i++ {
		lowestFrequenciesBlock[i] = make([]float64, 8)
		for j := 0; j < 8; j++ {
			lowestFrequenciesBlock[i][j] = dct[i][j]
			mean += dct[i][j]
		}
	}
	mean /= float64(size)

	// Create a binary representation of the top left block
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if dct[i][j] > mean {
				lowestFrequenciesBlock[i][j] = 1
			} else {
				lowestFrequenciesBlock[i][j] = 0
			}
		}
	}

	// Convert the binary representation to an uint64
	var hash uint64
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			hash |= uint64(lowestFrequenciesBlock[i][j]) << uint(63-(i*8+j))
		}
	}

	return hash, nil
}
