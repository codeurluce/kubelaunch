package analyzer

func DetectEntrypoint(runtime Runtime) string {
	switch runtime {

	case Node:
		return "npm start"

	case Python:
		return "uvicorn main:app --host 0.0.0.0 --port 8000"

	case Go:
		return "go build -o app && ./app"

	default:
		return ""
	}
}
