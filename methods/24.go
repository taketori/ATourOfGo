package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewNRGBA(image.Rect(100, 100, 0, 0))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
