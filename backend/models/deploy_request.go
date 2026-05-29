package models

type DeployRequest struct {
	RepoURL   string            `json:"repoUrl" binding:"required"`
	Branch    string            `json:"branch"`
	AppName   string            `json:"appName" binding:"required"`
	Namespace string            `json:"namespace"`
	Replicas  int32             `json:"replicas"`
	Port      int32             `json:"port"`
	Stack     StackType         `json:"stack"`
	EnvVars   map[string]string `json:"envVars"`

	// 🔥 IMPORTANT POUR PIPELINE BUILD
	BuildContext string `json:"buildContext"` // ex: "/backend"
	BuildCommand string `json:"buildCommand"` // ex: "npm run build"
	StartCommand string `json:"startCommand"` // ex: "npm start"

	// 🔥 DOCKER CONTROL
	DockerfilePath string `json:"dockerfilePath"` // option custom
	AutoDetect     bool   `json:"autoDetect"`     // true = analyzer
}
