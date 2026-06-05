package k8s

import (
	"context"
	"fmt"
	"net"
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

	case cluster.Kind:
		return resolvePortForwardURL(
			clientset,
			namespace,
			serviceName,
		)

	case cluster.Minikube:
		return resolveMinikubeURL(serviceName)

	case cluster.DockerDesktop:
		return resolveNodePortURL(
			clientset,
			namespace,
			serviceName,
		)

	case cluster.Cloud:
		return resolveNodePortURL(
			clientset,
			namespace,
			serviceName,
		)

	default:
		return "", fmt.Errorf(
			"unsupported kubernetes provider",
		)
	}
}

func resolvePortForwardURL(
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

	targetPort := svc.Spec.Ports[0].Port

	localPort, err := findFreePort()
	if err != nil {
		return "", err
	}

	err = StartPortForward(
		serviceName,
		fmt.Sprintf("%d", localPort),
		fmt.Sprintf("%d", targetPort),
	)

	if err != nil {
		return "", err
	}

	fmt.Println("=================================")
	fmt.Println("PORT FORWARD STARTED")
	fmt.Println("SERVICE:", serviceName)
	fmt.Println("LOCAL:", localPort)
	fmt.Println("TARGET:", targetPort)
	fmt.Println("=================================")

	return fmt.Sprintf(
		"http://localhost:%d",
		localPort,
	), nil
}

func resolveMinikubeURL(
	serviceName string,
) (string, error) {

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

	url := strings.TrimSpace(
		string(output),
	)

	if url == "" {
		return "", fmt.Errorf(
			"empty service url",
		)
	}

	lines := strings.Split(
		url,
		"\n",
	)

	return strings.TrimSpace(
		lines[0],
	), nil
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
		return "", fmt.Errorf(
			"service has no ports",
		)
	}

	nodePort := svc.Spec.Ports[0].NodePort

	return fmt.Sprintf(
		"http://localhost:%d",
		nodePort,
	), nil
}

func findFreePort() (int, error) {

	listener, err := net.Listen(
		"tcp",
		":0",
	)

	if err != nil {
		return 0, err
	}

	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)

	return addr.Port, nil
}
