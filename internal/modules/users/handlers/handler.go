package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shehabalqudiry/alqudiry_gocore/internal/modules/users/services"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(
	service *services.UserService,
) *UserHandler {

	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Index(
	c *fiber.Ctx,
) error {

	return c.JSON(fiber.Map{
		"success": true,
	})
}
