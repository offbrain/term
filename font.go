package term

import (
	"image"
	"io"

	"github.com/hajimehoshi/ebiten"
)

type Font struct {
	image         *ebiten.Image
	width, height int
}

func NewFont(r io.Reader, w, h int) (f *Font, err error) {
	var fontimg *ebiten.Image
	img, _, err := image.Decode(r)
	if err == nil {
		fontimg, err = ebiten.NewImageFromImage(img, ebiten.FilterNearest)
	}
	if err == nil {
		f = &Font{fontimg, w, h}
	}
	return
}

func (f Font) Image() *ebiten.Image { return f.image }
func (f Font) Size() (int, int)     { return f.width, f.height }
