package distance

import (
	"imghash/pkg/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HammingDistance(t *testing.T) {
	t.Parallel()

	t.Run("Max distance", func(t *testing.T) {
		t.Parallel()

		a := model.Hash(0b0000000000000000000000000000000000000000000000000000000000000000)
		b := model.Hash(0b1111111111111111111111111111111111111111111111111111111111111111)

		assert.Equal(t, 64, hammingDistance(a, b))
		assert.Equal(t, 0.0, HammingDistanceSimilarity(a, b))
	})

	t.Run("Min distance", func(t *testing.T) {
		t.Parallel()

		a := model.Hash(0b0000000000000000000000000000000000000000000000000000000000000000)
		b := model.Hash(0b0000000000000000000000000000000000000000000000000000000000000000)

		assert.Equal(t, 0, hammingDistance(a, b))
		assert.Equal(t, 1.0, HammingDistanceSimilarity(a, b))
	})

	t.Run("Between distance", func(t *testing.T) {
		t.Parallel()

		a := model.Hash(0b0000000000000000000000000000000000000000000000000000000000000000)
		b := model.Hash(0b1010101010101010101010101010101010101010101010101010101010101010)

		assert.Equal(t, 32, hammingDistance(a, b))
		assert.Equal(t, 0.5, HammingDistanceSimilarity(a, b))
	})
}
