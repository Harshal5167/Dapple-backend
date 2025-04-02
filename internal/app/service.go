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
		a.Repositories.EvaluationRepo,
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
		a.Repositories.TestRepo,
	)
	a.Services.LevelService = service.NewLevelService(
		a.Repositories.LevelRepo,
		a.Services.SectionService,
	)
	a.Services.EvaluationService = service.NewEvaluationService(
		a.Repositories.EvaluationRepo,
		a.Repositories.QuestionRepo,
		a.Services.GeminiService,
		a.Repositories.SectionRepo,
		a.Services.UserCourseService,
		a.Repositories.UserRepo,
		a.Repositories.TestRepo,
	)
	a.Services.UserService = service.NewUserService(
		a.Repositories.UserRepo,
	)
	a.Services.TestService = service.NewTestService(
		a.Services.GeminiService,
		a.Repositories.TestRepo,
		a.Repositories.SectionRepo,
		a.Repositories.QuestionRepo,
		a.Repositories.EvaluationRepo,
		a.Services.UserCourseService,
	)
	a.Services.SocketService = service.NewSocketService(
		a.Services.TestService,
	)
	a.Services.ExpertService = service.NewExpertService(
		a.Repositories.ExpertRepo,
	)

}
