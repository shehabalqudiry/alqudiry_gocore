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
		case "help":
			help()
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

	fmt.Printf(`
		[1;36m
		╔══════════════════════════════════════════════════════╗
		║                                                      ║
		║              ALQUDIRY GOCORE FRAMEWORK              ║
		║                                                      ║
		║        Enterprise Backend Framework for Go          ║
		║                                                      ║
		╚══════════════════════════════════════════════════════╝
		[0m

		[1;33mUsage:[0m

		gocore <command> [arguments]



		[1;32mProject Commands[0m
		──────────────────────────────────────────────────────

		[1;36mnew <project-name>[0m

			Create a brand new GoCore project structure.

			Example:
			gocore new crm

			Generates:
			• Modular architecture
			• Internal framework structure
			• Routes
			• Storage
			• Config system
			• Bootstrap system



		[1;32mModule Commands[0m
		──────────────────────────────────────────────────────

		[1;36mmake:module <module-name>[0m

			Generate a clean modular structure.

			Example:
			gocore make:module users

			Generates:
			• handlers
			• services
			• repositories
			• entities
			• dto
			• validators
			• routes
			• swagger
			• tests



		[1;32mAPI Commands[0m
		──────────────────────────────────────────────────────

		[1;36mmake:api <api-name>[0m

			Generate full REST API scaffolding.

			Example:
			gocore make:api products

			Generates:
			• REST handlers
			• validators
			• swagger docs
			• repositories
			• routes
			• mappers



		[1;32mModel Commands[0m
		──────────────────────────────────────────────────────

		[1;36mmake:model <model-name> --fields=""[0m

			Generate smart schema-driven models.

			Example:
			gocore make:model users --fields="name:string,email:string,age:int"

			Supported Types:
			• string
			• int
			• bool
			• float
			• time

			Generates:
			• entity
			• DTOs
			• migration
			• repository interface
			• validators
			• swagger docs



		[1;32mExamples[0m
		──────────────────────────────────────────────────────

		gocore new ecommerce

		gocore make:module users

		gocore make:api auth

		gocore make:model products --fields="title:string,price:float,in_stock:bool"



		[1;35mFramework Architecture[0m
		──────────────────────────────────────────────────────

		• Modular Architecture
		• Layered Architecture
		• Clean Architecture Ready
		• DDD Ready
		• Enterprise Ready
		• REST + gRPC Ready
		• Event Driven Ready



		[1;30mVersion: v0.1.0[0m
`)
}