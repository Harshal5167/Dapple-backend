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
	geminiService interfaces.GeminiService
    geminiRoute   *routes.GeminiRoutes
}

func NewApp(config *config.Config) (app *App) {
	app = &App{
		config: config,
		Fiber:  fiber.New(),
	}
	app.setupRepositories()
	app.setupServices()
	app.setupRoutes()
	return
}

func (a *App) setupRepositories() {
	a.authRepo = auth.NewAuthRepository(a.config.FirebaseApp)
}

func (a *App) setupServices() {
	a.authService = service.NewAuthService(a.authRepo)
	a.geminiService = service.NewGeminiService(a.config.GeminiModel)

}


func (a *App) setupRoutes() {
    authHandler := handler.NewAuthHandler(a.authService)
    geminiHandler := handler.NewGeminiHandler(a.geminiService)
    
    a.authRoute = routes.NewAuthRoute(authHandler)
    a.geminiRoute = routes.NewGeminiRoutes(geminiHandler)
    
    a.authRoute.AuthRoutes(a.Fiber)
    a.geminiRoute.SetupRoutes(a.Fiber)
}