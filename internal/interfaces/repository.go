package interfaces

import (
	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type AuthRepository interface {
	VerifyFirebaseToken(token string) (bool, string, error)
}

type LevelRepository interface {
	AddLevel(level model.Level) (string, error)
	AddSectionToLevel(levelId string, sectionId string) error
	GetAllLevels() (map[string]model.Level, error)
	GetLevelById(levelId string) (*model.Level, error)
}

type SectionRepository interface {
	AddSection(section model.Section) (string, error)
	AddQuestionToSection(sectionId string, questionId string) error
	AddLessonToSection(sectionId string, lessonId string) error
	GetQuestionsAndLessons(sectionId string) ([]string, []string, error)
	GetNoOfItems(sectionId string, itemType string) (int, error)
	StoreSectionProgress(userId string, sectionId string) (*model.SectionProgress, error)
	UpdateSectionProgress(userId string, sectionId string, xp int) (int, int, error)
	GetNextSectionId(sectionId string) (string, error)
	GetSectionById(sectionId string) (*response.Section, error)
	DeleteSectionProgress(userId string, sectionId string) error
	GetTimeStamp(userId string, sectionId string) (int64, error)
}

type QuestionRepository interface {
	AddQuestion(question model.Question) (string, error)
	GetQuestionById(questionId string) (*model.Question, error)
	GetHint(questionId string) (string, error)
	GetEvaluationByQuestionId(questionId string) (string, error)
}

type LessonRepository interface {
	AddLesson(lesson model.Lesson) (string, error)
	GetLessonById(lessonId string) (*model.Lesson, error)
}

type UserCourseRepository interface {
	AddUserCourse(userId string, levelsForUser *response.LevelsForUser) error
	GetUserCourse(userId string) (*model.UserCourse, error)
	UpdateUserProgress(userId string, levelInc bool) error
	GetUserProgress(userId string) (*model.UserProgress, error)
}

type UserRepository interface {
	CreateNewUser(user model.User) (string, error)
	GetUserDetailsFromEmail(email string) (*model.User, error)
	GetUserById(userId string) (*model.User, error)
	UpdateUserXP(userId string, xp int) error
	GetXP(userId string) (int, error)
}

type EvaluationRepository interface {
	AddVoiceEvaluation(voiceEvaluation model.VoiceEvaluation) (string, error)
	GetVoiceEvaluationById(evaluationId string) (*model.VoiceEvaluation, error)
	GetVideoEvaluationById(evaluationId string) (*model.Emotion, error)
	AddVideoEvaluation(videoEvaluation model.Emotion) (string, error)
}

type TestRepository interface {
	StoreTestSession(sessionId string, sectionId string) error
	StoreQuestionResult(sessionId string, sectionId string, testEval *model.TestAnswerEval) error
	GetTestSession(sessionId string, sectionId string) (*model.TestSession, error)
	GetAllQuestionResults(sessionId string, sectionId string) ([]model.TestAnswerEval, error)
	ClearTestSession(sessionId string, sectionId string) error
}

type ExpertRepository interface {
	AddExpert(expert *model.Expert) (string, error)
	GetExpertById(expertId string) (*model.Expert, error)
	GetAllExperts() (map[string]*model.Expert, error)
	UpdateExpert(expertId string, schedule []model.Schedule) error
}

type AppointmentRepository interface {
	AddTimeSlot(timeSlot *model.TimeSlot) (string, error)
	GetTimeSlotById(timeSlotId string) (*model.TimeSlot, error)
	GetTimeSlotsByExpertId(expertId string) ([]*model.TimeSlot, error)
	CreateAppointment(appointment *model.Appointment) (string, error)
	UpdateTimeSlotAvailability(timeSlotId string, available bool) error
	GetAllAppointments(userId string) (map[string]model.Appointment, error)
	GetAppointmentById(appointmentId string) (*model.Appointment, error)
}
