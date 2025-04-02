package repository

import (
	"context"

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

	ref, err := r.firebaseDB.NewRef("experts").Push(ctx, expert)
	if err != nil {
		return "", err
	}

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

func (r *ExpertRepository) AddTimeSlot(timeSlot *model.TimeSlot) (string, error) {
	ctx := context.Background()

	ref, err := r.firebaseDB.NewRef("time-slots").Push(ctx, timeSlot)
	if err != nil {
		return "", err
	}

	return ref.Key, nil
}

func (r *ExpertRepository) UpdateExpert(expertId string, schedule []*model.Schedule) error {
	ctx := context.Background()

	if err := r.firebaseDB.NewRef("experts/"+expertId+"/schedule").Set(ctx, schedule); err != nil {
		return err
	}

	return nil
}

func (r *ExpertRepository) GetTimeSlotById(timeSlotId string) (*model.TimeSlot, error) {
	ctx := context.Background()

	var timeSlot model.TimeSlot
	if err := r.firebaseDB.NewRef("time-slots/"+timeSlotId).Get(ctx, &timeSlot); err != nil {
		return nil, err
	}

	return &timeSlot, nil
}
