package templates

const MapperTemplate = `package {{ .Package }}

func To{{ .Module }}Response(
	model *{{ .Module }},
) map[string]any {

	return map[string]any{
		"id": model.ID,
	}
}
`