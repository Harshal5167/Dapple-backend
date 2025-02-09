package main

import (
	"github.com/Harshal5167/Dapple-backend/config"
	"github.com/Harshal5167/Dapple-backend/internal/app"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	firebaseApp, _ := config.InitializeFirebaseApp()

	config := config.NewConfig(firebaseApp)

	app := app.NewApp(config)
	app.Fiber.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Accept",
	}))

	app.Fiber.Listen(":8000")
}
