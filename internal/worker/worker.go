package worker

import (
	"errors"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"sync"

	"imghash/internal/types"
	"imghash/pkg/distance"
	"imghash/pkg/model"
	"imghash/pkg/utils"
)

type HashFunction = func(img image.Image) (uint64, error)
type CompleteHandler = func()
type ProgressHandler = func(label string, progress float64)
type MatchHandler = func(ic types.ImageCollection)
type ErrorHandler = func(err error)

type Worker struct {
	hashFunc    HashFunction
	sensitivity float64
	directories []string
	stopCh      chan bool
	flagLock    sync.Mutex
	running     bool
	onReady     CompleteHandler
	onProgress  ProgressHandler
	onMatch     MatchHandler
	onError     ErrorHandler
}

func New() *Worker {
	return &Worker{
		hashFunc:    nil,
		sensitivity: 1.0,
		directories: []string{},
		stopCh:      make(chan bool),
	}
}

func (w *Worker) OnReady(handler CompleteHandler) {
	w.onReady = handler
}

func (w *Worker) OnProgress(handler ProgressHandler) {
	w.onProgress = handler
}

func (w *Worker) OnMatch(handler MatchHandler) {
	w.onMatch = handler
}

func (w *Worker) OnError(handler ErrorHandler) {
	w.onError = handler
}

func (w *Worker) Run(hashFunc HashFunction, sensitivity float64, directories ...string) {
	w.flagLock.Lock()
	if w.running {
		w.flagLock.Unlock()
		return
	}
	w.flagLock.Unlock()

	if len(directories) == 0 {
		w.onReady()
		return
	}

	w.running = true
	w.stopCh = make(chan bool)
	w.hashFunc = hashFunc
	w.sensitivity = sensitivity
	w.directories = directories

	go func() {
		imageStacks := []types.ImageStack{}
		imageCollection := map[int]types.ImageCollection{}
		totalDirectories := len(directories)
		totalImages := 0
		totalComparisons := 0
		stage := byte(0)
		primaryIndex := 0
		secondaryIndex := 0
		tertiaryIndex := 0

		w.onProgress("Check directories...", 0)
		for {
			select {
			case <-w.stopCh:
				w.running = false
				return
			default:
				switch stage {
				case 0:
					files, err := os.ReadDir(directories[primaryIndex])
					if err != nil {
						w.onError(err)
						w.running = false
						return
					}

					for i, file := range files {
						if !file.IsDir() {
							absPath := filepath.Join(directories[primaryIndex], file.Name())
							img, err := utils.ReadImage(absPath)
							if errors.Is(err, utils.ErrUnsupportedFileFormat) {
								fmt.Println("Unsupported file format:", absPath)
								continue
							} else if err != nil {
								w.onError(err)
								continue
							}
							w.onProgress(fmt.Sprintf("Check directories (found %d images)...", i+1), float64(primaryIndex)/float64(totalDirectories)*100)

							width, height := img.Bounds().Dx(), img.Bounds().Dy()
							img = utils.ResizeRGBInterAreaWithRatio(img, 200)
							imageStacks = append(imageStacks, types.ImageStack{Path: absPath, Image: img, Width: uint(width), Height: uint(height)})
						}
					}
					w.onProgress("Check directories...", float64(primaryIndex+1)/float64(totalDirectories)*100)
					primaryIndex++
					if primaryIndex >= totalDirectories {
						stage++
						totalImages = len(imageStacks)
						primaryIndex = 0
						w.onProgress("Calculate hashes...", 0)
					}
				case 1:
					hash, err := w.hashFunc(imageStacks[primaryIndex].Image)
					if err != nil {
						w.onError(err)
						w.running = false
						return
					}
					imageStacks[primaryIndex].Hash = model.Hash(hash)

					w.onProgress("Calculate hashes...", float64(primaryIndex+1)/float64(totalImages)*100)
					primaryIndex++
					if primaryIndex >= totalImages {
						stage++
						primaryIndex = 0
						totalComparisons = (totalImages * (totalImages - 1)) / 2
						w.onProgress("Compare hashes...", 0)
					}
				case 2:
					if secondaryIndex+1 >= totalImages && primaryIndex+1 < totalImages {
						if ic, exists := imageCollection[primaryIndex]; exists {
							w.onMatch(ic)
						}
						primaryIndex++
						secondaryIndex = primaryIndex
						continue
					} else if primaryIndex < totalImages && imageStacks[primaryIndex].Matched {
						primaryIndex++
						secondaryIndex = primaryIndex
						tertiaryIndex += totalImages - primaryIndex
						w.onProgress("Compare hashes...", float64(tertiaryIndex)/float64(totalComparisons)*100)
						continue
					}
					if tertiaryIndex >= totalComparisons {
						stage++
						continue
					}
					secondaryIndex++
					tertiaryIndex++

					if distance.HammingDistanceSimilarity(imageStacks[primaryIndex].Hash, imageStacks[secondaryIndex].Hash) >= w.sensitivity {
						if _, exists := imageCollection[primaryIndex]; !exists {
							imageCollection[primaryIndex] = types.ImageCollection{
								Images: []types.ImageStack{
									imageStacks[primaryIndex],
									imageStacks[secondaryIndex],
								},
							}
							imageStacks[secondaryIndex].Matched = true
						} else {
							ic := imageCollection[primaryIndex]
							ic.Images = append(ic.Images, imageStacks[secondaryIndex])
							imageCollection[primaryIndex] = ic
							imageStacks[secondaryIndex].Matched = true
						}
					}

					w.onProgress("Compare hashes...", float64(tertiaryIndex)/float64(totalComparisons)*100)
				default:
					w.running = false
					w.onReady()
					return
				}
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.flagLock.Lock()
	defer w.flagLock.Unlock()
	if !w.running {
		return
	}

	w.stopCh <- true
	w.running = false
}
