package dockerfile

func GenerateDockerfile(runtime string) string {

	switch runtime {

	case "nodejs":
		return `FROM node:20-alpine

WORKDIR /app

COPY . .

RUN npm install

EXPOSE 3000

CMD ["npm","start"]`

	case "python":
		return `FROM python:3.12-alpine

WORKDIR /app

COPY . .

RUN pip install -r requirements.txt

EXPOSE 8000

CMD ["uvicorn","main:app","--host","0.0.0.0","--port","8000"]`

	case "go":
		return `FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o app .

EXPOSE 8080

CMD ["./app"]`

	case "php":
		return `FROM php:8.3-apache

WORKDIR /var/www/html

COPY . .

EXPOSE 80

CMD ["apache2-foreground"]`

	case "static":
		return `FROM nginx:alpine

WORKDIR /usr/share/nginx/html

COPY . /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]`

	default:
		return `FROM alpine:latest`
	}
}
