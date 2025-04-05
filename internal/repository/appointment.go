package repository

import (
	"context"

	"firebase.google.com/go/v4/db"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type AppointmentRepository struct {
	firebaseDB *db.Client
}

func NewAppointmentRepository(db *db.Client) *AppointmentRepository {
	return &AppointmentRepository{
		firebaseDB: db,
	}
}

func (r *AppointmentRepository) AddTimeSlot(timeSlot *model.TimeSlot) (string, error) {
	ctx := context.Background()

	ref, err := r.firebaseDB.NewRef("time-slots").Push(ctx, timeSlot)
	if err != nil {
		return "", err
	}

	return ref.Key, nil
}

func (r *AppointmentRepository) GetTimeSlotById(timeSlotId string) (*model.TimeSlot, error) {
	ctx := context.Background()

	var timeSlot model.TimeSlot
	if err := r.firebaseDB.NewRef("time-slots/"+timeSlotId).Get(ctx, &timeSlot); err != nil {
		return nil, err
	}

	return &timeSlot, nil
}

func (r *AppointmentRepository) GetTimeSlotsByExpertId(expertId string) ([]*model.TimeSlot, error) {
	ctx := context.Background()

	var allTimeSlots map[string]model.TimeSlot
	err := r.firebaseDB.NewRef("time-slots").
		OrderByChild("expertId").
		EqualTo(expertId).
		Get(ctx, &allTimeSlots)

	if err != nil {
		return nil, err
	}

	var result []*model.TimeSlot
	for id, ts := range allTimeSlots {
		tsCopy := ts   // create a new variable to avoid referencing loop var
		tsCopy.Id = id // if you need to include the Firebase key
		result = append(result, &tsCopy)
	}
	return result, nil
}

func (r *AppointmentRepository) CreateAppointment(appointment *model.Appointment) (string, error) {
	ctx := context.Background()

	ref, err := r.firebaseDB.NewRef("appointments").Push(ctx, appointment)
	if err != nil {
		return "", err
	}

	return ref.Key, nil
}

func (r *AppointmentRepository) UpdateTimeSlotAvailability(timeSlotId string, available bool) error {
	ctx := context.Background()

	if err := r.firebaseDB.NewRef("time-slots/"+timeSlotId).Update(ctx, map[string]interface{}{"available": available}); err != nil {
		return err
	}

	return nil
}

func (r *AppointmentRepository) GetAllAppointments(userId string) (map[string]model.Appointment, error) {
	ctx := context.Background()

	var appointments map[string]model.Appointment
	err := r.firebaseDB.NewRef("appointments").OrderByChild("userId").EqualTo(userId).Get(ctx, &appointments)
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

func (r *AppointmentRepository) GetAppointmentById(appointmentId string) (*model.Appointment, error) {
	ctx := context.Background()

	var appointment model.Appointment
	if err := r.firebaseDB.NewRef("appointments/"+appointmentId).Get(ctx, &appointment); err != nil {
		return nil, err
	}

	return &appointment, nil
}
