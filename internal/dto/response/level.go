package response

type AddLevelResponse struct {
	LevelId string `json:"levelId"`
}

type LevelsForUser struct {
	SelectedLevelIds []string `json:"selectedLevelIds"`
}