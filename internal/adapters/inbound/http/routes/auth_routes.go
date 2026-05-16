package routes

import (
	handler "userAuth/internal/adapters/inbound/http/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(app *fiber.App, handler *handler.AuthHandler) {
	auth := app.Group("/api/v1/auth")

	auth.Post("/login", handler.HandleLogin)
}
