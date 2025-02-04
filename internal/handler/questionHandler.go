package handler

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

type QuestionHandler struct {
	questionService interfaces.QuestionService	
}

func NewQuestionHandler(questionService interfaces.QuestionService) *QuestionHandler {
	return &QuestionHandler{
		questionService: questionService,
	}
}

func (h *QuestionHandler) AddQuestion(c *fiber.Ctx) error {
	var req *dto.AddQuestionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.QuestionText == "" || req.XP == 0 || req.SectionId == "" || (req.Type != model.Objective && req.Type != model.Subjective) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing or wrong fields",
		})
	}

	resp, err := h.questionService.AddQuestion(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}