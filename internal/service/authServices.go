package service

import (
	"errors"

	"github.com/Harshal5167/Dapple-backend/internal/dto"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
	"github.com/Harshal5167/Dapple-backend/internal/utils"
)

type AuthService struct {
	authRepository    interfaces.AuthRepository
	userCourseService interfaces.UserCourseService
}

func NewAuthService(authRepository interfaces.AuthRepository, userCourseService interfaces.UserCourseService) *AuthService {
	return &AuthService{
		authRepository:    authRepository,
		userCourseService: userCourseService,
	}
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
			FirstName: user.FirstName},
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

	newUser := model.User{
		Email:                   reqBody.Email,
		FirstName:               reqBody.FirstName,
		LastName:                reqBody.LastName,
		Age:                     reqBody.Age,
		Profession:              reqBody.Profession,
		SocialChallenges:        reqBody.SocialChallenges,
		StrugglingSocialSetting: reqBody.StrugglingSocialSetting,
		Gender:                  reqBody.Gender,
		XP:                      0,
	}
	userId, err := c.authRepository.CreateNewUser(newUser)
	if err != nil {
		return nil, err
	}

	if err = c.userCourseService.TailorUserCourse(userId, newUser); err != nil {
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
		},
		nil
}
