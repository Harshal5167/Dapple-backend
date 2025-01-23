package interfaces

import "github.com/Harshal5167/Dapple/internal/model"

type AuthRepository interface {
	CheckExistingEmail(email string) (bool, error)
	CreateNewUser(email string, password string) (string, error)
	GenerateCustomToken(uid string, user model.User) (string, error)
	CheckPassword(email string, password string) error
	GetUserIdFromEmail(email string) (string, error)
}
