package ui

import (
	"fmt"
	"time"
)

func Info(message string) {

	fmt.Printf(
		"\033[1;36m[INFO]\033[0m %s\n",
		message,
	)

	time.Sleep(60 * time.Millisecond)
}

func Warning(message string) {

	fmt.Printf(
		"\033[1;33m[WARNING]\033[0m %s\n",
		message,
	)

	time.Sleep(60 * time.Millisecond)
}

func Debug(message string) {

	fmt.Printf(
		"\033[1;35m[DEBUG]\033[0m %s\n",
		message,
	)

	time.Sleep(60 * time.Millisecond)
}