package loader

type Bar interface {
	Render(percent uint) string
}

type BarStyle interface {
	BarChar() rune
	EmptyChar() rune
	CursorChar() string
	CursorLen() uint
}
