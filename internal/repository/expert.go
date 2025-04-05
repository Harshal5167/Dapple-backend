package repository

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/db"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type ExpertRepository struct {
	firebaseDB *db.Client
}

func NewExpertRepository(db *db.Client) *ExpertRepository {
	return &ExpertRepository{
		firebaseDB: db,
	}
}

func (r *ExpertRepository) AddExpert(expert *model.Expert) (string, error) {
	ctx := context.Background()
	fmt.Println("Adding expert to Firebase DB:")

	ref, err := r.firebaseDB.NewRef("experts").Push(ctx, expert)
	if err != nil {
		return "", err
	}
	fmt.Println("Expert added with ID:", ref.Key)

	return ref.Key, nil
}

func (r *ExpertRepository) GetExpertById(expertId string) (*model.Expert, error) {
	ctx := context.Background()

	var expert model.Expert
	if err := r.firebaseDB.NewRef("experts/"+expertId).Get(ctx, &expert); err != nil {
		return nil, err
	}

	return &expert, nil
}

func (r *ExpertRepository) GetAllExperts() (map[string]*model.Expert, error) {
	ctx := context.Background()

	var experts map[string]*model.Expert
	if err := r.firebaseDB.NewRef("experts").Get(ctx, &experts); err != nil {
		return nil, err
	}

	return experts, nil
}

func (r *ExpertRepository) UpdateExpert(expertId string, schedule []model.Schedule) error {
	ctx := context.Background()

	if err := r.firebaseDB.NewRef("experts/"+expertId+"/schedule").Set(ctx, schedule); err != nil {
		return err
	}

	return nil
}
