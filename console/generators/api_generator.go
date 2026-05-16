package generators

import (
	"fmt"
	"path/filepath"

	"github.com/shehabalqudiry/alqudiry_gocore/console/templates"
	"github.com/shehabalqudiry/alqudiry_gocore/console/utils"
)

func GenerateAPI(name string) {

	name = utils.Normalize(name)

	module := utils.ModuleName(name)

	modulePath := filepath.Join(
		"internal",
		"modules",
		name,
	)

	utils.CreateDir(modulePath)

	data := TemplateData{
		Package: name,
		Module:  module,
	}

	files := map[string]string{

		"handler.go": templates.HandlerTemplate,

		"service.go": templates.ServiceTemplate,

		"repository.go": templates.RepositoryTemplate,

		"model.go": templates.ModelTemplate,

		"dto.go": templates.DTOTemplate,

		"routes.go": templates.RoutesTemplate,

		"validator.go": templates.ValidatorTemplate,

		"swagger.go": templates.SwaggerTemplate,

		"mapper.go": templates.MapperTemplate,
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
		"✔ API [%s] generated\n",
		name,
	)
}