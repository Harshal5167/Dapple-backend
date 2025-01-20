package app

import (
	"github.com/Harshal5167/Dapple/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func Initialize() *fiber.App {
	app := fiber.New()

	routes.AuthRoutes(app)

	return app
}
