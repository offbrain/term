package term

type Cell struct {
	R    rune
	Fg   Color
	Bg   Color
	Data interface{}
}

var EmptyCell = Cell{R: ' ', Fg: White, Bg: Black, Data: nil}
