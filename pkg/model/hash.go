package model

import (
	"fmt"
)

// Hash is a 64-bit unsigned integer that represents a hash.
type Hash uint64

// ToBinaryString returns the binary representation of the hash.
func (v Hash) ToBinaryString() string {
	return fmt.Sprintf("%064b", v)
}

// ToHexString returns the hexadecimal representation of the hash.
func (v Hash) ToHexString() string {
	return fmt.Sprintf("%016x", v)
}
