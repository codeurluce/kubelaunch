package handlers

import (
	"fmt"
	"net/http"

	"github.com/codeurluce/kubelaunch/backend/core/pipeline"
	"github.com/codeurluce/kubelaunch/backend/k8s"
	"github.com/codeurluce/kubelaunch/backend/models"
	"github.com/codeurluce/kubelaunch/backend/utils"
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

		req.AppName = utils.SanitizeK8sName(
			req.AppName,
		)
		fmt.Println("APP NAME:", req.AppName)

		// =========================
		// PIPELINE EXECUTION
		// =========================

		resp, err := pipeline.Run(clientset, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.DeployResponse{
				Success: false,
				Error:   err.Error(),
			})
			return
		}

		// =========================
		// RUNTIME VERIFICATION
		// =========================

		err = k8s.WaitForHTTP(resp.URL, 10)

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.DeployResponse{
				Success: false,
				Error:   "Deployment failed at runtime: app not reachable",
			})
			return
		}

		// =========================
		// RESPONSE
		// =========================

		c.JSON(http.StatusOK, resp)
	}
}
