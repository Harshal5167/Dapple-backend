package handler

import (
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"sessionId": sessionId,
		"sectionId": sectionId,
	})
	// userId, ok := c.Locals("userId").(string)
	// if !ok {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": "userId is required",
	// 	})
	// }

	// result, err := h.testService.GetTestResult(userId, sessionId, sectionId)
	// if err != nil {
	// 	return err
	// }
	// return c.Status(fiber.StatusOK).JSON(result)
}
