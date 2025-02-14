package repository

import (
	"context"

	"firebase.google.com/go/v4/db"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type LessonRepository struct {
	firebaseDB *db.Client
}

func NewLessonRepository(db *db.Client) *LessonRepository {
	return &LessonRepository{
		firebaseDB: db,
	}
}

func (r *LessonRepository) AddLesson(lesson model.Lesson) (string, error) {
	ctx := context.Background()

	ref, err := r.firebaseDB.NewRef("lessons").Push(ctx, lesson)
	if err != nil {
		return "", err
	}

	return ref.Key, nil
}

func (r *LessonRepository) GetLessonById(lessonId string) (*model.Lesson, error) {
	ctx := context.Background()

	var lesson model.Lesson
	if err := r.firebaseDB.NewRef("lessons/"+lessonId).Get(ctx, &lesson); err != nil {
		return nil, err
	}

	return &lesson, nil
}
