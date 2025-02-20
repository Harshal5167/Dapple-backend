package routes

import (
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type UserRoute struct {
	handler interfaces.UserHandler
}

func NewUserRoute(handler interfaces.UserHandler) *UserRoute {
	return &UserRoute{handler}
}

func (r *UserRoute) UserRoutes(app *fiber.App) {
	api := app.Group("/api")
	auth := api.Group("/user", middleware.IsAuth)
	auth.Post("/xp", r.handler.GetXP)
}
