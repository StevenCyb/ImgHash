package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"imghash/pkg/distance"
	"imghash/pkg/hash"
	"imghash/pkg/model"
	"log"
	"os"
	"strings"

	"golang.org/x/image/webp"

	"github.com/urfave/cli/v2"
)

var ErrUnsupportedFileFormat = fmt.Errorf("unsupported file format")

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "get hash of given image",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "format",
						Aliases: []string{"f"},
						Value:   "hex",
						Usage:   "Output format (hex, binary), default is hex",
					},
				},
				Subcommands: []*cli.Command{
					{
						Name:    "ahash",
						Aliases: []string{"a"},
						Usage:   "use Average Hashing (aHash) algorithm",
						Action: func(cCtx *cli.Context) error {
							hash, err := HashWith(hash.AHash, cCtx.Args().First(), cCtx.String("format"))
							if err == nil {
								fmt.Print(hash)
							}

							return err
						},
					},
					{
						Name:    "bhash",
						Aliases: []string{"b"},
						Usage:   "use Block Hashing (bHash) algorithm",
						Action: func(cCtx *cli.Context) error {
							hash, err := HashWith(hash.BHash, cCtx.Args().First(), cCtx.String("format"))
							if err == nil {
								fmt.Print(hash)
							}

							return err
						},
					},
					{
						Name:    "cmhash",
						Aliases: []string{"c"},
						Usage:   "use Color Moment Hashing (cmHash) algorithm",
						Action: func(cCtx *cli.Context) error {
							hash, err := HashWith(hash.CMHash, cCtx.Args().First(), cCtx.String("format"))
							if err == nil {
								fmt.Print(hash)
							}

							return err
						},
					},
					{
						Name:    "dhash",
						Aliases: []string{"d"},
						Usage:   "use Difference Hashing (dHash) algorithm",
						Action: func(cCtx *cli.Context) error {
							hash, err := HashWith(hash.DHash, cCtx.Args().First(), cCtx.String("format"))
							if err == nil {
								fmt.Print(hash)
							}

							return err
						},
					},
					{
						Name:    "mhash",
						Aliases: []string{"m"},
						Usage:   "use Median Hashing (mHash) algorithm",
						Action: func(cCtx *cli.Context) error {
							hash, err := HashWith(hash.MHash, cCtx.Args().First(), cCtx.String("format"))
							if err == nil {
								fmt.Print(hash)
							}

							return err
						},
					},
					{
						Name:    "phash",
						Aliases: []string{"p"},
						Usage:   "use Perceptual Hashing (pHash) algorithm",
						Action: func(cCtx *cli.Context) error {
							wrapped := func() func(img image.Image) (uint64, error) {
								return func(img image.Image) (uint64, error) {
									return hash.PHash(img, hash.HashSize32)
								}
							}

							hash, err := HashWith(wrapped(), cCtx.Args().First(), cCtx.String("format"))
							if err == nil {
								fmt.Print(hash)
							}

							return err
						},
					},
					{
						Name:    "whash",
						Aliases: []string{"w"},
						Usage:   "use Wavelet Hashing (wHash) algorithm",
						Action: func(cCtx *cli.Context) error {
							hash, err := HashWith(hash.WHash, cCtx.Args().First(), cCtx.String("format"))
							if err == nil {
								fmt.Print(hash)
							}

							return err
						},
					},
				},
			},
			{
				Name:    "compare",
				Aliases: []string{"c"},
				Usage:   "compare two images",
				Flags: []cli.Flag{
					&cli.UintFlag{
						Name:    "format",
						Aliases: []string{"f"},
						Value:   2,
						Usage:   "Rounding decimal to decimal places, default 2",
					},
				},
				Subcommands: []*cli.Command{
					{
						Name:    "ahash",
						Aliases: []string{"a"},
						Usage:   "use Average Hashing (aHash) algorithm",
						Action: func(cCtx *cli.Context) error {
							imageA := cCtx.Args().Get(0)
							imageB := cCtx.Args().Get(1)

							score, err := CompareWith(hash.AHash, imageA, imageB, cCtx.Uint("format"))
							if err == nil {
								fmt.Print(score)
							}

							return err
						},
					},
					{
						Name:    "bhash",
						Aliases: []string{"b"},
						Usage:   "use Block Hashing (bHash) algorithm",
						Action: func(cCtx *cli.Context) error {
							imageA := cCtx.Args().Get(0)
							imageB := cCtx.Args().Get(1)

							score, err := CompareWith(hash.BHash, imageA, imageB, cCtx.Uint("format"))
							if err == nil {
								fmt.Print(score)
							}

							return err
						},
					},
					{
						Name:    "cmhash",
						Aliases: []string{"c"},
						Usage:   "use Color Moment Hashing (cmHash) algorithm",
						Action: func(cCtx *cli.Context) error {
							imageA := cCtx.Args().Get(0)
							imageB := cCtx.Args().Get(1)

							score, err := CompareWith(hash.CMHash, imageA, imageB, cCtx.Uint("format"))
							if err == nil {
								fmt.Print(score)
							}

							return err
						},
					},
					{
						Name:    "dhash",
						Aliases: []string{"d"},
						Usage:   "use Difference Hashing (dHash) algorithm",
						Action: func(cCtx *cli.Context) error {
							imageA := cCtx.Args().Get(0)
							imageB := cCtx.Args().Get(1)

							score, err := CompareWith(hash.DHash, imageA, imageB, cCtx.Uint("format"))
							if err == nil {
								fmt.Print(score)
							}

							return err
						},
					},
					{
						Name:    "mhash",
						Aliases: []string{"m"},
						Usage:   "use Median Hashing (mHash) algorithm",
						Action: func(cCtx *cli.Context) error {
							imageA := cCtx.Args().Get(0)
							imageB := cCtx.Args().Get(1)

							score, err := CompareWith(hash.MHash, imageA, imageB, cCtx.Uint("format"))
							if err == nil {
								fmt.Print(score)
							}

							return err
						},
					},
					{
						Name:    "phash",
						Aliases: []string{"p"},
						Usage:   "use Perceptual Hashing (pHash) algorithm",
						Action: func(cCtx *cli.Context) error {
							wrapped := func() func(img image.Image) (uint64, error) {
								return func(img image.Image) (uint64, error) {
									return hash.PHash(img, hash.HashSize32)
								}
							}

							imageA := cCtx.Args().Get(0)
							imageB := cCtx.Args().Get(1)

							score, err := CompareWith(wrapped(), imageA, imageB, cCtx.Uint("format"))
							if err == nil {
								fmt.Print(score)
							}

							return err
						},
					},
					{
						Name:    "whash",
						Aliases: []string{"w"},
						Usage:   "use Wavelet Hashing (wHash) algorithm",
						Action: func(cCtx *cli.Context) error {
							imageA := cCtx.Args().Get(0)
							imageB := cCtx.Args().Get(1)

							score, err := CompareWith(hash.WHash, imageA, imageB, cCtx.Uint("format"))
							if err == nil {
								fmt.Print(score)
							}

							return err
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

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

func HashWith(hashFunc func(img image.Image) (uint64, error), imagePath string, format string) (string, error) {
	imageData, err := ReadImage(imagePath)
	if err != nil {
		return "", err
	}

	h, err := hashFunc(imageData)
	if err != nil {
		return "", err
	}

	hash := model.Hash(h)
	if format == "binary" {
		return hash.ToBinaryString(), nil
	}

	return hash.ToHexString(), nil
}

func CompareWith(hashFunc func(img image.Image) (uint64, error), imageAPath string, imageBPath string, format uint) (string, error) {
	imageAData, err := ReadImage(imageAPath)
	if err != nil {
		return "", err
	}

	imageBData, err := ReadImage(imageBPath)
	if err != nil {
		return "", err
	}

	hashA, err := hashFunc(imageAData)
	if err != nil {
		return "", err
	}

	hashB, err := hashFunc(imageBData)
	if err != nil {
		return "", err
	}

	similarityScore := distance.HammingDistanceSimilarity(model.Hash(hashA), model.Hash(hashB))

	return fmt.Sprintf("%"+fmt.Sprintf(".%df", format), similarityScore), nil
}
