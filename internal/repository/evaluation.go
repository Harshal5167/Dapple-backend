package repository

import (
	"context"

	"firebase.google.com/go/v4/db"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type EvaluationRepository struct {
	firebaseDB *db.Client
}

func NewEvaluation(db *db.Client) *EvaluationRepository {
	return &EvaluationRepository{
		firebaseDB: db,
	}
}

func (c *EvaluationRepository) AddVoiceEvaluation(voiceEvaluation model.VoiceEvaluation) (string, error) {
	ref, err := c.firebaseDB.NewRef("evaluations").Child("voice").Push(context.Background(), voiceEvaluation)
	if err != nil {
		return "", err
	}

	return ref.Key, nil
}

func (c *EvaluationRepository) GetVoiceEvaluationById(evaluationId string) (*model.VoiceEvaluation, error) {
	var voiceEvaluation model.VoiceEvaluation
	err := c.firebaseDB.NewRef("evaluations").Child("voice").Child(evaluationId).Get(context.Background(), &voiceEvaluation)
	if err != nil {
		return nil, err
	}

	return &voiceEvaluation, nil
}
