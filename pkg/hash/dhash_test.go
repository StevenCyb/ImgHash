package hash

import (
	"bytes"
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dHash(t *testing.T) {
	t.Parallel()

	img, _, err := image.Decode(bytes.NewBuffer(testImage))
	assert.NoError(t, err)

	hash, err := DHash(img)
	assert.NoError(t, err)

	expected := uint64(0x150767cf878b030b)
	assert.Equal(t, expected, hash)
}
