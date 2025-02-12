package request

type RegisterRequest struct {
	Email                   string   `json:"email"`
	FirstName               string   `json:"firstName"`
	LastName                string   `json:"lastName"`
	FirebaseToken           string   `json:"firebaseToken"`
	Age                     int      `json:"age"`
	Gender                  string   `json:"gender"`
	Profession              string   `json:"profession"`
	SocialChallenges        []string `json:"socialChallenges"`
	StrugglingSocialSetting []string `json:"strugglingSocialSetting"`
}

type LoginRequest struct {
	Email         string `json:"email"`
	FirebaseToken string `json:"firebaseToken"`
}
