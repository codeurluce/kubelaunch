package models

type DetectRequest struct {
	RepoURL string `json:"repoUrl" binding:"required"`
}
