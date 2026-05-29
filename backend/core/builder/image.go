package builder

import (
	"strings"
)

// NormalizeAppName = sécurise le nom image Docker
func NormalizeAppName(name string) string {
	return strings.ToLower(name)
}

// BuildImageTag = standard KubeLaunch
func BuildImageTag(appName string) string {
	clean := NormalizeAppName(appName)
	return "kubelaunch/" + clean + ":latest"
}
