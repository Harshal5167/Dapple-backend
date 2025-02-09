package handler

import (
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

type UserCourseHandler struct {
	UserCourseService interfaces.UserCourseService
}

func NewUserCourseHandler(service interfaces.UserCourseService) interfaces.UserCourseHandler {
	return &UserCourseHandler{
		UserCourseService: service,
	}
}

func (h *UserCourseHandler) GetUserCourse(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "userId not found",
		})
	}

	userCourse, err := h.UserCourseService.GetUserCourse(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(userCourse)
}
