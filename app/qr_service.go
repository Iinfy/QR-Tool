package app

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"qrgen/utils"
	"time"

	"github.com/gen2brain/beeep"
	"golang.design/x/clipboard"
)

func UrlToQR() {
	urlBytes := clipboard.Read(clipboard.FmtText)
	url := string(urlBytes)
	filename := fmt.Sprintf("%d-%d-%d_%d-%d-%d",
		time.Now().Day(),
		time.Now().Month(),
		time.Now().Year(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second())
	filepath := "QRs/" + filename + ".png"
	utils.GenerateQR(url, filepath)
	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("Failed to decode image: %v", err)
	}
	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		log.Fatalf("Failed to encode image to PNG: %v", err)
	}
	pngData := buf.Bytes()
	clipboard.Write(clipboard.FmtImage, pngData)
	beeep.Notify("QR Generated", "QR Code successfully generated and copied into clipboard", filepath)
	fmt.Println("qr in clipboard")
}

func QRToUrl() {
	screenshot, err := utils.CaptureScreenshot()
	if err != nil {
		log.Println(err)
	}
	url, err := utils.ScanQRFromImage(screenshot)
	if err != nil {
		log.Println(err)
		beeep.Notify("Scan error", "No QR found", "qrlogo.ico")
		return
	}
	clipboard.Write(clipboard.FmtText, []byte(url))
	fmt.Println(url)
	err = beeep.Notify("QR scanned", "QR successfully scanned, url copied to clipboard", "qrlogo.ico")
	fmt.Println(err)

}
