package pipeline

import (
	"github.com/codeurluce/kubelaunch/backend/core/analyzer"
	"github.com/codeurluce/kubelaunch/backend/models"
)

func BuildDeploymentSpec(
	files []string,
	req models.DeployRequest,
) models.DeploymentSpec {

	runtime := analyzer.DetectRuntime(files)

	port := analyzer.DetectPort(runtime)

	if req.Port != 0 {
		port = int(req.Port)
	}

	namespace := req.Namespace

	if namespace == "" {
		namespace = "default"
	}

	replicas := req.Replicas

	if replicas == 0 {
		replicas = 1
	}

	return models.DeploymentSpec{
		Name:        req.AppName,
		Runtime:     string(runtime),
		Port:        int32(port),
		Entrypoint:  analyzer.DetectEntrypoint(runtime),
		Image:       "kubelaunch/" + req.AppName,
		Tag:         "latest",
		Namespace:   namespace,
		Replicas:    replicas,
		ServiceType: "NodePort",
		EnvVars:     req.EnvVars,
	}
}
