package examples_test

import (
	"fmt"
	"github.com/gol4ng/cli/loader/bar"
	"github.com/gol4ng/cli/loader/bar/styles"
)

func Example_loader_bar(){
	progress := bar.New(10)

	fmt.Println(progress.Render(0))
	fmt.Println(progress.Render(50))
	fmt.Println(progress.Render(100))

	// Output:
	// ░░░░░░░░░░
	// █████▒░░░░
	// ██████████
}

func Example_loader_bar_style(){
	progress := bar.NewSimple(10, styles.NewSimple('=', '-', ">"))

	fmt.Println(progress.Render(0))
	fmt.Println(progress.Render(50))
	fmt.Println(progress.Render(100))

	// Output:
	// ----------
	// =====>----
	// ==========
}
