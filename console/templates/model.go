package templates

const ModelTemplate = `package entities

type {{ .Module }} struct {
	ID uint ` + "`gorm:\"primaryKey\"`" + `
{{ range .Fields }}
	{{ .PublicName }} {{ .Type }} ` + "`json:\"{{ .Name }}\"`" + `
{{ end }}
}
`