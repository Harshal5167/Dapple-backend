package app

import (
	"github.com/Harshal5167/Dapple-backend/internal/handler"
)

func (a *App) InitializeHandlers() {
	a.Handler.AuthHandler = handler.NewAuthHandler(a.Services.AuthService)
	a.Handler.LevelHandler = handler.NewLevelHandler(a.Services.LevelService)
	a.Handler.SectionHandler = handler.NewSectionHandler(a.Services.SectionService)
	a.Handler.QuestionHandler = handler.NewQuestionHandler(a.Services.QuestionService)
	a.Handler.EvaluationHandler = handler.NewEvaluationHandler(a.Services.EvaluationService, a.Services.QuestionService)
	a.Handler.LessonHandler = handler.NewLessonHandler(a.Services.LessonService)
	a.Handler.UserCourseHandler = handler.NewUserCourseHandler(a.Services.UserCourseService)
	a.Handler.EvaluationHandler = handler.NewEvaluationHandler(a.Services.EvaluationService, a.Services.QuestionService)
	a.Handler.UserHandler = handler.NewUserHandler(a.Services.UserService)
	// a.Handler.SocketHandler = handler.NewSocketHandler(a.Services.SocketService)
	a.Handler.TestHandler = handler.NewTestHandler(a.Services.TestService)
	a.Handler.ExpertHandler = handler.NewExpertHandler(a.Services.ExpertService)
	a.Handler.AppointmentHandler = handler.NewAppointmentHandler(a.Services.AppointmentService,)
}
