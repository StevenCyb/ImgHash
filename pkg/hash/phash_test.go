package hash

import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pHash(t *testing.T) {
	t.Parallel()

	img, _, err := image.Decode(bytes.NewBuffer(testImage))
	assert.NoError(t, err)

	t.Run("HashSize8", func(t *testing.T) {
		hash, err := PHash(img, HashSize8)
		assert.NoError(t, err)

		expected := uint64(0x8000208000000000)
		assert.Equal(t, expected, hash)
	})

	t.Run("HashSize16", func(t *testing.T) {
		hash, err := PHash(img, HashSize16)
		assert.NoError(t, err)

		expected := uint64(0x8040628000000000)
		assert.Equal(t, expected, hash)
	})

	t.Run("HashSize32", func(t *testing.T) {
		hash, err := PHash(img, HashSize32)
		assert.NoError(t, err)

		expected := uint64(0x90606a8000089001)
		assert.Equal(t, expected, hash)
	})

	t.Run("HashSize64", func(t *testing.T) {
		hash, err := PHash(img, HashSize64)
		assert.NoError(t, err)

		expected := uint64(0x90606a8444189541)
		assert.Equal(t, expected, hash)
	})
}
