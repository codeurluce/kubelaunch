package models

type DeployStep struct {
	ID      string `json:"id"`
	Label   string `json:"label"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

type DeployResponse struct {
	Success        bool         `json:"success"`
	Status         string       `json:"status"`
	AppName        string       `json:"appName"`
	Namespace      string       `json:"namespace"`
	DeploymentName string       `json:"deploymentName"`
	URL            string       `json:"url,omitempty"`
	Message        string       `json:"message,omitempty"`
	Error          string       `json:"error,omitempty"`
	Logs           []string     `json:"logs,omitempty"`
	Steps          []DeployStep `json:"steps,omitempty"`
}
