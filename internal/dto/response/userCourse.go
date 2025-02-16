package response

import (
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type UserCourseResponse struct {
	Levels      []Level            `json:"levels"`
	UserProgess model.UserProgress `json:"userProgress"`
}

type Level struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"imageUrl,omitempty"`
	Sections    []Section `json:"sections"`
}

type Section struct {
	Name      string `json:"name"`
	TotalXP   int    `json:"totalXP"`
	SectionId string `json:"sectionId"`
}
