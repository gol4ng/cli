package loader

import (
	"github.com/gol4ng/cli/unicode"
)

var ClockSquare = []rune{
	unicode.WhiteSquareWithUpperRightQuadrant,
	unicode.WhiteSquareWithLowerRightQuadrant,
	unicode.WhiteSquareWithLowerLeftQuadrant,
	unicode.WhiteSquareWithUpperLeftQuadrant,
}

var Clock = []rune{
	unicode.WhiteCircleWithUpperRightQuadrant,
	unicode.WhiteCircleWithLowerRightQuadrant,
	unicode.WhiteCircleWithLowerLeftQuadrant,
	unicode.WhiteCircleWithUpperLeftQuadrant,
}

type Spinner struct {
	chars []rune
	frame uint
}

func (s *Spinner) String() string {
	value := string(s.chars[int(s.frame)%len(s.chars)])
	s.frame++
	return value
}

func NewCharSpinner(chars []rune) *Spinner {
	if len(chars) == 0 {
		panic("you must provide chars")
	}
	return &Spinner{chars: chars, frame: 0}
}
