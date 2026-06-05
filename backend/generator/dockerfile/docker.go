package dockerfile

func Generate(input Input) string {

	if input.Runtime == "static" {
		return Static(input)
	}

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
