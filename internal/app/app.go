package app

import (
	"github.com/Harshal5167/Dapple-backend/config"
	"github.com/Harshal5167/Dapple-backend/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type App struct {
	config       *config.Config
	rdb          *redis.Client
	Fiber        *fiber.App
	Repositories *types.Repositories
	Services     *types.Services
	Handler      *types.Handler
}

func NewApp(config *config.Config, rdb *redis.Client) *App {
	app := &App{
		config:       config,
		rdb:          rdb,
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
