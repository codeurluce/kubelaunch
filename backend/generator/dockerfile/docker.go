package dockerfile

import "github.com/codeurluce/kubelaunch/backend/core/analyzer"

func GenerateDockerfile(runtime analyzer.Runtime) string {

	switch runtime {

	case analyzer.Node:
		return `FROM node:20-alpine
WORKDIR /app
COPY . .
RUN npm install
CMD ["npm","start"]`

	case analyzer.Python:
		return `FROM python:3.12-alpine
WORKDIR /app
COPY . .
RUN pip install -r requirements.txt
CMD ["uvicorn","main:app","--host","0.0.0.0","--port","8000"]`

	case analyzer.Go:
		return `FROM golang:1.21
WORKDIR /app
COPY . .
RUN go build -o app
CMD ["./app"]`

	default:
		return `FROM alpine:latest`
	}
}
