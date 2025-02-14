package routes

import (
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type QuestionRoute struct {
	questionHandler interfaces.QuestionHandler
}

func NewQuestionRoute(questionHandler interfaces.QuestionHandler) *QuestionRoute {
	return &QuestionRoute{
		questionHandler: questionHandler,
	}
}

func (r *QuestionRoute) QuestionRoutes(app *fiber.App) {
	api := app.Group("/api")
	question := api.Group("/question")
	question.Post("/", r.questionHandler.AddQuestion)
	question.Get("/:questionId/hint", middleware.IsAuth, r.questionHandler.GetHint)
	evaluateAnswer := question.Group("/evaluate-answer", middleware.IsAuth)
	evaluateAnswer.Post("/objective", r.questionHandler.EvaluateObjectiveAnswer)
	evaluateAnswer.Post("/subjective", r.questionHandler.EvaluateSubjectiveAnswer)
}
