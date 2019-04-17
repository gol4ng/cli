package loader

import (
	"strings"

	"github.com/gol4ng/cli"
	"github.com/gol4ng/cli/loader/style"
	"github.com/gol4ng/cli/unicode"
)

var VerticalChar = []rune{
	' ',
	unicode.LowerOneEighthBlock,
	unicode.LowerOneQuarterBlock,
	unicode.LowerThreeEighthsBlock,
	unicode.LowerHalfBlock,
	unicode.LowerFiveEighthsBlock,
	unicode.LowerThreeQuartersBlock,
	unicode.LowerSevenEighthsBlock,
	unicode.FullBlock,
}

var HorizontalChar = []rune{
	' ',
	unicode.LeftOneEighthBlock,
	unicode.LeftOneQuarterBlock,
	unicode.LeftThreeEighthsBlock,
	unicode.LeftHalfBlock,
	unicode.LeftFiveEighthsBlock,
	unicode.LeftThreeQuartersBlock,
	unicode.LeftSevenEighthsBlock,
	unicode.FullBlock,
}

var DotChar = []rune{
	' ',
	unicode.BraillePatternDots_1,
	unicode.BraillePatternDots_14,
	unicode.BraillePatternDots_124,
	unicode.BraillePatternDots_1245,
	unicode.BraillePatternDots_12345,
	unicode.BraillePatternDots_123456,
	unicode.BraillePatternDots_1234567,
	unicode.BraillePatternDots_12345678,
}

type BarStyle interface {
	BarChar() string
	BarLen() uint
	EmptyChar() string
	EmptyLen() uint
	CursorChar() string
	CursorLen() uint
}

type Bar struct {
	BarStyle
	length uint
}

func (d *Bar) Render(progress Progress) string {
	return d.RenderIn(progress, &strings.Builder{})
}

func (d *Bar) RenderIn(progress Progress, buffer cli.Buffer) string {
	percent := progress.Percent()
	barLen := int(d.length * uint(percent) / 100 * d.BarStyle.BarLen())
	emptyLen := int(d.length * d.BarStyle.EmptyLen())
	if percent > 0 {
		buffer.WriteString(strings.Repeat(string(d.BarStyle.BarChar()), barLen))
		emptyLen -= barLen

		if percent < 100 {
			cursor := d.BarStyle.CursorChar()
			if emptyLen < int(d.BarStyle.CursorLen()) {
				cursor = cursor[:emptyLen]
			}
			buffer.WriteString(cursor)
			emptyLen -= int(d.BarStyle.CursorLen())
		}
	}
	if emptyLen > 0 {
		buffer.WriteString(strings.Repeat(string(d.BarStyle.EmptyChar()), int(emptyLen)))
	}

	return buffer.String()
}

func NewDefaultBarRenderer(length uint) *Bar {
	return NewBarRenderer(length, style.NewBar())
}

func NewBarRenderer(length uint, style BarStyle) *Bar {
	return &Bar{
		length:   length,
		BarStyle: style,
	}
}

type SliceBar struct {
	chars []rune
}

func (d *SliceBar) Render(progress Progress) string {
	return d.RenderIn(progress, &strings.Builder{})
}

func (d *SliceBar) RenderIn(progress Progress, buffer cli.Buffer) string {
	percent := progress.Percent()
	threshold := float64(100) / float64(len(d.chars))
	for i := 0; i < len(d.chars); i++ {
		if percent < threshold*float64(i+1) {
			return string(d.chars[i])
		}
	}
	return string(d.chars[len(d.chars)-1])
}

func NewSliceBar(chars []rune) *SliceBar {
	return &SliceBar{chars: chars}
}
