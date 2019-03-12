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
