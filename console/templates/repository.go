package templates

const RepositoryTemplate = `package repositories

type {{ .Module }}Repository struct{}

func New{{ .Module }}Repository() *{{ .Module }}Repository {

	return &{{ .Module }}Repository{}
}
`