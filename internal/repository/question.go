package repository

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type QuestionRepository struct {
	firebaseApp *firebase.App
}

func NewQuestionRepository(firebase *firebase.App) *QuestionRepository {
	return &QuestionRepository{firebaseApp: firebase}
}

func (c *QuestionRepository) AddQuestion(question model.Question) (string, error) {
	ctx := context.Background()

	client, err := c.firebaseApp.Database(ctx)
	if err != nil {
		return "", err
	}

	ref, err := client.NewRef("questions").Push(ctx, question)
	if err != nil {
		return "", err
	}

	return ref.Key, nil
}

func (c *QuestionRepository) GetQuestionById(questionId string) (*model.Question, error) {
	ctx := context.Background()

	client, err := c.firebaseApp.Database(ctx)
	if err != nil {
		return nil, err
	}

	var question model.Question
	if err := client.NewRef("questions").Child(questionId).Get(ctx, &question); err != nil {
		return nil, err
	}

	return &question, nil
}
