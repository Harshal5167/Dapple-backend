package model

type Expert struct {
	Name       string     `json:"name"`
	ImageURL   string     `json:"imageUrl"`
	Bio        string     `json:"bio"`
	Schedule   []Schedule `json:"schedule,omitEmpty"`
	XpRequired int        `json:"xpRequired"`
	Rating     float64    `json:"rating"`
}

type Schedule struct {
	Date        string   `json:"date"`
	TimeSlotIds []string `json:"timeSlotIds"`
}

type TimeSlot struct {
	ExpertId  string `json:"expertId"`
	Date      string `json:"date"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Available bool   `json:"available"`
}
