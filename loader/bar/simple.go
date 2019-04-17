package bar

import (
	"strings"

	"github.com/gol4ng/cli/loader"
	"github.com/gol4ng/cli/loader/bar/styles"
)

type simple struct {
	loader.BarStyle
	length uint
}

func (d *simple) Render(percent uint) string {
	if percent > 100 {
		percent = 100
	}
	builder := strings.Builder{}
	barLen := d.length * percent / 100
	emptyLen := int(d.length)
	if percent > 0 {
		builder.WriteString(strings.Repeat(string(d.BarStyle.BarChar()), int(barLen)))
		emptyLen -= int(barLen)

		if percent < 100 {
			cursor := d.BarStyle.CursorChar()
			if emptyLen < int(d.BarStyle.CursorLen()) {
				cursor = cursor[:emptyLen]
			}
			builder.WriteString(cursor)
			emptyLen -= int(d.BarStyle.CursorLen())
		}
	}
	if emptyLen > 0 {
		builder.WriteString(strings.Repeat(string(d.BarStyle.EmptyChar()), int(emptyLen)))
	}

	return builder.String()
}

func New(length uint) *simple {
	return NewSimple(length, styles.New())
}

func NewSimple(length uint, style loader.BarStyle) *simple {
	return &simple{
		length:   length,
		BarStyle: style,
	}
}
