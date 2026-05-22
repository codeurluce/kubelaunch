package models

type DeployRequest struct {
	RepoURL   string            `json:"repoUrl" binding:"required"`
	AppName   string            `json:"appName" binding:"required"`
	Namespace string            `json:"namespace"`
	Replicas  int32             `json:"replicas"`
	Port      int32             `json:"port"`
	Stack     StackType         `json:"stack"`
	EnvVars   map[string]string `json:"envVars"`
}
