package model

type User struct {
	UserId         string `json:"userId"`
	Email          string `json:"email"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Age            int    `json:"age"`
	Role           string `json:"role"`
	XP             int    `json:"XP"`
	Badges         []int  `json:"badges"`
	Level          int    `json:"level"`
	Section        int    `json:"section"`
}
