package hash

import "errors"

type HashSize byte

var (
	ErrInvalidHashSize = errors.New("invalid hash size")
	ErrEmptyImage      = errors.New("empty image")
)

const (
	HashSize8  HashSize = 8
	HashSize16 HashSize = 16
	HashSize32 HashSize = 32
	HashSize64 HashSize = 64
)
