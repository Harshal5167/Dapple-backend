package repository

import (
	"context"

	"firebase.google.com/go/v4/db"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type LevelRepository struct {
	firebaseDB *db.Client
}

func NewLevelRepository(db *db.Client) *LevelRepository {
	return &LevelRepository{
		firebaseDB: db,
	}
}

func (c *LevelRepository) AddLevel(level model.Level) (string, error) {
	ctx := context.Background()

	var ref *db.Ref
	ref, err := c.firebaseDB.NewRef("levels").Push(ctx, level)
	if err != nil {
		return "", err
	}

	return ref.Key, nil
}

func (c *LevelRepository) AddSectionToLevel(levelId string, sectionId string) error {
	ctx := context.Background()

	ref := c.firebaseDB.NewRef("levels").Child(levelId).Child("sections")
	var sections []string
	err := ref.Get(ctx, &sections)
	if err != nil {
		return err
	}

	var lastSectionId string
	if len(sections) > 0 {
		lastSectionId = sections[len(sections)-1]
	} else {
		lastSectionId = ""
	}

	if lastSectionId != "" {
		err := c.firebaseDB.NewRef("sections").Child(lastSectionId).Child("nextSectionId").Set(ctx, sectionId)
		if err != nil {
			return err
		}
	}

	sections = append(sections, sectionId)
	err = ref.Set(ctx, sections)
	if err != nil {
		return err
	}
	return nil
}

func (c *LevelRepository) GetAllLevels() (map[string]model.Level, error) {
	ctx := context.Background()

	var levels map[string]model.Level
	err := c.firebaseDB.NewRef("levels").Get(ctx, &levels)
	if err != nil {
		return nil, err
	}

	return levels, nil
}

func (c *LevelRepository) GetLevelById(levelId string) (*model.Level, error) {
	ctx := context.Background()

	var level *model.Level
	err := c.firebaseDB.NewRef("levels").Child(levelId).Get(ctx, &level)
	if err != nil {
		return nil, err
	}

	return level, nil
}
