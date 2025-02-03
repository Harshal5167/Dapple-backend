package service

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto"
	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type SectionService struct {
	sectionRepo interfaces.SectionRepository
	levelRepo  interfaces.LevelRepository
}

func NewSectionService(sectionRepo interfaces.SectionRepository, levelRepo interfaces.LevelRepository) *SectionService {
	return &SectionService{
		sectionRepo: sectionRepo,
		levelRepo:  levelRepo,
	}
}

func (s *SectionService) AddSection(req *dto.AddSectionRequest) (*dto.AddSectionResponse, error) {
	sectionId, err := s.sectionRepo.AddSection(model.Section{
		Name:    req.Name,
		LevelId: req.LevelId,
		TotalXP: req.TotalXP,
	})
	if err != nil {
		return nil, err
	}
	
	err = s.levelRepo.AddSectionToLevel(req.LevelId, sectionId)
	if err != nil {
		return nil, err
	}

	return &dto.AddSectionResponse{
		SectionId: sectionId,
	}, nil
}