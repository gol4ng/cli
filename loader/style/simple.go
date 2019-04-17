package style

import (
	"github.com/gol4ng/cli/unicode"
	"unicode/utf8"
)

type Simple struct {
	Bar          string
	BarLength    uint
	Empty        string
	EmptyLength  uint
	Cursor       string
	CursorLength uint
}

func (d *Simple) BarChar() string {
	return d.Bar
}

func (d *Simple) BarLen() uint {
	return d.BarLength
}

func (d *Simple) EmptyChar() string {
	return d.Empty
}

func (d *Simple) EmptyLen() uint {
	return d.EmptyLength
}

func (d *Simple) CursorChar() string {
	return d.Cursor
}

func (d *Simple) CursorLen() uint {
	return d.CursorLength
}

func NewBar() *Simple {
	return NewBarSimple(string(unicode.FullBlock), string(unicode.LightShade), string(unicode.MediumShade))
}

func NewBarSimple(bar string, empty string, cursor string) *Simple {
	return &Simple{
		Bar:          bar,
		BarLength:    uint(utf8.RuneCountInString(bar)),
		Empty:        empty,
		EmptyLength:  uint(utf8.RuneCountInString(empty)),
		Cursor:       cursor,
		CursorLength: uint(utf8.RuneCountInString(cursor)),
	}
}
