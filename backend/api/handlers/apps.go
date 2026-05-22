package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/codeurluce/kubelaunch/backend/models"
	"github.com/gin-gonic/gin"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ListApps retourne toutes les apps déployées par KubeLaunch
func ListApps(clientset *kubernetes.Clientset) gin.HandlerFunc {
	return func(c *gin.Context) {

		deployments, err := clientset.AppsV1().
			Deployments("").
			List(
				context.Background(),
				metav1.ListOptions{
					LabelSelector: "managed-by=kubelaunch",
				},
			)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("erreur kubernetes: %v", err),
			})
			return
		}

		apps := []models.AppStatus{}

		for _, d := range deployments.Items {

			// =========================
			// REPLICAS
			// =========================

			var replicas int32 = 1
			if d.Spec.Replicas != nil {
				replicas = *d.Spec.Replicas
			}

			// =========================
			// STATUS
			// =========================

			status := "pending"

			switch {
			case d.Status.ReadyReplicas == replicas:
				status = "running"

			case d.Status.ReadyReplicas > 0:
				status = "starting"

			case d.Status.ReadyReplicas == 0:
				status = "pending"
			}

			// =========================
			// PORT
			// =========================

			var port int32 = 3000

			if len(d.Spec.Template.Spec.Containers) > 0 {
				container := d.Spec.Template.Spec.Containers[0]

				if len(container.Ports) > 0 {
					port = container.Ports[0].ContainerPort
				}
			}

			// =========================
			// STACK
			// =========================

			stack := models.StackUnknown

			if value, ok := d.Labels["stack"]; ok {
				stack = models.StackType(value)
			}

			// =========================
			// APP
			// =========================

			app := models.AppStatus{
				Name:          d.Name,
				Namespace:     d.Namespace,
				Stack:         stack,
				Port:          port,
				Replicas:      replicas,
				ReadyReplicas: d.Status.ReadyReplicas,
				Status:        status,
				CPUUsage:      "—",
				MemoryUsage:   "—",
				CreatedAt:     d.CreationTimestamp.String(),
			}

			apps = append(apps, app)
		}

		c.JSON(http.StatusOK, gin.H{
			"apps":  apps,
			"total": len(apps),
		})
	}
}
