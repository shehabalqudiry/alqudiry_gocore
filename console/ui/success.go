package ui

import "fmt"

func Success(message string) {

	fmt.Printf(
		"\n\033[1;32m✔ SUCCESS\033[0m %s\n\n",
		message,
	)
}

func Done(message string) {

	fmt.Printf(
		"\033[1;32m[DONE]\033[0m %s\n",
		message,
	)
}