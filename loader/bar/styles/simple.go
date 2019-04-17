package styles

import "unicode/utf8"

const (
	defaultBar    rune   = '█'
	defaultEmpty         = '░'
	defaultCursor string = "▒"
)

type Simple struct {
	Bar          rune
	Empty        rune
	Cursor       string
	CursorLength uint
}

func (d *Simple) BarChar() rune {
	return d.Bar
}

func (d *Simple) EmptyChar() rune {
	return d.Empty
}

func (d *Simple) CursorChar() string {
	return d.Cursor
}

func (d *Simple) CursorLen() uint {
	return d.CursorLength
}

func New() *Simple {
	return NewSimple(defaultBar, defaultEmpty, defaultCursor)
}

func NewSimple(bar rune, empty rune, cursor string) *Simple {
	return &Simple{
		Bar:          bar,
		Empty:        empty,
		Cursor:       cursor,
		CursorLength: uint(utf8.RuneCountInString(cursor)),
	}
}
