package hash

import (
	"bytes"
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mHash(t *testing.T) {
	t.Parallel()

	img, _, err := image.Decode(bytes.NewBuffer(testImage))
	assert.NoError(t, err)

	hash, err := MHash(img)
	assert.NoError(t, err)

	expected := uint64(0x200e5f7eff7f01)
	assert.Equal(t, expected, hash)
}
