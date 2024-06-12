package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Progress struct {
	*fyne.Container
	progress *widget.ProgressBar
}

func NewProgress() *Progress {
	progressContainer := container.New(layout.NewVBoxLayout())
	progress := widget.NewProgressBar()
	progressContainer.Add(progress)

	return &Progress{
		Container: progressContainer,
		progress:  progress,
	}
}

func (p *Progress) Reset() {
	p.progress.TextFormatter = func() string {
		return ""
		}
	p.progress.SetValue(0)
}

func (p *Progress) Init(min, max float64) {
	p.progress.Min = min
	p.progress.Max = max
}

func (p *Progress) SetProgress(stage string, value float64) {
	p.progress.TextFormatter = func() string {
		return fmt.Sprintf("%s: %.0f%%", stage, value)
	}
	p.progress.SetValue(value)
}
