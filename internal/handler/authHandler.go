package handler

import (
	"bytes"
	"encoding/json"
	"fmt"

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


func (h *AuthHandler) Login(c *fiber.Ctx) error{
	var user = model.User{}
	if err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	if (user.Username == "" || user.Email == "" || user.Password == "") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username, Email and Password are required",
		})
	}
	h.authService.Login(user)
	fmt.Print(user)
	return c.SendString("Login")
}