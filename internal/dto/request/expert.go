package request

type AddExpertRequest struct {
	Name       string     `json:"name"`
	ImageURL   string     `json:"imageUrl,omitempty"`
	Bio        string     `json:"bio"`
	Schedule   []Schedule `json:"schedule"`
	XpRequired int        `json:"xpRequired"`
	Rating     float64    `json:"rating"`
}

type Schedule struct {
	Date      string     `json:"date"`
	TimeSlots []TimeSlot `json:"timeSlots"`
}

type TimeSlot struct {
	ExpertId   string `json:"expertId"`
	TimeSlotId string `json:"timeSlotId"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
	Available  bool   `json:"available"`
}
