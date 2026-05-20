package generators

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/shehabalqudiry/alqudiry_gocore/console/templates"
	"github.com/shehabalqudiry/alqudiry_gocore/console/ui"
	"github.com/shehabalqudiry/alqudiry_gocore/console/utils"
)

type TemplateData struct {
	Package   string
	Module    string
	ModulePath string
}

func GenerateModule(name string) {

	ui.Banner()

	name = utils.Normalize(name)

	module := utils.ModuleName(name)

	modulePath := filepath.Join(
		"internal",
		"modules",
		name,
	)

	fmt.Println()
	logStep("Initializing module generator", "DONE")

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

	fmt.Println()
	fmt.Println("📁 Creating module directories")
	fmt.Println("────────────────────────────────────────────")

	for _, dir := range dirs {

		path := filepath.Join(
			modulePath,
			dir,
		)

		utils.CreateDir(path)

		logStep(
			fmt.Sprintf(
				"Created %s",
				path,
			),
			"OK",
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

	fmt.Println()
	fmt.Println("⚙️ Generating module files")
	fmt.Println("────────────────────────────────────────────")

	total := len(files)
	current := 0

	for file, tmpl := range files {

		current++

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

		ui.Progress(
			current,
			total,
			file,
		)
	}

	fmt.Println()
	fmt.Println()

	fmt.Printf(
		"\033[1;32m✔ Module [%s] generated successfully\033[0m\n",
		module,
	)

	fmt.Println()

	fmt.Println("📦 Generated Architecture")
	fmt.Println("────────────────────────────────────────────")

	for _, dir := range dirs {

		fmt.Printf(
			"\033[1;36m•\033[0m %s\n",
			dir,
		)
	}

	fmt.Println()

	fmt.Println("\033[1;35m🚀 Ready for development\033[0m")
	fmt.Println()
}


func logStep(message string, status string) {

	color := "\033[1;32m"

	if status != "DONE" && status != "OK" {
		color = "\033[1;33m"
	}

	fmt.Printf(
		"%s[%s]\033[0m %s\n",
		color,
		status,
		message,
	)

	time.Sleep(80 * time.Millisecond)
}

