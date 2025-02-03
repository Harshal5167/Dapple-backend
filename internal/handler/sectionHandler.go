package handler

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

type SectionHandler struct {
	sectionService interfaces.SectionService
}

func NewSectionHandler(sectionService interfaces.SectionService) *SectionHandler {
	return &SectionHandler{
		sectionService: sectionService,
	}
}

func (h *SectionHandler) AddSection(c *fiber.Ctx) error {
	var req *dto.AddSectionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Name == "" || req.LevelId == "" || req.TotalXP == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing required fields",
		})
	}

	resp, err := h.sectionService.AddSection(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

