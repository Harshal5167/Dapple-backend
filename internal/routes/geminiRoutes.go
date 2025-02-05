package routes

import (
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type GeminiRoutes struct {
	geminiHandler interfaces.GeminiHandler
}

func NewGeminiRoutes(handler interfaces.GeminiHandler) *GeminiRoutes {
	return &GeminiRoutes{
		geminiHandler: handler,
	}
}

func (r *GeminiRoutes) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	gemini := api.Group("/gemini", middleware.IsAuth)
	gemini.Post("/evaluate-answer", r.geminiHandler.EvaluateAnswer)
}
