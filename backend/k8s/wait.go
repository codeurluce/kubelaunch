package k8s

import (
	"context"
	"fmt"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

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
			return fmt.Errorf("deployment timeout: %s not ready", name)

		case <-tick:

			deploy, err := clientset.
				AppsV1().
				Deployments(namespace).
				Get(context.Background(), name, metav1.GetOptions{})

			if err != nil {
				return fmt.Errorf("failed to get deployment: %w", err)
			}

			// =========================
			// CRITICAL FAILURE CHECK
			// =========================
			if deploy.Status.UnavailableReplicas > 0 {
				continue // still stabilizing
			}

			// =========================
			// READY CHECK (STRICT)
			// =========================
			if deploy.Status.Replicas > 0 &&
				deploy.Status.ReadyReplicas == deploy.Status.Replicas &&
				deploy.Status.UnavailableReplicas == 0 {
				return nil
			}

			// =========================
			// HARD FAILURE DETECTION
			// =========================
			for _, c := range deploy.Status.Conditions {
				if c.Type == appsv1.DeploymentReplicaFailure {
					return fmt.Errorf("deployment failed: %s", c.Message)
				}
			}
		}
	}
}
