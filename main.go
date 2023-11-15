package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	// Open an image file
	file, err := os.Open("./images/ascii_image.jpeg")
	if err != nil {
		fmt.Println("Error opening image:", err)
		return
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	// Get image bounds
	bounds := img.Bounds()

	width := bounds.Dx()
	height := bounds.Dy()

	fmt.Println("Successifully loaded image!")
	fmt.Printf("Width is %d and height is %d", width, height)
}
