package routes

import (
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type SocketRoutes struct {
	socketHandler interfaces.SocketHandler
}

func NewSocketRoutes(sh interfaces.SocketHandler) *SocketRoutes {
	return &SocketRoutes{
		socketHandler: sh,
	}
}

func (s *SocketRoutes) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	ws := api.Group("/ws", middleware.WebsocketUpgradeMiddleware, middleware.IsAuth)

	ws.Use("/", s.socketHandler.HandleWebSocket())
}
