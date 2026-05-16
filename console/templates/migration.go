package templates

const MigrationTemplate = `CREATE TABLE {{ .Table }} (

	id BIGSERIAL PRIMARY KEY,

{{ range .Fields }}
	{{ .Name }} {{ .SQLType }},
{{ end }}

	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW()
);
`