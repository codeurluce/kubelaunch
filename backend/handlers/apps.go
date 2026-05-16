package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/codeurluce/kubelaunch/models"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ListApps retourne toutes les apps déployées par KubeLaunch
func ListApps(clientset *kubernetes.Clientset) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lister les Deployments avec le label managed-by=kubelaunch
		deployments, err := clientset.AppsV1().Deployments("").List(
			context.Background(),
			metav1.ListOptions{
				LabelSelector: "managed-by=kubelaunch",
			},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("erreur K8s: %v", err)})
			return
		}

		var apps []models.AppStatus
		for _, d := range deployments.Items {
			status := "pending"
			if d.Status.ReadyReplicas == *d.Spec.Replicas {
				status = "running"
			} else if d.Status.ReadyReplicas == 0 {
				status = "pending"
			}

			port := int32(3000)
			if len(d.Spec.Template.Spec.Containers) > 0 &&
				len(d.Spec.Template.Spec.Containers[0].Ports) > 0 {
				port = d.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort
			}

			apps = append(apps, models.AppStatus{
				Name:          d.Name,
				Namespace:     d.Namespace,
				Port:          port,
				Replicas:      *d.Spec.Replicas,
				ReadyReplicas: d.Status.ReadyReplicas,
				Status:        status,
				CPUUsage:      "—",
				MemoryUsage:   "—",
				CreatedAt:     d.CreationTimestamp.String(),
			})
		}

		if apps == nil {
			apps = []models.AppStatus{}
		}

		c.JSON(http.StatusOK, gin.H{"apps": apps, "total": len(apps)})
	}
}
