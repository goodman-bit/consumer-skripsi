package generalRepository

import (
	"consumerskripsi/models"
	"consumerskripsi/repositories"
)

type UsersRepository struct {
	RepoDB repositories.Repository
}

func NewUsersRepository(repoDB repositories.Repository) UsersRepository {
	return UsersRepository{
		RepoDB: repoDB,
	}
}

func (ctx UsersRepository) GetUserById(userId int64) (models.Users, error) {
	var result models.Users

	return result, nil
}
