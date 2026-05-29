package models

type DeploymentSpec struct {
	Name        string
	Runtime     string
	Image       string
	Tag         string
	Port        int32
	Replicas    int32
	Entrypoint  string
	Namespace   string
	ServiceType string

	EnvVars map[string]string
}
