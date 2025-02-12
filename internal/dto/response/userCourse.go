package response

import (
	"github.com/Harshal5167/Dapple-backend/data"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type UserCourseResponse struct {
	Levels      []model.Level            `json:"levels"`
	SectionData []data.StaticSectionData `json:"sectionData"`
	UserProgess model.UserProgress       `json:"userProgress"`
}
