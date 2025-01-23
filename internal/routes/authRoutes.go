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
	auth.Post("/login", r.handler.Login)
	auth.Post("/register", r.handler.Register)

	googleAuth := auth.Group("/google")
	googleAuth.Post("/register", r.handler.RegisterWithGoogle)
	googleAuth.Post("/login", r.handler.LoginWithGoogle)
}
