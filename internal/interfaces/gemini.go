// internal/interfaces/gemini.go
package interfaces

import (
	"context"

	"github.com/Harshal5167/Dapple-backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

type GeminiService interface {
	EvaluateAnswer(ctx context.Context, req *dto.EvaluationRequest) (*dto.EvaluationResponse, error)
}

type GeminiHandler interface {
	EvaluateAnswer(c *fiber.Ctx) error
}

type GeminiRoutes interface {
	SetupRoutes(app *fiber.App)
}
