package loader

import (
	"github.com/gol4ng/cli"
	"strconv"
	"strings"
)

type ProgressRenderer interface {
	Render(progress Progress) string
	RenderIn(progress Progress, buffer cli.Buffer) string
}

type Progress struct {
	step uint
	max  uint
}

func (p *Progress) GetStep() uint {
	return p.step
}

func (p *Progress) SetStep(step uint) *Progress {
	if step > p.max {
		step = p.max
	}
	p.step = step
	return p
}

func (p *Progress) GetMax() uint {
	return p.step
}

func (p *Progress) SetMax(max uint) *Progress {
	p.max = max
	if p.step > p.max {
		p.step = p.max
	}
	return p
}

func (p *Progress) Increase() *Progress {
	if p.step < p.max {
		p.step++
	}
	return p
}

func (p *Progress) Decrease() *Progress {
	if p.step > 0 {
		p.step--
	}
	return p
}

func (p *Progress) Fill() *Progress {
	p.step = p.max
	return p
}

func (p *Progress) Reset() *Progress {
	p.step = 0
	return p
}

func (p *Progress) Percent() float64 {
	return float64(p.step) / float64(p.max) * 100
}

func (p *Progress) String() string {
	b := strings.Builder{}
	b.WriteString(strconv.Itoa(int(p.step)))
	b.WriteString("/")
	b.WriteString(strconv.Itoa(int(p.max)))
	return b.String()
}

func NewProgress(max uint) *Progress {
	return &Progress{step: 0, max: max}
}
