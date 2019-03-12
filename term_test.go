package term_test

import (
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/hajimehoshi/ebiten"
	"github.com/offbrain/term"
)

var (
	/*updateCnt  = 0
	maxUpdates = 60 * 10*/
	width  = 40
	height = 25
)

func update() {
	/*updateCnt++
	if updateCnt >= maxUpdates {
		os.Exit(0)
	}*/

	if ebiten.IsDrawingSkipped() {
		return
	}

	if term.Keyp(ebiten.KeyR) {
		bg := term.NewColorFromInt8(uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255)
		fg := term.NewColorFromInt8(uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255)
		term.SetCell(rand.Intn(width), rand.Intn(height), term.Cell{R: rune(rand.Intn(255)), Fg: fg, Bg: bg})
	}

}

type testData struct {
	v int
}

func TestTermClear(t *testing.T) {
	term.Clear(term.Black)
	term.SetCell(1, 2, term.Cell{'Y', term.NewColorFromHex(0xAAAAAAFF), term.NewColorFromHex(0x121212FF), testData{22}})
	term.Clear(term.Black)
	cells := term.Cells()
	cell := cells[term.CellIndex(1, 2)]
	if cell.R != ' ' {
		t.Errorf("Cell R should be ' ' on initialisation but contains '%v'", cell.R)
	}
	if !cell.Fg.Equal(term.White) {
		t.Errorf("Cell Fg should be White on initialisation but contains [%v]", cell.Fg)
	}
	if !cell.Bg.Equal(term.Black) {
		t.Errorf("Cell Fg should be Black on initialisation but contains [%v]", cell.Bg)
	}
	if cell.Data != nil {
		t.Errorf("Cell Data should be nil on initialisation but contains '%v'", cell.Data)
	}
}

func TestTermCell(t *testing.T) {
	term.Clear(term.Black)
	cells := term.Cells()
	cell := cells[term.CellIndex(1, 2)]
	cell.R = 'X'
	fg := term.NewColorFromHex(0x000000FF)
	bg := term.NewColorFromHex(0xFFFFFFFF)
	cell.Fg = fg
	cell.Bg = bg
	data := testData{33}
	cell.Data = data
	term.SetCell(1, 2, cell)
	cells = term.Cells()
	cell = cells[term.CellIndex(1, 2)]
	if cell.R != 'X' {
		t.Errorf("Cell R should contain 'X' but is %v", cell.R)
	}
	if cell.Data.(testData).v != 33 {
		t.Errorf("Cell Data should contain '33' but is %v", cell.Data.(testData).v)
	}
}

func TestTermData(t *testing.T) {
	term.Clear(term.Black)
	term.SetData(2, 3, testData{98})
	term.SetCell(2, 3, term.Cell{R: '!', Fg: term.NewColorFromHex(0x000000FF), Bg: term.NewColorFromHex(0xFFFFFFFF)})
	if term.Data(2, 3).(testData).v != 98 {
		t.Errorf("Cell Data should contain '98' but is %v", term.Data(2, 3).(testData).v)
	}
}

func TestTermRainbowBg(t *testing.T) {
	term.Clear(term.Black)
	var r, g, b int = 0, 0, 0
	var dx, dy int = 255 / width, 255 / height
	for y := 0; y < height; y++ {
		g += dy
		r = 0
		for x := 0; x < width; x++ {
			r += dx
			c := term.NewColorFromInt8(uint8(r), uint8(g), uint8(b), 255)
			term.SetCell(x, y, term.Cell{' ', term.Transparent, c, nil})
		}
	}
	//t.Errorf("WTF?")
	//os.Exit(0)
}

func TestTermRainbowFg(t *testing.T) {
	term.Clear(term.Black)
	var r, g, b int = 0, 0, 0
	var dx, dy int = 255 / width, 255 / height
	for y := 0; y < height; y++ {
		g += dy
		r = 0
		for x := 0; x < width; x++ {
			r += dx
			c := term.NewColorFromInt8(uint8(r), uint8(g), uint8(b), 255)
			term.SetCell(x, y, term.Cell{'#', c, term.Black, nil})
		}
	}
	//t.Errorf("WTF?")
	//os.Exit(0)
}

func TestTermRainbowBoth(t *testing.T) {
	term.Clear(term.Black)
	var r, g, b int = 0, 0, 0
	var dx, dy int = 255 / width, 255 / height
	for y := 0; y < height; y++ {
		g += dy
		r = 0
		for x := 0; x < width; x++ {
			r += dx
			fg := term.NewColorFromInt8(uint8(r), uint8(g), uint8(b), 255)
			bg := term.NewColorFromInt8(uint8(255-r), uint8(255-g), uint8(255-b), 255)
			term.SetCell(x, y, term.Cell{rune(254), fg, bg, nil})
		}
	}
	//t.Errorf("WTF?")
	//os.Exit(0)
}

func TestTermPrint(t *testing.T) {
	TestTermRainbowBg(t)
	term.Print(1, 1, "Hello, World!", term.NewColorFromHex(0xFFFF77FF), term.NewColorFromHex(0x7777FFFF))
	term.Print(2, 3, "Hello, \nWorld!", term.NewColorFromHex(0x777777FF), term.Transparent)
}

func TestTermScroll(t *testing.T) {
	term.Clear(term.Black)
	var r, g, b int = 0, 0, 0
	var dx, dy int = 255 / width, 255 / height
	irune := 0
	var rval rune
	for y := 0; y < height; y++ {
		g += dy
		r = 0
		for x := 0; x < width; x++ {
			r += dx
			fg := term.NewColorFromInt8(uint8(r), uint8(g), uint8(b), 255)
			bg := term.NewColorFromInt8(uint8(255-r), uint8(255-g), uint8(255-b), 255)
			term.SetCell(x, y, term.Cell{rune(irune), fg, bg, nil})
			irune++
			if irune >= 255 {
				irune = 0
			}
		}
	}
	rval = term.Cells()[term.CellIndex(11, 12)].R
	term.Scroll(0, 1)
	if rval == term.Cells()[term.CellIndex(11, 12)].R {
		t.Errorf("Cell has not moved")
	}
	if rval != term.Cells()[term.CellIndex(11, 13)].R {
		t.Errorf("Cell not found after Scroll()")
	}
}

func TestMain(m *testing.M) {
	if err := term.Init(width, height, 1, "test"); err != nil {
		log.Fatal(err)
	}
	term.Debug = true

	ret := m.Run()

	if err := term.Run(update); err != nil {
		log.Fatal(err)
	}

	os.Exit(ret)
}
