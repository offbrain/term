package term

import (
	"image"
	"io"

	"github.com/hajimehoshi/ebiten"
)

type TermFont struct {
	image         *ebiten.Image
	width, height int
}

func NewFont(r io.Reader, w, h int) (f *TermFont, err error) {
	var fontimg *ebiten.Image
	img, _, err := image.Decode(r)
	if err == nil {
		fontimg, err = ebiten.NewImageFromImage(img, ebiten.FilterNearest)
	}
	if err == nil {
		f = &TermFont{fontimg, w, h}
	}
	return
}

func (f TermFont) Image() *ebiten.Image { return f.image }
func (f TermFont) Size() (int, int)     { return f.width, f.height }
