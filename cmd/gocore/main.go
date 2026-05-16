package main

import (
	"fmt"
	"os"

	"github.com/shehabalqudiry/alqudiry_gocore/console/commands"
)

func main() {

	if len(os.Args) < 2 {
		help()
		return
	}

	command := os.Args[1]

	switch command {

	case "new":

		if len(os.Args) < 3 {
			fmt.Println("Project name required")
			os.Exit(1)
		}

		commands.NewProject(os.Args[2])

	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}

func help() {

	fmt.Println(`
╔══════════════════════════════╗
║      ALQUDIRY GOCORE        ║
╚══════════════════════════════╝

Commands:
    new <project-name>
`)
}