package app

import (
	"github.com/Harshal5167/Dapple-backend/internal/repository"
)

func (a *App) InitializeRepositories() {
	a.Repositories.AuthRepo = repository.NewAuthRepository(a.config.FirebaseApp)
	a.Repositories.LevelRepo = repository.NewLevelRepository(a.config.FirebaseApp)
	a.Repositories.SectionRepo = repository.NewSectionRepository(a.config.FirebaseApp, a.rdb)
	a.Repositories.QuestionRepo = repository.NewQuestionRepository(a.config.FirebaseApp)
	a.Repositories.LessonRepo = repository.NewLessonRepository(a.config.FirebaseApp)
	a.Repositories.UserCourseRepo = repository.NewUserCourseRepository(a.config.FirebaseApp)
	a.Repositories.UserRepo = repository.NewUserRepository(a.config.FirebaseApp)
}
