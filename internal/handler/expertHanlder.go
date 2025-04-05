package handler

import (
	"fmt"

	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/gofiber/fiber/v2"

	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
)

type ExpertHandler struct {
	ExpertService interfaces.ExpertService
}

func NewExpertHandler(expertService interfaces.ExpertService) *ExpertHandler {
	return &ExpertHandler{
		ExpertService: expertService,
	}
}

func (h *ExpertHandler) AddExpert(c *fiber.Ctx) error {
	var req *request.AddExpertRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Name == "" || req.XpRequired == 0 || req.Bio == "" || req.Rating == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing or wrong fields",
		})
	}

	resp, err := h.ExpertService.AddExpert(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *ExpertHandler) GetExpertById(c *fiber.Ctx) error {
	expertId := c.Params("expertId")
	if expertId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing expertId",
		})
	}

	resp, err := h.ExpertService.GetExpertById(expertId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *ExpertHandler) GetExpertSchedule(c *fiber.Ctx) error {
	expertId := c.Params("expertId")
	if expertId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing expertId",
		})
	}
	fmt.Println("expertid: ", expertId)

	resp, err := h.ExpertService.GetExpertSchedule(expertId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *ExpertHandler) GetAllExperts(c *fiber.Ctx) error {
	resp, err := h.ExpertService.GetAllExperts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
