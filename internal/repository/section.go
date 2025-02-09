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

func (c *SectionRepository) GetNoOfItems(sectionId string, itemType string) (int, error) {
	ctx := context.Background()

	client, err := c.firebaseApp.Database(ctx)
	if err != nil {
		return 0, err
	}
	ref := client.NewRef("sections").Child(sectionId).Child(itemType)
	var items []string
	err = ref.Get(ctx, &items)
	if err != nil {
		return 0, err
	}
	return len(items), nil
}

func (c *SectionRepository) AddQuestionToSection(sectionId string, questionId string) error {
	ctx := context.Background()

	client, err := c.firebaseApp.Database(ctx)
	if err != nil {
		return err
	}

	ref := client.NewRef("sections").Child(sectionId).Child("questions")
	var questions []string
	err = ref.Get(ctx, &questions)
	if err != nil {
		return err
	}

	questions = append(questions, questionId)
	err = ref.Set(ctx, questions)
	if err != nil {
		return err
	}
	return nil
}

func (c *SectionRepository) AddLessonToSection(sectionId string, lessonId string) error {
	ctx := context.Background()

	client, err := c.firebaseApp.Database(ctx)
	if err != nil {
		return err
	}

	ref := client.NewRef("sections").Child(sectionId).Child("lessons")
	var lessons []string
	err = ref.Get(ctx, &lessons)
	if err != nil {
		return err
	}

	lessons = append(lessons, lessonId)
	err = ref.Set(ctx, lessons)
	if err != nil {
		return err
	}
	return nil
}

func (c *SectionRepository) GetQuestionsAndLessons(sectionId string) ([]string, []string, error) {
	ctx := context.Background()

	client, err := c.firebaseApp.Database(ctx)
	if err != nil {
		return nil, nil, err
	}

	var questions []string
	ref := client.NewRef("sections").Child(sectionId)
	if err = ref.Child("questions").Get(ctx, &questions); err != nil {
		return nil, nil, err
	}

	var lessons []string
	if err = ref.Child("lessons").Get(ctx, &lessons); err != nil {
		return nil, nil, err
	}

	return questions, lessons, nil
}
