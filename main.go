package main

import (
	"github.com/Harshal5167/Dapple-backend/config"
	"github.com/Harshal5167/Dapple-backend/internal/app"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	firebaseApp, _ := config.InitializeFirebaseApp()

	rdb := config.InitializeRedis()
	config := config.NewConfig(firebaseApp)
	app := app.NewApp(config, rdb)
	app.Fiber.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Accept",
	}))

	app.Fiber.Listen(":8000")
}
