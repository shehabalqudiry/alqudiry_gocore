package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func NewProject(name string) {

	root := name

	dirs := []string{
		"cmd/api",
		"internal/modules",
		"routes",
		"storage/logs",
		"storage/cache",
	}

	files := map[string]string{

		"README.md": "# " + name,

		".env": `APP_NAME=Alqudiry
APP_PORT=8000
`,

		"cmd/api/main.go": mainStub(),

		"routes/api.go": routeStub(),
	}

	for _, dir := range dirs {

		path := filepath.Join(root, dir)

		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			panic(err)
		}
	}

	for file, content := range files {

		path := filepath.Join(root, file)

		if err := os.WriteFile(
			path,
			[]byte(content),
			0644,
		); err != nil {
			panic(err)
		}
	}

	fmt.Printf(
		"✔ Project [%s] created successfully\n",
		name,
	)
}

func mainStub() string {

	return `package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Listen(":8000")
}
`
}

func routeStub() string {

	return `package routes

func RegisterRoutes() {

}
`
}