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

	case "make:module":
		if len(os.Args) < 3 {
			panic("Module name required")
		}

		commands.MakeModule(os.Args[2])
	case "make:api":
		if len(os.Args) < 3 {
			panic("API name required")
		}

		commands.MakeAPI(os.Args[2])

	case "make:model":

		if len(os.Args) < 4 {
			panic("Usage: make:model name --fields=")
		}

		name := os.Args[2]

		fields := ""

		for _, arg := range os.Args {

			if len(arg) > 9 &&
				arg[:9] == "--fields" {

				fields = arg[10:]
			}
		}

		commands.MakeModel(
			name,
			fields,
		)
	case "make":

		if len(os.Args) < 4 || os.Args[2] != "module" {
			fmt.Println("Usage: make module <module-name>")
			os.Exit(1)
		}

		commands.MakeModule(os.Args[3])

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
	make:module <module-name>
	make module <module-name>
`)
}
