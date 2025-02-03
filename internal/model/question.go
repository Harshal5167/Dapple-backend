package model

type QuestionType string

const (
	Subjective QuestionType = "subjective"
	Objective  QuestionType = "objective"
)

type Question struct {
	Question      string       `json:"question"`
	ImageUrl      string       `json:"imageUrl,omitempty"`
	Type          QuestionType `json:"type"`
	Options       []string     `json:"options,omitempty"`
	CorrectOption int          `json:"correctOption,omitempty"`
	BestAnswer    []string     `json:"bestAnswer,omitempty"`
	SectionId     string       `json:"sectionId"`
	Explanation   []string     `json:"explanation,omitempty"`
	XP            int          `json:"xp"`
}
