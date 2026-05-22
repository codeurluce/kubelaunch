package dockerfile

import "fmt"

func Flask(i Input) string {
	return fmt.Sprintf(`
FROM python:3.12-alpine

WORKDIR /app

COPY requirements.txt .
RUN pip install -r requirements.txt

COPY . .

EXPOSE %d

CMD ["python", "app.py"]
`, i.Port)
}
