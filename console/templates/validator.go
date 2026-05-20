package templates

const ValidatorTemplate = `package validators

import (
	"github.com/go-playground/validator/v10"
	"{{ .ModulePath }}/dto"
)

var validate = validator.New()

func ValidateCreate{{ .Module }}(
	dto dto.Create{{ .Module }}DTO,
) error {

	return validate.Struct(dto)
}
`