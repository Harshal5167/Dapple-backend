package interfaces

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type AuthService interface {
	Login(reqBody *dto.LoginRequest) (*dto.AuthResponse, error)
	Register(reqBody *dto.RegisterRequest) (*dto.AuthResponse, error)
}

type LevelService interface {
	AddLevel(req *dto.AddLevelRequest) (*dto.AddLevelResponse, error)
	AddCompleteLevel(req *dto.AddCompleteLevelRequest) (*dto.AddLevelResponse, error)
}

type SectionService interface {
	AddSection(req *dto.AddSectionRequest) (*dto.AddSectionResponse, error)
	GetSectionData(sectionId string) (*dto.SectionData, error)
	AddCompleteSection(section *model.SectionData, levelId string) error
}

type QuestionService interface {
	AddQuestion(req *dto.AddQuestionRequest) (*dto.AddQuestionResponse, error)
}

type LessonService interface {
	AddLesson(req *dto.AddLessonRequest) (*dto.AddLessonResponse, error)
}

type UserCourseService interface {
	TailorUserCourse(userId string, user model.User) error
	GetUserCourse(userId string) (*dto.UserCourseResponse, error)
}
