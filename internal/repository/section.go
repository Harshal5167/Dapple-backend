package repository

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"github.com/Harshal5167/Dapple-backend/internal/model"
	"github.com/redis/go-redis/v9"
)

type SectionRepository struct {
	firebaseApp *firebase.App
	rdb         *redis.Client
}

func NewSectionRepository(firebase *firebase.App, rdb *redis.Client) *SectionRepository {
	return &SectionRepository{
		firebaseApp: firebase,
		rdb:         rdb,
	}
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

func (c *SectionRepository) StoreSectionProgress(userId string, sectionId string) (*model.SectionProgress, error) {
	ctx := context.Background()
	key := fmt.Sprintf("user:%s:section:%s", userId, sectionId)

	val, err := c.rdb.Exists(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if val != 0 {
		var sectionProgress *model.SectionProgress
		err = c.rdb.HGetAll(ctx, key).Scan(&sectionProgress)
		if err != nil {
			return nil, err
		}
		return sectionProgress, nil
	}

	pipe := c.rdb.TxPipeline()

	pipe.HSet(ctx, key, []string{
		"progress", "0",
		"xp", "0",
	}).Err()
	pipe.Expire(ctx, key, 86400)
	_, err = pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &model.SectionProgress{
		Progress: 0,
		XP:       0,
	}, nil
}

func (c *SectionRepository) UpdateSectionProgress(userId string, sectionId string, xp int) error {
	ctx := context.Background()
	key := fmt.Sprintf("user:%s:section:%s", userId, sectionId)

	err := c.rdb.HIncrBy(ctx, key, "progress", 1).Err()
	if err != nil {
		return err
	}

	err = c.rdb.HIncrBy(ctx, key, "xp", int64(xp)).Err()
	if err != nil {
		return err
	}
	return nil
}
