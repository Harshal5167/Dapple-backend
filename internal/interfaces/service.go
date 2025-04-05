package interfaces

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/Harshal5167/Dapple-backend/internal/model"
	"github.com/gofiber/contrib/socketio"
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
	AddCompleteSection(section *request.SectionData, levelId string) error
	UpdateSectionProgress(userId string, lessonId string) error
	GetTestData(sectionId string) (*response.SectionData, error)
}

type QuestionService interface {
	AddQuestion(req *request.AddQuestionRequest) (*response.AddQuestionResponse, error)
	GetHint(questionId string) (*response.GetHintResponse, error)
}

type LessonService interface {
	AddLesson(req *request.AddLessonRequest) (*response.AddLessonResponse, error)
}

type UserCourseService interface {
	TailorUserCourse(userId string, user model.User) error
	GetUserCourse(userId string) (*response.UserCourseResponse, error)
	UpdateUserProgress(userId string, sectionId string, xp int) error
	GetUserProgress(userId string) (*model.UserProgress, error)
}

type GeminiService interface {
	GenerateUserCourse(user model.User, levelDetails []map[string]string) (*response.LevelsForUser, error)
	EvaluateUserAnswer(user *model.User, question *model.Question, userAnswer []string) (*model.UserAnswerEvalutaion, error)
	FormatVoiceEvaluationResponse(obtainedVoiceEvaluation *response.VoiceEvaluation, desiredVoiceEvaluation *model.VoiceEvaluation) (*model.UserAnswerEvalutaion, error)
	EvaluateTestAnswer(Answer string, question *model.Question, obtainedVideoEvaluation *response.QuestionResult, desiredVideoEvaluation *model.Emotion) (*model.TestAnswerEval, error)
}

type UserService interface {
	GetXP(userId string) (*response.GetXP, error)
}

type EvaluationService interface {
	EvaluateSubjectiveAnswer(userId string, req *request.EvaluateSubjectiveAnswerReq) (*response.EvaluateSubjectiveAnswerResponse, error)
	EvaluateObjectiveAnswer(userId string, req *request.EvaluateObjectiveAnswerReq) (*response.EvaluateObjectiveAnswerResponse, error)
	EvaluateVoiceAnswer(userId string, req *request.EvaluateVoiceAnswerReq, buf []byte) (*response.EvaluateVoiceAnswerResponse, error)
}

type SocketService interface {
	HandleConnection(kws *socketio.Websocket)
}

type TestService interface {
	EvaluateTestAnswer(message *request.TestData) (bool, error)
	EvaluateImageAnswer(message *request.TestData) error
	GetTestResult(userId string, sessionId string, sectionId string) (*response.TestResultResponse, error)
	RetryQuestion(sessionId string, questionId string) error
}

type ExpertService interface {
	AddExpert(req *request.AddExpertRequest) (*response.AddExpertResponse, error)
	GetExpertById(expertId string) (*response.GetExpertResponse, error)
	GetAllExperts() ([]*response.GetExpertResponse, error)
	GetExpertSchedule(expertId string) (*response.GetExpertScheduleResponse, error)
}

type AppointmentService interface {
	CreateAppointment(timeSlotId string, userId string) (*response.CreateAppointmentResponse, error)
	GetAllAppointments(userId string) ([]response.GetAllAppointmentsResponse, error)
	GetAppointmentById(appointmentId string) (*response.GetAppointmentByIdResponse, error)
}
