package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shehabalqudiry/alqudiry_gocore/internal/modules/users/handlers"
)

func RegisterRoutes(
	app fiber.Router,
	handler *handlers.UserHandler,
) {

	group := app.Group("/users")

	group.Get("/", handler.Index)
}
