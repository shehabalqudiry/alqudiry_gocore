package templates

const EntityTemplate = `package entities

type {{ .Module }} struct {
	ID uint ` + "`gorm:\"primaryKey\"`" + `
}
`