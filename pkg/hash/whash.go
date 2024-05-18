package hash

import (
	"image"
	"imghash/pkg/utils"
	"math"
	"sort"
)

// Wavelet Hashing (wHash).
func WHash(img image.Image) (uint64, error) {
	if img == nil {
		return 0, ErrEmptyImage
	}

	// Resize the image to a power of 2 size for the Haar wavelet transform
	img = utils.ResizeRGBInterArea(img, 64, 64)

	// Convert the image to grayscale
	grayImg := utils.ToGrayscale(img)

	// Perform the Haar wavelet transform
	coefficients := haarWaveletTransform(grayImg)

	// Extract the most significant coefficients for hashing
	significantCoefficients := extractSignificantCoefficients(coefficients)

	// Generate the hash from the significant coefficients
	var hash uint64
	for _, coefficient := range significantCoefficients {
		if coefficient > 0 {
			hash = (hash << 1) | 1
		} else {
			hash <<= 1
		}
	}

	return hash, nil
}

func haarWaveletTransform(img *image.Gray) [][]float64 {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	// Perform the Haar wavelet transform horizontally
	horizontalCoefficients := make([][]float64, height)
	for y := 0; y < height; y++ {
		horizontalCoefficients[y] = make([]float64, width)
		for x := 0; x < width; x += 2 {
			pixel1 := float64(img.GrayAt(x, y).Y)
			pixel2 := float64(img.GrayAt(x+1, y).Y)
			average := (pixel1 + pixel2) / 2.0
			difference := pixel1 - average
			horizontalCoefficients[y][x/2] = average
			horizontalCoefficients[y][width/2+x/2] = difference
		}
	}

	// Perform the Haar wavelet transform vertically
	coefficients := make([][]float64, height/2) // Change this line
	for y := 0; y < height; y += 2 {
		for x := 0; x < width; x++ {
			average := (horizontalCoefficients[y][x] + horizontalCoefficients[y+1][x]) / 2.0
			difference := horizontalCoefficients[y][x] - average
			coefficients[y/2] = append(coefficients[y/2], average)    // Append to existing slice
			coefficients[y/2] = append(coefficients[y/2], difference) // Append to existing slice
		}
	}

	return coefficients
}

func extractSignificantCoefficients(coefficients [][]float64) []float64 {
	// Flatten the coefficient matrix and sort in descending order of magnitude
	flattened := make([]float64, len(coefficients)*len(coefficients[0]))
	index := 0
	for _, row := range coefficients {
		for _, value := range row {
			flattened[index] = math.Abs(value)
			index++
		}
	}
	sort.Slice(flattened, func(i, j int) bool {
		return flattened[i] > flattened[j]
	})

	// Select the most significant coefficients (e.g., top 64 coefficients)
	numSignificant := 64
	if len(flattened) < numSignificant {
		numSignificant = len(flattened)
	}
	significantCoefficients := make([]float64, numSignificant)
	copy(significantCoefficients, flattened[:numSignificant])

	return significantCoefficients
}
