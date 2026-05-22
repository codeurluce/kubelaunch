package handlers

import (
	"net/http"

	k8sgenerator "github.com/codeurluce/kubelaunch/backend/generator/kubernetes"
	"github.com/codeurluce/kubelaunch/backend/models"
	"github.com/gin-gonic/gin"
)

// PreviewDeploy génère un aperçu YAML avant déploiement
func PreviewDeploy(c *gin.Context) {

	var req models.DeploymentSpec

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "payload invalide",
		})
		return
	}

	deployment := k8sgenerator.GenerateDeployment(req)
	service := k8sgenerator.GenerateService(req)

	c.JSON(http.StatusOK, gin.H{
		"stack":      req.Runtime,
		"port":       req.Port,
		"deployment": deployment,
		"service":    service,
	})
}
