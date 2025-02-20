package model

type VoiceEvaluation struct {
	Emotions         []Emotion        `json:"emotions"`
	SpeechRate       SpeechRate       `json:"speechRate"`
	VolumeMean       VolumeMean       `json:"volumeMean"`
	SpectralCentroid SpectralCentroid `json:"spectralCentroid"`
	Tempo            Tempo            `json:"tempo"`
}

type Emotion struct {
	Emotion    string  `json:"emotion"`
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
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type Tempo struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}
