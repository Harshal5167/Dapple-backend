package handler

import (
	"bytes"
	"fmt"
	// "fmt"
	"io"
	// "slices"

	"github.com/Harshal5167/Dapple-backend/config"
	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

type EvaluationHandler struct {
	EvaluationService interfaces.EvaluationService
	QuestionService   interfaces.QuestionService
}

func NewEvaluationHandler(evaluationService interfaces.EvaluationService, questionService interfaces.QuestionService) *EvaluationHandler {
	return &EvaluationHandler{
		EvaluationService: evaluationService,
		QuestionService:   questionService,
	}
}

func (h *EvaluationHandler) EvaluateObjectiveAnswer(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid userId",
		})
	}

	var req *request.EvaluateObjectiveAnswerReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.QuestionId == "" || req.SelectedOption < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing required fields",
		})
	}

	resp, err := h.EvaluationService.EvaluateObjectiveAnswer(userId, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *EvaluationHandler) EvaluateSubjectiveAnswer(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid userId",
		})
	}

	var req *request.EvaluateSubjectiveAnswerReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.QuestionId == "" || len(req.UserAnswer) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing required fields",
		})
	}

	resp, err := h.EvaluationService.EvaluateSubjectiveAnswer(userId, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *EvaluationHandler) EvaluateVoiceAnswer(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid userId",
		})
	}

	var req = &request.EvaluateVoiceAnswerReq{}
	req.QuestionId = c.FormValue("questionId")
	if req.QuestionId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing QuestionId",
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid file",
		})
	}

	fmt.Println(file.Header.Get("Content-Type"))
	if req.QuestionId == "" || int(file.Size) > config.MaxFileSize {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Fields",
		})
	}

	var buf = &bytes.Buffer{}
	fileReader, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read file",
		})
	}
	defer fileReader.Close()

	if _, err := io.Copy(buf, fileReader); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read file",
		})
	}

	// resp, err := clients.VoiceEvaluation(buf.Bytes())
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": err.Error(),
	// 	})
	// }
	// return c.Status(fiber.StatusOK).JSON(resp)

	resp, err := h.EvaluationService.EvaluateVoiceAnswer(userId, req, buf.Bytes())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
