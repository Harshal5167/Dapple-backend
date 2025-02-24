package request

import "github.com/Harshal5167/Dapple-backend/internal/model"

type AddQuestionRequest struct {
	QuestionText    string                `json:"question"`
	ImageUrl        string                `json:"imageUrl,omitempty"`
	Type            model.QuestionType    `json:"type"`
	Options         []string              `json:"options,omitempty"`
	Hint            string                `json:"hint,omitempty"`
	CorrectOption   int                   `json:"correctOption,omitzero"`
	BestAnswer      []string              `json:"bestAnswer,omitempty"`
	SectionId       string                `json:"sectionId"`
	Explanation     []string              `json:"explanation,omitempty"`
	XP              int                   `json:"xp"`
	VoiceEvaluation model.VoiceEvaluation `json:"voiceEvaluation,omitempty"`
	VideoEvaluation model.Emotion         `json:"videoEvaluation,omitempty"`
}

type EvaluateSubjectiveAnswerReq struct {
	QuestionId string   `json:"questionId"`
	UserAnswer []string `json:"userAnswer"`
}

type EvaluateObjectiveAnswerReq struct {
	QuestionId     string `json:"questionId"`
	SelectedOption int    `json:"selectedOption"`
}

type EvaluateVoiceAnswerReq struct {
	QuestionId string `json:"questionId"`
}
