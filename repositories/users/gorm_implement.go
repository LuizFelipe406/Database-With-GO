package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/LuizFelipe406/Database-With-GO/entities"
	"github.com/jinzhu/gorm"
)

type userGormRepo struct {
	writer *gorm.DB
	reader *gorm.DB
}

func NewGormRepo(writer, reader *gorm.DB) UserRepositoryInterface {
	return &userGormRepo{writer, reader}
}

func (repo *userGormRepo) GetAll(ctx context.Context) ([]entities.User, error) {
	var user []entities.User
	repo.reader.Table("users").Find(&user)

	if repo.reader.Error != nil {
		fmt.Println(repo.reader.Error)
		return nil, errors.New("usuário não encontrado")
	}

	return user, nil
}
