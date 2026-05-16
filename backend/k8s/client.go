package k8s

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// NewClient crée un client Kubernetes
// Il essaie d'abord la config in-cluster (si on tourne dans K8s)
// Sinon il utilise le kubeconfig local (~/.kube/config)
func NewClient() (*kubernetes.Clientset, error) {
	// Essayer in-cluster d'abord
	config, err := rest.InClusterConfig()
	if err != nil {
		// Fallback : kubeconfig local
		kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")

		// Sur Windows, HOME peut être vide — utiliser USERPROFILE
		if os.Getenv("HOME") == "" {
			kubeconfig = filepath.Join(os.Getenv("USERPROFILE"), ".kube", "config")
		}

		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, fmt.Errorf("impossible de charger le kubeconfig: %w", err)
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("impossible de créer le client K8s: %w", err)
	}

	return clientset, nil
}
