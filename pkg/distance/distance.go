package distance

import (
	"imghash/pkg/model"
	"math/bits"
)

// hammingDistance returns the Hamming distance between two hashes.
func hammingDistance(a, b model.Hash) int {
	xor := a ^ b

	return bits.OnesCount64(uint64(xor))
}

// HammingDistanceSimilarity returns the similarity between two hashes using the Hamming distance.
func HammingDistanceSimilarity(a, b model.Hash) float64 {
	return (64.0 - float64(hammingDistance(a, b))) / 64.0
}
