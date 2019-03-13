package term

import "image"

//Size is a pair of width(W) and height(H)
type Size struct {
	W, H int
}

//A Point is an X, Y coordinate pair. The axes increase right and down.
type Point image.Point

//ZP is the zero Point.
var ZP Point

//Pt is shorthand for Point{x, y}.
func Pt(x, y int) Point {
	return Point{x, y}
}
