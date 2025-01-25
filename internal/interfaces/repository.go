package interfaces

import (
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type AuthRepository interface {
	CreateNewUser(user model.User) (string, error)
	GetUserDetailsFromEmail(email string) (model.User, error)
	VerifyFirebaseToken(token string) (bool, error)
}
