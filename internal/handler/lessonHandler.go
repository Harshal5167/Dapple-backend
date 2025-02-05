package handler

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

type LessonHandler struct {
	lessonService interfaces.LessonService
}

func NewLessonHandler(lessonService interfaces.LessonService) *LessonHandler {
	return &LessonHandler{
		lessonService: lessonService,
	}
}

func (h *LessonHandler) AddLesson(c *fiber.Ctx) error {
	var req *dto.AddLessonRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if req.Title == "" || len(req.Content) == 0 || req.SectionId == "" || req.XP == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	res, err := h.lessonService.AddLesson(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
