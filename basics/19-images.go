package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Define a custom Image type with width and height
type Image struct {
	w, h int
}

// To implement image.Image, we need to implement the following methods:
//   - Bounds - Return the dimensions of the image
//   - ColorModel - Return the color model (RGBA)
//   - At - Return the color at the given x, y pixel

// Implement the Bounds method
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

// Implement the ColorModel method
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Implement the At method
func (img Image) At(x, y int) color.Color {
	v := uint8(x ^ y) // same as previous picture generator
	return color.RGBA{v, v, 255, 255}
}

func images() {
	m := Image{w: 256, h: 256} // Set the image size
	pic.ShowImage(m)
}
