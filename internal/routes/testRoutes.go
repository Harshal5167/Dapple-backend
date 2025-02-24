package routes

import (
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

type TestRoutes struct {
	testHandler interfaces.TestHandler
}

func NewTestRoutes() *TestRoutes {
	return &TestRoutes{}
}

func (r *TestRoutes) TestRoutes(app *fiber.App) {
	api := app.Group("/api")
	test := api.Group("/test")
	test.Get("/result", r.testHandler.GetTestResult)
}
