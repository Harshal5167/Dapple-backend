package routes

import (
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

type LevelRoute struct {
	handler interfaces.LevelHandler
}

func NewLevelRoute(handler interfaces.LevelHandler) *LevelRoute {
	return &LevelRoute{handler}
}

func (r *LevelRoute) LevelRoutes(app *fiber.App) {
	api := app.Group("/api")
	level:= api.Group("/level")
	level.Post("/", r.handler.AddLevel)
}
