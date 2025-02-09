// internal/handler/gemini.go
package handler

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

type GeminiHandler struct {
	geminiService interfaces.GeminiService
}

func NewGeminiHandler(geminiService interfaces.GeminiService) *GeminiHandler {
	return &GeminiHandler{
		geminiService: geminiService,
	}
}

func (h *GeminiHandler) EvaluateAnswer(c *fiber.Ctx) error {
	var req *dto.EvaluationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Question == "" || req.UserAnswer == "" || len(req.EvaluationCriteria) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing required fields",
		})
	}

	resp, err := h.geminiService.EvaluateAnswer(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(resp)
}
