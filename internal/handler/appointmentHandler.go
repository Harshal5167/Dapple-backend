package handler

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

type AppointmentHandler struct {
	appointmentService interfaces.AppointmentService
}

func NewAppointmentHandler(appointmentService interfaces.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{
		appointmentService: appointmentService,
	}
}

func (h *AppointmentHandler) CreateAppointment(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid userId",
		})
	}

	var req *request.CreateAppointmentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.TimeSlotId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing or wrong fields",
		})
	}

	resp, err := h.appointmentService.CreateAppointment(req.TimeSlotId, userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *AppointmentHandler) GetAllAppointments(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid userId",
		})
	}

	resp, err := h.appointmentService.GetAllAppointments(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": resp,
	})
}

func (h *AppointmentHandler) GetAppointmentByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing or wrong fields",
		})
	}

	resp, err := h.appointmentService.GetAppointmentById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
