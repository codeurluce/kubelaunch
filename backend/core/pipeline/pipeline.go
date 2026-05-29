package pipeline

import (
	"context"
	"fmt"
	"time"

	"github.com/codeurluce/kubelaunch/backend/core/analyzer"
	"github.com/codeurluce/kubelaunch/backend/core/builder"
	cluster "github.com/codeurluce/kubelaunch/backend/core/cluster"
	"github.com/codeurluce/kubelaunch/backend/generator/dockerfile"
	k8sgen "github.com/codeurluce/kubelaunch/backend/generator/kubernetes"
	kube "github.com/codeurluce/kubelaunch/backend/k8s"
	"github.com/codeurluce/kubelaunch/backend/models"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Run = FULL KubeLaunch pipeline
func Run(clientset *kubernetes.Clientset, req models.DeployRequest) (models.DeployResponse, error) {

	// =========================
	// 1. DETECT PROVIDER
	// =========================
	provider := cluster.DetectProvider()
	if provider == cluster.Unknown {
		return models.DeployResponse{}, fmt.Errorf("no supported Kubernetes provider found")
	}

	// =========================
	// 2. CLONE REPO
	// =========================
	repoPath, err := cloneRepo(req.RepoURL, req.Branch)
	if err != nil {
		return models.DeployResponse{}, fmt.Errorf("clone failed: %w", err)
	}

	// =========================
	// 3. ANALYZE PROJECT
	// =========================
	files := scanFiles(repoPath)
	// fmt.Println("FILES FOUND:")
	// for _, f := range files {
	// 	fmt.Println("-", f)
	// }

	runtime := analyzer.DetectRuntime(files)
	port := analyzer.DetectPort(runtime)
	entry := analyzer.DetectEntrypoint(runtime)

	fmt.Println("=================================")
	fmt.Println("RUNTIME:", runtime)
	fmt.Println("PORT:", port)
	fmt.Println("ENTRY:", entry)
	fmt.Println("=================================")

	// =========================
	// 4. KIND SETUP
	// =========================
	if provider == cluster.Kind {
		if err := cluster.EnsureKindCluster(); err != nil {
			return models.DeployResponse{}, fmt.Errorf("kind init failed: %w", err)
		}
	}

	// =========================
	// 5. BUILD SPEC
	// =========================
	spec := builder.BuildSpec(req, runtime, port, entry)

	// =========================
	// 6. GENERATE DOCKERFILE
	// =========================
	dockerfileContent := dockerfile.GenerateDockerfile(runtime)

	fmt.Println("=================================")
	fmt.Println("DOCKERFILE GENERATED:")
	fmt.Println(dockerfileContent)
	fmt.Println("=================================")

	if err := writeFile(repoPath+"/Dockerfile", dockerfileContent); err != nil {
		return models.DeployResponse{}, fmt.Errorf("dockerfile write failed: %w", err)
	}

	// =========================
	// 7. BUILD IMAGE
	// =========================
	imageTag := builder.BuildImageTag(req.AppName)

	if err := builder.BuildImage(repoPath, imageTag); err != nil {
		return models.DeployResponse{}, fmt.Errorf("docker build failed: %w", err)
	}

	// =========================
	// 8. LOAD IMAGE INTO CLUSTER
	// =========================
	switch provider {

	case cluster.Kind:
		if err := builder.LoadImageToKind(imageTag, cluster.DefaultKindCluster); err != nil {
			return models.DeployResponse{}, fmt.Errorf("kind load failed: %w", err)
		}

	case cluster.Minikube:
		if err := builder.LoadImageToMinikube(imageTag); err != nil {
			return models.DeployResponse{}, fmt.Errorf("minikube load failed: %w", err)
		}

	case cluster.DockerDesktop:
		// shared docker daemon

	case cluster.Cloud:
		// future registry push
	}

	// =========================
	// 9. GENERATE K8S YAML
	// =========================
	deploymentYAML := k8sgen.GenerateDeployment(spec)
	serviceYAML := k8sgen.GenerateService(spec)

	// =========================
	// 10. APPLY TO CLUSTER
	// =========================
	if err := kubectlApply(deploymentYAML); err != nil {
		return models.DeployResponse{}, fmt.Errorf("deployment apply failed: %w", err)
	}

	if err := kubectlApply(serviceYAML); err != nil {
		return models.DeployResponse{}, fmt.Errorf("service apply failed: %w", err)
	}

	// =========================
	// 11. WAIT FOR READINESS (CRITICAL FIX)
	// =========================
	if err := WaitForDeploymentReady(
		clientset,
		req.Namespace,
		req.AppName,
	); err != nil {
		return models.DeployResponse{}, fmt.Errorf("deployment not ready: %w", err)
	}

	// =========================
	// 12. RESOLVE URL
	// =========================
	url, err := kube.ResolveServiceURL(
		clientset,
		req.Namespace,
		req.AppName+"-service",
	)

	if err != nil {
		return models.DeployResponse{}, fmt.Errorf("url resolution failed: %w", err)
	}

	// =========================
	// SUCCESS RESPONSE
	// =========================
	return models.DeployResponse{
		Status:         "deployed",
		AppName:        req.AppName,
		Namespace:      req.Namespace,
		DeploymentName: req.AppName,
		URL:            url,
		Message:        "Deployment successful",
	}, nil
}

// =====================================================
// WAIT FOR DEPLOYMENT READY (REAL CHECK)
// =====================================================
func WaitForDeploymentReady(
	clientset *kubernetes.Clientset,
	namespace string,
	name string,
) error {

	timeout := time.After(120 * time.Second)
	tick := time.Tick(2 * time.Second)

	for {
		select {

		case <-timeout:
			return fmt.Errorf("timeout waiting for deployment ready")

		case <-tick:

			deploy, err := clientset.
				AppsV1().
				Deployments(namespace).
				Get(context.Background(), name, metav1.GetOptions{})

			if err != nil {
				return fmt.Errorf("failed to get deployment: %w", err)
			}

			// =========================
			// SUCCESS CONDITION
			// =========================
			if deploy.Status.ReadyReplicas >= 1 &&
				deploy.Status.UnavailableReplicas == 0 {
				return nil
			}

			// =========================
			// FAILURE DETECTION
			// =========================
			for _, c := range deploy.Status.Conditions {
				if c.Type == appsv1.DeploymentReplicaFailure {
					return fmt.Errorf("deployment replica failure: %s", c.Message)
				}
			}

			// CrashLoopBackOff detection (heuristic)
			if deploy.Status.Replicas > 0 &&
				deploy.Status.ReadyReplicas == 0 &&
				deploy.Status.UnavailableReplicas > 0 {
				// continue waiting, not instantly fail (avoid false positives)
			}
		}
	}
}
