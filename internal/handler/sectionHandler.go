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

func (h *SectionHandler) GetSection(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid userId",
		})
	}

	sectionId := c.Params("sectionId")
	if sectionId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing sectionId",
		})
	}

	resp, err := h.sectionService.GetSectionData(userId, sectionId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *SectionHandler) UpdateSectionProgress(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid userId",
		})
	}

	var req *dto.UpdateSectionProgress
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.LessonId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing lessonId",
		})
	}

	err := h.sectionService.UpdateSectionProgress(userId, req.LessonId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Section progress updated successfully",
	})
}
