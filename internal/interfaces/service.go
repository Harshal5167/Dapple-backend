package interfaces

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type AuthService interface {
	Login(reqBody *request.LoginRequest) (*response.AuthResponse, error)
	Register(reqBody *request.RegisterRequest) (*response.AuthResponse, error)
}

type LevelService interface {
	AddLevel(req *request.AddLevelRequest) (*response.AddLevelResponse, error)
	AddCompleteLevel(req *request.AddCompleteLevelRequest) (*response.AddLevelResponse, error)
}

type SectionService interface {
	AddSection(req *request.AddSectionRequest) (*response.AddSectionResponse, error)
	GetSectionData(userId string, sectionId string) (*response.SectionData, error)
	AddCompleteSection(section *model.SectionData, levelId string) error
	UpdateSectionProgress(userId string, lessonId string) error
}

type QuestionService interface {
	AddQuestion(req *request.AddQuestionRequest) (*response.AddQuestionResponse, error)
	EvaluateSubjectiveAnswer(userId string, req *request.EvaluateSubjectiveAnswerReq) (*response.EvaluateSubjectiveAnswerResponse, error)
	EvaluateObjectiveAnswer(userId string, req *request.EvaluateObjectiveAnswerReq) (*response.EvaluateObjectiveAnswerResponse, error)
	GetHint(questionId string) (*response.GetHintResponse, error)
}

type LessonService interface {
	AddLesson(req *request.AddLessonRequest) (*response.AddLessonResponse, error)
}

type UserCourseService interface {
	TailorUserCourse(userId string, user model.User) error
	GetUserCourse(userId string) (*response.UserCourseResponse, error)
	UpdateUserProgress(userId string, sectionId string, xp int) error
}

type GeminiService interface {
	GenerateUserCourse(user model.User, levelDetails []map[string]string) (*response.LevelsForUser, error)
	EvaluateUserAnswer(user *model.User, question *model.Question, userAnswer []string) (*model.UserAnswerEvalutaion, error)
}
