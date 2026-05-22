package config

import (
	"log"
)

type Config struct {
	ServerPort     string
	GitHubToken    string
	DockerRegistry string
	DockerUsername string
	DockerPassword string
	KubeConfigPath string
	Namespace      string
	Environment    string
}

var AppConfig *Config

func Load() {
	AppConfig = &Config{
		ServerPort:     getEnv("SERVER_PORT", "8080"),
		GitHubToken:    getEnv("GITHUB_TOKEN", ""),
		DockerRegistry: getEnv("DOCKER_REGISTRY", "docker.io"),
		DockerUsername: getEnv("DOCKER_USERNAME", ""),
		DockerPassword: getEnv("DOCKER_PASSWORD", ""),
		KubeConfigPath: getEnv("KUBECONFIG_PATH", ""),
		Namespace:      getEnv("KUBE_NAMESPACE", "default"),
		Environment:    getEnv("ENV", "development"),
	}

	log.Println("Config loaded successfully")
	log.Printf("Env: %s | Port: %s\n", AppConfig.Environment, AppConfig.ServerPort)
}

func IsProduction() bool {
	return AppConfig.Environment == "production"
}
