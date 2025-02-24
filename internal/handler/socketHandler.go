package handler

import (
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/fiber/v2"
)

type SocketHandler struct {
	socketService interfaces.SocketService
}

func NewSocketHandler(socketService interfaces.SocketService) *SocketHandler {
	return &SocketHandler{
		socketService: socketService,
	}
}

func (h *SocketHandler) HandleWebSocket() func(c *fiber.Ctx) error {
	return socketio.New(func(kws *socketio.Websocket) {
		h.socketService.HandleConnection(kws)
	})
}
