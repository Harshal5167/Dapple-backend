package routes

import (
	"github.com/Harshal5167/Dapple-backend/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type GeminiRoutes struct {
	geminiHandler *handler.GeminiHandler
}

func NewGeminiRoutes(handler *handler.GeminiHandler) *GeminiRoutes {
	return &GeminiRoutes{
		geminiHandler: handler,
	}
}

func (r *GeminiRoutes) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	gemini := api.Group("/gemini")
	gemini.Post("/evaluate-answer", r.geminiHandler.EvaluateAnswer)
}
