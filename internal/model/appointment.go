package model

import (
	"time"
)

type Appointment struct {
	ExpertID           string `json:"expertId"`
	TimeSlotId         string `json:"timeSlotId"`
	MeetingLink        string `json:"meetingLink"`
	GoogleCalendarLink string `json:"googleCalendarLink"`
	UserId             string `json:"userId"`
}

type TimeSlot struct {
	Id        string    `json:"id"` // Firebase key
	ExpertId  string    `json:"expertId"`
	Date      time.Time `json:"date"`      // date only (00:00:00)
	StartTime time.Time `json:"startTime"` // full ISO 8601 timestamp
	EndTime   time.Time `json:"endTime"`   // unchanged (you can update similarly if needed)
	Available bool      `json:"available"`
}
