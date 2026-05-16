package templates

const ServiceTemplate = `package services

import (
	"{{ .ModulePath }}/repositories"
)

type {{ .Module }}Service struct {
	repo *repositories.{{ .Module }}Repository
}

func New{{ .Module }}Service(
	repo *repositories.{{ .Module }}Repository,
) *{{ .Module }}Service {

	return &{{ .Module }}Service{
		repo: repo,
	}
}
`