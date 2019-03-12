package term

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png" // needed for loading default font

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	font          *Font
	update        func()
	screen        *ebiten.Image
	runeBlk       *ebiten.Image
	backbuffer    *ebiten.Image
	width, height int
	scale         float64
	title         string
	buffer        []Cell
	Debug         bool
)

//from: https://graphics.stanford.edu/~seander/bithacks.html#RoundUpPowerOf2
func nextPowerOf2(v int) int {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}

func Size() (int, int) { return width, height }

//func Font() *Font      { return font }

func Init(w, h int, s float64, t string) (err error) {
	width = w
	height = h
	buffer = make([]Cell, width*height)
	title = t
	scale = s
	font, err = NewFont(bytes.NewReader(Font16x16SbASCII), 16, 16)
	if err == nil {
		runeBlk, err = ebiten.NewImage(nextPowerOf2(font.width), nextPowerOf2(font.height), ebiten.FilterNearest)
	}
	if err == nil {
		err = runeBlk.Fill(White)
	}
	w, h = font.Size()
	w, h = nextPowerOf2(width*w), nextPowerOf2(height*h)
	if err == nil {
		backbuffer, err = ebiten.NewImage(w, h, ebiten.FilterNearest)
	}
	if err == nil {
		Clear(Black)
	}
	return
}

func __update(s *ebiten.Image) error {
	screen = s
	update()
	screen.DrawImage(backbuffer, nil)
	if Debug {
		w, h := font.Size()
		ebitenutil.DebugPrint(s, fmt.Sprintf("%vx%v %0.2f", width*w, height*h, ebiten.CurrentTPS()))
	}
	return nil
}

func Run(upd func()) error {
	update = upd
	w, h := font.Size()
	return ebiten.Run(__update, width*w, height*h, scale, title)
}

func Clear(bg Color) {
	for i := 0; i < len(buffer); i++ {
		buffer[i] = Cell{' ', White, bg, nil}
	}
	backbuffer.Fill(bg)
}

func SetCell(x, y int, c Cell) {
	idx := y*width + x
	ocell := buffer[idx]
	if c.Fg.Equal(Transparent) {
		c.Fg = ocell.Fg
	}
	if c.Bg.Equal(Transparent) {
		c.Bg = ocell.Bg
	}
	if c.Data == nil {
		c.Data = ocell.Data
	}

	fw := font.width
	fh := font.height
	ffw := float64(fw)
	ffh := float64(fh)

	buffer[idx] = c

	co := c.Bg
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x)*ffw, float64(y)*ffh)
	op.ColorM.Scale(0, 0, 0, 1)
	op.ColorM.Translate(co.r, co.g, co.b, 0)
	backbuffer.DrawImage(runeBlk, op)
	co = c.Fg
	op.ColorM.Scale(0, 0, 0, 1)
	op.ColorM.Translate(co.r, co.g, co.b, 0)
	ch := int(c.R)
	xs := ch % 16 * fw //font maps are 16*16 characters
	ys := ch / 16 * fh
	rect := image.Rect(xs, ys, xs+fw, ys+fh)
	backbuffer.DrawImage(font.Image().SubImage(rect).(*ebiten.Image), op)
}

func CellAt(x, y int) Cell {
	return buffer[y*width+x]
}

//do not modify cell content. Use SetCell
func Cells() []Cell {
	return buffer
}

func CellIndex(x, y int) int {
	return y*width + x
}

func SetData(x, y int, data interface{}) {
	idx := y*width + x
	cell := buffer[idx]
	cell.Data = data
	buffer[idx] = cell
}

func Data(x, y int) interface{} {
	return buffer[y*width+x].Data
}

func Scroll(x, y int) {
	_, fh := font.Size()
	var xs, ys, xe, ye, xd, yd int
	if x == 0 {
		if y == 0 {
			return
		}
		if y > 0 && y < height-1 { // scroll down
			//pre := slices.Repeat(EmptyCell, y*width)
			//buffer = append(pre,  buffer[:(width*height)-(y*width)] )
			buffer = buffer[:(width*height)-(y*width)]
			yd = y * fh
			ys = 0
			ye = height*fh - y*fh - 1
		} else { //scroll up
			buffer = buffer[y*width:]
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(xd), float64(yd))
	rect := image.Rect(xs, ys, xe, ye)
	w, h := backbuffer.Size()
	newbuffer, err := ebiten.NewImage(w, h, ebiten.FilterNearest)
	//todo: handle error correctly
	if err == nil {
		//todo: check visual
		//newbuffer.Fill(Black)
		fmt.Printf("%v, %v %v %v", xd, yd, rect, height*fh)
		//newbuffer.DrawImage(backbuffer.SubImage(rect).(*ebiten.Image), op)
		newbuffer.DrawImage(backbuffer, op)
		backbuffer = newbuffer
	}
}

//todo: tester Cells() + CellIndex()  + Data() et SetData() + check data pas modifiée apres SetCell
//todo: scroll
//todo: overlays
//todo: voir si color dans son propre package
//todo: puis changer code dans main et map
//todo: Add Font() + SetFont() ?
//todo: encapsulate ebiten keys

func Print(x int, y int, s string, fg Color, bg Color) {
	xo := x
	for _, r := range s {
		if r == '\n' {
			y++
			x = xo
		} else {
			SetCell(x, y, Cell{r, fg, bg, nil})
			x++
		}
	}
}

func Key(key ebiten.Key) bool {
	return ebiten.IsKeyPressed(key)
}

var keymap [ebiten.KeyMax + 1]int

func Keyp(key ebiten.Key) bool {
	const (
		delay    = 10
		interval = 4
	)
	if keymap[key] > 0 {
		if ebiten.IsKeyPressed(key) {
			d := keymap[key]
			if d >= delay && (d-delay)%interval == 0 {
				return true
			}
			keymap[key]++
		} else {
			keymap[key] = 0
		}
		return false
	} else {
		if ebiten.IsKeyPressed(key) {
			keymap[key] = 1
		}
		return keymap[key] > 0
	}
}

func Axis() (x int, y int) {
	if Keyp(ebiten.KeyUp) {
		y--
	}
	if Keyp(ebiten.KeyDown) {
		y++
	}
	if Keyp(ebiten.KeyLeft) {
		x--
	}
	if Keyp(ebiten.KeyRight) {
		x++
	}
	return
}
