package response

import "github.com/Harshal5167/Dapple-backend/internal/model"

type AddQuestionResponse struct {
	QuestionId string `json:"questionId"`
}

type EvaluateObjectiveAnswerResponse struct {
	CorrectOption int      `json:"correctOption"`
	Explanation   []string `json:"explanation"`
	XP            int      `json:"xp"`
}

type EvaluateSubjectiveAnswerResponse struct {
	Evaluation []model.Evaluation `json:"evaluation"`
	BestAnswer []string           `json:"bestAnswer"`
	UserAnswer []string           `json:"userAnswer"`
	XP         int                `json:"xp"`
}
