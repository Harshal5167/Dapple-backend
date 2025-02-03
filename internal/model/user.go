package model

type User struct {
	UserId           string   `json:"userId"`
	Email            string   `json:"email"`
	FirstName        string   `json:"firstName"`
	LastName         string   `json:"lastName"`
	Age              int      `json:"age"`
	Role             string   `json:"role"`
	Profession       string   `json:"profession"`
	SocialChallenges []string `json:"socialChallenges"`
	SocialSettings   []string `json:"socialSettings"`
	XP               int      `json:"XP"`
	Gender           string   `json:"gender"`
	Level            int      `json:"level"`
	Section          int      `json:"section"`
}
