package analyzer

import "strings"

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

func DetectRuntime(files []string) Runtime {

	has := func(name string) bool {
		for _, f := range files {
			if strings.EqualFold(f, name) {
				return true
			}
		}
		return false
	}

	// PHP FIRST
	if has("composer.json") || has("artisan") {
		return PHP
	}

	// Node.js
	if has("package.json") ||
		has("pnpm-workspace.yaml") ||
		has("turbo.json") {
		return Node
	}

	// Python
	if has("requirements.txt") ||
		has("pyproject.toml") ||
		has("setup.py") {
		return Python
	}

	// Go
	if has("go.mod") {
		return Go
	}

	// Rust
	if has("cargo.toml") {
		return Rust
	}

	// Static
	if has("index.html") {
		return Static
	}

	return Unknown
}
