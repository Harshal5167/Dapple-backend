package routes

import (
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

type SectionRoutes struct {
	sectionHandler interfaces.SectionHandler	
}

func NewSectionRoutes(handler interfaces.SectionHandler) *SectionRoutes {
	return &SectionRoutes{handler}
}

func (r *SectionRoutes) SectionRoutes(app *fiber.App) {
	api := app.Group("/api")
	section := api.Group("/section")
	section.Post("/", r.sectionHandler.AddSection)
}