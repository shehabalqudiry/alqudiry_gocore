package templates

const ModelTemplate = `package {{ .Package }}

type {{ .Module }} struct {
	ID uint ` + "`gorm:\"primaryKey\"`" + `
}
`