package main

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	_ = image.YCbCr{}
	_ = color.Alpha{}
}
