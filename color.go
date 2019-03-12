package term

import (
	"encoding/binary"
	ic "image/color"
)

type Color struct {
	r, g, b, a float64
}

var (
	White       = Color{1, 1, 1, 1}
	Black       = Color{0, 0, 0, 1}
	Transparent = Color{0, 0, 0, 0}
)

func NewColor(r, g, b, a float64) Color {
	return Color{r, g, b, a}
}

// stored value can be slightly different due to float conversion
func NewColorFromInt8(r, g, b, a uint8) Color {
	rf := float64(r) / 0xff
	gf := float64(g) / 0xff
	bf := float64(b) / 0xff
	af := float64(a) / 0xff
	return Color{rf, gf, bf, af}
}

// stored value can be slightly different due to float conversion
func NewColorFromInt32(r, g, b, a uint32) Color {
	return NewColorFromInt8(uint8(r), uint8(g), uint8(b), uint8(a))
}

// stored value can be slightly different due to float conversion
func NewColorFromHex(c uint32) Color {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, c)
	return NewColorFromInt8(uint8(b[0]), uint8(b[1]), uint8(b[2]), uint8(b[3]))
}

// stored value can be slightly different due to float conversion
func NewColorFromColor(c ic.Color) Color {
	r, g, b, a := c.RGBA()
	return NewColorFromInt8(uint8(r), uint8(g), uint8(b), uint8(a))
}

func (c Color) RGBA() (r, g, b, a uint32) {
	aa := uint32(c.a * 0xff)
	r = uint32(c.r * 0xff)
	r |= r << 8
	r *= aa
	r /= 0xff
	g = uint32(c.g * 0xff)
	g |= g << 8
	g *= aa
	g /= 0xff
	b = uint32(c.b * 0xff)
	b |= b << 8
	b *= aa
	b /= 0xff
	a = aa
	a |= a << 8
	return
}

func (c Color) FRGBA() (r, g, b, a float64) {
	return c.r, c.g, c.b, c.a
}

func (c Color) Equal(o Color) bool {
	return c.r == o.r && c.g == o.g && c.b == o.b && c.a == o.a
}
