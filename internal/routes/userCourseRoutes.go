package routes

import (
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type UserCourseRoutes struct {
	handler interfaces.UserCourseHandler
}

func NewUserCourseRoutes(handler interfaces.UserCourseHandler) *UserCourseRoutes {
	return &UserCourseRoutes{
		handler: handler,
	}
}

func (r *UserCourseRoutes) UserCourseRoutes(app *fiber.App) {
	api := app.Group("/api")
	userCourse := api.Group("/userCourse", middleware.IsAuth)
	userCourse.Get("/", r.handler.GetUserCourse)
}
