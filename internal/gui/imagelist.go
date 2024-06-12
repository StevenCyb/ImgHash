package gui

import (
	"fmt"
	"net/url"
	"os"

	"imghash/internal/types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type ImageList struct {
	*container.Scroll
	view   *fyne.Container
	window fyne.Window
}

func NewImageList(window fyne.Window) *ImageList {
	view := container.NewVBox()

	imageList := &ImageList{
		Scroll: container.NewScroll(view),
		view:   view,
		window: window,
	}

	return imageList
}

func (il *ImageList) Clear() {
	il.view.RemoveAll()
	il.view.Refresh()
}

func (il *ImageList) Add(ic types.ImageCollection) {
	ic.Sort()
	var itemRow *fyne.Container

	skipButton := widget.NewButton("Skip", func() {
		il.view.Remove(itemRow)
	})
	keepBestButton := widget.NewButton("Keep best", func() {
		deleteCount := len(ic.Images) - 1
		plural := ""

		if deleteCount > 1 {
			plural = "s"
		}

		dialog.NewConfirm("Keep best",
			fmt.Sprintf("Are you sure you want to delete %d image%s?", deleteCount, plural),
			func(b bool) {
				if !b {
					return
				}

				for i := 1; i < len(ic.Images); i++ {
					os.Remove(ic.Images[i].Path)
				}

				il.view.Remove(itemRow)
			}, il.window).Show()
	})

	buttonContainer := container.NewVBox(skipButton, keepBestButton)

	imagesContainer := container.NewHBox()
	for _, image := range ic.Images {
		imageContainer := container.NewVBox()

		img := &canvas.Image{Image: image.Image}
		img.FillMode = canvas.ImageFillOriginal

		displayPath := image.Path
		if len(displayPath) > 30 {
			displayPath = displayPath[:15] + "..." + displayPath[len(displayPath)-15:]
		}

		sizeLabel := widget.NewLabel(fmt.Sprintf("%dx%d", image.Width, image.Height))
		sizeLabel.Alignment = fyne.TextAlignCenter
		pathLabel := widget.NewRichText(&widget.HyperlinkSegment{
			Text:      displayPath,
			URL:       &url.URL{Path: image.Path},
			Alignment: fyne.TextAlignCenter,
		})

		deleteButton := widget.NewButton("Delete", func() {
			dialog.NewConfirm("Keep best",
				fmt.Sprintf("Are you sure you want to delete %s?", image.Path),
				func(b bool) {
					if !b {
						return
					}

					os.Remove(image.Path)
					if len(imagesContainer.Objects) <= 2 {
						il.view.Remove(itemRow)
					} else {
						imagesContainer.Remove(imageContainer)
					}
				}, il.window).Show()
		})

		imageContainer.Add(img)
		imageContainer.Add(deleteButton)
		imageContainer.Add(sizeLabel)
		imageContainer.Add(pathLabel)
		imagesContainer.Add(imageContainer)
	}

	itemRow = container.NewHBox(buttonContainer, imagesContainer)
	itemRow = container.NewVBox(itemRow, widget.NewSeparator())

	il.view.Add(itemRow)
	il.view.Refresh()
}
