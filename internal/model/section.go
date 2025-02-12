package model

type Section struct {
	Name          string   `json:"name"`
	LevelId       string   `json:"levelId"`
	TotalXP       int      `json:"totalXP"`
	Lessons       []string `json:"lessons"`
	Questions     []string `json:"questions"`
	NextSectionId string   `json:"nextSectionId"`
}

type SectionData struct {
	Name      string     `json:"name"`
	TotalXP   int        `json:"totalXP"`
	Questions []Question `json:"questions"`
	Lessons   []Lesson   `json:"lessons,omitempty"`
}
type SectionProgress struct {
	Progress int `json:"progress" redis:"progress"`
	XP       int `json:"xp" redis:"xp"`
}
