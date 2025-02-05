package model

type Lesson struct {
	Title     string   `json:"title"`
	Content   []string `json:"content"`
	SectionId string   `json:"sectionId"`
	XP        int      `json:"xp"`
	ImageUrl  string   `json:"imageUrl,omitempty"`
}
