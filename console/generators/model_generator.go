package generators

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/shehabalqudiry/alqudiry_gocore/console/templates"
	"github.com/shehabalqudiry/alqudiry_gocore/console/utils"
)

type ModelField struct {
	Name       string
	PublicName string
	Type       string
	SQLType    string
}

type ModelTemplateData struct {
	Package    string
	Module     string
	Table      string
	Fields     []ModelField
	ModulePath string
}

func GenerateModel(
	name string,
	fieldString string,
) {

	name = utils.Normalize(name)

	module := utils.ModuleName(name)

	modulePath := filepath.Join(
		"internal",
		"modules",
		name,
	)

	parsed := utils.ParseFields(fieldString)

	fields := []ModelField{}

	for _, field := range parsed {

		fields = append(fields, ModelField{
			Name:       field.Name,
			PublicName: utils.ToPublicName(field.Name),
			Type:       mapGoType(field.Type),
			SQLType:    mapSQLType(field.Type),
		})
	}

	data := ModelTemplateData{
		Package: name,
		Module:  module,
		Table:   name,
		Fields:  fields,
		ModulePath: "github.com/shehabalqudiry/alqudiry_gocore/internal/modules/" + name,
	}

	dirs := []string{

		"entities",
		"dto",
		"repositories",
		"validators",
		"migrations",
		"swagger",
		"mappers",
		"tests",
	}

	for _, dir := range dirs {

		utils.CreateDir(
			filepath.Join(
				modulePath,
				dir,
			),
		)
	}

	migrationName := fmt.Sprintf(
		"%d_create_%s_table.sql",
		time.Now().Unix(),
		name,
	)

	files := map[string]string{

		"entities/entity.go":
		templates.ModelTemplate,

		"dto/create.go":
		templates.DTOTemplate,

		"repositories/interface.go":
		templates.RepositoryInterfaceTemplate,

		"validators/validator.go":
		templates.ValidatorTemplate,

		"swagger/swagger.go":
		templates.SwaggerTemplate,

		"mappers/mapper.go":
		templates.MapperTemplate,

		filepath.Join(
			"migrations",
			migrationName,
		):
		templates.MigrationTemplate,
	}

	for file, tmpl := range files {

		content, err := utils.ParseTemplate(
			tmpl,
			data,
		)

		if err != nil {
			panic(err)
		}

		path := filepath.Join(
			modulePath,
			file,
		)

		if err := utils.WriteFile(
			path,
			content,
		); err != nil {

			panic(err)
		}
	}

	fmt.Printf(
		"✔ Model [%s] generated successfully\n",
		name,
	)
}

func mapGoType(t string) string {

	switch t {

	case "string":
		return "string"

	case "int":
		return "int"

	case "bool":
		return "bool"

	case "float":
		return "float64"

	case "time":
		return "time.Time"

	default:
		return "string"
	}
}

func mapSQLType(t string) string {

	switch t {

	case "string":
		return "TEXT"

	case "int":
		return "INTEGER"

	case "bool":
		return "BOOLEAN"

	case "float":
		return "DOUBLE PRECISION"

	case "time":
		return "TIMESTAMP"

	default:
		return "TEXT"
	}
}