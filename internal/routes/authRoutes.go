package routes

import (
	"github.com/Harshal5167/Dapple/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

type AuthRoute struct {
	handler interfaces.AuthHandler
}

func NewAuthRoute(handler interfaces.AuthHandler) *AuthRoute {
	return &AuthRoute{handler}
}

func (r *AuthRoute) AuthRoutes(app *fiber.App) {


	auth := app.Group("/auth")
	// auth.Post("/register", func (c *fiber.Ctx) error {
		
	// })
	auth.Post("/login", r.handler.Login)

	// googleAuth := auth.Group("/google")
	// googleAuth.Post("/register", )
	// googleAuth.Post("/login", )
}
