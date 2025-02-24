package middleware

import (
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func WebsocketUpgradeMiddleware(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		fmt.Println("Websocket upgrade")
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}
