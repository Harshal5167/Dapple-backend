package repository

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"github.com/Harshal5167/Dapple-backend/internal/model"
)

type UserRepository struct {
	FirebaseApp *firebase.App
}

func NewUserRepository(firebase *firebase.App) *UserRepository {
	return &UserRepository{FirebaseApp: firebase}
}

func (c *UserRepository) CreateNewUser(user model.User) (string, error) {
	ctx := context.Background()

	client, err := c.FirebaseApp.Database(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get Realtime Database client: %v", err)
	}

	ref, err := client.NewRef("users").Push(ctx, user)
	userId := ref.Key

	if err != nil {
		fmt.Printf("Failed adding user: %v", err)
		return "", err
	}
	return userId, nil
}

func (c *UserRepository) GetUserDetailsFromEmail(email string) (*model.User, error) {
	ctx := context.Background()

	client, err := c.FirebaseApp.Database(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get Realtime Database client: %v", err)
	}

	ref := client.NewRef("users")
	var users map[string]model.User

	err = ref.OrderByChild("email").EqualTo(email).LimitToFirst(1).Get(ctx, &users)
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

	client, err := c.FirebaseApp.Database(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get Realtime Database client: %v", err)
	}

	var user *model.User
	err = client.NewRef("users").Child(userId).Get(ctx, &user)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	return user, nil
}
