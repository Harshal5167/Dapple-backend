package handler

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

type TestHandler struct {
	testService interfaces.TestService
}

func NewTestHandler(testService interfaces.TestService) *TestHandler {
	return &TestHandler{
		testService: testService,
	}
}

func (h *TestHandler) GetTestResult(c *fiber.Ctx) error {
	sessionId := c.Query("sessionId")
	if sessionId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "sessionId is required",
		})
	}

	sectionId := c.Query("sectionId")
	if sectionId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "sectionId is required",
		})
	}

	userId, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "userId is required",
		})
	}

	result, err := h.testService.GetTestResult(userId, sessionId, sectionId)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (h *TestHandler) UploadImage(c *fiber.Ctx) error {
	var req *request.TestData
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.ImageUrl == "" || req.SessionId == "" || req.QuestionId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing or wrong fields",
		})
	}

	err := h.testService.EvaluateImageAnswer(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Image uploaded successfully",
	})
}

func (h *TestHandler) RetryQuestion(c *fiber.Ctx) error {
	var req *request.TestData
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.SessionId == "" || req.QuestionId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing or wrong fields",
		})
	}

	err := h.testService.RetryQuestion(req.SessionId, req.QuestionId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Question retried successfully",
	})
}

func (h *TestHandler) UploadText(c *fiber.Ctx) error {
	var req *request.TestData
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Answer == "" || req.SessionId == "" || req.QuestionId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing or wrong fields",
		})
	}

	_, err := h.testService.EvaluateTestAnswer(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Answer uploaded successfully",
	})
}
