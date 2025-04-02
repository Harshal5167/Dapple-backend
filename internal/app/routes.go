package app

import (
	"github.com/Harshal5167/Dapple-backend/internal/routes"
)

func (a *App) InitializeRoutes() {
	routes.NewAuthRoute(a.Handler.AuthHandler).AuthRoutes(a.Fiber)
	routes.NewLevelRoute(a.Handler.LevelHandler).LevelRoutes(a.Fiber)
	routes.NewSectionRoutes(a.Handler.SectionHandler).SectionRoutes(a.Fiber)
	routes.NewQuestionRoute(a.Handler.QuestionHandler, a.Handler.EvaluationHandler).QuestionRoutes(a.Fiber)
	routes.NewLessonRoutes(a.Handler.LessonHandler).LessonRoutes(a.Fiber)
	routes.NewUserCourseRoutes(a.Handler.UserCourseHandler).UserCourseRoutes(a.Fiber)
	routes.NewUserRoute(a.Handler.UserHandler).UserRoutes(a.Fiber)
	routes.NewSocketRoutes(a.Handler.SocketHandler).SetupRoutes(a.Fiber)
	routes.NewTestRoutes(a.Handler.TestHandler).TestRoutes(a.Fiber)
	routes.NewExpertRoutes(a.Handler.ExpertHandler).ExpertRoutes(a.Fiber)
}
