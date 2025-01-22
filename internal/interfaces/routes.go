package interfaces

import "github.com/gofiber/fiber/v2"

type AuthRoutes interface {
	AuthRoutes(app *fiber.App)
}