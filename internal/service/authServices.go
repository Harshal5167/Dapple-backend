package service

import (
	"github.com/Harshal5167/Dapple/internal/interfaces"
	"github.com/Harshal5167/Dapple/internal/model"
)

type AuthService struct {
	authRepository interfaces.AuthRepository
}

func NewAuthService(authRepository interfaces.AuthRepository) *AuthService {
	return &AuthService{authRepository}
}

func (c *AuthService) Login(user model.User) error {
	c.authRepository.CheckExistingUser(user.Email, user.Username)
	return nil
}