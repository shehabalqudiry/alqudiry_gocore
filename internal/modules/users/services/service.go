package services

import (
	"github.com/shehabalqudiry/alqudiry_gocore/internal/modules/users/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(
	repo *repositories.UserRepository,
) *UserService {

	return &UserService{
		repo: repo,
	}
}
