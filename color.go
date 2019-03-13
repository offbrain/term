//This is free and unencumbered software released into the public domain.

package term

import (
	"encoding/binary"
	"image/color"

	"github.com/offbrain/math/clamp"
)

//Predefined colors
var (
	White = NewColorHex(0xFFFFFFFF)
	Black = NewColorHex(0x000000FF)
)

// Color represents a non-alpha-premultiplied 32-bit color
// It can convert itself to alpha-premultiplied 16-bits per channel RGBA.
// The conversion may be lossy.
type Color color.NRGBA

// RGBA returns the alpha-premultiplied red, green, blue and alpha values
// for the color. Each value ranges within [0, 0xffff], but is represented
// by a uint32 so that multiplying by a blend factor up to 0xffff will not
// overflow.
//
// An alpha-premultiplied color component c has been scaled by alpha (a),
// so has valid values 0 <= c <= a.
func (c Color) RGBA() (r, g, b, a uint32) {
	return color.NRGBA(c).RGBA()
}

// FNRGBA returns the non-alpha-premultiplied red, green, blue and alpha values
// for the color as float64. Each value ranges within [0.0, 1.0].
func (c Color) FNRGBA() (r, g, b, a float64) {
	r = clamp.Float64(float64(c.R)/255.0, 0, 1)
	g = clamp.Float64(float64(c.G)/255.0, 0, 1)
	b = clamp.Float64(float64(c.B)/255.0, 0, 1)
	a = clamp.Float64(float64(c.A)/255.0, 0, 1)
	return
}

//NewColor return a new Color containing non-alpha-premultiplied red, green, blue and alpha values
func NewColor(r, g, b, a uint8) Color {
	return Color{R: r, G: g, B: b, A: a}
}

//NewColorHex returns a new color created with the non-alpha-premultiplied nrgba value
//The nrgba value is not requided to be passed as an hexadecimal value but it's more convenient.
func NewColorHex(nrgba uint32) Color {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, nrgba)
	return Color{R: uint8(b[0]), G: uint8(b[1]), B: uint8(b[2]), A: uint8(b[3])}
}
