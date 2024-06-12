package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"

	"imghash/internal/worker"
)

type GUI struct {
	app    fyne.App
	window fyne.Window
	worker *worker.Worker
}

func New() *GUI {
	app := app.New()
	window := app.NewWindow("ImgHash")
	window.Resize(fyne.NewSize(800, 600))

	return &GUI{
		app:    app,
		window: window,
		worker: worker.New(),
	}
}

func (u *GUI) Run() *GUI {
	progressContainer := NewProgress()
	progressContainer.Reset()
	progressContainer.Init(0, 100)

	imageContainer := NewImageList(u.window)
	var sidebarContainer *Sidebar
	sidebarContainer = NewSidebar(u.window,
		func() {
			imageContainer.Clear()
			u.worker.Run(
				sidebarContainer.HashFunc,
				sidebarContainer.Sensitivity,
				sidebarContainer.Directories...,
			)
		},
		func() {
			u.worker.Stop()
			progressContainer.Reset()
		})
	u.worker.OnReady(func() {
		sidebarContainer.Stop()
		progressContainer.Reset()
	})
	u.worker.OnProgress(progressContainer.SetProgress)
	u.worker.OnMatch(imageContainer.Add)
	u.worker.OnError(func(err error) {
		dialog.NewError(err, u.window).Show()
	})

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
