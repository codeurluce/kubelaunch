package analyzer

func DetectPort(runtime string) int {
	switch runtime {
	case "nodejs":
		return 3000
	case "python":
		return 8000
	case "go":
		return 8080
	case "static":
		return 80
	default:
		return 3000
	}
}
