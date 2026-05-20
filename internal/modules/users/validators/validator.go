package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/shehabalqudiry/alqudiry_gocore/internal/modules/users/dto"
)

var validate = validator.New()

func ValidateCreateUser(
	dto dto.CreateUserDTO,
) error {

	return validate.Struct(dto)
}
