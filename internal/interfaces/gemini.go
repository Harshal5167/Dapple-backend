// internal/interfaces/gemini.go
package interfaces

import (
    "context"
    "github.com/gofiber/fiber/v2"
)

type EvaluationRequest struct {
    Question           string   `json:"question"`
    UserAnswer         string   `json:"user_answer"`
    EvaluationCriteria []string `json:"evaluation_criteria"`
}

type EvaluationResponse struct {
    Evaluation map[string]string `json:"evaluation"`
    Feedback   map[string]string `json:"feedback"`
    Error      string           `json:"error,omitempty"`
}

type GeminiService interface {
    EvaluateAnswer(ctx context.Context, req *EvaluationRequest) (*EvaluationResponse, error)
}

type GeminiHandler interface {
    EvaluateAnswer(c *fiber.Ctx) error
}

type GeminiRoutes interface {
    SetupRoutes(app *fiber.App)
}