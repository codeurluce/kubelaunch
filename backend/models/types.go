package models

// StackType représente le type de stack détecté
type StackType string

const (
	StackNodeJS  StackType = "nodejs"
	StackNextJS  StackType = "nextjs"
	StackPython  StackType = "python"
	StackGo      StackType = "go"
	StackPHP     StackType = "php"
	StackRuby    StackType = "ruby"
	StackJava    StackType = "java"
	StackRust    StackType = "rust"
	StackDocker  StackType = "docker"
	StackStatic  StackType = "static"
	StackUnknown StackType = "unknown"
)
