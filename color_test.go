package term_test

import (
	"encoding/binary"
	ic "image/color"
	"testing"

	"github.com/offbrain/term"
)

func TestColorFromInt8(t *testing.T) {
	rf, gf, bf, af := 0.25, 0.50, 0.75, 1.0
	ri, gi, bi, ai := uint8(rf*0xff), uint8(gf*0xff), uint8(bf*0xff), uint8(af*0xff)
	rf, gf, bf, af = float64(ri)/0xff, float64(gi)/0xff, float64(bi)/0xff, float64(ai)/0xff
	c := term.NewColorFromInt8(ri, gi, bi, ai)
	r, g, b, a := c.FRGBA()
	if r != rf || g != gf || b != bf || a != af {
		t.Errorf("NewColorFromInt8 returned incorrect values : %#f, %#f, %#f, %#f", r, g, b, a)
	}
}

func TestColorFromColor(t *testing.T) {
	rf, gf, bf, af := 0.25, 0.50, 0.75, 1.0
	ri, gi, bi, ai := uint8(rf*0xff), uint8(gf*0xff), uint8(bf*0xff), uint8(af*0xff)
	rf, gf, bf, af = float64(ri)/0xff, float64(gi)/0xff, float64(bi)/0xff, float64(ai)/0xff
	icc := ic.NRGBA{ri, gi, bi, ai}
	c := term.NewColorFromColor(icc)
	r, g, b, a := c.FRGBA()
	if r != rf || g != gf || b != bf || a != af {
		t.Errorf("NewColorFromColor returned incorrect values : %#f, %#f, %#f, %#f", r, g, b, a)
	}
}

func TestColorFromHex(t *testing.T) {
	rf, gf, bf, af := 0.25, 0.50, 0.75, 1.0
	ri, gi, bi, ai := uint8(rf*0xff), uint8(gf*0xff), uint8(bf*0xff), uint8(af*0xff)
	by := make([]byte, 4)
	by[0], by[1], by[2], by[3] = byte(ri), byte(gi), byte(bi), byte(ai)
	ui32 := binary.BigEndian.Uint32(by)
	rf, gf, bf, af = float64(ri)/0xff, float64(gi)/0xff, float64(bi)/0xff, float64(ai)/0xff
	c := term.NewColorFromHex(ui32)
	r, g, b, a := c.FRGBA()
	if r != rf || g != gf || b != bf || a != af {
		t.Errorf("NewColorFromHex returned incorrect values : %#f, %#f, %#f, %#f", r, g, b, a)
	}
}

func TestColorRGBA(t *testing.T) {
	rf, gf, bf, af := 0.25, 0.50, 0.75, 1.0
	ri, gi, bi, ai := uint8(rf*0xff), uint8(gf*0xff), uint8(bf*0xff), uint8(af*0xff)
	by := make([]byte, 4)
	by[0], by[1], by[2], by[3] = byte(ri), byte(gi), byte(bi), byte(ai)
	r2, g2, b2, a2 := uint32(by[0]), uint32(by[1]), uint32(by[2]), uint32(by[3])
	r3, g3, b3, a3 := r2<<8|r2, g2<<8|g2, b2<<8|b2, a2<<8|a2
	c := term.NewColorFromInt8(ri, gi, bi, ai)
	r, g, b, a := c.RGBA()
	if r != r3 || g != g3 || b != b3 || a != a3 {
		t.Errorf("RGBA should return 16 bits values : %#04x, %#04x, %#04x, %#04x", r, g, b, a)
	}
}

func TestColorFRGBA(t *testing.T) {
	c := term.NewColor(0.1, 0.2, 0.3, 0.4)
	r, g, b, a := c.FRGBA()
	if r != 0.1 || g != 0.2 || b != 0.3 || a != 0.4 {
		t.Errorf("FRGBA should return the stored values : %#0.1f, %#0.1f, %#0.1f, %#0.1f", r, g, b, a)
	}
}
