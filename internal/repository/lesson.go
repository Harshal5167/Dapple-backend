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
