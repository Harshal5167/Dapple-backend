package service

import (
	"errors"

	"github.com/Harshal5167/Dapple/internal/interfaces"
	"github.com/Harshal5167/Dapple/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepository interfaces.AuthRepository
}

func NewAuthService(authRepository interfaces.AuthRepository) *AuthService {
	return &AuthService{authRepository}
}

func (c *AuthService) Login(user model.User) (string,error) {
	isRegisteredEmail,err:=c.authRepository.CheckExistingEmail(user.Email); 
	if err!=nil {
		return "", err
	}
	if !isRegisteredEmail {
		return "", errors.New("email is not registered")
	}

	if err:=c.authRepository.CheckPassword(user.Email,user.Password); err!=nil {
		return "", err
	}
	
	userId,err:=c.authRepository.GetUserIdFromEmail(user.Email)
	if err!=nil {
		return "",err
	}

	token,err:=c.authRepository.GenerateCustomToken(userId, user)
	if err!=nil {
		return "", err
	}
	return token,nil
}

func (c *AuthService) Register(user model.User) (string, error) {
	isRegisteredEmail,err:=c.authRepository.CheckExistingEmail(user.Email); 
	if err!=nil {
		return "", err
	}
	if isRegisteredEmail {
		return "", errors.New("email is already registered")
	}

	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err!=nil{
		return "", err
	}
	user.Password = string(hashedPassword)
	userId, err := c.authRepository.CreateNewUser(user.Email, user.Password)
	if err != nil {
		return "", err
	}
	token,err:=c.authRepository.GenerateCustomToken(userId, user)
	if err!=nil {
		return "", err
	}
	return token,nil
}