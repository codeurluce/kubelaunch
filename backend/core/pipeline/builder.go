package pipeline

import (
	"github.com/codeurluce/kubelaunch/backend/core/analyzer"
	"github.com/codeurluce/kubelaunch/backend/models"
)

func BuildDeploymentSpec(files []string, appName string) models.DeploymentSpec {

	runtime := analyzer.DetectRuntime(files)

	return models.DeploymentSpec{
		Name:       appName,
		Runtime:    string(runtime),
		Port:       analyzer.DetectPort(runtime),
		Entrypoint: analyzer.DetectEntrypoint(runtime),
		Image:      "kubelaunch/" + appName,
		Namespace:  "default",
	}
}
