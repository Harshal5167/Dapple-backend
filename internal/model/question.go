package model

type QuestionType string

const (
	Subjective QuestionType = "subjective"
	Objective  QuestionType = "objective"
	Voice      QuestionType = "voice"
	Test       QuestionType = "test"
)

type Question struct {
	QuestionText    string          `json:"question"`
	ImageUrl        string          `json:"imageUrl,omitempty"`
	Type            QuestionType    `json:"type"`
	Hint            string          `json:"hint,omitempty"`
	Options         []string        `json:"options,omitempty"`
	CorrectOption   int             `json:"correctOption,omitzero"`
	BestAnswer      []string        `json:"bestAnswer,omitempty"`
	SectionId       string          `json:"sectionId,omitempty"`
	Explanation     []string        `json:"explanation,omitempty"`
	XP              int             `json:"xp"`
	EvaluationId    string          `json:"evaluationId,omitempty"`
}
