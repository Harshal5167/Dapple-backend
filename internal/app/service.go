package app

import (
	"github.com/Harshal5167/Dapple-backend/internal/service"
)

func (a *App) InitializeServices() {
	a.Services.GeminiService = service.NewGeminiService(a.config.Genai)
	a.Services.UserCourseService = service.NewUserCourseService(
		a.Repositories.UserCourseRepo,
		a.Repositories.LevelRepo,
		a.Services.GeminiService,
		a.Repositories.SectionRepo,
		a.Repositories.UserRepo,
	)
	a.Services.QuestionService = service.NewQuestionService(
		a.Repositories.QuestionRepo,
		a.Repositories.SectionRepo,
		a.Services.GeminiService,
		a.Repositories.UserRepo,
		a.Services.UserCourseService,
	)
	a.Services.LessonService = service.NewLessonService(
		a.Repositories.LessonRepo,
		a.Repositories.SectionRepo,
	)
	a.Services.AuthService = service.NewAuthService(
		a.Repositories.AuthRepo,
		a.Services.UserCourseService,
		a.Repositories.UserRepo,
	)
	a.Services.SectionService = service.NewSectionService(
		a.Repositories.SectionRepo,
		a.Repositories.LevelRepo,
		a.Repositories.QuestionRepo,
		a.Repositories.LessonRepo,
		a.Services.QuestionService,
		a.Services.LessonService,
	)
	a.Services.LevelService = service.NewLevelService(
		a.Repositories.LevelRepo,
		a.Services.SectionService,
	)
}
