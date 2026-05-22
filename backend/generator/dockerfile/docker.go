package dockerfile

func Generate(input Input) string {
	switch input.Framework {

	case "nextjs":
		return NextJS(input)

	case "nestjs":
		return NestJS(input)

	case "flask":
		return Flask(input)

	case "fastapi":
		return FastAPI(input)

	case "laravel":
		return Laravel(input)

	default:
		return GenericNode(input)
	}
}
