package repository

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type SectionRepository struct {
	firebaseApp *firebase.App
}

func NewSectionRepository(firebase *firebase.App) *SectionRepository {
	return &SectionRepository{firebaseApp: firebase}
}

func (c *SectionRepository) AddSection(section model.Section) (string, error) {
	ctx := context.Background()

	client, err := c.firebaseApp.Database(ctx)
	if err != nil {
		return "", err
	}

	var ref *db.Ref
	ref, err = client.NewRef("sections").Push(ctx, section)
	if err != nil {
		return "", err
	}
	return ref.Key, nil
}

func (c *SectionRepository) LinkSectionsById(sectionIds []string) error {
	ctx := context.Background()

	client, err := c.firebaseApp.Database(ctx)
	if err != nil {
		return err
	}

	ref := client.NewRef("sections")
	for _, sectionId := range sectionIds {
		err = ref.Child(sectionId).Set(ctx, true)
		if err != nil {
			return err
		}
	}
	return nil
}