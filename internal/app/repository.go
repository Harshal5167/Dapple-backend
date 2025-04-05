package app

import (
	"github.com/Harshal5167/Dapple-backend/internal/repository"
)

func (a *App) InitializeRepositories() {
	a.Repositories.AuthRepo = repository.NewAuthRepository(a.config.FirebaseAuth)
	a.Repositories.LevelRepo = repository.NewLevelRepository(a.config.FirebaseDB)
	a.Repositories.SectionRepo = repository.NewSectionRepository(a.config.FirebaseDB, a.config.Redis)
	a.Repositories.QuestionRepo = repository.NewQuestionRepository(a.config.FirebaseDB)
	a.Repositories.LessonRepo = repository.NewLessonRepository(a.config.FirebaseDB)
	a.Repositories.UserCourseRepo = repository.NewUserCourseRepository(a.config.FirebaseDB)
	a.Repositories.UserRepo = repository.NewUserRepository(a.config.FirebaseDB)
	a.Repositories.EvaluationRepo = repository.NewEvaluation(a.config.FirebaseDB)
	a.Repositories.TestRepo = repository.NewTestRepository(a.config.Redis)
	a.Repositories.ExpertRepo = repository.NewExpertRepository(a.config.FirebaseDB)
	a.Repositories.AppointmentRepo = repository.NewAppointmentRepository(a.config.FirebaseDB)
}
