package term

import "image"

//Dimension is a pair of width(W) and height(H)
type Dimension struct {
	W, H int
}

//Dim is shorthand for Dimension{w, h}.
func Dim(w, h int) Dimension {
	return Dimension{w, h}
}

//A Point is an X, Y coordinate pair. The axes increase right and down.
type Point image.Point

//ZP is the zero Point.
var ZP Point

//Pt is shorthand for Point{x, y}.
func Pt(x, y int) Point {
	return Point{x, y}
}
