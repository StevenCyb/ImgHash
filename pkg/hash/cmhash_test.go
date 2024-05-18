package hash

import (
	"bytes"
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cmHash(t *testing.T) {
	t.Parallel()

	img, _, err := image.Decode(bytes.NewBuffer(testImage))
	assert.NoError(t, err)

	hash, err := CMHash(img)
	assert.NoError(t, err)

	expected := uint64(0xffffffffffefff88)
	assert.Equal(t, expected, hash)
}
