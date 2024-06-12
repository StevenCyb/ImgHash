package gui

import (
	"image"
	"imghash/pkg/hash"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type HashFunction = func(img image.Image) (uint64, error)

const defaultAlgorithm = "Perceptual Hashing (pHash)"

var hashAlgorithms = map[string]HashFunction{
	"Average Hashing (aHash)":       hash.AHash,
	"Block Hashing (bHash)":         hash.BHash,
	"Color Moment Hashing (cmHash)": hash.CMHash,
	"Difference Hashing (dHash)":    hash.DHash,
	"Median Hashing (mHash)":        hash.MHash,
	"Perceptual Hashing (pHash)": func(img image.Image) (uint64, error) {
		return hash.PHash(img, hash.HashSize32)
	},
	"Wavelet Hashing (wHash)": hash.WHash,
}

type Sidebar struct {
	*fyne.Container
	Algorithm       HashFunction
	Sensitivity     float64
	Directories     []string
	StartFunc       func()
	StopFunc        func()
	startStopButton *widget.Button
}

func NewSidebar(window fyne.Window, startFunc, stopFunc func()) *Sidebar {
	sidebar := &Sidebar{
		Algorithm: hashAlgorithms[defaultAlgorithm],
		StartFunc: startFunc,
		StopFunc:  stopFunc,
	}

	sidebar.startStopButton = widget.NewButton("Start", func() {
		if sidebar.startStopButton.Text == "Start" {
			sidebar.startStopButton.SetText("Stop")
			sidebar.StartFunc()
			return
		}

		sidebar.startStopButton.SetText("Start")
		sidebar.StopFunc()
	})

	algorithmOptions := make([]string, 0, len(hashAlgorithms))
	for k := range hashAlgorithms {
		algorithmOptions = append(algorithmOptions, k)
	}
	algorithmSelect := widget.NewSelect(algorithmOptions, func(value string) {
		sidebar.Algorithm = hashAlgorithms[value]
	})
	algorithmSelect.Selected = defaultAlgorithm

	sensitivity := widget.NewSlider(0, 1.0)
	sensitivity.Step = 0.1
	sensitivity.Value = 0.9
	sensitivity.OnChanged = func(newSensitivity float64) {
		sidebar.Sensitivity = newSensitivity
	}

	var dirList *widget.List
	dirList = widget.NewList(
		func() int {
			return len(sidebar.Directories)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("template", func() {})
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Button).SetText(sidebar.Directories[i])
			o.(*widget.Button).OnTapped = func() {
				sidebar.Directories = append(sidebar.Directories[:i], sidebar.Directories[i+1:]...)
				dirList.Refresh()
			}
		},
	)

	dirSelect := widget.NewButton("Add Paths", func() {
		dialogInput := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil || uri == nil {
				return
			}
			path := strings.TrimPrefix(uri.String(), "file://")
			sidebar.Directories = append(sidebar.Directories, path)
			dirList.Refresh()
		}, window)
		dialogInput.Show()
	})

	topContainer := container.NewVBox(
		sidebar.startStopButton,
		algorithmSelect,
		sensitivity,
		dirSelect,
	)
	sidebar.Container = container.NewBorder(
		topContainer,
		nil,
		nil,
		nil,
		dirList,
	)

	return sidebar
}

func (s *Sidebar) Stop() {
	s.startStopButton.SetText("Start")
	s.StopFunc()
}
