package repository

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type LessonRepository struct {
	firebaseApp *firebase.App
}

func NewLessonRepository(firebaseApp *firebase.App) *LessonRepository {
	return &LessonRepository{
		firebaseApp: firebaseApp,
	}
}

func (r *LessonRepository) AddLesson(lesson model.Lesson) (string, error) {
	ctx := context.Background()

	client, err := r.firebaseApp.Database(ctx)
	if err != nil {
		return "", err
	}

	ref, err := client.NewRef("lessons").Push(ctx, lesson)
	if err != nil {
		return "", err
	}

	return ref.Key, nil
}

func (r *LessonRepository) GetLessonById(lessonId string) (*map[string]interface{}, error) {
	ctx := context.Background()

	client, err := r.firebaseApp.Database(ctx)
	if err != nil {
		return nil, err
	}

	var lesson map[string]interface{}
	if err := client.NewRef("lessons/"+lessonId).Get(ctx, &lesson); err != nil {
		return nil, err
	}
	lesson["lessonId"] = lessonId

	return &lesson, nil
}
