package dto

import "github.com/Harshal5167/Dapple-backend/internal/model"

type RegisterRequest struct {
	Email                   string   `json:"email"`
	FirstName               string   `json:"firstName"`
	LastName                string   `json:"lastName"`
	FirebaseToken           string   `json:"firebaseToken"`
	Age                     int      `json:"age"`
	Gender                  string   `json:"gender"`
	Profession              string   `json:"profession"`
	SocialChallenges        []string `json:"socialChallenges"`
	StrugglingSocialSetting []string `json:"strugglingSocialSetting"`
}

type LoginRequest struct {
	Email         string `json:"email"`
	FirebaseToken string `json:"firebaseToken"`
}

type EvaluationRequest struct {
	Question           string   `json:"question"`
	UserAnswer         string   `json:"user_answer"`
	EvaluationCriteria []string `json:"evaluation_criteria"`
}

type AddLevelRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl,omitempty"`
}

type AddSectionRequest struct {
	Name    string `json:"name"`
	LevelId string `json:"levelId"`
	TotalXP int    `json:"totalXP"`
}

type AddQuestionRequest struct {
	QuestionText  string             `json:"question"`
	ImageUrl      string             `json:"imageUrl,omitempty"`
	Type          model.QuestionType `json:"type"`
	Options       []string           `json:"options,omitempty"`
	CorrectOption int                `json:"correctOption,omitzero"`
	BestAnswer    []string           `json:"bestAnswer,omitempty"`
	SectionId     string             `json:"sectionId"`
	Explanation   []string           `json:"explanation,omitempty"`
	XP            int                `json:"xp"`
}

type AddLessonRequest struct {
	Title     string   `json:"title"`
	Content   []string `json:"content"`
	SectionId string   `json:"sectionId"`
	XP        int      `json:"xp"`
	ImageUrl  string   `json:"imageUrl,omitempty"`
}

type AddCompleteLevelRequest struct {
	Level    model.Level         `json:"level"`
	Sections []model.SectionData `json:"sections"`
}

type UpdateSectionProgress struct {
	LessonId string `json:"lessonId"`
}

type EvaluateSubjectiveAnswerReq struct {
	QuestionId string   `json:"questionId"`
	UserAnswer []string `json:"userAnswer"`
}

type EvaluateObjectiveAnswerReq struct {
	QuestionId     string `json:"questionId"`
	SelectedOption int    `json:"selectedOption"`
}
