package templates

const RepositoryTemplate = `package {{ .Package }}

type {{ .Module }}Repository struct{}

func New{{ .Module }}Repository() *{{ .Module }}Repository {

	return &{{ .Module }}Repository{}
}
`