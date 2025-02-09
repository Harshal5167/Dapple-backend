package app

import (
	"github.com/Harshal5167/Dapple-backend/config"
	"github.com/Harshal5167/Dapple-backend/internal/handler"
	"github.com/Harshal5167/Dapple-backend/internal/repository"
	"github.com/Harshal5167/Dapple-backend/internal/routes"
	"github.com/Harshal5167/Dapple-backend/internal/service"
	"github.com/Harshal5167/Dapple-backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	config       *config.Config
	Fiber        *fiber.App
	Repositories *types.Repositories
	Services     *types.Services
	Handler      *types.Handler
}

func NewApp(config *config.Config) (app *App) {
	app = &App{
		config:       config,
		Fiber:        fiber.New(),
		Repositories: &types.Repositories{},
		Services:     &types.Services{},
		Handler:      &types.Handler{},
	}
	app.setupRepositories()
	app.setupServices()
	app.setupHandlers()
	app.setupRoutes()
	return
}

func (a *App) setupRepositories() {
	a.Repositories.AuthRepo = repository.NewAuthRepository(a.config.FirebaseApp)
	a.Repositories.LevelRepo = repository.NewLevelRepository(a.config.FirebaseApp)
	a.Repositories.SectionRepo = repository.NewSectionRepository(a.config.FirebaseApp)
	a.Repositories.QuestionRepo = repository.NewQuestionRepository(a.config.FirebaseApp)
	a.Repositories.LessonRepo = repository.NewLessonRepository(a.config.FirebaseApp)
	a.Repositories.UserCourseRepo = repository.NewUserCourseRepository(a.config.FirebaseApp)
}

func (a *App) setupServices() {
	a.Services.GeminiService = service.NewGeminiService(a.config.GeminiModel)
	a.Services.UserCourseService = service.NewUserCourseService(a.Repositories.UserCourseRepo, a.Repositories.LevelRepo, a.Services.GeminiService)
	a.Services.AuthService = service.NewAuthService(a.Repositories.AuthRepo, a.Services.UserCourseService)
	a.Services.LevelService = service.NewLevelService(a.Repositories.LevelRepo)
	a.Services.SectionService = service.NewSectionService(a.Repositories.SectionRepo, a.Repositories.LevelRepo, a.Repositories.QuestionRepo, a.Repositories.LessonRepo)
	a.Services.QuestionService = service.NewQuestionService(a.Repositories.QuestionRepo, a.Repositories.SectionRepo)
	a.Services.LessonService = service.NewLessonService(a.Repositories.LessonRepo, a.Repositories.SectionRepo)
}

func (a *App) setupHandlers() {
	a.Handler.AuthHandler = handler.NewAuthHandler(a.Services.AuthService)
	a.Handler.GeminiHandler = handler.NewGeminiHandler(a.Services.GeminiService)
	a.Handler.LevelHandler = handler.NewLevelHandler(a.Services.LevelService)
	a.Handler.SectionHandler = handler.NewSectionHandler(a.Services.SectionService, )
	a.Handler.QuestionHandler = handler.NewQuestionHandler(a.Services.QuestionService)
	a.Handler.LessonHandler = handler.NewLessonHandler(a.Services.LessonService)
	a.Handler.UserCourseHandler = handler.NewUserCourseHandler(a.Services.UserCourseService)
}

func (a *App) setupRoutes() {
	routes.NewAuthRoute(a.Handler.AuthHandler).AuthRoutes(a.Fiber)
	routes.NewGeminiRoutes(a.Handler.GeminiHandler).SetupRoutes(a.Fiber)
	routes.NewLevelRoute(a.Handler.LevelHandler).LevelRoutes(a.Fiber)
	routes.NewSectionRoutes(a.Handler.SectionHandler).SectionRoutes(a.Fiber)
	routes.NewQuestionRoute(a.Handler.QuestionHandler).QuestionRoutes(a.Fiber)
	routes.NewLessonRoutes(a.Handler.LessonHandler).LessonRoutes(a.Fiber)
	routes.NewUserCourseRoutes(a.Handler.UserCourseHandler).UserCourseRoutes(a.Fiber)
}
