package main

import (
	"log"
	"net/http"

	"github.com/codeurluce/kubelaunch/backend/api/handlers"
	"github.com/codeurluce/kubelaunch/backend/k8s"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connexion au cluster Kubernetes
	clientset, err := k8s.NewClient()
	if err != nil {
		log.Fatalf("Impossible de se connecter à Kubernetes: %v", err)
	}
	log.Println("✅ Connecté au cluster Kubernetes")

	// Créer le router Gin
	r := gin.Default()

	// CORS — permet au frontend Next.js de parler au backend
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// Routes
	api := r.Group("/api")
	{
		// Health check
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "KubeLaunch backend is running 🚀",
			})
		})

		// Détecter le stack d'un repo GitHub
		api.POST("/detect", handlers.Detect)

		// Déployer une application
		api.POST("/deploy", handlers.Deploy(clientset))

		// Lister les applications déployées
		api.GET("/apps", handlers.ListApps(clientset))
	}

	log.Println("🚀 KubeLaunch backend démarré sur http://localhost:8080")
	r.Run(":8080")
}
