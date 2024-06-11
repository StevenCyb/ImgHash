package utils

import (
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"

	"golang.org/x/image/webp"
)

var ErrUnsupportedFileFormat = errors.New("unsupported file format")

func ReadImage(path string) (image.Image, error) {
	imageFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer imageFile.Close()

	lowerPath := strings.ToLower(path)
	if strings.HasSuffix(lowerPath, ".webp") {
		imageData, err := webp.Decode(imageFile)
		if err != nil {
			return nil, err
		}

		return imageData, nil
	} else if strings.HasSuffix(lowerPath, ".png") || strings.HasSuffix(lowerPath, ".jpg") || strings.HasSuffix(lowerPath, ".jpeg") {
		imageData, _, err := image.Decode(imageFile)
		if err != nil {
			return nil, err
		}

		return imageData, nil
	}

	return nil, ErrUnsupportedFileFormat
}
