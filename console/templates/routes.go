package templates

const RoutesTemplate = `package {{ .Package }}

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(
	app fiber.Router,
	handler *{{ .Module }}Handler,
) {

	group := app.Group("/{{ .Package }}")

	group.Get("/", handler.Index)
}
`