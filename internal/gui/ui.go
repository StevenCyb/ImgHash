package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

type GUI struct {
	app    fyne.App
	window fyne.Window
}

func New() *GUI {
	app := app.New()
	window := app.NewWindow("ImgHash")
	window.Resize(fyne.NewSize(800, 600))

	return &GUI{
		app:    app,
		window: window,
	}
}

func (u *GUI) Run() *GUI {
	progressContainer := NewProgress()
	progressContainer.Reset()

	sidebarContainer := NewSidebar(u.window,
		func() {
			progressContainer.Init(0, 100)
			// TODO worker
		},
		func() {
			progressContainer.Reset()
			// TODO worker
		})

	imageContainer := NewImageList()

	splitContainer := container.NewHSplit(sidebarContainer.Container, imageContainer.Scroll)
	splitContainer.SetOffset(0.3125)

	u.window.SetContent(container.NewBorder(progressContainer.Container, nil, nil, nil, splitContainer))

	u.window.SetOnClosed(func() {
		sidebarContainer.Stop()
	})

	return u
}

func (g *GUI) ShowAndRun() {
	g.window.ShowAndRun()
}
