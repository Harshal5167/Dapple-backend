package utils

import (
	"os"
	"time"

	"github.com/Harshal5167/Dapple-backend/internal/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func GenerateJWTToken(user model.User) (string, error) {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":  user.Email,
			"userId": user.UserId,
			"exp":    time.Now().AddDate(0, 1, 0).Unix(),
		},
	)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
