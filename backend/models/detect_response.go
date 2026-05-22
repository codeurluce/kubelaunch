package models

type DetectResponse struct {
	Stack         StackType `json:"stack"`
	Port          int       `json:"port"`
	Confidence    string    `json:"confidence"`
	DetectedFiles []string  `json:"detected_files"`
	Framework     string    `json:"framework"`
}
