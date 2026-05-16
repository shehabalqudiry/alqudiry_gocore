package templates

const RepositoryInterfaceTemplate = `package {{ .Package }}

type {{ .Module }}RepositoryInterface interface {

	Create(data *{{ .Module }}) error

	FindAll() ([]{{ .Module }}, error)

	FindByID(id uint) (*{{ .Module }}, error)

	Update(data *{{ .Module }}) error

	Delete(id uint) error
}
`