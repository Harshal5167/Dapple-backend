package routes

import (
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type TestRoutes struct {
	testHandler interfaces.TestHandler
}

func NewTestRoutes(testHandler interfaces.TestHandler) *TestRoutes {
	return &TestRoutes{
		testHandler: testHandler,
	}
}

func (r *TestRoutes) TestRoutes(app *fiber.App) {
	api := app.Group("/api")
	test := api.Group("/test", middleware.IsAuth)
	test.Get("/result", r.testHandler.GetTestResult)
}
