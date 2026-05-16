package templates

const HandlerTemplate = `package handlers

import (
	"github.com/gofiber/fiber/v2"

	"{{ .ModulePath }}/services"
)

type {{ .Module }}Handler struct {
	service *services.{{ .Module }}Service
}

func New{{ .Module }}Handler(
	service *services.{{ .Module }}Service,
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