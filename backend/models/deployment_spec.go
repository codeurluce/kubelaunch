package models

type DeploymentSpec struct {
	Name       string
	Runtime    string
	Image      string
	Port       int
	Entrypoint string
	Namespace  string
}
