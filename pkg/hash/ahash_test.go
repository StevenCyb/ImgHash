package hash

import (
	"bytes"
	"image"
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
)

//go:embed test_image.png
var testImage []byte

func Test_aHash(t *testing.T) {
	t.Parallel()

	img, _, err := image.Decode(bytes.NewBuffer(testImage))
	assert.NoError(t, err)

	hash, err := AHash(img)
	assert.NoError(t, err)

	expected := uint64(0xe5e7e7f3d00)
	assert.Equal(t, expected, hash)
}
