package service

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type LevelService struct {
	levelRepo interfaces.LevelRepository
}

func NewLevelService(levelRepo interfaces.LevelRepository) *LevelService {
	return &LevelService{
		levelRepo: levelRepo,
	}
}

func (s *LevelService) AddLevel(req *dto.AddLevelRequest) (*dto.AddLevelResponse, error) {
	levelId, err := s.levelRepo.AddLevel(model.Level{
		Name:        req.Name,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
	})
	if err != nil {
		return nil, err
	}

	return &dto.AddLevelResponse{
		LevelId: levelId,
	}, nil
}
