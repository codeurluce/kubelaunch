package analyzer

import (
	"path/filepath"
	"strings"
)

func DetectFramework(files []string) string {

	has := func(name string) bool {
		for _, f := range files {
			base := strings.ToLower(filepath.Base(f))

			if base == strings.ToLower(name) {
				return true
			}
		}
		return false
	}

	// NestJS
	if has("nest-cli.json") ||
		has("tsconfig.spec.json") &&
			has("package.json") {
		return "nestjs"
	}

	// Next.js
	if has("next.config.js") ||
		has("next.config.mjs") {
		return "nextjs"
	}

	// Flask / FastAPI
	if has("pyproject.toml") {

		for _, f := range files {

			if strings.Contains(strings.ToLower(f), "fastapi") {
				return "fastapi"
			}
		}

		return "flask"
	}

	// Laravel
	if has("artisan") ||
		has("composer.json") {
		return "laravel"
	}

	// Gin
	if has("go.mod") {

		for _, f := range files {
			if strings.Contains(f, "gin") {
				return "gin"
			}
		}

		return "go"
	}

	return "unknown"
}
