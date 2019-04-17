package loader

import (
	"github.com/gol4ng/cli"
)

type ProgressBar struct {
	Progress
	renderer ProgressRenderer
}

func (pb *ProgressBar) Render() string {
	return pb.renderer.Render(pb.Progress)
}

func (pb *ProgressBar) RenderIn(buffer cli.Buffer) string {
	return pb.renderer.RenderIn(pb.Progress, buffer)
}

func (pb *ProgressBar) String() string {
	return pb.Render()
}

func NewDefaultProgressBar(max uint, length uint) *ProgressBar {
	return &ProgressBar{
		Progress: *NewProgress(max),
		renderer: NewDefaultBarRenderer(length),
	}
}

func NewProgressBar(max uint, renderer ProgressRenderer) *ProgressBar {
	return &ProgressBar{
		Progress: *NewProgress(max),
		renderer: renderer,
	}
}
