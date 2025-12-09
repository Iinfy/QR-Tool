package utils

import (
	"fmt"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func GenerateQR(url, filename string) {
	qrc, err := qrcode.New(url)
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
		return
	}

	w, err := standard.New(filename)
	if err != nil {
		fmt.Printf("standard.New failed: %v", err)
		return
	}

	// save file
	if err = qrc.Save(w); err != nil {
		fmt.Printf("could not save image: %v", err)
	}
}
