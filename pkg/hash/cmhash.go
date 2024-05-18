package hash

import (
	"image"
	"imghash/pkg/utils"
	"math"
)

// ColorMoment Hashing (cmHash).
func CMHash(img image.Image) (uint64, error) {
	if img == nil {
		return 0, ErrEmptyImage
	}

	// Resize the image
	img = utils.ResizeRGBInterArea(img, 64, 64)

	rgba := utils.ToRGBA(img)

	// Calculate color moments for each channel (R, G, B)
	moments := make([][3]float64, 3) // One set of moments per channel
	for c := 0; c < 3; c++ {         // Loop over channels: 0=R, 1=G, 2=B
		moments[c] = calculateColorMoments(rgba, c)
	}

	// Combine the moments into a single hash
	var hash uint64
	for _, channelMoments := range moments {
		for i := 0; i < 3; i++ { // Loop over moments (mean, variance, skewness)
			hash <<= 8
			hash |= uint64(channelMoments[i] * 255) // Scale to fit into 8 bits
		}
	}

	return hash, nil
}

func calculateColorMoments(img *image.RGBA, channel int) [3]float64 {
	var mean, variance, skewness float64

	// Calculate mean
	var total float64
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			var value uint32
			switch channel {
			case 0:
				value = r
			case 1:
				value = g
			case 2:
				value = b
			}
			total += float64(value)
		}
	}
	pixels := float64(bounds.Dx() * bounds.Dy())
	mean = total / pixels

	// Calculate variance
	total = 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			var value uint32
			switch channel {
			case 0:
				value = r
			case 1:
				value = g
			case 2:
				value = b
			}
			deviation := float64(value) - mean
			total += deviation * deviation
		}
	}
	variance = total / pixels

	// Calculate skewness
	total = 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			var value uint32
			switch channel {
			case 0:
				value = r
			case 1:
				value = g
			case 2:
				value = b
			}
			deviation := float64(value) - mean
			total += deviation * deviation * deviation
		}
	}
	cubeRootPixels := math.Pow(pixels, 1.0/3.0)
	skewness = total / (cubeRootPixels * math.Sqrt(variance) * variance)

	return [3]float64{mean, variance, skewness}
}
