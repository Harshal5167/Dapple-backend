// internal/interfaces/gemini.go
package interfaces

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto"
	"github.com/Harshal5167/Dapple-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

type GeminiService interface {
	EvaluateAnswer(req *dto.EvaluationRequest) (*dto.EvaluationResponse, error)
	GenerateUserCourse(user model.User, levelDetails []map[string]string) (*dto.LevelsForUser, error)
}

type GeminiHandler interface {
	EvaluateAnswer(c *fiber.Ctx) error
}

type GeminiRoutes interface {
	SetupRoutes(app *fiber.App)
}
