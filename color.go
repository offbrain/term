//This is free and unencumbered software released into the public domain.

package term

import (
	"encoding/binary"
	"image/color"

	"github.com/offbrain/math/clamp"
)

//Predefined colors
var (
	White = ColorFromHex(0xFFFFFFFF)
	Black = ColorFromHex(0x000000FF)
)

// Color can convert itself to alpha-premultiplied 16-bits per channel RGBA.
// The conversion may be lossy.
type Color color.Color

// NRGBA represents a non-alpha-premultiplied 32-bit color
type NRGBA color.NRGBA

// RGBA returns the alpha-premultiplied red, green, blue and alpha values
// for the color. Each value ranges within [0, 0xffff], but is represented
// by a uint32 so that multiplying by a blend factor up to 0xffff will not
// overflow.
//
// An alpha-premultiplied color component c has been scaled by alpha (a),
// so has valid values 0 <= c <= a.
func (c NRGBA) RGBA() (r, g, b, a uint32) {
	return color.NRGBA(c).RGBA()
}

// FNRGBA returns the non-alpha-premultiplied red, green, blue and alpha values
// for the color as float64. Each value ranges within [0.0, 1.0].
func (c NRGBA) FNRGBA() (r, g, b, a float64) {
	r = clamp.Float64(float64(c.R)/255.0, 0, 1)
	g = clamp.Float64(float64(c.G)/255.0, 0, 1)
	b = clamp.Float64(float64(c.B)/255.0, 0, 1)
	a = clamp.Float64(float64(c.A)/255.0, 0, 1)
	return
}

//ColorFromHex returns the NRGBA color created with the nrgba hexadecimal value
func ColorFromHex(nrgba uint32) Color {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, nrgba)
	return NRGBA{uint8(b[0]), uint8(b[1]), uint8(b[2]), uint8(b[3])}
}
