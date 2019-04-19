package examples_test

import (
	"fmt"
	"github.com/gol4ng/cli/loader"
	"github.com/gol4ng/cli/loader/style"
)

func Example_progress_bar() {
	progress := loader.NewDefaultProgressBar(10, 10)

	fmt.Println(progress.Progress.String(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.SetStep(5), progress)

	fmt.Println(progress.Fill(), progress)

	fmt.Println(progress.SetMax(20), progress)
	fmt.Println(progress.SetStep(15), progress)
	fmt.Println(progress.SetMax(10), progress)

	// Output:

	// 0/10 ░░░░░░░░░░
	// 1/10 █▒░░░░░░░░
	// 5/10 █████▒░░░░
	// 10/10 ██████████
	// 10/20 █████▒░░░░
	// 15/20 ███████▒░░
	// 10/10 ██████████
}

func Example_progress_bar_custom() {
	progress := loader.NewProgressBar(10, loader.NewBarRenderer(10, style.NewBarSimple("=", "-", ">")))

	fmt.Println(progress.Progress.String(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.SetStep(5), progress)

	fmt.Println(progress.Fill(), progress)

	fmt.Println(progress.SetMax(20), progress)
	fmt.Println(progress.SetStep(15), progress)
	fmt.Println(progress.SetMax(10), progress)

	// Output:

	// 0/10 ----------
	// 1/10 =>--------
	// 5/10 =====>----
	// 10/10 ==========
	// 10/20 =====>----
	// 15/20 =======>--
	// 10/10 ==========
}

func Example_progress_bar_vertical_char() {
	progress := loader.NewProgressBar(8, loader.NewSliceBar(loader.VerticalChar))

	fmt.Println(progress.Progress.String(), progress, "|")
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)

	// Output:
	// 0/8   |
	// 1/8 ▁
	// 2/8 ▂
	// 3/8 ▃
	// 4/8 ▄
	// 5/8 ▅
	// 6/8 ▆
	// 7/8 ▇
	// 8/8 █
}

func Example_progress_bar_horizontal_char() {
	progress := loader.NewProgressBar(8, loader.NewSliceBar(loader.HorizontalChar))

	fmt.Println(progress.Progress.String(), progress, "|")
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)

	// Output:
	// 0/8   |
	// 1/8 ▏
	// 2/8 ▎
	// 3/8 ▍
	// 4/8 ▌
	// 5/8 ▋
	// 6/8 ▊
	// 7/8 ▉
	// 8/8 █
}

func Example_progress_dot_char() {
	progress := loader.NewProgressBar(8, loader.NewSliceBar(loader.DotChar))

	fmt.Println(progress.Progress.String(), progress, "|")
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)
	fmt.Println(progress.Increase(), progress)

	// Output:
	//0/8   |
	//1/8 ⠁
	//2/8 ⠉
	//3/8 ⠋
	//4/8 ⠛
	//5/8 ⠟
	//6/8 ⠿
	//7/8 ⡿
	//8/8 ⣿
}
