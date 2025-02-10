package model

type Lesson struct {
	Title     string   `json:"title"`
	Content   []string `json:"content"`
	SectionId string   `json:"sectionId,omitempty"`
	XP        int      `json:"xp"`
	ImageUrl  string   `json:"imageUrl,omitempty"`
}
