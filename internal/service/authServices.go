package service

import (
	"errors"
	"fmt"

	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
	"github.com/Harshal5167/Dapple-backend/internal/utils"
)

type AuthService struct {
	authRepository    interfaces.AuthRepository
	userCourseService interfaces.UserCourseService
	userRepo          interfaces.UserRepository
}

func NewAuthService(authRepository interfaces.AuthRepository, userCourseService interfaces.UserCourseService, userRepo interfaces.UserRepository) *AuthService {
	return &AuthService{
		authRepository:    authRepository,
		userCourseService: userCourseService,
		userRepo:          userRepo,
	}
}

func (c *AuthService) Login(reqBody *request.LoginRequest) (*response.AuthResponse, error) {
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

	user, err := c.userRepo.GetUserDetailsFromEmail(reqBody.Email)
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
	return &response.AuthResponse{
			Token:     token,
			FirstName: user.FirstName,
			XP:        user.XP,
		},
		nil
}

func (c *AuthService) Register(reqBody *request.RegisterRequest) (*response.AuthResponse, error) {
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
	fmt.Println(newUser)
	userId, err := c.userRepo.CreateNewUser(newUser)
	if err != nil {
		return nil, err
	}

	fmt.Println(newUser)
	if err = c.userCourseService.TailorUserCourse(userId, newUser); err != nil {
		return nil, err
	}
	fmt.Println(newUser)

	token, err := utils.GenerateJWTToken(model.User{
		UserId: userId,
		Email:  reqBody.Email,
	})
	if err != nil {
		return nil, err
	}
	return &response.AuthResponse{
			Token:     token,
			FirstName: reqBody.FirstName,
			XP:        0,
		},
		nil
}
