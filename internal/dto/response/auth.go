package response

type AuthResponse struct {
	Token     string `json:"token"`
	FirstName string `json:"firstName"`
	XP        int    `json:"xp"`
}