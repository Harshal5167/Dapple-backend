package utils

import (
	"fmt"
	"strings"

	"github.com/Harshal5167/Dapple-backend/internal/model"
)

func BuildStringForEmotions(emotions []model.Emotion) string {
	var builder strings.Builder

	builder.WriteString("Emotions:\n")

	for i, emotion := range emotions {	
		builder.WriteString(fmt.Sprintf(
			"%d. Emotion: %s\n   Confidence: %.2f%%\n\n",
			i+1,
			emotion.Emotion,
			emotion.Confidence*100,
		))
	}
	return builder.String()
}
