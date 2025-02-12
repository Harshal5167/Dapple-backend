package app

import (
	"github.com/Harshal5167/Dapple-backend/internal/handler"
)

func (a *App) InitializeHandlers() {
	a.Handler.AuthHandler = handler.NewAuthHandler(a.Services.AuthService)
	a.Handler.LevelHandler = handler.NewLevelHandler(a.Services.LevelService)
	a.Handler.SectionHandler = handler.NewSectionHandler(a.Services.SectionService)
	a.Handler.QuestionHandler = handler.NewQuestionHandler(a.Services.QuestionService)
	a.Handler.LessonHandler = handler.NewLessonHandler(a.Services.LessonService)
	a.Handler.UserCourseHandler = handler.NewUserCourseHandler(a.Services.UserCourseService)
}
