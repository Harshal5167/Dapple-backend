package main

import (
	"github.com/Harshal5167/Dapple-backend/config"
	"github.com/Harshal5167/Dapple-backend/internal/app"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config := config.NewConfig()
	app := app.NewApp(config)
	app.Fiber.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Accept",
	}))

	app.Fiber.Listen("127.0.0.1:8000")
}
