package examples_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/gol4ng/cli"
)

func Test_ansi_sgr(t *testing.T) {
	w := cli.NewStringBuilder().

		// add some content
		Write("first content\n").

		// add some content with style
		Write("stylize content\n", cli.Italic, cli.Yellow).
		Write("style resetted\n").
		Red("style red\n").
		Bold("style bold\n").

		// begin style and write content
		Style(cli.Underline, cli.Red).WriteString("apply a style\n").
		WriteString("previous style continue until Reset was called or another style apply\n").

		// new style began
		Style(cli.Bold, cli.Blue).WriteString("new style\n")

	// manually reset style
	w.Reset().Write("style manually resetted\n")

	assert.Equal(
		t,
		"first content\n\x1b[m\x1b[3;33mstylize content\n\x1b[mstyle resetted\n\x1b[m\x1b[31mstyle red\n\x1b[m\x1b[1mstyle bold\n\x1b[m\x1b[4;31mapply a style\nprevious style continue until Reset was called or another style apply\n\x1b[1;34mnew style\n\x1b[mstyle manually resetted\n\x1b[m",
		w.String(),
	)
}
