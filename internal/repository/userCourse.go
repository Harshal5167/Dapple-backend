package repository

import (
	"context"

	"firebase.google.com/go/v4/db"
	"github.com/Harshal5167/Dapple-backend/internal/dto/response"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type UserCourseRepository struct {
	firebaseDB *db.Client
}

func NewUserCourseRepository(db *db.Client) *UserCourseRepository {
	return &UserCourseRepository{
		firebaseDB: db,
	}
}

func (r *UserCourseRepository) AddUserCourse(userId string, levelsForUser *response.LevelsForUser) error {
	ctx := context.Background()

	ref := r.firebaseDB.NewRef("userCourses").Child(userId)
	if err := ref.Set(ctx, model.UserCourse{
		Levels: levelsForUser.SelectedLevelIds,
		UserProgress: model.UserProgress{
			CompletedLevels:   0,
			CompletedSections: 0,
		},
	}); err != nil {
		return err
	}
	return nil
}

func (c *UserCourseRepository) GetUserCourse(userId string) (*model.UserCourse, error) {
	ctx := context.Background()

	ref := c.firebaseDB.NewRef("userCourses").Child(userId)
	var userCourse model.UserCourse
	if err := ref.Get(ctx, &userCourse); err != nil {
		return nil, err
	}
	return &userCourse, nil
}

func (c *UserCourseRepository) UpdateUserProgress(userId string, levelInc bool) error {
	ctx := context.Background()

	ref := c.firebaseDB.NewRef("userCourses").Child(userId).Child("userProgress")

	var updates = make(map[string]interface{})
	if levelInc {
		updates["completedLevels"] = map[string]interface{}{".sv": map[string]interface{}{"increment": 1}}
		updates["completedSections"] = 0
	} else {
		updates["completedSections"] = map[string]interface{}{".sv": map[string]interface{}{"increment": 1}}
	}

	if len(updates) > 0 {
		if err := ref.Update(ctx, updates); err != nil {
			return err
		}
	}
	return nil
}

func (c *UserCourseRepository) GetUserProgress(userId string) (*model.UserProgress, error) {
	ctx := context.Background()

	ref := c.firebaseDB.NewRef("userCourses").Child(userId).Child("userProgress")
	var userProgress model.UserProgress
	if err := ref.Get(ctx, &userProgress); err != nil {
		return nil, err
	}
	return &userProgress, nil
}