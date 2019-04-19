package cli

import (
	"strings"
	"unsafe"
)

type Attribute string

const (
	Reset        = "0"
	Bold         = "1"
	Faint        = "2"
	Italic       = "3"
	Underline    = "4"
	BlinkSlow    = "5"
	BlinkRapid   = "6"
	ReverseVideo = "7"
	Concealed    = "8"
	CrossedOut   = "9"

	Default   = "39" // Default was here for reference
	Black     = "30"
	Red       = "31"
	Green     = "32"
	Yellow    = "33"
	Blue      = "34"
	Magenta   = "35"
	Cyan      = "36"
	White     = "37"
	HiBlack   = "90"
	HiRed     = "91"
	HiGreen   = "92"
	HiYellow  = "93"
	HiBlue    = "94"
	HiMagenta = "95"
	HiCyan    = "96"
	HiWhite   = "97"

	BgDefault   = "49" // BgDefault was here for reference
	BgBlack     = "40"
	BgRed       = "41"
	BgGreen     = "42"
	BgYellow    = "43"
	BgBlue      = "44"
	BgMagenta   = "45"
	BgCyan      = "46"
	BgWhite     = "47"
	BgHiBlack   = "100"
	BgHiRed     = "101"
	BgHiGreen   = "102"
	BgHiYellow  = "103"
	BgHiBlue    = "104"
	BgHiMagenta = "105"
	BgHiCyan    = "106"
	BgHiWhite   = "107"
)

func (c *Builder) attribute(attributes ...Attribute) *Builder {
	return c.WriteString(strings.Join(*(*[]string)(unsafe.Pointer(&attributes)), SepControlChar))
}

func (c *Builder) Write(s string, attributes ...Attribute) *Builder {
	return c.Style(attributes...).WriteString(s).Reset()
}

func (c *Builder) Style(attributes ...Attribute) *Builder {
	if len(attributes) > 0 {
		c.
			WriteString(CSI).
			attribute(attributes...).
			WriteString(SGREnd)
	}
	return c
}

func (c *Builder) Reset() *Builder {
	return c.WriteString(SGRReset)
}

func (c *Builder) Bold(s string) *Builder {
	return c.Style(Bold).WriteString(s).Reset()
}

func (c *Builder) Faint(s string) *Builder {
	return c.Style(Faint).WriteString(s).Reset()
}

func (c *Builder) Italic(s string) *Builder {
	return c.Style(Italic).WriteString(s).Reset()
}

func (c *Builder) Underline(s string) *Builder {
	return c.Style(Underline).WriteString(s).Reset()
}

func (c *Builder) BlinkSlow(s string) *Builder {
	return c.Style(BlinkSlow).WriteString(s).Reset()
}

func (c *Builder) BlinkRapid(s string) *Builder {
	return c.Style(BlinkRapid).WriteString(s).Reset()
}

func (c *Builder) ReverseVideo(s string) *Builder {
	return c.Style(ReverseVideo).WriteString(s).Reset()
}

func (c *Builder) Concealed(s string) *Builder {
	return c.Style(Concealed).WriteString(s).Reset()
}

func (c *Builder) CrossedOut(s string) *Builder {
	return c.Style(CrossedOut).WriteString(s).Reset()
}

func (c *Builder) Black(value string) *Builder {
	return c.Style(Black).WriteString(value).Reset()
}

func (c *Builder) Red(value string) *Builder {
	return c.Style(Red).WriteString(value).Reset()
}

func (c *Builder) Green(value string) *Builder {
	return c.Style(Green).WriteString(value).Reset()
}

func (c *Builder) Yellow(value string) *Builder {
	return c.Style(Yellow).WriteString(value).Reset()
}

func (c *Builder) Blue(value string) *Builder {
	return c.Style(Blue).WriteString(value).Reset()
}

func (c *Builder) Magenta(value string) *Builder {
	return c.Style(Magenta).WriteString(value).Reset()
}

func (c *Builder) Cyan(value string) *Builder {
	return c.Style(Cyan).WriteString(value).Reset()
}

func (c *Builder) White(value string) *Builder {
	return c.Style(White).WriteString(value).Reset()
}

func (c *Builder) HiBlack(value string) *Builder {
	return c.Style(HiBlack).WriteString(value).Reset()
}

func (c *Builder) HiRed(value string) *Builder {
	return c.Style(HiRed).WriteString(value).Reset()
}

func (c *Builder) HiGreen(value string) *Builder {
	return c.Style(HiGreen).WriteString(value).Reset()
}

func (c *Builder) HiYellow(value string) *Builder {
	return c.Style(HiYellow).WriteString(value).Reset()
}

func (c *Builder) HiBlue(value string) *Builder {
	return c.Style(HiBlue).WriteString(value).Reset()
}

func (c *Builder) HiMagenta(value string) *Builder {
	return c.Style(HiMagenta).WriteString(value).Reset()
}

func (c *Builder) HiCyan(value string) *Builder {
	return c.Style(HiCyan).WriteString(value).Reset()
}

func (c *Builder) HiWhite(value string) *Builder {
	return c.Style(HiWhite).WriteString(value).Reset()
}

func (c *Builder) BgBlack(value string) *Builder {
	return c.Style(BgBlack).WriteString(value).Reset()
}

func (c *Builder) BgRed(value string) *Builder {
	return c.Style(BgRed).WriteString(value).Reset()
}

func (c *Builder) BgGreen(value string) *Builder {
	return c.Style(BgGreen).WriteString(value).Reset()
}

func (c *Builder) BgYellow(value string) *Builder {
	return c.Style(BgYellow).WriteString(value).Reset()
}

func (c *Builder) BgBlue(value string) *Builder {
	return c.Style(BgBlue).WriteString(value).Reset()
}

func (c *Builder) BgMagenta(value string) *Builder {
	return c.Style(BgMagenta).WriteString(value).Reset()
}

func (c *Builder) BgCyan(value string) *Builder {
	return c.Style(BgCyan).WriteString(value).Reset()
}

func (c *Builder) BgWhite(value string) *Builder {
	return c.Style(BgWhite).WriteString(value).Reset()
}

func (c *Builder) BgHiBlack(value string) *Builder {
	return c.Style(BgHiBlack).WriteString(value).Reset()
}

func (c *Builder) BgHiRed(value string) *Builder {
	return c.Style(BgHiRed).WriteString(value).Reset()
}

func (c *Builder) BgHiGreen(value string) *Builder {
	return c.Style(BgHiGreen).WriteString(value).Reset()
}

func (c *Builder) BgHiYellow(value string) *Builder {
	return c.Style(BgHiYellow).WriteString(value).Reset()
}

func (c *Builder) BgHiBlue(value string) *Builder {
	return c.Style(BgHiBlue).WriteString(value).Reset()
}

func (c *Builder) BgHiMagenta(value string) *Builder {
	return c.Style(BgHiMagenta).WriteString(value).Reset()
}

func (c *Builder) BgHiCyan(value string) *Builder {
	return c.Style(BgHiCyan).WriteString(value).Reset()
}

func (c *Builder) BgHiWhite(value string) *Builder {
	return c.Style(BgHiWhite).WriteString(value).Reset()
}
