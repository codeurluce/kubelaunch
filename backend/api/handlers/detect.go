package handlers

import (
	"net/http"

	"github.com/codeurluce/kubelaunch/backend/core/pipeline"
	"github.com/codeurluce/kubelaunch/backend/models"
	"github.com/gin-gonic/gin"
)

// Detect analyse un repo GitHub et détecte le stack
func Detect(c *gin.Context) {

	var req models.DetectRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "repoUrl est requis",
		})
		return
	}

	result, err := pipeline.RunDetect(req.RepoURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
