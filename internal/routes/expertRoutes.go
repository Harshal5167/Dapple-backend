package routes

import (
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

type ExpertRoutes struct {
	handler interfaces.ExpertHandler
}

func NewExpertRoutes(handler interfaces.ExpertHandler) *ExpertRoutes {
	return &ExpertRoutes{handler}
}

func (r *ExpertRoutes) ExpertRoutes(app *fiber.App) {
	api := app.Group("/api")
	expert := api.Group("/expert")
	expert.Post("/", r.handler.AddExpert)
	expert.Get("/:expertId", r.handler.GetExpertById)
	expert.Get("/:expertId/schedule", r.handler.GetExpertSchedule)
	expert.Get("/", r.handler.GetAllExperts)
}
