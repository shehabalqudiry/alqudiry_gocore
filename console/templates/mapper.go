package templates

const MapperTemplate = `package mappers
import "{{ .ModulePath }}/entities"

func To{{ .Module }}Response(
	model *entities.{{ .Module }},
) map[string]any {

	return map[string]any{
		"id": model.ID,
	}
}
`