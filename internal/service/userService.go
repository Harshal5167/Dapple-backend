package service

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
)

type UserService struct{
	userRepo interfaces.UserRepository	
}

func NewUserService(userRepo interfaces.UserRepository) interfaces.UserService{
	return &UserService{userRepo}
}

func (s *UserService) GetXP(userId string) (*response.GetXP, error){
	xp, err := s.userRepo.GetXP(userId)
	if err != nil{
		return nil, err
	}

	return &response.GetXP{
		XP: xp,
	}, nil
}