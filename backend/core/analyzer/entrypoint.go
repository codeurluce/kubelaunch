package analyzer

func DetectEntrypoint(runtime string) string {
	switch runtime {

	case "nodejs":
		return "npm start"

	case "python":
		return "uvicorn main:app --host 0.0.0.0 --port 8000"

	case "go":
		return "go build -o app && ./app"

	default:
		return ""
	}
}
