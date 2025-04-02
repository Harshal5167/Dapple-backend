package response

import "github.com/Harshal5167/Dapple-backend/internal/model"

type EvaluateObjectiveAnswerResponse struct {
	CorrectOption     int      `json:"correctOption"`
	Explanation       []string `json:"explanation"`
	XP                int      `json:"xp"`
	TotalXP           int      `json:"totalXP,omitempty"`
	TotalTimeTaken    int64    `json:"totalTimeTaken,omitempty"`
	CompletedLevels   int      `json:"completedLevels,omitempty"`
	CompletedSections int      `json:"completedSections,omitempty"`
}

type EvaluateSubjectiveAnswerResponse struct {
	Evaluation        []model.Evaluation `json:"evaluation"`
	BestAnswer        []string           `json:"bestAnswer"`
	UserAnswer        []string           `json:"userAnswer"`
	XP                int                `json:"xp"`
	TotalXP           int                `json:"totalXP,omitempty"`
	TotalTimeTaken    int64              `json:"totalTimeTaken,omitempty"`
	CompletedLevels   int                `json:"completedLevels,omitempty"`
	CompletedSections int                `json:"completedSections,omitempty"`
}

type EvaluateVoiceAnswerResponse struct {
	Evaluation        []model.Evaluation `json:"evaluation"`
	XP                int                `json:"xp"`
	TotalXP           int                `json:"totalXP,omitempty"`
	TotalTimeTaken    int64              `json:"totalTimeTaken,omitempty"`
	CompletedLevels   int                `json:"completedLevels,omitempty"`
	CompletedSections int                `json:"completedSections,omitempty"`
}
