package analyzer

func DetectPort(
	framework string,
	repoPath string,
) int {

	switch framework {

	case "express":
		return DetectNodePort(repoPath)

	case "nestjs":
		return DetectNodePort(repoPath)

	case "nextjs":
		return DetectNodePort(repoPath)

	case "fastapi":
		return 8000

	case "flask":
		return 5000

	case "gin":
		return 8080

	default:
		return 3000
	}
}
