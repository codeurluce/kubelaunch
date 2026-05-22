package models

type DeployResponse struct {
	Status         string `json:"status"`
	AppName        string `json:"appName"`
	Namespace      string `json:"namespace"`
	DeploymentName string `json:"deploymentName"`
	Message        string `json:"message"`
}
