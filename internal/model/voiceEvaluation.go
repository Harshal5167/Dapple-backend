package model

type VoiceEvaluation struct {
	Emotions         []Emotion        `json:"emotions"`
	SpeechRate       SpeechRate       `json:"speechRate"`
	VolumeMean       VolumeMean       `json:"volumeMean"`
	SpectralCentroid SpectralCentroid `json:"spectralCentroid"`
	Tempo            Tempo            `json:"tempo"`
}

type Emotion struct {
	Type       string  `json:"type"`
	Confidence float64 `json:"confidence"`
}

type SpeechRate struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type VolumeMean struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type SpectralCentroid struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type Tempo struct {
	Min int `json:"min"`
	Max int `json:"max"`
}
