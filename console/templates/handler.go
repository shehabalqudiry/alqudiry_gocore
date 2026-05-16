package templates

const HandlerTemplate = `package {{ .Package }}

import "github.com/gofiber/fiber/v2"

type {{ .Module }}Handler struct {
	service *{{ .Module }}Service
}

func New{{ .Module }}Handler(
	service *{{ .Module }}Service,
) *{{ .Module }}Handler {

	return &{{ .Module }}Handler{
		service: service,
	}
}

func (h *{{ .Module }}Handler) Index(
	c *fiber.Ctx,
) error {

	return c.JSON(fiber.Map{
		"success": true,
	})
}
`