
package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/Harshal5167/Dapple/internal/handler"
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
    api.Post("/evaluate-answer", r.geminiHandler.EvaluateAnswer)
}