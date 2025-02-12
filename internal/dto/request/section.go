package request

type AddSectionRequest struct {
	Name    string `json:"name"`
	LevelId string `json:"levelId"`
	TotalXP int    `json:"totalXP"`
}

type UpdateSectionProgress struct {
	LessonId string `json:"lessonId"`
}