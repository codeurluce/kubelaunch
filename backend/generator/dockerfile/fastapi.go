package dockerfile

import "fmt"

func FastAPI(i Input) string {
	return fmt.Sprintf(`
FROM python:3.12-alpine

WORKDIR /app

COPY requirements.txt .
RUN pip install -r requirements.txt

COPY . .

EXPOSE %d

CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "%d"]
`, i.Port, i.Port)
}
