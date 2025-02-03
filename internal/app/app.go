package app

import (
	"github.com/Harshal5167/Dapple-backend/config"
	"github.com/Harshal5167/Dapple-backend/internal/handler"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	auth "github.com/Harshal5167/Dapple-backend/internal/repository"
	"github.com/Harshal5167/Dapple-backend/internal/routes"
	"github.com/Harshal5167/Dapple-backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	config         *config.Config
	Fiber          *fiber.App
	authRepo       interfaces.AuthRepository
	authService    interfaces.AuthService
	authRoute      interfaces.AuthRoutes
	geminiService  interfaces.GeminiService
	geminiRoute    *routes.GeminiRoutes
	levelRoute     interfaces.LevelRoutes
	levelRepo      interfaces.LevelRepository
	levelService   interfaces.LevelService
	sectionRoute   interfaces.SectionRoutes
	sectionService interfaces.SectionService
	sectionRepo    interfaces.SectionRepository
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
	a.levelRepo = auth.NewLevelRepository(a.config.FirebaseApp)
	a.sectionRepo = auth.NewSectionRepository(a.config.FirebaseApp)
}

func (a *App) setupServices() {
	a.authService = service.NewAuthService(a.authRepo)
	a.geminiService = service.NewGeminiService(a.config.GeminiModel)
	a.levelService = service.NewLevelService(a.levelRepo)
	a.sectionService = service.NewSectionService(a.sectionRepo, a.levelRepo)
}

func (a *App) setupRoutes() {
	authHandler := handler.NewAuthHandler(a.authService)
	geminiHandler := handler.NewGeminiHandler(a.geminiService)
	levelHandler := handler.NewLevelHandler(a.levelService)
	sectionHandler := handler.NewSectionHandler(a.sectionService)

	a.authRoute = routes.NewAuthRoute(authHandler)
	a.geminiRoute = routes.NewGeminiRoutes(geminiHandler)
	a.levelRoute = routes.NewLevelRoute(levelHandler)
	a.sectionRoute = routes.NewSectionRoutes(sectionHandler)

	a.authRoute.AuthRoutes(a.Fiber)
	a.geminiRoute.SetupRoutes(a.Fiber)
	a.levelRoute.LevelRoutes(a.Fiber)
	a.sectionRoute.SectionRoutes(a.Fiber)
}
