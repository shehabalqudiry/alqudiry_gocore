package mappers
import "github.com/shehabalqudiry/alqudiry_gocore/internal/modules/users/entities"

func ToUserResponse(
	model *entities.User,
) map[string]any {

	return map[string]any{
		"id": model.ID,
	}
}
