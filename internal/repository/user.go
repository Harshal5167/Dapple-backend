package repository

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/db"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type UserRepository struct {
	firebaseDB *db.Client
}

func NewUserRepository(db *db.Client) *UserRepository {
	return &UserRepository{
		firebaseDB: db,
	}
}

func (c *UserRepository) CreateNewUser(user model.User) (string, error) {
	ctx := context.Background()

	ref, err := c.firebaseDB.NewRef("users").Push(ctx, user)
	userId := ref.Key

	if err != nil {
		fmt.Printf("Failed adding user: %v", err)
		return "", err
	}
	return userId, nil
}

func (c *UserRepository) GetUserDetailsFromEmail(email string) (*model.User, error) {
	ctx := context.Background()

	ref := c.firebaseDB.NewRef("users")
	var users map[string]model.User

	err := ref.OrderByChild("email").EqualTo(email).LimitToFirst(1).Get(ctx, &users)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}

	for user := range users {
		userDetails := users[user]
		userDetails.UserId = user
		return &userDetails, nil
	}
	return nil, fmt.Errorf("unknown error occurred")
}

func (c *UserRepository) GetUserById(userId string) (*model.User, error) {
	ctx := context.Background()

	var user *model.User
	err := c.firebaseDB.NewRef("users").Child(userId).Get(ctx, &user)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	return user, nil
}

func (c *UserRepository) UpdateUserXP(userId string, xp int) error {
	ctx := context.Background()

	var updates = make(map[string]interface{})
	updates["xp"] = map[string]interface{}{".sv": map[string]interface{}{"increment": xp}}

	if len(updates) > 0 {
		if err := c.firebaseDB.NewRef("users").Child(userId).Update(ctx, updates); err != nil {
			return err
		}
	}
	return nil
}

func (c *UserRepository) GetXP(userId string) (int, error) {
	var xp int
	err := c.firebaseDB.NewRef("users").Child(userId).Child("xp").Get(context.Background(), &xp)
	if err != nil {
		return 0, fmt.Errorf("error querying database: %v", err)
	}
	return xp, nil
}
