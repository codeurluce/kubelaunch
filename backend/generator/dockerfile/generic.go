package dockerfile

import "fmt"

func GenericNode(i Input) string {
	return fmt.Sprintf(`
FROM node:18-alpine

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

EXPOSE %d

CMD ["npm", "start"]
`, i.Port)
}
