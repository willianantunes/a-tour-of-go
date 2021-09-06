package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

/*
Exercise: Images
Remember the picture generator[1] you wrote earlier? Let's write another one, but this time it will return an implementation of image.Image instead of a slice of data.

Define your own Image type, implement the necessary methods[2], and call pic.ShowImage.

Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).

ColorModel should return color.RGBAModel.

At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.

[1] https://tour.golang.org/moretypes/18
[2] https://pkg.go.dev/image#Image
*/

type Image struct{ width, height int }

func (customImage Image) ColorModel() color.Model {
	// ColorModel should return color.RGBAModel.
	return color.RGBAModel
}

func (customImage Image) Bounds() image.Rectangle {
	// Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
	return image.Rect(0, 0, customImage.width, customImage.height)
}

func (customImage Image) At(x, y int) color.Color {
	// fmt.Println(x, y)
	// At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

// You should run this code at: https://tour.golang.org/methods/25
func main() {
	m := Image{300, 50}
	pic.ShowImage(m)
}
