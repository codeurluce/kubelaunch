package analyzer

func DetectPort(runtime Runtime) int {
	switch runtime {
	case Node:
		return 3000
	case Python:
		return 8000
	case Go:
		return 8080
	case Static:
		return 80
	default:
		return 3000
	}
}
