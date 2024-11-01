package main

import (
	"fmt"
	"log"

	"github.com/skip2/go-qrcode"
)

func main() {
	url := "https://www.linkedin.com/in/dfigueredo/"
	filename := "qrcode.png"

	err := qrcode.WriteFile(url, qrcode.High, 512, filename)
	if err != nil {
		log.Fatalf("Failed to generate QR code: %v", err)
	}

	fmt.Printf("QR code saved as %s", filename)
}
