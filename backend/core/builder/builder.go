package builder

import (
	"github.com/codeurluce/kubelaunch/backend/models"
)

// BuildSpec = convertit request + runtime en spec Kubernetes propre
func BuildSpec(req models.DeployRequest, runtime string, port int, entrypoint string) models.DeploymentSpec {

	namespace := req.Namespace
	if namespace == "" {
		namespace = "default"
	}

	replicas := req.Replicas
	if replicas == 0 {
		replicas = 1
	}

	finalPort := req.Port

	if runtime == "static" {
		finalPort = 80
	} else if finalPort == 0 {
		finalPort = int32(port)
	}

	return models.DeploymentSpec{
		Name:        req.AppName,
		Runtime:     string(runtime),
		Image:       "kubelaunch/" + req.AppName,
		Tag:         "latest",
		Port:        finalPort,
		Replicas:    replicas,
		Entrypoint:  entrypoint,
		Namespace:   namespace,
		ServiceType: "NodePort",
		EnvVars:     req.EnvVars,
	}
}
