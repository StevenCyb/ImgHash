package hash

import (
	"bytes"
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bHash(t *testing.T) {
	t.Parallel()

	img, _, err := image.Decode(bytes.NewBuffer(testImage))
	assert.NoError(t, err)

	hash, err := BHash(img)
	assert.NoError(t, err)

	expected := uint64(0x92025a081bd9f024)
	assert.Equal(t, expected, hash)
}
