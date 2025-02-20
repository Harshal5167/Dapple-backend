package handler

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
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
	var req *request.AddQuestionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Type != model.Objective && req.Type != model.Subjective {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid question type",
		})
	}

	if req.QuestionText == "" || req.XP == 0 || req.SectionId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing or wrong fields",
		})
	}

	if req.Type == model.Objective && (req.CorrectOption < 0 || req.CorrectOption >= len(req.Options) || len(req.Options) < 4 || len(req.Explanation) == 0) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid objective question",
		})
	}

	if (req.Type == model.Subjective || req.Type == model.Test || req.Type == model.Voice) && (len(req.BestAnswer) == 0) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid subjective question",
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

func (h *QuestionHandler) GetHint(c *fiber.Ctx) error {
	questionId := c.Params("questionId")
	if questionId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing questionId",
		})
	}

	resp, err := h.questionService.GetHint(questionId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
