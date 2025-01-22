package app

import (
	"github.com/Harshal5167/Dapple/config"
	"github.com/Harshal5167/Dapple/internal/handler"
	"github.com/Harshal5167/Dapple/internal/interfaces"
	"github.com/Harshal5167/Dapple/internal/repository/auth"
	"github.com/Harshal5167/Dapple/internal/routes"
	"github.com/Harshal5167/Dapple/internal/service"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	config      *config.Config
	Fiber       *fiber.App
	authRepo    interfaces.AuthRepository
	authService interfaces.AuthService
	authRoute   interfaces.AuthRoutes
}

func NewApp(config *config.Config) *App {
	app := &App{
		config: config,
		Fiber:  fiber.New(),
	}
	app.setupRepositories()
	app.setupServices()
	app.setupRoutes()
	return app
}

func (a *App) setupRepositories() {
	a.authRepo = auth.NewAuthRepository(a.config.FirebaseApp)
}

func (a *App) setupServices() {
	a.authService = service.NewAuthService(a.authRepo)
}

func (a *App) setupRoutes() {
	authHandler := handler.NewAuthHandler(a.authService)
	a.authRoute = routes.NewAuthRoute(authHandler)
	a.authRoute.AuthRoutes(a.Fiber)
}