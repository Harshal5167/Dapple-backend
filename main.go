package main

import (
	"github.com/Harshal5167/Dapple/config"
	"github.com/Harshal5167/Dapple/internal/app"
)

func main() {
	app := app.Initialize()
	config.InitializeFirebaseApp()
	app.Listen(":5000")
}
