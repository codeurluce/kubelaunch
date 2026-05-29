package k8s

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/codeurluce/kubelaunch/backend/core/cluster"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ResolveServiceURL(
	clientset *kubernetes.Clientset,
	namespace string,
	serviceName string,
) (string, error) {

	provider := cluster.DetectProvider()

	switch provider {

	case cluster.Minikube:
		return resolveMinikubeURL(serviceName)

	case cluster.Kind:
		err := StartPortForward(
			serviceName,
			"9090",
			"80",
		)

		if err != nil {
			return "", err
		}

		return "http://localhost:9090", nil

	default:
		return "", fmt.Errorf("unsupported kubernetes provider")
	}
}

func resolveMinikubeURL(serviceName string) (string, error) {

	cmd := exec.Command(
		"minikube",
		"service",
		serviceName,
		"--url",
	)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	url := strings.TrimSpace(string(output))

	if url == "" {
		return "", fmt.Errorf("empty service url")
	}

	lines := strings.Split(url, "\n")

	return strings.TrimSpace(lines[0]), nil
}

func resolveNodePortURL(
	clientset *kubernetes.Clientset,
	namespace string,
	serviceName string,
) (string, error) {

	svc, err := clientset.
		CoreV1().
		Services(namespace).
		Get(
			context.Background(),
			serviceName,
			metav1.GetOptions{},
		)

	if err != nil {
		return "", err
	}

	if len(svc.Spec.Ports) == 0 {
		return "", fmt.Errorf("service has no ports")
	}

	nodePort := svc.Spec.Ports[0].NodePort

	return fmt.Sprintf("http://localhost:%d", nodePort), nil
	// for production
	// return fmt.Sprintf("http://127.0.0.1:%d", nodePort), nil
}
