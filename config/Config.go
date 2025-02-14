package config

import (
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/db"
	"github.com/google/generative-ai-go/genai"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	FirebaseDB   *db.Client
	FirebaseAuth *auth.Client
	Redis        *redis.Client
	Genai        *genai.Client
}

func NewConfig() *Config {

	db, auth := InitializeFirebaseClient()
	redis := InitializeRedis()
	genai := ConfigureGenai()

	return &Config{
		FirebaseDB:   db,
		FirebaseAuth: auth,
		Redis:        redis,
		Genai:        genai,
	}
}
