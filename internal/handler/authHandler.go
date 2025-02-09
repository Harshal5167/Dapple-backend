package handler

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService interfaces.AuthService
}

func NewAuthHandler(authService interfaces.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var reqBody *dto.LoginRequest

	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	if reqBody.Email == "" || reqBody.FirebaseToken == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Required fields are missing",
		})
	}
	response, err := h.authService.Login(reqBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var reqBody *dto.RegisterRequest

	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	if reqBody.Email == "" ||
		reqBody.FirstName == "" ||
		reqBody.LastName == "" ||
		reqBody.FirebaseToken == "" ||
		reqBody.Age == 0 ||
		reqBody.Gender == "" ||
		reqBody.Profession == "" ||
		len(reqBody.SocialChallenges) == 0 ||
		len(reqBody.StrugglingSocialSetting) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Required fields are missing",
		})
	}

	response, err := h.authService.Register(reqBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
