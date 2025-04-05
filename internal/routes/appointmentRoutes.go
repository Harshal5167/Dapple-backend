package routes

import (
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type AppointmentRoutes struct {
	handler interfaces.AppointmentHandler
}

func NewAppointmentRoutes(handler interfaces.AppointmentHandler) *AppointmentRoutes {
	return &AppointmentRoutes{handler}
}

func (r *AppointmentRoutes) AppointmentRoutes(app *fiber.App) {
	api := app.Group("/api")
	appointment := api.Group("/appointment")
	appointment.Post("/create", middleware.IsAuth, r.handler.CreateAppointment)
	appointment.Get("/", middleware.IsAuth, r.handler.GetAllAppointments)
	appointment.Get("/:id", middleware.IsAuth, r.handler.GetAppointmentByID)
}
