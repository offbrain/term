package term_test

import (
	"image/color"
	"math"
	"testing"

	"github.com/offbrain/term"
)

func TestColorRGBA(t *testing.T) {
	ic := color.NRGBA{0x12, 0x34, 0x56, 0x78}
	c := term.Color{R: 0x12, G: 0x34, B: 0x56, A: 0x78}
	w, x, y, z := ic.RGBA()
	r, g, b, a := c.RGBA()
	if r != w || g != x || b != y || a != z {
		t.Errorf("RGBA returned incorrect values: %#x, %#x, %#x, %#x", r, g, b, a)
	}
}

func TestNewColor(t *testing.T) {
	ic := color.NRGBA{0x12, 0x34, 0x56, 0x78}
	c := term.NewColor(0x12, 0x34, 0x56, 0x78)
	w, x, y, z := ic.RGBA()
	r, g, b, a := c.RGBA()
	if r != w || g != x || b != y || a != z {
		t.Errorf("NewColor initialized incorrect values: %#x, %#x, %#x, %#x", r, g, b, a)
	}
}

func TestNewColorHex(t *testing.T) {
	c := term.NewColorHex(0x12345678)
	if c.R != 0x12 || c.G != 0x34 || c.B != 0x56 || c.A != 0x78 {
		t.Errorf("NewColorHex initialized incorrect values: %#x, %#x, %#x, %#x", c.R, c.G, c.B, c.A)
	}
}

func TestColorFromImageColor(t *testing.T) {
	ic := color.NRGBA{0x12, 0x34, 0x56, 0x78}
	c := term.Color(ic)
	w, x, y, z := ic.RGBA()
	r, g, b, a := c.RGBA()
	if r != w || g != x || b != y || a != z {
		t.Errorf("Color convertion created incorrect values: %#x, %#x, %#x, %#x", r, g, b, a)
	}
}

func TestColorFNRGBA(t *testing.T) {
	c := term.NewColor(0, 64, 128, 255)
	r, g, b, a := c.FNRGBA()
	g = math.Round(g*100) / 100
	b = math.Round(b*100) / 100
	if r != 0.0 || g != 0.25 || b != 0.50 || a != 1.0 {
		t.Errorf("FNRGBA returned incorrect values: %#f, %#f, %#f, %#f", r, g, b, a)
	}
}
