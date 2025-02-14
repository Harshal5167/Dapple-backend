package repository

import (
	"context"

	"firebase.google.com/go/v4/db"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type QuestionRepository struct {
	firebaseDB *db.Client
}

func NewQuestionRepository(db *db.Client) *QuestionRepository {
	return &QuestionRepository{
		firebaseDB: db,
	}
}

func (c *QuestionRepository) AddQuestion(question model.Question) (string, error) {
	ctx := context.Background()

	ref, err := c.firebaseDB.NewRef("questions").Push(ctx, question)
	if err != nil {
		return "", err
	}

	return ref.Key, nil
}

func (c *QuestionRepository) GetQuestionById(questionId string) (*model.Question, error) {
	ctx := context.Background()

	var question model.Question
	if err := c.firebaseDB.NewRef("questions").Child(questionId).Get(ctx, &question); err != nil {
		return nil, err
	}

	return &question, nil
}

func (c *QuestionRepository) GetHint(questionId string) (string, error) {
	ctx := context.Background()

	var hint string
	if err := c.firebaseDB.NewRef("questions").Child(questionId).Child("hint").Get(ctx, &hint); err != nil {
		return "", err
	}
	return hint, nil
}
