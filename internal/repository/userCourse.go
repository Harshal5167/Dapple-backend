package repository

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"github.com/Harshal5167/Dapple-backend/internal/dto"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type UserCourseRepository struct {
	firebaseApp *firebase.App
}

func NewUserCourseRepository(firebaseApp *firebase.App) *UserCourseRepository {
	return &UserCourseRepository{
		firebaseApp: firebaseApp,
	}
}

func (r *UserCourseRepository) AddUserCourse(userId string, levelsForUser *dto.LevelsForUser) error {
	ctx := context.Background()

	client, err := r.firebaseApp.Database(ctx)
	if err != nil {
		return err
	}

	ref := client.NewRef("userCourses").Child(userId)
	if err = ref.Set(ctx, model.UserCourse{
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

	client, err := c.firebaseApp.Database(ctx)
	if err != nil {
		return nil, err
	}

	ref := client.NewRef("userCourses").Child(userId)
	var userCourse model.UserCourse
	if err := ref.Get(ctx, &userCourse); err != nil {
		return nil, err
	}
	return &userCourse, nil
}

func (c *UserCourseRepository) UpdateUserProgress(userId string, levelInc bool) error {
	ctx := context.Background()
	client, err := c.firebaseApp.Database(ctx)
	if err != nil {
		return err
	}

	ref := client.NewRef("userCourses").Child(userId).Child("userProgress")

	var updates = make(map[string]interface{})
	if levelInc {
		updates["completedLevels"] = map[string]interface{}{".sv": "increment", "value": 1}
		updates["completedSections"] = 0
	} else {
		updates["completedSections"] = map[string]interface{}{".sv": "increment", "value": 1}
	}

	if len(updates) > 0 {
		if err := ref.Update(ctx, updates); err != nil {
			return err
		}
	}
	return nil
}