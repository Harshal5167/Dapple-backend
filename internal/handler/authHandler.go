package handler

import (
	"bytes"
	"encoding/json"
	// "fmt"

	"github.com/Harshal5167/Dapple/internal/interfaces"
	"github.com/Harshal5167/Dapple/internal/model"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService interfaces.AuthService
}

func NewAuthHandler(authService interfaces.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var user = model.User{}
	if err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	if user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and Password are required",
		})
	}
	token, err := h.authService.Login(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var user = model.User{}
	if err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	if user.Email == "" || user.Password == "" || user.FirstName == "" || user.LastName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "First Name, Last Name, Email and Password are required",
		})
	}
	token, err := h.authService.Register(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (h *AuthHandler) LoginWithGoogle(c *fiber.Ctx) error {
	var user = model.User{}
	if err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	if user.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email is required",
		})
	}
	token, err := h.authService.LoginWithGoogle(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (h *AuthHandler) RegisterWithGoogle(c *fiber.Ctx) error {
	var user = model.User{}
	if err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	if (user.Email == "" || user.FirstName == "" || user.LastName == "") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "First Name, Last Name and Email are required",
		})
	}
	token, err := h.authService.RegisterWithGoogle(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
