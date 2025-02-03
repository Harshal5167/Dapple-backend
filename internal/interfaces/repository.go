package interfaces

import (
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type AuthRepository interface {
	CreateNewUser(user model.User) (string, error)
	GetUserDetailsFromEmail(email string) (model.User, error)
	VerifyFirebaseToken(token string) (bool, string, error)
}

type LevelRepository interface {
	AddLevel(level model.Level) (string, error)
	AddSectionToLevel(levelId string, sectionId string) error
}

type SectionRepository interface {
	AddSection(section model.Section) (string, error)
}
