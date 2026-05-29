package handlers

import (
	"net/http"

	"github.com/codeurluce/kubelaunch/backend/core/pipeline"
	"github.com/codeurluce/kubelaunch/backend/models"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
)

// Deploy API handler (KubeLaunch platform entrypoint)
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
		// DEFAULTS (safe fallback)
		// =========================

		if req.Namespace == "" {
			req.Namespace = "default"
		}

		if req.Replicas == 0 {
			req.Replicas = 1
		}

		// =========================
		// PIPELINE EXECUTION
		// =========================

		resp, err := pipeline.Run(clientset, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// =========================
		// RESPONSE
		// =========================

		c.JSON(http.StatusOK, resp)
	}
}
