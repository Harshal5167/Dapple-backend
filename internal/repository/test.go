package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Harshal5167/Dapple-backend/internal/model"
	"github.com/redis/go-redis/v9"
)

type TestRepository struct {
	rdb *redis.Client
}

func NewTestRepository(rdb *redis.Client) *TestRepository {
	return &TestRepository{
		rdb: rdb,
	}
}

func (c *TestRepository) StoreTestSession(sessionId string, sectionId string) error {
	ctx := context.Background()
	key := fmt.Sprintf("testsession:%s:section:%s", sessionId, sectionId)

	pipe := c.rdb.Pipeline()
	pipe.HSet(ctx, key, "totalXP", "0", "timestamp", fmt.Sprintf("%d", time.Now().Unix()))
	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (c *TestRepository) StoreQuestionResult(sessionId string, sectionId string, testEval *model.TestAnswerEval) error {
	ctx := context.Background()
	key := fmt.Sprintf("testsession:%s:sectionId:%s:result", sessionId, sectionId)

	jsonData, err := json.Marshal(testEval)
	if err != nil {
		return err
	}

	err = c.rdb.SAdd(ctx, key, jsonData).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *TestRepository) GetTestSession(sessionId string, sectionId string) (*model.TestSession, error) {
	ctx := context.Background()
	key := fmt.Sprintf("testsession:%s:section:%s", sessionId, sectionId)

	var testSession = &model.TestSession{}
	err := c.rdb.HGetAll(ctx, key).Scan(testSession)
	if err != nil {
		return nil, err
	}
	return testSession, nil
}

func (c *TestRepository) GetAllQuestionResults(sessionId string, sectionId string) ([]model.TestAnswerEval, error) {
	ctx := context.Background()
	key := fmt.Sprintf("testsession:%s:section:%s:result", sessionId, sectionId)

	jsonDataSlice, err := c.rdb.SMembers(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve set members: %v", err)
	}

	results := make([]model.TestAnswerEval, 0, len(jsonDataSlice))
	for _, jsonData := range jsonDataSlice {
		var testEval model.TestAnswerEval
		if err := json.Unmarshal([]byte(jsonData), &testEval); err != nil {
			return nil, fmt.Errorf("failed to unmarshal set member: %v", err)
		}
		results = append(results, testEval)
	}

	return results, nil
}

func (c *TestRepository) ClearTestSession(sessionId string, sectionId string) error {
	ctx := context.Background()
	key := fmt.Sprintf("testsession:%s:section:%s", sessionId, sectionId)

	err := c.rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	key = fmt.Sprintf("testsession:%s:section:%s:result", sessionId, sectionId)
	err = c.rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}
