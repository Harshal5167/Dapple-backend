package response

import "github.com/Harshal5167/Dapple-backend/internal/model"

type VoiceEvaluation struct {
	Status              string             `json:"status"`
	EmotionDistribution map[string]float64 `json:"emotion_distribution"`
	AudioFeatures       AudioFeatures      `json:"audio_features"`
	Top3Emotions        []model.Emotion    `json:"top3_emotions"`
}

type AudioFeatures struct {
	VolumeMean       float64 `json:"volume_mean"`
	SpeechRate       float64 `json:"speech_rate"`
	SpectralCentroid float64 `json:"spectral_centroid"`
	Tempo            float64 `json:"tempo"`
}
