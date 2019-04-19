package examples_test

import (
	"fmt"
	"github.com/gol4ng/cli/loader"
)

func Example_progress() {
	progress := loader.NewProgress(10)

	fmt.Println(progress.Decrease(), progress.Percent())
	fmt.Println(progress.Fill(), progress.Percent())
	fmt.Println(progress.Increase(), progress.Percent())

	fmt.Println(progress.SetStep(5), progress.Percent())
	fmt.Println(progress.Decrease(), progress.Percent())
	fmt.Println(progress.Reset(), progress.Percent())

	fmt.Println(progress.SetMax(20), progress.Percent())
	fmt.Println(progress.SetStep(15), progress.Percent())
	fmt.Println(progress.SetMax(10), progress.Percent())

	// Output:
	// 0/10 0
	// 10/10 100
	// 10/10 100
	// 5/10 50
	// 4/10 40
	// 0/10 0
	// 0/20 0
	// 15/20 75
	// 10/10 100
}
