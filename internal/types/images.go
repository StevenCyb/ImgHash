package types

import (
	"image"
	"imghash/pkg/model"
	"sort"
)

type ImageStack struct {
	Image   image.Image
	Path    string
	Width   uint
	Height  uint
	Hash    model.Hash
	Matched bool
}

type ImageCollection struct {
	Images []ImageStack
}

func (ic ImageCollection) Sort() {
	sort.Slice(ic.Images, func(i, j int) bool {
		return (ic.Images[i].Width * ic.Images[i].Height) > (ic.Images[j].Width * ic.Images[j].Height)
	})
}
