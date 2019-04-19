package cli

// https://en.wikipedia.org/wiki/ANSI_escape_code#Sequence_elements
const (
	ESC string = "\x1b"

	//Control Sequence Introducer
	CSI = ESC + "["

	//Select Graphic Rendition
	SGREnd   = "m"
	SGRReset = CSI + SGREnd

	SepControlChar = ";"
)
