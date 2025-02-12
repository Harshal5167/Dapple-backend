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
	GetSectionData(userId string, sectionId string) (*dto.SectionData, error)
	AddCompleteSection(section *model.SectionData, levelId string) error
	UpdateSectionProgress(userId string, lessonId string) error
}

type QuestionService interface {
	AddQuestion(req *dto.AddQuestionRequest) (*dto.AddQuestionResponse, error)
	EvaluateSubjectiveAnswer(userId string, req *dto.EvaluateSubjectiveAnswerReq) (*dto.EvaluateSubjectiveAnswerResponse, error)
	EvaluateObjectiveAnswer(userId string, req *dto.EvaluateObjectiveAnswerReq) (*dto.EvaluateObjectiveAnswerResponse, error)
}

type LessonService interface {
	AddLesson(req *dto.AddLessonRequest) (*dto.AddLessonResponse, error)
}

type UserCourseService interface {
	TailorUserCourse(userId string, user model.User) error
	GetUserCourse(userId string) (*dto.UserCourseResponse, error)
	UpdateUserProgress(userId string, sectionId string) error
}

type GeminiService interface {
	GenerateUserCourse(user model.User, levelDetails []map[string]string) (*dto.LevelsForUser, error)
	EvaluateUserAnswer(user *model.User, question *model.Question, userAnswer []string) (*model.UserAnswerEvalutaion, error)
}
