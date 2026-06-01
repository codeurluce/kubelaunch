package pipeline

import (
	"fmt"

	"github.com/codeurluce/kubelaunch/backend/core/analyzer"
	"github.com/codeurluce/kubelaunch/backend/models"
	"github.com/codeurluce/kubelaunch/backend/services"
)

func RunDetect(repoURL string) (models.DetectResponse, error) {

	// 1. fetch repo
	repoPath, err := services.ExtractRepoPath(repoURL)
	if err != nil {
		return models.DetectResponse{}, err
	}

	files, err := services.FetchRootFiles(repoPath)
	if err != nil {
		return models.DetectResponse{}, err
	}

	// 2. detect runtime
	runtime := analyzer.DetectRuntime(files)
	framework := analyzer.DetectFramework(files)

	fmt.Println("FRAMEWORK:", framework)
	fmt.Println("REPO:", repoPath)

	// 3. response
	return models.DetectResponse{
		Stack:         models.StackType(runtime),
		Port:          analyzer.DetectPort(framework, repoPath),
		Framework:     framework,
		DetectedFiles: files,
		Confidence:    "high",
	}, nil
}
