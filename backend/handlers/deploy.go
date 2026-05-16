package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/codeurluce/kubelaunch/models"
	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Deploy déploie une application sur Kubernetes
func Deploy(clientset *kubernetes.Clientset) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.DeployRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Valeurs par défaut
		if req.Namespace == "" {
			req.Namespace = "default"
		}
		if req.Replicas == 0 {
			req.Replicas = 1
		}
		if req.Port == 0 {
			req.Port = 3000
		}

		// Image Docker à utiliser (basée sur le stack)
		image := getDefaultImage(req.Stack)

		// 1. Créer le Secret pour les variables d'environnement
		if len(req.EnvVars) > 0 {
			if err := createSecret(clientset, req); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("erreur Secret: %v", err)})
				return
			}
		}

		// 2. Créer le Deployment
		if err := createDeployment(clientset, req, image); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("erreur Deployment: %v", err)})
			return
		}

		// 3. Créer le Service
		if err := createService(clientset, req); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("erreur Service: %v", err)})
			return
		}

		c.JSON(http.StatusOK, models.DeployResponse{
			Status:         "deploying",
			AppName:        req.AppName,
			Namespace:      req.Namespace,
			DeploymentName: req.AppName,
			Message:        fmt.Sprintf("App %s en cours de déploiement dans le namespace %s", req.AppName, req.Namespace),
		})
	}
}

// createDeployment crée un Deployment Kubernetes
func createDeployment(clientset *kubernetes.Clientset, req models.DeployRequest, image string) error {
	replicas := req.Replicas
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      req.AppName,
			Namespace: req.Namespace,
			Labels:    map[string]string{"app": req.AppName, "managed-by": "kubelaunch"},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": req.AppName},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": req.AppName},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  req.AppName,
							Image: image,
							Ports: []corev1.ContainerPort{
								{ContainerPort: req.Port},
							},
							// Charger les env vars depuis le Secret
							EnvFrom: func() []corev1.EnvFromSource {
								if len(req.EnvVars) > 0 {
									return []corev1.EnvFromSource{{
										SecretRef: &corev1.SecretEnvSource{
											LocalObjectReference: corev1.LocalObjectReference{
												Name: req.AppName + "-secrets",
											},
										},
									}}
								}
								return nil
							}(),
						},
					},
				},
			},
		},
	}

	_, err := clientset.AppsV1().Deployments(req.Namespace).Create(
		context.Background(), deployment, metav1.CreateOptions{},
	)
	return err
}

// createService crée un Service Kubernetes
func createService(clientset *kubernetes.Clientset, req models.DeployRequest) error {
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      req.AppName,
			Namespace: req.Namespace,
			Labels:    map[string]string{"app": req.AppName, "managed-by": "kubelaunch"},
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{"app": req.AppName},
			Ports: []corev1.ServicePort{
				{Port: req.Port, Protocol: corev1.ProtocolTCP},
			},
			Type: corev1.ServiceTypeClusterIP,
		},
	}

	_, err := clientset.CoreV1().Services(req.Namespace).Create(
		context.Background(), service, metav1.CreateOptions{},
	)
	return err
}

// createSecret crée un Secret pour les variables d'environnement
func createSecret(clientset *kubernetes.Clientset, req models.DeployRequest) error {
	stringData := make(map[string]string)
	for k, v := range req.EnvVars {
		stringData[k] = v
	}

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      req.AppName + "-secrets",
			Namespace: req.Namespace,
			Labels:    map[string]string{"app": req.AppName, "managed-by": "kubelaunch"},
		},
		StringData: stringData,
	}

	_, err := clientset.CoreV1().Secrets(req.Namespace).Create(
		context.Background(), secret, metav1.CreateOptions{},
	)
	return err
}

// getDefaultImage retourne une image Docker de base selon le stack
func getDefaultImage(stack models.StackType) string {
	switch stack {
	case models.StackNodeJS, models.StackNextJS:
		return "node:20-alpine"
	case models.StackPython:
		return "python:3.12-slim"
	case models.StackGo:
		return "golang:1.21-alpine"
	default:
		return "nginx:alpine"
	}
}
