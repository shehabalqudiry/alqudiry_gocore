package generators

import (
	"fmt"
	"path/filepath"

	"github.com/shehabalqudiry/alqudiry_gocore/console/templates"
	"github.com/shehabalqudiry/alqudiry_gocore/console/utils"
)

type TemplateData struct {
	Package string
	Module  string
	ModulePath string
}

func GenerateModule(name string) {

	name = utils.Normalize(name)

	module := utils.ModuleName(name)

	modulePath := filepath.Join(
		"internal",
		"modules",
		name,
	)

	dirs := []string{

		"entities",
		"repositories",
		"services",
		"handlers",
		"dto",
		"validators",
		"routes",
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

	data := TemplateData{
		Package: name,
		Module:  module,
		ModulePath: "github.com/shehabalqudiry/alqudiry_gocore/internal/modules/" + name,
	}

	files := map[string]string{

		"handlers/handler.go":
		templates.HandlerTemplate,

		"services/service.go":
		templates.ServiceTemplate,

		"repositories/repository.go":
		templates.RepositoryTemplate,

		"entities/entity.go":
		templates.EntityTemplate,

		"dto/create.go":
		templates.DTOTemplate,

		"routes/routes.go":
		templates.RoutesTemplate,

		"validators/validator.go":
		templates.ValidatorTemplate,

		"swagger/swagger.go":
		templates.SwaggerTemplate,

		"mappers/mapper.go":
		templates.MapperTemplate,
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
		"✔ Module [%s] generated\n",
		name,
	)
}