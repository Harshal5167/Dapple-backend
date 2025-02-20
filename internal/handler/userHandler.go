package handler

import (
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct{
	userService interfaces.UserService
}

func NewUserHandler(userService interfaces.UserService) *UserHandler{
	return &UserHandler{userService}
}

func (h *UserHandler) GetXP(c *fiber.Ctx) error{
	userId, ok := c.Locals("userId").(string)
	if !ok{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	resp, err := h.userService.GetXP(userId)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}