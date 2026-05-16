package templates

const ValidatorTemplate = `package {{ .Package }}

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func ValidateCreate{{ .Module }}(
	dto Create{{ .Module }}DTO,
) error {

	return validate.Struct(dto)
}
`