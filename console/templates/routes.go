package templates

const RoutesTemplate = `package routes

import (
	"github.com/gofiber/fiber/v2"

	"{{ .ModulePath }}/handlers"
)

func RegisterRoutes(
	app fiber.Router,
	handler *{{ .Module }}Handler,
) {

	group := app.Group("/{{ .Package }}")

	group.Get("/", handler.Index)
}
`