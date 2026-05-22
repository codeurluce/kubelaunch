package dockerfile

import "fmt"

func NextJS(i Input) string {
	return fmt.Sprintf(`
FROM node:18-alpine

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

RUN npm run build

EXPOSE %d

CMD ["npm", "start"]
`, i.Port)
}
