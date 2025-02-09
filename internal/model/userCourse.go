package model

type UserCourse struct {
	Levels       []string     `json:"levels"`
	UserProgress UserProgress `json:"userProgress"`
}

type UserProgress struct {
	CompletedLevels   int `json:"completedLevels"`
	CompletedSections int `json:"completedSections"`
}
