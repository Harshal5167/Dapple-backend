package app

import (
	"github.com/Harshal5167/Dapple-backend/config"
	"github.com/Harshal5167/Dapple-backend/internal/app/types"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	config       *config.Config
	Fiber        *fiber.App
	Repositories *types.Repositories
	Services     *types.Services
	Handler      *types.Handler
}

func NewApp(config *config.Config) *App {
	app := &App{
		config:       config,
		Fiber:        fiber.New(),
		Repositories: &types.Repositories{},
		Services:     &types.Services{},
		Handler:      &types.Handler{},
	}

	app.InitializeRepositories()
	app.InitializeServices()
	app.InitializeHandlers()
	app.InitializeRoutes()

	return app
}
