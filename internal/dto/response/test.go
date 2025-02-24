package response

import "github.com/Harshal5167/Dapple-backend/internal/model"

type TestResultResponse struct {
	TotalXP        int                    `json:"totalXP"`
	TotalTimeTaken int                    `json:"totalTimeTaken"`
	QuestionResult []model.TestAnswerEval `json:"questionResult"`
}
