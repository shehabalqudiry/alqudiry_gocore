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
}

func GenerateModule(name string) {

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