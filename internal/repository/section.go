package repository

import (
	"context"
	"fmt"
	"time"

	"firebase.google.com/go/v4/db"
	"github.com/Harshal5167/Dapple-backend/internal/model"
	"github.com/redis/go-redis/v9"
)

type SectionRepository struct {
	firebaseDB *db.Client
	rdb        *redis.Client
}

func NewSectionRepository(db *db.Client, rdb *redis.Client) *SectionRepository {
	return &SectionRepository{
		firebaseDB: db,
		rdb:        rdb,
	}
}

func (c *SectionRepository) AddSection(section model.Section) (string, error) {
	ctx := context.Background()

	var ref *db.Ref
	ref, err := c.firebaseDB.NewRef("sections").Push(ctx, section)
	if err != nil {
		return "", err
	}
	return ref.Key, nil
}

func (c *SectionRepository) GetNoOfItems(sectionId string, itemType string) (int, error) {
	ctx := context.Background()

	ref := c.firebaseDB.NewRef("sections").Child(sectionId).Child(itemType)
	var items []string
	err := ref.Get(ctx, &items)
	if err != nil {
		return 0, err
	}
	return len(items), nil
}

func (c *SectionRepository) AddQuestionToSection(sectionId string, questionId string) error {
	ctx := context.Background()

	ref := c.firebaseDB.NewRef("sections").Child(sectionId).Child("questions")
	err := ref.Transaction(ctx, func(node db.TransactionNode) (interface{}, error) {
		var questions []string
		if err := node.Unmarshal(&questions); err != nil {
			return nil, err
		}
		questions = append(questions, questionId)
		return questions, nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *SectionRepository) AddLessonToSection(sectionId string, lessonId string) error {
	ctx := context.Background()

	ref := c.firebaseDB.NewRef("sections").Child(sectionId).Child("lessons")

	err := ref.Transaction(ctx, func(node db.TransactionNode) (interface{}, error) {
		var lessons []string
		if err := node.Unmarshal(&lessons); err != nil {
			return nil, err
		}
		lessons = append(lessons, lessonId)
		return lessons, nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *SectionRepository) GetQuestionsAndLessons(sectionId string) ([]string, []string, error) {
	ctx := context.Background()

	var section model.Section
	if err := c.firebaseDB.NewRef("sections").Child(sectionId).Get(ctx, &section); err != nil {
		return nil, nil, err
	}

	return section.Questions, section.Lessons, nil
}

func (c *SectionRepository) StoreSectionProgress(userId string, sectionId string) (*model.SectionProgress, error) {
    ctx := context.Background()
    key := fmt.Sprintf("user:%s:section:%s", userId, sectionId)
    
    val, err := c.rdb.Exists(ctx, key).Result()
    if err != nil {
        return nil, err
    }
    if val != 0 {
        sectionProgress := &model.SectionProgress{}
        err = c.rdb.HGetAll(ctx, key).Scan(sectionProgress)
        if err != nil {
            return nil, err
        }
        return sectionProgress, nil
    }
    
    pipe := c.rdb.Pipeline()
    pipe.HSet(ctx, key, "progress", "0", "xp", "0")
    pipe.Expire(ctx, key, 86400*time.Second)
    _, err = pipe.Exec(ctx)
    if err != nil {
        return nil, err
    }
    return &model.SectionProgress{
        Progress: 0, 
        XP:       0, 
    }, nil
}

func (c *SectionRepository) UpdateSectionProgress(userId string, sectionId string, xp int) (int, int, error) {
	ctx := context.Background()
	key := fmt.Sprintf("user:%s:section:%s", userId, sectionId)

	pipe := c.rdb.TxPipeline()

	progressCmd := pipe.HIncrBy(ctx, key, "progress", 1)
	xpCmd := pipe.HIncrBy(ctx, key, "xp", int64(xp))

	_, err := pipe.Exec(ctx)
	if err != nil {
		return 0, 0, err
	}

	progress, err := progressCmd.Result()
	if err != nil {
		return 0, 0, err
	}

	xpGained, err := xpCmd.Result()
	if err != nil {
		return 0, 0, err
	}

	return int(progress), int(xpGained), nil
}

func (c *SectionRepository) GetNextSectionId(sectionId string) (string, error) {
	ctx := context.Background()
	var nextSectionId string
	err := c.firebaseDB.NewRef("sections").Child(sectionId).Child("nextSectionId").Get(ctx, &nextSectionId)
	if err != nil {
		return "", err
	}

	return nextSectionId, nil
}
