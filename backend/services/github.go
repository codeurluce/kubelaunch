package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type githubFile struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// ExtractRepoPath extrait owner/repo depuis une URL GitHub
func ExtractRepoPath(url string) (string, error) {

	url = strings.TrimSpace(url)

	url = strings.TrimSuffix(url, "/")

	url = strings.TrimPrefix(url, "https://github.com/")
	url = strings.TrimPrefix(url, "http://github.com/")
	url = strings.TrimPrefix(url, "github.com/")

	url = strings.TrimSuffix(url, ".git")

	parts := strings.Split(url, "/")

	if len(parts) < 2 {
		return "", fmt.Errorf("URL GitHub invalide")
	}

	return parts[0] + "/" + parts[1], nil
}

// FetchRootFiles récupère les fichiers racine du repo
func FetchRootFiles(repoPath string) ([]string, error) {

	apiURL := fmt.Sprintf(
		"https://api.github.com/repos/%s/contents/",
		repoPath,
	)

	req, err := http.NewRequest("GET", apiURL, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "KubeLaunch/0.1")
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	switch resp.StatusCode {

	case http.StatusOK:

	case http.StatusNotFound:
		return nil, fmt.Errorf("repo introuvable")

	case http.StatusForbidden:
		return nil, fmt.Errorf("GitHub rate limit atteint")

	default:
		return nil, fmt.Errorf(
			"GitHub API error (%d)",
			resp.StatusCode,
		)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var githubFiles []githubFile

	if err := json.Unmarshal(body, &githubFiles); err != nil {
		return nil, err
	}

	files := []string{}

	for _, f := range githubFiles {
		files = append(files, strings.ToLower(f.Name))
	}

	return files, nil
}

func FetchFileContent(repoPath string, filePath string) (string, error) {

	apiURL := fmt.Sprintf(
		"https://api.github.com/repos/%s/contents/%s",
		repoPath,
		filePath,
	)

	req, err := http.NewRequest("GET", apiURL, nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "KubeLaunch/0.1")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
