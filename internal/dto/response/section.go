package response

import "github.com/Harshal5167/Dapple-backend/internal/model"

type AddSectionResponse struct {
	SectionId string `json:"sectionId"`
}

type SectionData struct {
	Data            []map[string]interface{} `json:"data"`
	SectionProgress model.SectionProgress    `json:"sectionProgress"`
}