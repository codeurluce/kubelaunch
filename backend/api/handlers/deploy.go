package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/codeurluce/kubelaunch/backend/models"
	"github.com/gin-gonic/gin"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/intstr"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Deploy déploie une application Kubernetes
func Deploy(clientset *kubernetes.Clientset) gin.HandlerFunc {

	return func(c *gin.Context) {

		var req models.DeployRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// =========================
		// DEFAULTS
		// =========================

		applyDefaults(&req)

		// =========================
		// IMAGE
		// =========================

		image := getDefaultImage(req.Stack)

		// =========================
		// SECRET
		// =========================

		if len(req.EnvVars) > 0 {

			if err := createSecret(clientset, req); err != nil {

				c.JSON(http.StatusInternalServerError, gin.H{
					"error": fmt.Sprintf("secret error: %v", err),
				})

				return
			}
		}

		// =========================
		// DEPLOYMENT
		// =========================

		if err := createDeployment(clientset, req, image); err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("deployment error: %v", err),
			})

			return
		}

		// =========================
		// SERVICE
		// =========================

		if err := createService(clientset, req); err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("service error: %v", err),
			})

			return
		}

		// =========================
		// RESPONSE
		// =========================

		c.JSON(http.StatusOK, models.DeployResponse{
			Status:         "deploying",
			AppName:        req.AppName,
			Namespace:      req.Namespace,
			DeploymentName: req.AppName,
			Message: fmt.Sprintf(
				"%s deployment started in namespace %s",
				req.AppName,
				req.Namespace,
			),
		})
	}
}

// =========================
// DEFAULTS
// =========================

func applyDefaults(req *models.DeployRequest) {

	if req.Namespace == "" {
		req.Namespace = "default"
	}

	if req.Replicas == 0 {
		req.Replicas = 1
	}

	if req.Port == 0 {
		req.Port = 3000
	}
}

// =========================
// LABELS
// =========================

func buildLabels(req models.DeployRequest) map[string]string {

	return map[string]string{
		"app":        req.AppName,
		"managed-by": "kubelaunch",
		"stack":      string(req.Stack),
	}
}

// =========================
// DEPLOYMENT
// =========================

func createDeployment(
	clientset *kubernetes.Clientset,
	req models.DeployRequest,
	image string,
) error {

	labels := buildLabels(req)

	replicas := req.Replicas

	container := corev1.Container{
		Name:            req.AppName,
		Image:           image,
		ImagePullPolicy: corev1.PullIfNotPresent,

		Ports: []corev1.ContainerPort{
			{
				ContainerPort: req.Port,
			},
		},

		ReadinessProbe: &corev1.Probe{
			InitialDelaySeconds: 5,
			PeriodSeconds:       10,
			ProbeHandler: corev1.ProbeHandler{
				TCPSocket: &corev1.TCPSocketAction{
					Port: intstrFromInt32(req.Port),
				},
			},
		},

		LivenessProbe: &corev1.Probe{
			InitialDelaySeconds: 15,
			PeriodSeconds:       20,
			ProbeHandler: corev1.ProbeHandler{
				TCPSocket: &corev1.TCPSocketAction{
					Port: intstrFromInt32(req.Port),
				},
			},
		},
	}

	// ENV VARS
	if len(req.EnvVars) > 0 {

		container.EnvFrom = []corev1.EnvFromSource{
			{
				SecretRef: &corev1.SecretEnvSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: req.AppName + "-secrets",
					},
				},
			},
		}
	}

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      req.AppName,
			Namespace: req.Namespace,
			Labels:    labels,
		},

		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,

			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": req.AppName,
				},
			},

			Template: corev1.PodTemplateSpec{

				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},

				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						container,
					},
				},
			},
		},
	}

	_, err := clientset.
		AppsV1().
		Deployments(req.Namespace).
		Create(
			context.Background(),
			deployment,
			metav1.CreateOptions{},
		)

	if apierrors.IsAlreadyExists(err) {
		return fmt.Errorf("deployment already exists")
	}

	return err
}

// =========================
// SERVICE
// =========================

func createService(
	clientset *kubernetes.Clientset,
	req models.DeployRequest,
) error {

	labels := buildLabels(req)

	service := &corev1.Service{

		ObjectMeta: metav1.ObjectMeta{
			Name:      req.AppName,
			Namespace: req.Namespace,
			Labels:    labels,
		},

		Spec: corev1.ServiceSpec{

			Selector: map[string]string{
				"app": req.AppName,
			},

			Ports: []corev1.ServicePort{
				{
					Port:       req.Port,
					TargetPort: intstrFromInt32(req.Port),
					Protocol:   corev1.ProtocolTCP,
				},
			},

			Type: corev1.ServiceTypeClusterIP,
		},
	}

	_, err := clientset.
		CoreV1().
		Services(req.Namespace).
		Create(
			context.Background(),
			service,
			metav1.CreateOptions{},
		)

	if apierrors.IsAlreadyExists(err) {
		return fmt.Errorf("service already exists")
	}

	return err
}

// =========================
// SECRET
// =========================

func createSecret(
	clientset *kubernetes.Clientset,
	req models.DeployRequest,
) error {

	labels := buildLabels(req)

	secret := &corev1.Secret{

		ObjectMeta: metav1.ObjectMeta{
			Name:      req.AppName + "-secrets",
			Namespace: req.Namespace,
			Labels:    labels,
		},

		StringData: req.EnvVars,
	}

	_, err := clientset.
		CoreV1().
		Secrets(req.Namespace).
		Create(
			context.Background(),
			secret,
			metav1.CreateOptions{},
		)

	if apierrors.IsAlreadyExists(err) {
		return nil
	}

	return err
}

// =========================
// DEFAULT IMAGE
// =========================

func getDefaultImage(stack models.StackType) string {

	switch stack {

	case models.StackNodeJS:
		return "node:20-alpine"

	case models.StackNextJS:
		return "node:20-alpine"

	case models.StackPython:
		return "python:3.12-slim"

	case models.StackGo:
		return "golang:1.21-alpine"

	case models.StackPHP:
		return "php:8.3-apache"

	case models.StackJava:
		return "eclipse-temurin:21"

	case models.StackRuby:
		return "ruby:3.3-alpine"

	default:
		return "nginx:alpine"
	}
}

// =========================
// HELPERS
// =========================

func intstrFromInt32(port int32) intstr.IntOrString {
	return intstr.FromInt(int(port))
}
