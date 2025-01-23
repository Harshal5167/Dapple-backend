package model

type User struct {
	UserId         string `json:"userId"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lirstName"`
	Age            int    `json:"age"`
	Role           string `json:"role"`
	TypeOfDisorder string `json:"typeOfDisorder"`
	XP             int    `json:"XP"`
	Badges         []int  `json:"badges"`
}
