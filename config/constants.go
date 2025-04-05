package config

import (
	"os"

	"github.com/joho/godotenv"
)

var MaxFileSize = 10 * 1024 * 1024 // 10 MB
var AllowedFileExtensions = []string{"audio/x-wav"}

var MaxNoOfLessons int = 4
var MaxNoOfQuestions int = 4

var VoiceModelAPI string
var ImageModelAPI string

var ModelName = "gemini-2.0-flash"

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	VoiceModelAPI = os.Getenv("VOICE_MODEL_API")
	ImageModelAPI = os.Getenv("IMAGE_MODEL_API")
}
