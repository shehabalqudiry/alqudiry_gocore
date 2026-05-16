package templates

const DTOTemplate = `package dto

type Create{{ .Module }}DTO struct {
{{ range .Fields }}
	{{ .PublicName }} {{ .Type }} ` + "`json:\"{{ .Name }}\" validate:\"required\"`" + `
{{ end }}
}
`