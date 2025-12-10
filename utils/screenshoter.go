package utils

import (
	"errors"
	"image"
	"log"

	"github.com/kbinani/screenshot"
)

var activeDisplay int = 0

func CaptureScreenshot() (image.Image, error) {
	displayBounds := screenshot.GetDisplayBounds(activeDisplay)
	screenshot, err := screenshot.CaptureRect(displayBounds)
	if err != nil {
		log.Println(err)
	}
	return screenshot, nil
}

func SetActiveDisplay(newActiveDisplay int) error {
	n := screenshot.NumActiveDisplays()
	if n < newActiveDisplay {
		return errors.New("incorrect display")
	}
	activeDisplay = newActiveDisplay
	return nil
}

func GetActiveDisplays() int {
	n := screenshot.NumActiveDisplays()
	return n
}
