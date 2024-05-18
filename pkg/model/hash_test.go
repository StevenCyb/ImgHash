package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HashToString(t *testing.T) {
	t.Parallel()

	t.Run("ToBinaryString", func(t *testing.T) {
		t.Parallel()

		v := Hash(0b1010101010101010101010101010101010101010101010101010101010101010)
		expect := "1010101010101010101010101010101010101010101010101010101010101010"
		actual := v.ToBinaryString()

		assert.Equal(t, expect, actual)
	})

	t.Run("ToHexString", func(t *testing.T) {
		t.Parallel()

		v := Hash(0b1010101010101010101010101010101010101010101010101010101010101010)
		expect := "aaaaaaaaaaaaaaaa"
		actual := v.ToHexString()

		assert.Equal(t, expect, actual)
	})
}
