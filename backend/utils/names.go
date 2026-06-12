package utils

import (
	"regexp"
	"strings"
)

func SanitizeK8sName(name string) string {

	name = strings.ToLower(name)

	name = strings.ReplaceAll(name, "_", "-")
	name = strings.ReplaceAll(name, " ", "-")

	reg := regexp.MustCompile(`[^a-z0-9\-]`)
	name = reg.ReplaceAllString(name, "")

	// Enforce 63-character limit
    if len(name) > 63 {
        name = name[:63]
    }
	
	return strings.Trim(name, "-")
}
