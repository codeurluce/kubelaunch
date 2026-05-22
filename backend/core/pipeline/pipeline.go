package pipeline

import (
	"github.com/codeurluce/kubelaunch/backend/core/analyzer"
	"github.com/codeurluce/kubelaunch/backend/generator/dockerfile"
	k8sgenerator "github.com/codeurluce/kubelaunch/backend/generator/kubernetes"
	"github.com/codeurluce/kubelaunch/backend/models"
)

// Run = orchestrateur principal KubeLaunch
func Run(files []string, appName string) (models.DeployResponse, models.DeploymentSpec) {

	// 1. Detect runtime
	runtime := analyzer.DetectRuntime(files)
	port := analyzer.DetectPort(runtime)
	entry := analyzer.DetectEntrypoint(runtime)

	// 2. Build deployment spec
	spec := models.DeploymentSpec{
		Name:       appName,
		Runtime:    string(runtime),
		Image:      "kubelaunch/" + appName,
		Port:       port,
		Entrypoint: entry,
		Namespace:  "default",
	}

	// 3. Generate assets
	_ = dockerfile.GenerateDockerfile(runtime)
	_ = k8sgenerator.GenerateDeployment(spec)
	_ = k8sgenerator.GenerateService(spec)

	// 4. Response
	resp := models.DeployResponse{
		Status:         "success",
		AppName:        appName,
		Namespace:      "default",
		DeploymentName: appName,
		Message:        "Application generated successfully",
	}

	return resp, spec
}
