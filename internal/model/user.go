package model

type User struct {
	UserId                  string   `json:"userId"`
	Email                   string   `json:"email"`
	FirstName               string   `json:"firstName"`
	LastName                string   `json:"lastName"`
	Age                     int      `json:"age"`
	Role                    string   `json:"role"`
	Profession              string   `json:"profession"`
	SocialChallenges        []string `json:"socialChallenges"`
	StrugglingSocialSetting []string `json:"strugglingSocialSetting"`
	XP                      int      `json:"xp"`
	Gender                  string   `json:"gender"`
}
