package auth

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"github.com/Harshal5167/Dapple/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepository struct {
	FirebaseApp *firebase.App
}

func NewAuthRepository(firebase *firebase.App) *AuthRepository {
	return &AuthRepository{FirebaseApp: firebase}
}

func (c *AuthRepository) CheckExistingEmail(email string) (bool, error) {
	ctx := context.Background()

	client, err := c.FirebaseApp.Database(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to get Realtime Database client: %v", err)
	}

	ref := client.NewRef("users")
	var users map[string]interface{}

	err = ref.OrderByChild("email").EqualTo(email).Get(ctx, &users)
	if err != nil {
		return false, fmt.Errorf("error querying database: %v", err)
	}

	return len(users) > 0, nil
}

func (c *AuthRepository) CreateNewUser(email string, password string) (string, error) {
	ctx := context.Background()

	client, err := c.FirebaseApp.Database(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get Realtime Database client: %v", err)
	}

	ref, err := client.NewRef("users").Push(ctx, map[string]interface{}{
		"email":    email,
		"password": password,
	})

	userId := ref.Key

	if err != nil {
		fmt.Printf("Failed adding user: %v", err)
		return "", err
	}
	return userId, nil
}

func (c *AuthRepository) GenerateCustomToken(uid string, user model.User) (string, error) {
	ctx := context.Background()

	client, err := c.FirebaseApp.Auth(ctx)
	if err != nil {
		return "", err
	}

	token, err := client.CustomTokenWithClaims(ctx, uid, map[string]interface{}{
		"email":  user.Email,
		"userId": uid,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

func (c *AuthRepository) CheckPassword(email string, password string) error {
	ctx := context.Background()

	client, err := c.FirebaseApp.Database(ctx)
	if err != nil {
		return fmt.Errorf("failed to get Realtime Database client: %v", err)
	}

	ref := client.NewRef("users")
	var users map[string]map[string]interface{}

	err = ref.OrderByChild("email").EqualTo(email).Get(ctx, &users)
	if err != nil {
		return fmt.Errorf("error querying database: %v", err)
	}

	for user := range users {
		hashedPassword, ok := users[user]["password"].(string)
		if !ok {
			return fmt.Errorf("password field is not a string")
		}

		if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
			return err
		}
	}
	return nil
}

func (c *AuthRepository) GetUserIdFromEmail(email string) (string, error) {
	ctx := context.Background()

	client, err := c.FirebaseApp.Database(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get Realtime Database client: %v", err)
	}

	ref := client.NewRef("users")
	var users map[string]map[string]interface{}

	err = ref.OrderByChild("email").EqualTo(email).Get(ctx, &users)
	if err != nil {
		return "", fmt.Errorf("error querying database: %v", err)
	}
	for user := range users {
		return user, nil
	}
	return "", nil
}
