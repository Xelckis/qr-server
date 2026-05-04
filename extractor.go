package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/png"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func main() {
	imgFlag := flag.String("in", "", "Path to the QR Code image")
	outFlag := flag.String("out", "webserver_final", "Executable file name")
	flag.Parse()

	if *imgFlag == "" {
		fmt.Println("[-] Error: Input image is required.")
		os.Exit(1)
	}

	fmt.Println("[*] Initializing Gambiarra Labs optical scanner...")

	file, err := os.Open(*imgFlag)
	if err != nil {
		fmt.Printf("[-] Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("[-] Error decoding PNG:", err)
		os.Exit(1)
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		fmt.Println("[-] Error in bitmap:", err)
		os.Exit(1)
	}

	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		fmt.Println("[-] Error decoding QR Code:", err)
		os.Exit(1)
	}

	var rawBinaryData []byte
	meta := result.GetResultMetadata()

	if segments, ok := meta[gozxing.ResultMetadataType_BYTE_SEGMENTS].([][]byte); ok && len(segments) > 0 {
		for _, seg := range segments {
			rawBinaryData = append(rawBinaryData, seg...)
		}
	} else {
		fmt.Println("[-] Failure: The QR Code does not contain byte segments in RAW mode.")
		os.Exit(1)
	}

	err = os.WriteFile(*outFlag, rawBinaryData, 0755)
	if err != nil {
		fmt.Println("[-] Error writing file:", err)
		os.Exit(1)
	}

	fmt.Printf("[+] OVERENGINEERING COMPLETE!\n")
	fmt.Printf("[+] %d bytes of pure Assembly extracted without corruption.\n", len(rawBinaryData))
	fmt.Printf("[+] Start the server: ./%s\n", *outFlag)
}
