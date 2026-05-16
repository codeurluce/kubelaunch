package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/codeurluce/kubelaunch/models"
	"github.com/gin-gonic/gin"
)

// githubFile représente un fichier retourné par l'API GitHub
type githubFile struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Detect analyse un repo GitHub et retourne le stack détecté
func Detect(c *gin.Context) {
	var req models.DetectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "repoUrl est requis"})
		return
	}

	// Extraire owner/repo depuis l'URL GitHub
	// Ex: https://github.com/codeurluce/api-service → codeurluce/api-service
	repoPath, err := extractRepoPath(req.RepoURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Appeler l'API GitHub pour lister les fichiers racine
	files, err := fetchRootFiles(repoPath)
	if err != nil {
		// Si l'API GitHub échoue, on fait une détection basée sur l'URL
		result := detectFromURL(req.RepoURL)
		c.JSON(http.StatusOK, result)
		return
	}

	// Détecter le stack depuis les fichiers
	result := detectStack(files)
	c.JSON(http.StatusOK, result)
}

// extractRepoPath extrait "owner/repo" depuis une URL GitHub
func extractRepoPath(url string) (string, error) {
	url = strings.TrimSuffix(url, "/")
	url = strings.TrimPrefix(url, "https://github.com/")
	url = strings.TrimPrefix(url, "http://github.com/")

	parts := strings.Split(url, "/")
	if len(parts) < 2 {
		return "", fmt.Errorf("URL GitHub invalide: %s", url)
	}

	return parts[0] + "/" + parts[1], nil
}

// fetchRootFiles appelle l'API GitHub Contents pour lister les fichiers racine
func fetchRootFiles(repoPath string) ([]string, error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/contents/", repoPath)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("erreur HTTP: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("repo introuvable ou privé (status %d)", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var githubFiles []githubFile
	if err := json.Unmarshal(body, &githubFiles); err != nil {
		return nil, err
	}

	// Extraire juste les noms de fichiers
	var fileNames []string
	for _, f := range githubFiles {
		if f.Type == "file" {
			fileNames = append(fileNames, f.Name)
		}
	}

	return fileNames, nil
}

// detectStack analyse les fichiers et retourne le stack
func detectStack(files []string) models.DetectResponse {
	fileSet := make(map[string]bool)
	for _, f := range files {
		fileSet[f] = true
	}

	// Next.js : package.json + next.config.js ou next.config.ts
	if fileSet["package.json"] && (fileSet["next.config.js"] || fileSet["next.config.ts"] || fileSet["next.config.mjs"]) {
		return models.DetectResponse{
			Stack: models.StackNextJS, Port: 3000,
			Confidence: "high", Framework: "Next.js",
			DetectedFiles: []string{"package.json", "next.config.js"},
		}
	}

	// React + Vite
	if fileSet["package.json"] && (fileSet["vite.config.ts"] || fileSet["vite.config.js"]) {
		return models.DetectResponse{
			Stack: models.StackNodeJS, Port: 5173,
			Confidence: "high", Framework: "React + Vite",
			DetectedFiles: []string{"package.json", "vite.config.ts"},
		}
	}

	// Node.js
	if fileSet["package.json"] {
		return models.DetectResponse{
			Stack: models.StackNodeJS, Port: 3000,
			Confidence: "high", Framework: "Node.js",
			DetectedFiles: []string{"package.json"},
		}
	}

	// Python
	if fileSet["requirements.txt"] || fileSet["pyproject.toml"] || fileSet["setup.py"] {
		detected := []string{}
		if fileSet["requirements.txt"] {
			detected = append(detected, "requirements.txt")
		}
		framework := "Python"
		if fileSet["main.py"] {
			framework = "FastAPI / Flask"
		}
		return models.DetectResponse{
			Stack: models.StackPython, Port: 8000,
			Confidence: "high", Framework: framework,
			DetectedFiles: detected,
		}
	}

	// Go
	if fileSet["go.mod"] {
		return models.DetectResponse{
			Stack: models.StackGo, Port: 8080,
			Confidence: "high", Framework: "Go",
			DetectedFiles: []string{"go.mod"},
		}
	}

	// Docker only
	if fileSet["Dockerfile"] {
		return models.DetectResponse{
			Stack: models.StackDocker, Port: 8080,
			Confidence: "medium", Framework: "Docker",
			DetectedFiles: []string{"Dockerfile"},
		}
	}

	return models.DetectResponse{
		Stack: models.StackUnknown, Port: 8080,
		Confidence: "low", Framework: "Unknown",
		DetectedFiles: []string{},
	}
}

// detectFromURL fallback si l'API GitHub est inaccessible
func detectFromURL(url string) models.DetectResponse {
	url = strings.ToLower(url)
	if strings.Contains(url, "next") {
		return models.DetectResponse{Stack: models.StackNextJS, Port: 3000, Confidence: "low", Framework: "Next.js", DetectedFiles: []string{}}
	}
	if strings.Contains(url, "python") || strings.Contains(url, "fastapi") {
		return models.DetectResponse{Stack: models.StackPython, Port: 8000, Confidence: "low", Framework: "Python", DetectedFiles: []string{}}
	}
	return models.DetectResponse{Stack: models.StackNodeJS, Port: 3000, Confidence: "low", Framework: "Node.js", DetectedFiles: []string{}}
}
