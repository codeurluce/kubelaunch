package dockerfile

import "fmt"

func NestJS(i Input) string {
	return fmt.Sprintf(`
FROM node:18-alpine

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

RUN npm run build

EXPOSE %d

CMD ["node dist/main.js"]
`, i.Port)
}
