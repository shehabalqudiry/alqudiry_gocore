package ui

import (
	"fmt"
	"strings"
	"time"
)

func Progress(
	current int,
	total int,
	file string,
) {

	width := 32

	progress := int(
		float64(current) /
			float64(total) *
			float64(width),
	)

	bar := strings.Repeat(
		"█",
		progress,
	)

	bar += strings.Repeat(
		"░",
		width-progress,
	)

	percent := int(
		float64(current) /
			float64(total) * 100,
	)

	fmt.Print("\r\033[K")

	fmt.Printf(
		"\033[1;36m[%s]\033[0m \033[1;32m%3d%%\033[0m (%d/%d) → \033[1;37m%s\033[0m",
		bar,
		percent,
		current,
		total,
		file,
	)

	time.Sleep(100 * time.Millisecond)

	if current == total {
		fmt.Println()
	}
}