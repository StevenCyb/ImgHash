package gui

import (
	"image"

	"imghash/pkg/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type ImageStack struct {
	Image  image.Image
	Path   string
	Width  uint
	Height uint
}

type ImageCollection struct {
	Images []ImageStack
}

type ImageList struct {
	*container.Scroll
	view *fyne.Container
}

func NewImageList() *ImageList {
	view := container.NewVBox()

	imageList := &ImageList{
		Scroll: container.NewScroll(view),
		view:   view,
	}

	testCase(imageList)

	return imageList
}

func (il *ImageList) Clear() {
	il.view.RemoveAll()
	il.view.Refresh()
}

func (il *ImageList) Add(ic ImageCollection) {
	il.view.Refresh()
	il.ScrollToBottom()
}

func testCase(il *ImageList) { // TODO remove
	img, err := utils.ReadImage("test/6293768-smile-emoticon-icon-for-world-happiness-vector-template-design-illustration-vektor.jpg")
	if err != nil {
		panic(err)
	}

	il.Add(ImageCollection{
		Images: []ImageStack{
			{
				Image:  img,
				Path:   "test",
				Width:  100,
				Height: 100,
			},
			{
				Image:  img,
				Path:   "test2",
				Width:  200,
				Height: 200,
			},
		},
	})
}
