package repository

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type LevelRepository struct {
	firebaseApp *firebase.App
}

func NewLevelRepository(firebase *firebase.App) *LevelRepository {
	return &LevelRepository{firebaseApp: firebase}
}

func (c *LevelRepository) AddLevel(level model.Level) (string, error) {
	ctx := context.Background()

	client, err := c.firebaseApp.Database(ctx)
	if err != nil {
		return "", err
	}

	var ref *db.Ref
	ref, err = client.NewRef("levels").Push(ctx, level)
	if err != nil {
		return "", err
	}

	return ref.Key, nil
}
