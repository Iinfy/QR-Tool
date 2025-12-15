package utils

import (
	"errors"
	"image"
	"log"

	"github.com/kbinani/screenshot"
)

var activeDisplay int

func CaptureScreenshot() (image.Image, error) {
	displayBounds := screenshot.GetDisplayBounds(activeDisplay)
	screenshot, err := screenshot.CaptureRect(displayBounds)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return screenshot, nil
}

func SetActiveDisplay(newActiveDisplay int) error {
	n := screenshot.NumActiveDisplays()
	if n < newActiveDisplay {
		return errors.New("incorrect display")
	}
	activeDisplay = newActiveDisplay - 1
	return nil
}

func GetActiveDisplays() int {
	n := screenshot.NumActiveDisplays()
	return n
}

func IsMainDisplay(display int) bool {
	if display-1 == activeDisplay {
		return true
	} else {
		return false
	}
}
