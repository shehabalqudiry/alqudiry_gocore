package ui

import (
	"fmt"
	"time"
)

var spinnerFrames = []string{
	"⠋",
	"⠙",
	"⠹",
	"⠸",
	"⠼",
	"⠴",
	"⠦",
	"⠧",
	"⠇",
	"⠏",
}

func Spinner(
	message string,
	duration time.Duration,
) {

	done := make(chan bool)

	go func() {

		for i := 0; ; i++ {

			select {

			case <-done:
				return

			default:

				frame := spinnerFrames[i%len(spinnerFrames)]

				fmt.Printf(
					"\r\033[1;36m%s\033[0m %s",
					frame,
					message,
				)

				time.Sleep(
					80 * time.Millisecond,
				)
			}
		}
	}()

	time.Sleep(duration)

	done <- true

	fmt.Print("\r\033[K")

	fmt.Printf(
		"\033[1;32m✔\033[0m %s\n",
		message,
	)
}