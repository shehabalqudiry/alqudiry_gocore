package templates

const ServiceTemplate = `package {{ .Package }}

type {{ .Module }}Service struct {
	repo *{{ .Module }}Repository
}

func New{{ .Module }}Service(
	repo *{{ .Module }}Repository,
) *{{ .Module }}Service {

	return &{{ .Module }}Service{
		repo: repo,
	}
}
`