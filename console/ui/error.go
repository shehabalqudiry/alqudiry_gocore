package ui

import (
	"fmt"
	"os"
)

func Error(message string) {

	fmt.Printf(
		"\n\033[1;31m✘ ERROR\033[0m %s\n\n",
		message,
	)
}

func Fatal(message string) {

	fmt.Printf(
		"\n\033[1;31m✘ FATAL\033[0m %s\n\n",
		message,
	)

	os.Exit(1)
}