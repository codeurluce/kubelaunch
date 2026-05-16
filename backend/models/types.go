package models

// StackType représente le type de stack détecté
type StackType string

const (
	StackNodeJS    StackType = "nodejs"
	StackNextJS    StackType = "nextjs"
	StackPython    StackType = "python"
	StackGo        StackType = "go"
	StackDocker    StackType = "docker"
	StackUnknown   StackType = "unknown"
)

// DetectRequest reçu du frontend
type DetectRequest struct {
	RepoURL string `json:"repoUrl" binding:"required"`
}

// DetectResponse renvoyé au frontend
type DetectResponse struct {
	Stack         StackType `json:"stack"`
	Port          int       `json:"port"`
	Confidence    string    `json:"confidence"`
	DetectedFiles []string  `json:"detected_files"`
	Framework     string    `json:"framework"`
}

// DeployRequest reçu du frontend
type DeployRequest struct {
	RepoURL   string            `json:"repoUrl" binding:"required"`
	AppName   string            `json:"appName" binding:"required"`
	Namespace string            `json:"namespace"`
	Replicas  int32             `json:"replicas"`
	Port      int32             `json:"port"`
	Stack     StackType         `json:"stack"`
	EnvVars   map[string]string `json:"envVars"`
}

// DeployResponse renvoyé au frontend
type DeployResponse struct {
	Status         string `json:"status"`
	AppName        string `json:"appName"`
	Namespace      string `json:"namespace"`
	DeploymentName string `json:"deploymentName"`
	Message        string `json:"message"`
}

// AppStatus représente le statut d'une app déployée
type AppStatus struct {
	Name            string    `json:"name"`
	Namespace       string    `json:"namespace"`
	Stack           StackType `json:"stack"`
	Port            int32     `json:"port"`
	Replicas        int32     `json:"replicas"`
	ReadyReplicas   int32     `json:"readyReplicas"`
	Status          string    `json:"status"`
	CPUUsage        string    `json:"cpu"`
	MemoryUsage     string    `json:"memory"`
	CreatedAt       string    `json:"createdAt"`
}