package interfaces

import "github.com/Harshal5167/Dapple/internal/model"

type AuthService interface {
	Login(user model.User) (string, error)
	Register(user model.User) (string, error)
	RegisterWithGoogle(user model.User) (string, error)
	LoginWithGoogle(user model.User) (string, error)
}
