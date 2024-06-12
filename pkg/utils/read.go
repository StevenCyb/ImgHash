package utils

import (
	"errors"
	"image"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"

	"golang.org/x/image/webp"
)

var ErrUnsupportedFileFormat = errors.New("unsupported file format")

func ReadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var imageData image.Image
	lowerPath := strings.ToLower(path)
	suffix := func(s string) bool {
		return strings.HasSuffix(lowerPath, s)
	}

	switch {
	case suffix(".webp"):
		imageData, err = webp.Decode(file)
	case suffix(".png"), suffix(".jpg"), suffix(".jpeg"):
		imageData, _, err = image.Decode(file)
	case suffix(".gif"):
		imageData, err = gif.Decode(file)
	default:
		return nil, ErrUnsupportedFileFormat
	}

	if err != nil {
		return nil, err
	}

	return imageData, nil
}
