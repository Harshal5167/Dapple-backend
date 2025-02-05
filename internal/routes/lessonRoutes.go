package routes

import (
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

type LessonRoutes struct {
	handler interfaces.LessonHandler
}

func NewLessonRoutes(handler interfaces.LessonHandler) *LessonRoutes {
	return &LessonRoutes{
		handler: handler,
	}
}

func (r *LessonRoutes) LessonRoutes(app *fiber.App) {
	api := app.Group("/api")
	lesson := api.Group("/lesson")
	lesson.Post("/", r.handler.AddLesson)
}
