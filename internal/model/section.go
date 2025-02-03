package model

type Section struct {
	Name          string   `json:"name"`
	LevelId       string   `json:"levelId"`
	TotalXP       int      `json:"totalXP"`
	Lessons       []string `json:"lessons"`
	Questions     []string `json:"questions"`
	NextSectionId string   `json:"nextSectionId"`
}
