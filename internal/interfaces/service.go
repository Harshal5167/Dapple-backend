package interfaces

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto"
)

type AuthService interface {
	Login(reqBody *dto.LoginRequest) (*dto.AuthResponse, error)
	Register(reqBody *dto.RegisterRequest) (*dto.AuthResponse, error)
}

type LevelService interface {
	AddLevel(req *dto.AddLevelRequest) (*dto.AddLevelResponse, error)
}
