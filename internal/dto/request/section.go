package request

import "github.com/Harshal5167/Dapple-backend/internal/model"

type AddSectionRequest struct {
	Name    string `json:"name"`
	LevelId string `json:"levelId"`
	TotalXP int    `json:"totalXP"`
}

type UpdateSectionProgress struct {
	LessonId string `json:"lessonId"`
}

type SectionData struct {
	Name      string               `json:"name"`
	TotalXP   int                  `json:"totalXP"`
	Questions []AddQuestionRequest `json:"questions"`
	Lessons   []model.Lesson       `json:"lessons,omitempty"`
}
