package term

import (
	"fmt"
	"image/color"
	_ "image/png" // needed for loading default font
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/inconsolata"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/goregular"
)

/*
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
	bg            []color.Color
	Debug         bool
)*/

/***** NEW API *****/
//see https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
//https://github.com/pkg/term
/*
tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	mplusBigFont = truetype.NewFace(tt, &truetype.Options{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
*/

//golang.org/x/image/font/gofont/gomono

//type TColor color.Color

//var White = TColor(color.NRGBA{255, 255, 255, 255})

type term struct {
	width, height int
	scale         int
	font          font.Face //todo: one font per layer
	bitmapFont    bool      //or monospace  + size if TTF and w+h if bmp
	color         Color     //todo: remove, its just for test
	fullscreen    bool
	title         string
}

var (
	f2, f3, f4, f5, f6 font.Face
	runeBlk            *ebiten.Image
	rw                 = 16
	rh                 = 32
)

var _term term

func Size(width, height int) func(*term) error {
	return func(t *term) error {
		t.width = width
		t.height = height
		return nil
	}
}

func Scale(scale int) func(*term) error {
	return func(t *term) error {
		t.scale = scale
		return nil
	}
}

func Colo(color Color) func(*term) error {
	return func(t *term) error {
		t.color = color
		return nil
	}
}

func Font(font font.Face) func(*term) error {
	return func(t *term) error {
		t.font = font
		return nil
	}
}

func Fullscreen() func(*term) error {
	return func(t *term) error {
		t.fullscreen = true
		return nil
	}
}

func Title(title string) func(*term) error {
	return func(t *term) error {
		t.title = title
		return nil
	}
}

func Dump() {
	fmt.Printf("%v,%v\n", _term.width, _term.height)
	fmt.Printf("Color:%v\n", _term.color)
}

//or use option struct

//chromeos terminal fonts :
//"DejaVu Sans Mono", "Noto Sans Mono", "Everson Mono", FreeMono, Menlo, Terminal, monospace
//size 15 and use noto
func Open(options ...func(*term) error) error {
	t := &_term
	for _, option := range options {
		option(t)

	}
	if t.width == 0 {
		t.width = 80
	}
	if t.height == 0 {
		t.height = 25
	}
	if t.scale == 0 {
		t.scale = 1
	}
	if t.font == nil {
		t.font = inconsolata.Regular8x16
	}
	f2 = inconsolata.Bold8x16

	tt, err := truetype.Parse(gomono.TTF)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	f3 = truetype.NewFace(tt, &truetype.Options{
		Size:    14,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	w, _ := f3.GlyphAdvance('@')
	h := f3.Metrics().Height
	fmt.Printf("%v,%v 3\n", w, h)
	f4 = truetype.NewFace(tt, &truetype.Options{
		Size:    16,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	w, _ = f4.GlyphAdvance('@')
	h = f4.Metrics().Height
	fmt.Printf("%v,%v 4\n", w, h)

	f5 = truetype.NewFace(tt, &truetype.Options{
		Size:    25,
		DPI:     dpi,
		Hinting: font.HintingVertical,
		//SubPixelsX: 1,
		//SubPixelsY: 1,
	})
	w, _ = f5.GlyphAdvance('@')
	h = f5.Metrics().Height
	fmt.Printf("%v,%v 5\n", w, h)

	tt, err = truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	f6 = truetype.NewFace(tt, &truetype.Options{
		Size:    14,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	w, _ = f6.GlyphAdvance('@')
	h = f6.Metrics().Height
	fmt.Printf("%v,%v 6\n", w, h)
	//todo: handle error
	runeBlk, _ = ebiten.NewImage(rw, rh, ebiten.FilterNearest)

	runeBlk.Fill(color.White)

	return nil
}

func Close() {

}

func Refresh() {

}

func __update(s *ebiten.Image) error {
	/*screen = s
	update()*/
	//screen.DrawImage(backbuffer, nil)

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(0, 0, 0, 1)
	op.ColorM.Translate(.2, .2, .8, 1.0)
	for y := 0; y < 10; y++ {
		//op.GeoM..Scale(0, 0)
		for x := 0; x < 40; x++ {
			s.DrawImage(runeBlk, op)
			op.GeoM.Translate(float64(rw+rw), 0)
		}
		op.GeoM.Translate(-40*float64(rw+rw), float64(rh+rh))
	}

	text.Draw(s, "Hello World! @ # 1", _term.font, 16, 16, White)
	text.Draw(s, "Hello World! @ # 2", f2, 16, 48, color.White)
	text.Draw(s, "Hello World! @ # 3", f3, 16, 48+32, color.White)
	text.Draw(s, "Hello World! @ # 4", f4, 16, 48+64, color.White)
	text.Draw(s, "Hello World! @ # 5", f5, 16, 48+64+32, color.White)
	text.Draw(s, "Hello World! @ # 6", f6, 16, 48+128, color.White)

	/*if Debug {
		w, h := font.Size()
		ebitenutil.DebugPrint(s, fmt.Sprintf("%vx%v %0.2f", width*w, height*h, ebiten.CurrentTPS()))
	}*/
	return nil
}

func Run(upd func()) error {
	//update = upd
	//w, h := font.Size()
	return ebiten.Run(__update, _term.width*8, _term.height*16, float64(_term.scale), _term.title)
}

/***** OLD API *****/
//func Size() (int, int) { return width, height }

//func Font() *Font      { return font }
/*
func Init(w, h int, s float64, t string) (err error) {
	width = w
	height = h
	buffer = make([]Cell, width*height)
	bg = make([]image.Color, width*height)
	title = t
	scale = s
	font, err = NewFont(bytes.NewReader(Font16x16SbASCII), 16, 16)
	if err == nil {
		runeBlk, err = ebiten.NewImage(ints.RoundUpPowerOf2(font.width), ints.RoundUpPowerOf2(font.height), ebiten.FilterNearest)
	}
	if err == nil {
		err = runeBlk.Fill(White)
	}
	w, h = font.Size()
	w, h = ints.RoundUpPowerOf2(width*w), ints.RoundUpPowerOf2(height*h)
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

//todo: methode At plutot que CellAt

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
}*/
