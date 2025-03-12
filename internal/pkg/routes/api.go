package routes

import (
	"learn-portal/internal/registry"

	"github.com/gofiber/fiber/v2"
)

func NewAPIRoutes(app *fiber.App, handlers registry.Handler) {
	user := app.Group("/users")
	user.Post("/", handlers.User.Register)
	user.Get("/", handlers.User.GetAllUser)
	user.Get("/:id", handlers.User.GetUserProfile)
}
