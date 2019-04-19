package cli_test

import (
	"github.com/gol4ng/cli"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Builder_SGR(t *testing.T) {
	tests := []struct {
		name       string
		attributes []cli.Attribute
		expected   string
	}{
		{name: "clean", expected: "test", attributes: []cli.Attribute{}},
		{name: "styles", expected: "\x1b[0;1;2;3;4;5;6;7;8;9mtest", attributes: []cli.Attribute{cli.Reset, cli.Bold, cli.Faint, cli.Italic, cli.Underline, cli.BlinkSlow, cli.BlinkRapid, cli.ReverseVideo, cli.Concealed, cli.CrossedOut}},
		{name: "colors", expected: "\x1b[39;30;31;32;33;34;35;36;37;90;91;92;93;94;95;96;97mtest", attributes: []cli.Attribute{cli.Default, cli.Black, cli.Red, cli.Green, cli.Yellow, cli.Blue, cli.Magenta, cli.Cyan, cli.White, cli.HiBlack, cli.HiRed, cli.HiGreen, cli.HiYellow, cli.HiBlue, cli.HiMagenta, cli.HiCyan, cli.HiWhite}},
		{name: "background colors", expected: "\x1b[49;40;41;42;43;44;45;46;47;100;101;102;103;104;105;106;107mtest", attributes: []cli.Attribute{cli.BgDefault, cli.BgBlack, cli.BgRed, cli.BgGreen, cli.BgYellow, cli.BgBlue, cli.BgMagenta, cli.BgCyan, cli.BgWhite, cli.BgHiBlack, cli.BgHiRed, cli.BgHiGreen, cli.BgHiYellow, cli.BgHiBlue, cli.BgHiMagenta, cli.BgHiCyan, cli.BgHiWhite}},
	}

	for _, test := range tests {
		builder := cli.NewStringBuilder()
		builder.Style(test.attributes...).WriteString("test")
		assert.Equal(t, test.expected, builder.String())
	}
}
