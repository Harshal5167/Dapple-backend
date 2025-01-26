package service

import (
	"errors"

	"github.com/Harshal5167/Dapple-backend/internal/dto"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
	"github.com/Harshal5167/Dapple-backend/internal/utils"
)

type AuthService struct {
	authRepository interfaces.AuthRepository
}

func NewAuthService(authRepository interfaces.AuthRepository) *AuthService {
	return &AuthService{authRepository}
}

func (c *AuthService) Login(reqBody *dto.LoginRequest) (*dto.AuthResponse, error) {
	isVerified, tokenEmail, err := c.authRepository.VerifyFirebaseToken(reqBody.FirebaseToken)
	if err != nil {
		return nil, err
	}
	if !isVerified {
		return nil, errors.New("firebase token is not verified")
	}
	if tokenEmail != reqBody.Email {
		return nil, errors.New("email in token does not match email in request body")
	}

	user, err := c.authRepository.GetUserDetailsFromEmail(reqBody.Email)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateJWTToken(model.User{
		UserId: user.UserId,
		Email:  user.Email,
	})
	if err != nil {
		return nil, err
	}
	return &dto.AuthResponse{
			Token:     token,
			FirstName: user.FirstName,
			Level:     user.Level,
			Section:   user.Section},
		nil
}

func (c *AuthService) Register(reqBody *dto.RegisterRequest) (*dto.AuthResponse, error) {
	isVerified, tokenEmail, err := c.authRepository.VerifyFirebaseToken(reqBody.FirebaseToken)
	if err != nil {
		return nil, err
	}
	if !isVerified {
		return nil, errors.New("firebase token is not verified")
	}
	if tokenEmail != reqBody.Email {
		return nil, errors.New("email in token does not match email in request body")
	}

	userId, err := c.authRepository.CreateNewUser(model.User{
		Email:     reqBody.Email,
		FirstName: reqBody.FirstName,
		LastName:  reqBody.LastName,
		Level:     0,
		Section:   0,
	})
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateJWTToken(model.User{
		UserId: userId,
		Email:  reqBody.Email,
	})
	if err != nil {
		return nil, err
	}
	return &dto.AuthResponse{
			Token:     token,
			FirstName: reqBody.FirstName,
			Level:     0,
			Section:   0},
		nil
}