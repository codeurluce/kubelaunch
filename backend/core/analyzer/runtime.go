package analyzer

import (
	"path/filepath"
	"strings"
)

type Runtime string

const (
	Node    Runtime = "nodejs"
	Python  Runtime = "python"
	Go      Runtime = "go"
	Rust    Runtime = "rust"
	PHP     Runtime = "php"
	Static  Runtime = "static"
	Unknown Runtime = "unknown"
)

func DetectRuntime(files []string) string {

	has := func(name string) bool {

		for _, f := range files {

			base := strings.ToLower(filepath.Base(f))

			if base == strings.ToLower(name) {
				return true
			}
		}

		return false
	}

	// =========================
	// PHP
	// =========================
	if has("composer.json") || has("artisan") {
		return "php"
	}

	// =========================
	// NODEJS
	// =========================
	if has("package.json") ||
		has("pnpm-workspace.yaml") ||
		has("turbo.json") {
		return "nodejs"
	}

	// =========================
	// PYTHON
	// =========================
	if has("requirements.txt") ||
		has("pyproject.toml") ||
		has("setup.py") {
		return "python"
	}

	// =========================
	// GO
	// =========================
	if has("go.mod") {
		return "go"
	}

	// =========================
	// RUST
	// =========================
	if has("cargo.toml") {
		return "rust"
	}

	// =========================
	// STATIC
	// =========================
	if has("index.html") {
		return "static"
	}

	return "unknown"
}
