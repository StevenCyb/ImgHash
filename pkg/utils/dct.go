package utils

import (
	"fmt"
	"image"
	"math"
)

// ForwardDCT computes the forward DCT-II of a block of size N x N
func ForwardDCT(img image.Gray) ([][]float64, error) {
	numRows := img.Bounds().Dy()
	numCols := img.Bounds().Dx()

	if numRows != numCols {
		return nil, fmt.Errorf("input image is not even")
	}

	N := numRows
	coef := make([][]float64, N)

	for u := 0; u < N; u++ {
		coef[u] = make([]float64, N)
		for v := 0; v < N; v++ {
			var sum float64
			for x := 0; x < N; x++ {
				for y := 0; y < N; y++ {
					value := float64(img.GrayAt(x, y).Y)
					sum += value * math.Cos((math.Pi/float64(N))*(float64(x)+0.5)*float64(u)) *
						math.Cos((math.Pi/float64(N))*(float64(y)+0.5)*float64(v))
				}
			}
			Cu := 1.0
			Cv := 1.0
			if u == 0 {
				Cu = 1 / math.Sqrt(2)
			}
			if v == 0 {
				Cv = 1 / math.Sqrt(2)
			}
			coef[u][v] = 0.25 * Cu * Cv * sum
		}
	}

	return coef, nil
}
