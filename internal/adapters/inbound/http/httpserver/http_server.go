package server

import (
	"github.com/gofiber/fiber/v2"
)

func NewServer() *fiber.App {
	fiber := fiber.New()
	return fiber
}
