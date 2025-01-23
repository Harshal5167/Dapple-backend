package main

import (
	"github.com/Harshal5167/Dapple/config"
	"github.com/Harshal5167/Dapple/internal/app"
)

func main() {
	firebaseApp,_ := config.InitializeFirebaseApp()

	config := config.NewConfig(firebaseApp)

	app:= app.NewApp(config)
	app.Fiber.Listen(":8000")
}
	