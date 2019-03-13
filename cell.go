//This is free and unencumbered software released into the public domain.

package term

import (
	"image"
	"image/color"
)

type Cell struct {
	R rune
	C color.Color
}

var EmptyCell = Cell{R: ' ', C: image.White}
